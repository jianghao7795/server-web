<template>
  <div>
    <!-- <warning-bar
      title="获取字典且缓存方法已在前端utils/dictionary 已经封装完成 不必自己书写 使用方法查看文件内注释"
    /> -->
    <div class="search-box">
      <el-form :inline="true" :model="searchInfo" @keyup.enter.native="onSubmit">
        <el-form-item label="字典名（中）">
          <el-input v-model="searchInfo.name" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="字典名（英）">
          <el-input v-model="searchInfo.type" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-select v-model="searchInfo.status" clear placeholder="请选择">
            <el-option key="true" label="是" value="true" />
            <el-option key="false" label="否" value="false" />
          </el-select>
        </el-form-item>
        <el-form-item label="描述">
          <el-input v-model="searchInfo.desc" placeholder="搜索条件" />
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
        <el-table-column align="left" label="日期" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="字典名（中）" prop="name" width="160" />
        <el-table-column align="left" label="字典名（英）" prop="type" width="120" />
        <el-table-column align="left" label="状态" prop="status" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.status) }}</template>
        </el-table-column>
        <el-table-column align="left" label="描述" prop="desc" width="280" />
        <el-table-column align="left" label="操作">
          <template #default="scope">
            <el-button size="small" icon="document" link type="primary" @click="toDetile(scope.row)">详情</el-button>
            <el-button size="small" icon="edit" link type="primary" @click="updateSysDictionaryFunc(scope.row)">编辑</el-button>
            <!-- v-model:visible="scope.row.visible" -->
            <el-popconfirm
              placement="top"
              width="160"
              confirm-button-text="确定"
              cancel-button-text="取消"
              title="确定要删除吗?"
              @confirm="deleteSysDictionaryFunc(scope.row)"
            >
              <template #reference>
                <el-button link type="primary" icon="delete" size="small" style="margin-left: 10px">删除</el-button>
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="字典">
      <el-form ref="dialogForm" :model="formData" :rules="rules" size="default" label-width="110px">
        <el-form-item label="字典名（中）" prop="name">
          <el-input v-model="formData.name" placeholder="请输入字典名（中）" clearable :style="{ width: '100%' }" />
        </el-form-item>
        <el-form-item label="字典名（英）" prop="type">
          <el-input v-model="formData.type" placeholder="请输入字典名（英）" clearable :style="{ width: '100%' }" />
        </el-form-item>
        <el-form-item label="状态" prop="status" required>
          <el-switch v-model="formData.status" active-text="开启" inactive-text="停用" />
        </el-form-item>
        <el-form-item label="描述" prop="desc">
          <el-input v-model="formData.desc" placeholder="请输入描述" clearable :style="{ width: '100%' }" />
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
export default {
  name: "SysDictionary",
};
</script>

<script setup>
import { createSysDictionary, deleteSysDictionary, updateSysDictionary, findSysDictionary, getSysDictionaryList } from "@/api/sysDictionary"; //  此处请自行替换地址
// import warningBar from "@/components/warningBar/warningBar.vue";
import { ref } from "vue";
import { useRouter } from "vue-router";
import { ElMessage } from "element-plus";
import { formatBoolean, formatDate } from "@/utils/format";

const router = useRouter();

const formData = ref({
  name: null,
  type: null,
  status: true,
  desc: null,
});
const rules = ref({
  name: [
    {
      required: true,
      message: "请输入字典名（中）",
      trigger: "blur",
    },
  ],
  type: [
    {
      required: true,
      message: "请输入字典名（英）",
      trigger: "blur",
    },
  ],
  desc: [
    {
      required: true,
      message: "请输入描述",
      trigger: "blur",
    },
  ],
});

const page = ref(1);
const total = ref(0);
const pageSize = ref(10);
const tableData = ref([]);
const searchInfo = ref({});
const loading = ref(false);

const onReset = () => {
  searchInfo.value = {};
};

// 条件搜索前端看此方法
const onSubmit = () => {
  page.value = 1;
  pageSize.value = 10;
  if (searchInfo.value.status === "") {
    searchInfo.value.status = null;
  }
  getTableData();
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
  loading.value = true;
  const table = await getSysDictionaryList({
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

getTableData();

const toDetile = (row) => {
  router.push({
    name: "dictionaryDetail",
    params: {
      id: row.ID,
    },
  });
};

const dialogFormVisible = ref(false);
const type = ref("");
const updateSysDictionaryFunc = async (row) => {
  const res = await findSysDictionary({ ID: row.ID });
  type.value = "update";
  if (res.code === 0) {
    formData.value = res.data.resysDictionary;
    dialogFormVisible.value = true;
  }
};
const closeDialog = () => {
  dialogFormVisible.value = false;
  formData.value = {
    name: null,
    type: null,
    status: true,
    desc: null,
  };
};
const deleteSysDictionaryFunc = async (row) => {
  row.visible = false;
  const res = await deleteSysDictionary({ ID: row.ID });
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

const dialogForm = ref(null);
const enterDialog = async () => {
  dialogForm.value.validate(async (valid) => {
    if (!valid) return;
    let res;
    switch (type.value) {
      case "create":
        res = await createSysDictionary(formData.value);
        break;
      case "update":
        res = await updateSysDictionary(formData.value);
        break;
      default:
        res = await createSysDictionary(formData.value);
        break;
    }
    if (res.code === 0) {
      closeDialog();
      getTableData();
    }
  });
};
const openDialog = () => {
  type.value = "create";
  dialogFormVisible.value = true;
};
</script>

<style></style>
