<template>
  <div>
    <!-- v-loading.fullscreen.lock="fullscreenLoading" -->
    <div class="table-box">
      <div class="btn-list">
        <upload-common v-model:imageCommon="imageCommon" class="upload-btn" @on-success="getTableData" />
        <upload-image v-model:imageUrl="imageUrl" :file-size="512" :max-w-h="1080" class="upload-btn" @on-success="getTableData" />
        <el-form ref="searchForm" :inline="true" :model="search" @keyup.enter.native="getTableData">
          <el-form-item label="">
            <el-input v-model="search.keyword" class="keyword" placeholder="请输入文件名或备注" />
          </el-form-item>

          <el-form-item>
            <el-button size="small" type="primary" icon="search" @click="getTableData">查询</el-button>
          </el-form-item>
        </el-form>
      </div>

      <el-table :data="tableData">
        <el-table-column align="left" label="预览" width="100">
          <template #default="scope">
            <div @click="() => changePicker(true, scope.row.url, scope.row.name)">
              <!-- <viewer :images="[scope.row.url]"><CustomPic pic-type="file" :pic-src="scope.row.url" /></viewer> -->
              <CustomPic pic-type="file" :pic-src="scope.row.url" />
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="日期" prop="UpdatedAt" width="180">
          <template #default="scope">
            <div>{{ formatDate(scope.row.UpdatedAt) }}</div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="文件名/备注" prop="name" width="180">
          <template #default="scope">
            <div class="name" @click="editFileNameFunc(scope.row)">
              {{ scope.row.name }}
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="链接" props="url" min-width="300">
          <template #default="scope">
            <div>{{ scope.row.url }}</div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="标签" props="tag" width="100">
          <template #default="scope">
            <el-tag :type="scope.row.tag === 'jpg' ? 'warning' : 'success'" disable-transitions>
              {{ scope.row.tag }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="类型" props="is_cropper" width="100">
          <template #default="scope">
            <el-tag :type="scope.row.is_cropper ? 'warning' : 'success'" disable-transitions>
              {{ imageType[scope.row.is_cropper] }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作" width="160">
          <template #default="scope">
            <el-button size="small" icon="download" link type="primary" @click="downloadFile(scope.row)">下载</el-button>
            <el-button size="small" icon="delete" link type="primary" @click="deleteFileFunc(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="pagination">
        <el-pagination background :current-page="page" :page-size="pageSize" :page-sizes="[10, 30, 50, 100]" :style="{ float: 'right', padding: '20px' }" :total="total" layout="total, sizes, prev, pager, next, jumper" @current-change="handleCurrentChange" @size-change="handleSizeChange" />
      </div>
      <!-- <VueViewer :images="tableData.map((i) => i.url)">
        <template slot-scope="scope">
          <img v-for="src in scope.images" :src="src" :key="src" class="image" />
        </template>
      </VueViewer> -->
      <el-dialog v-model="dialogFormVisible" title="裁剪图片">
        <div style="width: 100%; height: 400px">
          <vue-cropper
            ref="cropper"
            :img="option.img"
            :output-size="option.size"
            :output-type="option.outputType"
            :info="true"
            :full="option.full"
            :fixed="false"
            :fixed-number="[75, 34]"
            :can-move="option.canMove"
            :can-move-box="option.canMoveBox"
            :fixed-box="option.fixedBox"
            :original="option.original"
            :auto-crop="option.autoCrop"
            :auto-crop-width="option.autoCropWidth"
            :auto-crop-height="option.autoCropHeight"
            :center-box="option.centerBox"
            @real-time="realTime"
            :high="option.high"
            @img-load="imgLoad"
            mode="contain"
            :max-img-size="option.max"
            @crop-moving="cropMoving"
          ></vue-cropper>
        </div>
        <template #footer>
          <span class="dialog-footer">
            <el-button @click="oncancel">取消</el-button>
            <el-button type="primary" @click="onOk">确定</el-button>
          </span>
        </template>
      </el-dialog>
    </div>
  </div>
</template>

<script setup>
import { getFileList, deleteFile, editFileName, uploadFile } from "@/api/fileUploadAndDownload";
import { downloadImage } from "@/utils/downloadImg";
import CustomPic from "@/components/customPic/index.vue";
import UploadImage from "@/components/upload/image.vue";
import UploadCommon from "@/components/upload/common.vue";
import { formatDate } from "@/utils/format";
import { VueCropper } from "vue-cropper";
// import VueViewer from "v-viewer";
import { ref, onMounted } from "vue";
import { ElMessage, ElMessageBox } from "element-plus";

const cropper = ref(null);
const path = ref(import.meta.env.VITE_BASE_API + "/");
// const userStore = useUserStore();
const imageUrl = ref("");
const imageCommon = ref("");
const filename = ref("");

const dialogFormVisible = ref(false);
const option = ref({
  img: "",
  size: 1,
  full: false,
  outputType: "png",
  canMove: true,
  fixedBox: false,
  original: false,
  canMoveBox: true,
  autoCrop: true,
  // 只有自动截图开启 宽度高度才生效
  autoCropWidth: 750,
  autoCropHeight: 340,
  centerBox: true,
  high: true,
  max: 99999,
});

const imageType = {
  1: "后台图片",
  2: "后台截图",
  3: "前端图片",
  4: "移动端图片",
};

const page = ref(1);
const total = ref(0);
const pageSize = ref(10);
const search = ref({});
const tableData = ref([]);

const cropMoving = (e) => {
  console.log(e);
};

const realTime = (e) => {
  console.log(e);
};

const imgLoad = (e) => {
  console.log(e);
};

const oncancel = () => {
  changePicker(false, "");
};

const onOk = () => {
  // console.log(cropper.value);
  cropper.value.getCropBlob(async (data) => {
    const file = new window.File([data], filename.value, { type: data.type });
    const forms = new FormData();
    forms.append("file", file);
    const resp = await uploadFile(forms, 2);
    if (resp.code === 0) {
      ElMessage({
        type: "success",
        message: "截图成功 上传成功！",
      });
      getTableData();
      changePicker();
    }
  });
};

const changePicker = (status = false, pickerUrl = "", name = "") => {
  dialogFormVisible.value = status;
  if (pickerUrl) {
    filename.value = "screenshot_" + name;
    option.value.img = path.value + pickerUrl;
  } else {
    option.value.img = "";
    filename.value = "";
  }
};

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val;
  getTableData();
};

const handleCurrentChange = (val) => {
  page.value = val;
  getTableData();
};

// 查询
const getTableData = async () => {
  const table = await getFileList({
    page: page.value,
    pageSize: pageSize.value,
    ...search.value,
  });
  if (table.code === 0) {
    tableData.value = table.data.list;
    total.value = table.data.total;
    page.value = table.data.page;
    pageSize.value = table.data.pageSize;
    // console.log(table);
  }
};
onMounted(() => {
  getTableData();
});

const deleteFileFunc = async (row) => {
  ElMessageBox.confirm("此操作将永久文件, 是否继续?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  })
    .then(async () => {
      const res = await deleteFile(row);
      if (res.code === 0) {
        ElMessage({
          type: "success",
          message: "删除成功!",
        });
        if (tableData.value.length === 1 && page.value > 1) {
          page.value--;
        }
        getTableData();
      }
    })
    .catch(() => {
      ElMessage({
        type: "info",
        message: "已取消删除",
      });
    });
};

const downloadFile = (row) => {
  if (row.url.indexOf("http://") > -1 || row.url.indexOf("https://") > -1) {
    downloadImage(row.url, row.name);
  } else {
    downloadImage(path.value + row.url, row.name);
  }
};

/**
 * 编辑文件名或者备注
 * @param row
 * @returns {Promise<void>}
 */
const editFileNameFunc = async (row) => {
  // console.log(row);
  ElMessageBox.prompt("请输入文件名或者备注", "编辑", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    inputPattern: /\S/,
    inputErrorMessage: "不能为空",
    inputValue: row.name,
  })
    .then(async ({ value }) => {
      row.name = value;
      // console.log(value);
      const res = await editFileName(row);
      if (res.code === 0) {
        ElMessage({
          type: "success",
          message: "编辑成功!",
        });
        getTableData();
      }
    })
    .catch(() => {
      ElMessage({
        type: "warning",
        message: "取消修改",
      });
    });
};
</script>

<script>
export default {
  name: "Upload",
};
</script>
<style scoped>
.name {
  cursor: pointer;
  color: #4d70ff;
}
.upload-btn + .upload-btn {
  margin-left: 12px;
}
</style>
