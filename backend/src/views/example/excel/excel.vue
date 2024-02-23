<template>
  <div class="upload">
    <div class="table-box">
      <div class="btn-list">
        <el-upload class="excel-btn" :show-file-list="false" accept=".xls,.xlsx,.csv" v-bind:http-request="uploadFile" :loading="loadingUpload">
          <el-button size="small" type="primary" icon="upload">导入</el-button>
        </el-upload>
        <el-button class="excel-btn" size="small" type="primary" icon="download" @click="handleExcelExport('ExcelExport.xlsx')">导出</el-button>
        <!-- <el-button class="excel-btn" size="small" type="success" icon="download" @click="downloadExcelTemplate()">下载模板</el-button> -->
      </div>
      <el-table :data="tableData" row-key="ID">
        <el-table-column align="left" label="ID" width="50" prop="ID" />
        <el-table-column align="left" show-overflow-tooltip label="文件名" min-width="120" prop="filename" />
        <el-table-column align="left" show-overflow-tooltip label="文件名MD5" min-width="160" prop="filename_md5" />
        <el-table-column align="left" label="路径" min-width="100" show-overflow-tooltip prop="file_path"></el-table-column>
        <el-table-column align="left" label="大小（KB）" min-width="90" prop="file_size">
          <template #default="scope">
            <span>{{ fileSizeChange(scope.row.file_size, "KB") }}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="文件类型" min-width="70" prop="file_type" />
        <el-table-column align="left" label="上传时间" min-width="160" prop="CreatedAt">
          <template #default="scope">
            <span>{{ formatDate(scope.row.CreatedAt) }}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作" min-width="70">
          <template #default="{ row }">
            <el-space :size="2" spacer="|">
              <el-button link type="primary" @click="fileDownload(row.file_path)">下载</el-button>
              <el-button link type="primary" @click="deleteFileDownload(row)">删除</el-button>
            </el-space>
          </template>
        </el-table-column>
      </el-table>
      <div class="pagination">
        <el-pagination :current-page="page" :page-size="pageSize" :page-sizes="[10, 30, 50, 100]" :total="total" background layout="total, sizes, prev, pager, next, jumper" @current-change="handleCurrentChange" @size-change="handleSizeChange" />
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Excel",
};
</script>

<script setup>
import { ref, onMounted } from "vue";
import { importExcel, getFileList, deleteFile, exportExcel } from "@/api/excel";
import { fileSizeChange, formatDate } from "@/utils/format";
import { useUserStore } from "@/pinia/modules/user";

const loadingUpload = ref(false);
const tableData = ref([]);
const page = ref(1);
const pageSize = ref(10);
const total = ref(0);

// console.log(fileSizeChange(607, "KB"));

const uploadFile = (file) => {
  const formData = new FormData();
  formData.append("file", file.file);
  loadingUpload.value = true;
  importExcel(formData).then((res) => {
    loadingUpload.value = false;
    if (res.code === 0) {
      ElMessage({
        showClose: true,
        message: "导入成功",
        type: "success",
      });
      getTableData();
    }
  });
};
// 查询
const getTableData = async () => {
  const table = await getFileList({ page: page.value, pageSize: pageSize.value });
  if (table.code === 0) {
    tableData.value = table.data.list;
    total.value = table.data.total;
    page.value = table.data.page;
    pageSize.value = table.data.pageSize;
  }
};
onMounted(() => {
  getTableData();
});

const handleCurrentChange = (val) => {
  page.value = val;
  getTableData();
};
const handleSizeChange = (val) => {
  pageSize.value = val;
  getTableData();
};
const userStore = useUserStore();

const handleExcelExport = (fileName) => {
  if (!!fileName) {
    fileName = "ExcelExport.xlsx";
  }
  exportExcel(tableData.value, fileName);
};
const loadExcel = async () => {
  await getTableData();
};
const fileDownload = (path) => {
  window.open("/backend/" + path);
};

const deleteFileDownload = async (row) => {
  const res = await deleteFile({ id: row.ID });
  if (res.code === 0) {
    ElMessage({
      showClose: true,
      message: "删除成功",
      type: "success",
    });
    getTableData();
  }
};
</script>

<style lang="scss" scoped>
.btn-list {
  display: flex;
  margin-bottom: 12px;
  justify-content: flex-start;
}
.excel-btn + .excel-btn {
  margin-left: 10px;
}
</style>
