<template>
  <div>
    <div class="search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter.native="onSubmit">
        <el-form-item label="用户名:">
          <el-input v-model="searchInfo.name"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="table-box">
      <div class="btn-list">
        <el-button size="small" type="primary" icon="plus" @click="openDialog">新增</el-button>
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px">
            <el-button size="small" link type="primary" @click="deleteVisible = false">取消</el-button>
            <el-button size="small" type="primary" @click="onDelete">确定</el-button>
          </div>
          <template #reference>
            <el-button icon="delete" size="small" style="margin-left: 10px" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
          </template>
        </el-popover>
      </div>
      <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID" @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="ID" prop="ID" width="80" />
        <el-table-column align="left" label="用户名" prop="name" width="120" />
        <el-table-column align="left" label="背景图" prop="headImg" width="120">
          <template #default="scope">
            <div>
              <!-- <viewer :images="[scope.row.url]"><CustomPic pic-type="file" :pic-src="scope.row.url" /></viewer> -->
              <CustomPic pic-type="file" :pic-src="scope.row.headImg" />
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="简介" prop="introduction" width="120" />
        <el-table-column align="left" label="用户信息" prop="content" width="300" />
        <el-table-column align="left" label="创建时间" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="操作">
          <template #default="scope">
            <el-button link type="primary" icon="edit" size="small" class="table-button" @click="() => updateUserFunc(scope.row)">编辑</el-button>
            <el-button link type="primary" icon="delete" size="small" @click="() => deleteRow(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="pagination">
        <el-pagination layout="total, sizes, prev, pager, next, jumper" :current-page="page" :page-size="pageSize" :page-sizes="[10, 30, 50, 100]" :total="total" @current-change="handleCurrentChange" @size-change="handleSizeChange" />
      </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="用户">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="用户名:">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="密码:" v-show="!(type === 'update')">
          <el-input type="password" show-password v-model="formData.password" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="背景图:">
          <img v-if="formData.headImg" class="header-img-box" :src="formData.headImg && formData.headImg.slice(0, 4) !== 'http' ? path + `/${formData.headImg}` : formData.headImg" />
        </el-form-item>
        <el-form-item label="简介:">
          <el-input v-model="formData.introduction" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户信息:">
          <el-input v-model="formData.content" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="头像:" label-width="80px">
          <div style="display: inline-block" @click="openHeaderChange">
            <img v-if="formData.header" class="header-img-box" :src="formData.header && formData.header.slice(0, 4) !== 'http' ? path + `/${formData.header}` : formData.header" />
            <div v-else class="header-img-box">从媒体库选择</div>
          </div>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
    <ChooseImg ref="chooseImg" :target="formData" :target-key="`header`" />
  </div>
</template>

<script>
export default {
  name: "User",
};
</script>

<script setup>
import { createUser, deleteUser, deleteUserByIds, updateUser, findUser, getUserList } from "@/api/users";
import CustomPic from "@/components/customPic/index.vue";
// 全量引入格式化工具 请按需保留
import { formatDate } from "@/utils/format";
import { ElMessage, ElMessageBox } from "element-plus";
import { ref } from "vue";
import ChooseImg from "@/components/chooseImg/index.vue";

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  name: "",
  header: "",
  introduction: "",
  content: "",
  password: "",
  headImg: "",
});

const path = ref(import.meta.env.VITE_BASE_API);
const chooseImg = ref(null);

// =========== 表格控制部分 ===========
const page = ref(1);
const total = ref(0);
const pageSize = ref(10);
const tableData = ref([]);
const searchInfo = ref({});

const openHeaderChange = () => {
  chooseImg.value.open();
};

// 重置
const onReset = () => {
  searchInfo.value = {};
};

// 搜索
const onSubmit = () => {
  page.value = 1;
  pageSize.value = 10;
  getTableData();
};

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val;
  getTableData();
};

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val;
  getTableData();
};

// 查询
const getTableData = async () => {
  const table = await getUserList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value });
  if (table.code === 0) {
    tableData.value = table.data.list;
    total.value = table.data.total;
    page.value = table.data.page;
    pageSize.value = table.data.pageSize;
  }
};

getTableData();

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
// const setOptions = async () => {};

// 获取需要的字典 可能为空 按需保留
// setOptions();

// 多选数据
const multipleSelection = ref([]);
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val;
};

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm("确定要删除吗?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(() => {
    deleteUserFunc(row);
  });
};

// 批量删除控制标记
const deleteVisible = ref(false);

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
  const res = await deleteUserByIds({ ids });
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

// 行为控制标记（弹窗内部需要增还是改）
const type = ref("");

// 更新行
const updateUserFunc = async (row) => {
  const res = await findUser({ ID: row.ID });
  type.value = "update";
  if (res.code === 0) {
    formData.value = res.data.reuser;
    dialogFormVisible.value = true;
    console.log(formData.value);
  }
};

// 删除行
const deleteUserFunc = async (row) => {
  const res = await deleteUser({ ID: row.ID });
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: "删除成功",
    });
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--;
    }
    getTableData();
  }
};

// 弹窗控制标记
const dialogFormVisible = ref(false);

// 打开弹窗
const openDialog = () => {
  type.value = "create";
  dialogFormVisible.value = true;
};

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false;
  formData.value = {
    name: "",
    headImg: "",
    introduction: "",
    content: "",
  };
};
// 弹窗确定
const enterDialog = async () => {
  let res;
  switch (type.value) {
    case "create":
      res = await createUser(formData.value);
      break;
    case "update":
      res = await updateUser(formData.value);
      break;
    default:
      res = await createUser(formData.value);
      break;
  }
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: "创建/更改成功",
    });
    closeDialog();
    getTableData();
  }
};
</script>

<style scoped>
.header-img-box {
  width: 200px;
  height: 200px;
  border: 1px dashed #ccc;
  border-radius: 20px;
  text-align: center;
  line-height: 200px;
  cursor: pointer;
}
</style>
