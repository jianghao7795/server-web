<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searForm" :model="searchInfo" class="demo-form-inline" :inline="true">
        <el-form-item label="名称:">
          <el-input v-model="searchInfo.title"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="reset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" size="small" icon="plus" @click="openDialog">新增</el-button>
        <el-popconfirm title="确定要删除吗?" @confirm="onDelete" placement="top">
          <template #reference>
            <el-button icon="delete" size="small" style="margin-left: 10px" :disabled="!multipleSelection.length">删除</el-button>
          </template>
        </el-popconfirm>
      </div>
      <el-table
        :data="tableData"
        style="width: 100%"
        row-key="ID"
        tooltip-effect="dark"
        ref="multipleTable"
        @selection-change="handleSelectionChange"
        v-loading="loadingInit"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column props="ID" width="55" />
        <el-table-column label="标题" prop="title"></el-table-column>
        <el-table-column label="标签" prop="tag">
          <template #default="{ row }">
            <div>{{ row?.tag?.name }}</div>
          </template>
        </el-table-column>
        <el-table-column label="简化内容" prop="desc"></el-table-column>
        <el-table-column label="创建时间">
          <template #default="{ row }">
            {{ formatDate(row.CreatedAt) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作">
          <template #default="scope">
            <el-button link type="primary" icon="edit" size="small" class="table-button" @click="updateArticleFunc(scope.row)">编辑</el-button>
            <el-popconfirm title="确认删除？" placement="top" v-on:confirm="deleteRow(scope.row)">
              <template #reference><el-button type="primary" link icon="delete" size="small">删除</el-button></template>
            </el-popconfirm>
          </template>
        </el-table-column>
      </el-table>
      <div class="pagination">
        <el-pagination
          background
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="type === 'update' ? '更新文章' : '新建文章'" draggable :width="1100">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="标签:">
          <el-select v-model="formData.tag_id" placeholder="请选择" filterable remote reserve-keyword :remote-method="searchTag" :loading="loading">
            <el-option v-for="item in tags" :key="item.ID" :label="item.name" :value="item.ID"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="标题">
          <el-input v-model="formData.title" placeholder="请输入"></el-input>
        </el-form-item>
        <el-form-item label="简化内容">
          <el-input v-model="formData.desc" placeholder="请输入"></el-input>
        </el-form-item>
        <el-form-item label="内容">
          <MdEditor style="width: 1000px" v-model="text" @on-upload-img="onUploadImg"></MdEditor>
        </el-form-item>
        <el-form-item label="状态:">
          <el-radio-group v-model="formData.state">
            <el-radio-button :label="1">显示</el-radio-button>
            <el-radio-button :label="0">隐藏</el-radio-button>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default { name: "Article" };
</script>

<script setup>
import { formatDate } from "@/utils/format";
import { ElMessage } from "element-plus";
import { getArticleList, deleteArticle, findArticle, createArticle, updateArticle, uploadFile } from "@/api/article";
import { getAppTabList } from "@/api/appTab";
import { ref, onBeforeMount } from "vue";
import { useDebounceFn } from "@vueuse/core";
import MdEditor from "md-editor-v3";
import "md-editor-v3/lib/style.css";

const formData = ref({
  title: "",
  tag_id: undefined,
  desc: "",
  content: "",
  state: 1,
});

// search tag loading
const loading = ref(false);
const loadingInit = ref(false);
const text = ref("");

const tags = ref([]);

const page = ref(1);
const total = ref(0);
const pageSize = ref(10);
const tableData = ref([]);
const searchInfo = ref({});
const path = ref(import.meta.env.VITE_BASE_API);

const getTableData = async () => {
  loadingInit.value = true;
  const table = await getArticleList({
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo.value,
  });
  loadingInit.value = false;
  if (table.code === 0) {
    tableData.value = table.data.list;
    total.value = table.data.total;
    page.value = table.data.page;
    pageSize.value = table.data.pageSize;
  }
};

const onUploadImg = async (files, callback) => {
  const res = await Promise.all(
    files.map((file) => {
      return new Promise((rev, rej) => {
        const form = new FormData();
        form.append("file", file);
        uploadFile({
          form,
          headers: {
            "Content-Type": "multipart/form-data",
          },
        })
          .then((res) => rev(res))
          .catch((error) => rej(error));
      });
    }),
  );
  // console.log(res);
  callback(res.map((item) => `${path.value}/${item.data.file.url}`));
};

onBeforeMount(() => {
  getTableData();
});

const onSubmit = () => {
  page.value = 1;
  pageSize.value = 10;
  getTableData();
};

const reset = () => {
  page.value = 1;
  pageSize.value = 10;
  searchInfo.value = {};
  getTableData();
};

// 多选数据
const multipleSelection = ref([]);
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val;
};
// 多选删除
const onDelete = async () => {
  const ids = [];
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: "warning",
      message: "请选择要删除的数据",
    });
    return;
  }
  multipleSelection.value &&
    multipleSelection.value.map((item) => {
      ids.push(item.ID);
    });
  const res = await deleteAppTabByIds({ ids });
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: "删除成功",
    });
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--;
    }
    deleteVisible.value = false;
    getTableData();
  }
};

const handleCurrentChange = (val) => {
  page.value = val;
  getTableData();
};
const handleSizeChange = (val) => {
  pageSize.value = val;
  getTableData();
};

const type = ref("");
const dialogFormVisible = ref(false);
// 删除行
const deleteRow = (row) => {
  deleteArticle(row);
};

const openDialog = () => {
  type.value = "create";
  dialogFormVisible.value = true;
};

const closeDialog = () => {
  dialogFormVisible.value = false;
  formData.value = {
    title: "",
    tag_id: undefined,
    desc: "",
    content: "",
    state: 1,
  };
};

const updateArticleFunc = async (row) => {
  const res = await findArticle({ ID: row.ID });
  type.value = "update";
  if (res.code === 0) {
    tags.value = [res.data.rearticle.tag];
    formData.value = res.data.rearticle;
    formData.value.tag = undefined;
    text.value = formData.value.content;
    dialogFormVisible.value = true;
  }
};

const enterDialog = async () => {
  let res;
  formData.value.content = text.value;
  switch (type.value) {
    case "create":
      res = await createArticle(formData.value);
      break;
    case "update":
      res = await updateArticle(formData.value);
      break;
    default:
      res = await createArticle(formData.value);
      break;
  }
  console.log(formData.value);
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: "创建/更改成功",
    });
    closeDialog();
    getTableData();
  }
};

const search = (val) => {
  if (!!val) {
    loading.value = true;
    getAppTabList({ name: val }).then((resp) => {
      loading.value = false;
      if (resp.code === 0) {
        tags.value = resp.data.list;
      }
    });
  }
};

const searchTag = useDebounceFn(search, 500, { maxWait: 5000 });
</script>

<style></style>
