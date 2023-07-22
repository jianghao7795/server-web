<template>
  <div>
    <div class="search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="用户名">
          <el-input v-model="searchInfo.username" />
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
            <el-button size="small" type="primary" link @click="deleteVisible = false">取消</el-button>
            <el-button size="small" type="primary" @click="onDelete">确定</el-button>
          </div>
          <template #reference>
            <el-button icon="delete" size="small" style="margin-left: 10px" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
          </template>
        </el-popover>
      </div>
      <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID" @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="用户名" prop="username" width="120" />
        <el-table-column align="left" label="昵称" prop="nickname" width="120" />
        <el-table-column align="left" label="真实姓名" prop="realname" width="120" />
        <el-table-column label="头像" prop="avatar" width="120" align="center">
          <template #default="scope">
            <CustomPic style="margin-top: 8px; text-align: left" :pic-src="scope.row.avatar" />
          </template>
        </el-table-column>
        <el-table-column align="left" label="简介" prop="sign" width="120" />
        <el-table-column align="left" label="主页封面" prop="cover" width="120">
          <template #default="scope">
            <CustomPic picType="file" :pic-src="scope.row.cover" />
          </template>
        </el-table-column>
        <el-table-column align="left" label="用户信息" prop="content" width="120" />
        <!-- <el-table-column align="left" label="密码" prop="password" width="120" /> -->
        <el-table-column align="left" label="行业" prop="industry" width="120">
          <template #default="scope">{{ viewOption[scope.row.industry] }}</template>
        </el-table-column>
        <el-table-column align="left" label="性别" prop="gender" width="120">
          <template #default="scope">{{ viewSex[scope.row.gender] }}</template>
        </el-table-column>
        <el-table-column align="left" label="操作">
          <template #default="scope">
            <el-button link type="primary" icon="edit" size="small" class="table-button" @click="updateMoblieUserFunc(scope.row)">编辑</el-button>
            <el-button link type="primary" icon="delete" size="small" @click="deleteRow(scope.row)">删除</el-button>
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
          <el-input v-model="formData.username" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="昵称:">
          <el-input v-model="formData.nickname" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="真实姓名:">
          <el-input v-model="formData.realname" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="头像:">
          <div style="display: inline-block" @click="openHeaderChange">
            <img v-if="!!formData.avatar" class="header-img-box" :src="formData.avatar ? path + `/${formData.avatar}` : formData.avatar" />
            <div v-else class="header-img-box">从媒体库选择</div>
          </div>
        </el-form-item>
        <el-form-item label="简介:">
          <el-input v-model="formData.sign" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="主页封面:">
          <el-input v-model="formData.cover" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户信息:">
          <el-input v-model="formData.content" clearable placeholder="请输入" />
        </el-form-item>
        <!-- <el-form-item label="密码:">
          <el-input v-model="formData.password" clearable placeholder="请输入" />
        </el-form-item> -->
        <el-form-item label="行业:">
          <el-select v-model="formData.industry">
            <el-option v-for="item in options" :key="item.value" :label="item.text" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="性别:">
          <el-radio-group v-model="formData.gender">
            <el-radio :label="0">未知</el-radio>
            <el-radio :label="1">男</el-radio>
            <el-radio :label="2">女</el-radio>
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
    <ChooseImg ref="chooseImg" @enter-img="enterImage" />
  </div>
</template>

<script>
export default {
  name: "MobileUser",
};
</script>

<script setup>
import { createMobileUser, deleteMobileUser, deleteMobileUserByIds, updateMobileUser, findMobileUser, getMobileUserList } from "@/api/mobileUser";
import ChooseImg from "@/components/chooseImg/index.vue";
import CustomPic from "@/components/customPic/index.vue";

// 全量引入格式化工具 请按需保留
// import { formatDate, formatBoolean } from "@/utils/format";
import { ElMessage, ElMessageBox } from "element-plus";
import { ref } from "vue";

const path = ref(import.meta.env.VITE_BASE_API);
const chooseImg = ref(null);

const options = [
  { text: "不展示", value: 0 },
  { text: "学生", value: 1 },
  { text: "自由职业", value: 2 },
  { text: "IT/互联网/通信", value: 3 },
  { text: "金融", value: 4 },
  { text: "健康/医疗", value: 5 },
  { text: "工业/制造业", value: 6 },
  { text: "零售", value: 7 },
  { text: "贸易", value: 8 },
  { text: "教育/科研", value: 9 },
  { text: "培训", value: 10 },
  { text: "房地产/建筑", value: 11 },
  { text: "文化/艺术", value: 12 },
  { text: "影视/娱乐", value: 13 },
  { text: "法律/会计/咨询", value: 14 },
  { text: "媒体/广告/公关", value: 15 },
  { text: "体育/健身", value: 16 },
  { text: "企事业单位", value: 17 },
];

const viewSex = {
  0: "未知",
  1: "男",
  2: "女",
};

const viewOption = {
  0: "不展示",
  1: "学生",
  2: "自由职业",
  3: "IT/互联网/通信",
  4: "金融",
  5: "健康/医疗",
  6: "工业/制造业",
  7: "零售",
  8: "贸易",
  9: "教育/科研",
  10: "培训",
  11: "房地产/建筑",
  12: "文化/艺术",
  13: "影视/娱乐",
  14: "法律/会计/咨询",
  15: "媒体/广告/公关",
  16: "体育/健身",
  17: "企事业单位",
};

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  username: "",
  nickname: "",
  realname: "",
  avatar: "",
  sign: "",
  cover: "",
  content: "",
  password: "",
  industry: 0,
  gender: 0,
});

// =========== 表格控制部分 ===========
const page = ref(1);
const total = ref(0);
const pageSize = ref(10);
const tableData = ref([]);
const searchInfo = ref({});

const enterImage = (url, _item) => {
  // console.log(url, item);
  formData.value.avatar = url;
};

// 选择图片
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
  if (searchInfo.value.industry === "") {
    searchInfo.value.industry = null;
  }
  if (searchInfo.value.gender === "") {
    searchInfo.value.gender = null;
  }
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
  const table = await getMobileUserList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value });
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
const setOptions = async () => {};

// 获取需要的字典 可能为空 按需保留
setOptions();

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
    deleteMobileUserFunc(row);
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
  const res = await deleteMobileUserByIds({ ids });
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
const updateMoblieUserFunc = async (row) => {
  const res = await findMobileUser({ ID: row.ID });
  type.value = "update";
  if (res.code === 0) {
    formData.value = res.data.remoblieUser;
    dialogFormVisible.value = true;
  }
};

// 删除行
const deleteMobileUserFunc = async (row) => {
  const res = await deleteMobileUser({ ID: row.ID });
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
    username: "",
    nickname: "",
    realname: "",
    avatar: "",
    sign: "",
    cover: "",
    content: "",
    password: "",
    industry: 0,
    gender: 0,
  };
};
// 弹窗确定
const enterDialog = async () => {
  let res;
  switch (type.value) {
    case "create":
      res = await createMobileUser(formData.value);
      break;
    case "update":
      res = await updateMobileUser(formData.value);
      break;
    default:
      res = await createMobileUser(formData.value);
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
