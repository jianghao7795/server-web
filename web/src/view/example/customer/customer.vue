<template>
  <div>
    <!-- <warning-bar
      title="在资源权限中将此角色的资源权限清空 或者不包含创建者的角色 即可屏蔽此客户资源的显示"
    /> -->
    <div class="search-box">
      <el-form :inline="true" ref="searchForm" :model="searchInfo" @keyup.enter.native="onSubmit">
        <el-form-item label="姓名">
          <el-input placeholder="请输入" v-model="searchInfo.customerName"></el-input>
        </el-form-item>
        <el-form-item label="电话">
          <el-input placeholder="请输入" v-model="searchInfo.customerPhoneData"></el-input>
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
      </div>
      <el-table ref="multipleTable" :data="tableData" style="width: 100%" tooltip-effect="dark" row-key="ID" v-loading="loading">
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="ID" prop="ID" width="80"></el-table-column>
        <el-table-column align="left" label="接入日期" width="180">
          <template #default="scope">
            <span>{{ formatDate(scope.row.CreatedAt) }}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="姓名" prop="customerName" width="120" />
        <el-table-column align="left" label="电话" prop="customerPhoneData" width="120" />
        <el-table-column align="left" label="创建人" prop="sysUserId" width="120">
          <template #default="{ row }">
            <el-tag class="ml-2" :type="row.sysUser.name === 'admin' ? 'success' : 'info'">{{ row.sysUser.userName }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作" min-width="160">
          <template #default="scope">
            <el-button size="small" link type="primary" icon="edit" @click="updateCustomer(scope.row)">编辑</el-button>
            <el-popconfirm
              confirm-button-text="确定"
              cancel-button-text="取消"
              title="确定要删除吗？"
              @confirm="deleteCustomer(scope.row)"
              placement="top"
            >
              <!-- @click="deleteCustomer(scope.row)" -->
              <template #reference>
                <el-button link type="danger" icon="delete" size="small">删除</el-button>
              </template>
            </el-popconfirm>
          </template>
        </el-table-column>
      </el-table>
      <div class="pagination">
        <el-pagination
          background
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="客户">
      <el-form :inline="true" :model="form" label-width="80px">
        <el-form-item label="客户名">
          <el-input v-model="form.customerName" autocomplete="off" />
        </el-form-item>
        <el-form-item label="客户电话">
          <el-input v-model="form.customerPhoneData" autocomplete="off" />
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

<script setup>
import { createExaCustomer, updateExaCustomer, deleteExaCustomer, getExaCustomer, getExaCustomerList } from "@/api/customer";
// import warningBar from "@/components/warningBar/warningBar.vue";
import { ref, onMounted } from "vue";
import { ElMessage } from "element-plus";
import { formatDate } from "@/utils/format";

const searchInfo = ref({});

const form = ref({
  customerName: "",
  customerPhoneData: "",
});

const page = ref(1);
const total = ref(0);
const pageSize = ref(10);
const tableData = ref([]);
const loading = ref(false);

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val;
  getTableData();
};

const handleCurrentChange = (val) => {
  page.value = val;
  getTableData();
};

// 查询 分页
const getTableData = async () => {
  loading.value = true;
  const table = await getExaCustomerList({
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo.value,
  });
  loading.value = false;
  if (table.code === 0) {
    tableData.value = table.data.list;
    total.value = table.data.total;
    page.value = table.data.page;
    pageSize.value = table.data.pageSize;
  }
};
// 搜索 查询
const onSubmit = async () => {
  page.value = 1;
  pageSize.value = 10;
  getTableData();
};
// 重置
const onReset = () => {
  searchInfo.value = {};
  getTableData();
};
onMounted(() => {
  getTableData();
});

const dialogFormVisible = ref(false);
const type = ref("");
const updateCustomer = async (row) => {
  const res = await getExaCustomer({ ID: row.ID });
  type.value = "update";
  if (res.code === 0) {
    form.value = res.data.customer;
    dialogFormVisible.value = true;
  }
};
const closeDialog = () => {
  dialogFormVisible.value = false;
  form.value = {
    customerName: "",
    customerPhoneData: "",
  };
};
const deleteCustomer = async (row) => {
  row.visible = false;
  const res = await deleteExaCustomer({ ID: row.ID });
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
const enterDialog = async () => {
  let res;
  switch (type.value) {
    case "create":
      res = await createExaCustomer(form.value);
      break;
    case "update":
      res = await updateExaCustomer(form.value);
      break;
    default:
      res = await createExaCustomer(form.value);
      break;
  }

  if (res.code === 0) {
    closeDialog();
    getTableData();
  }
};
const openDialog = () => {
  type.value = "create";
  dialogFormVisible.value = true;
};
</script>

<script>
export default {
  name: "Customer",
};
</script>

<style scoped></style>
