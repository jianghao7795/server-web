<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searForm" :model="searchInfo" class="demo-form-inline" :inline="true">
        <el-form-item label="名称:">
          <el-input v-model="searchInfo.title"></el-input>
        </el-form-item>
        <el-form-item label="标签:">
          <el-select v-model="searchInfo.tag_id" placeholder="请选择" filterable remote reserve-keyword clearable>
            <el-option v-for="item in tags" :key="item.ID" :label="item.name" :value="item.ID"></el-option>
          </el-select>
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
        <el-popconfirm title="确定要删除吗?" @confirm="onDelete" placement="top" width="200">
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
        <el-table-column label="ID" prop="ID" width="55" />
        <el-table-column label="标题" prop="title"></el-table-column>
        <el-table-column label="标签" prop="tag">
          <template #default="{ row }">
            <div class="centerBg">
              <el-space>
                <el-tag v-for="item in row.tags" :key="item.id">{{ item.name }}</el-tag>
              </el-space>
            </div>
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
    <el-dialog
      :model-value="dialogFormVisible"
      :before-close="closeDialog"
      :title="type === 'update' ? '更新文章' : '新建文章'"
      draggable
      :width="1100"
    >
      <el-form
        :model="formData"
        ref="ruleFormRef"
        label-position="right"
        label-width="80px"
        :inline-message="true"
        :scroll-to-error="true"
        label-suffix=":"
        :rules="rules"
        status-icon
      >
        <el-form-item label="标签" prop="tags">
          <el-select
            v-model="formData.tags"
            placeholder="请选择"
            filterable
            multiple
            style="width: 100%"
            @change="changeTagsFunc"
            @remove-tag="removeTag"
          >
            <el-option v-for="item in tags" :key="item.ID" :label="item.name" :value="item.ID"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="标题" prop="title">
          <el-input v-model="formData.title" placeholder="请输入"></el-input>
        </el-form-item>
        <el-form-item label="简化内容" prop="desc">
          <el-input v-model="formData.desc" placeholder="请输入"></el-input>
        </el-form-item>
        <el-form-item label="内容" prop="content">
          <MdEditor style="width: 1000px" v-model="text" @on-upload-img="onUploadImg"></MdEditor>
        </el-form-item>
        <el-form-item label="状态:" prop="state">
          <el-radio-group v-model="formData.state">
            <el-radio-button :label="1">显示</el-radio-button>
            <el-radio-button :label="0">隐藏</el-radio-button>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog(ruleFormRef)">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog(ruleFormRef)">确 定</el-button>
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
import { getArticleList, deleteArticle, findArticle, createArticle, updateArticle, uploadFile, deleteArticleByIds } from "@/api/article";
import { getAppTabList } from "@/api/appTab";
import { ref, onBeforeMount, reactive } from "vue";
// import { useDebounceFn } from "@vueuse/core";
import MdEditor from "md-editor-v3";
import "md-editor-v3/lib/style.css";

const formData = ref({
  title: "",
  tags: [],
  desc: "",
  content: "",
  state: 1,
});

const ruleFormRef = ref(null); // form ref
const rules = reactive({
  tags: [{ required: true, message: "请查询并选择", trigger: "blur" }],
  title: [{ required: true, message: "请输入", trigger: "blur" }],
});

// search tag loading
// const loading = ref(false);
const loadingInit = ref(false);
const text = ref("");

const tags = ref([]);
const changeTags = ref([]);

const page = ref(1);
const total = ref(0);
const pageSize = ref(10);
const tableData = ref([]);
const searchInfo = ref({});
const path = ref(import.meta.env.VITE_BASE_API);

const changeTagsFunc = (row) => {
  const selectTag = tags.value.filter((i) => row.includes(i.ID));
  changeTags.value = selectTag;
};

const removeTag = (val) => {
  const selectTag = changeTags.value.filter((i) => val !== i.ID);
  changeTags.value = selectTag;
};

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
  getAppTabList({ page: 1, pageSize: 999 }).then((resp) => {
    if (resp.code === 0) {
      tags.value = resp.data.list;
    }
  });
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
    multipleSelection.value.forEach((item) => {
      ids.push(item.ID);
    });
  const res = await deleteArticleByIds({ ids });
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: "删除成功",
    });
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--;
    }
    // deleteVisible.value = false;
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
  if (type.value === "create") {
    ruleFormRef.resetFields();
  }
  dialogFormVisible.value = false;
  formData.value = {
    title: "",
    desc: "",
    content: "",
    state: 1,
  };
  text.value = "";
};

const updateArticleFunc = async (row) => {
  const res = await findArticle({ ID: row.ID });
  type.value = "update";
  if (res.code === 0) {
    changeTags.value = res.data.rearticle.tags;
    formData.value = res.data.rearticle;
    formData.value.tags = res.data.rearticle.tags.map((i) => i.ID);
    text.value = formData.value.content;
    dialogFormVisible.value = true;
  }
};

const enterDialog = async (formRules) => {
  if (!formRules) return;
  await formRules.validate(async (valid, fields) => {
    // console.log(valid, fields);
    if (valid) {
      // console.log("submit!");
      let res;
      formData.value.content = text.value;

      switch (type.value) {
        case "create":
          res = await createArticle({ ...formData.value, tags: changeTags.value });
          break;
        case "update":
          res = await updateArticle({ ...formData.value, tags: changeTags.value });
          break;
        default:
          res = await createArticle({ ...formData.value, tags: changeTags.value });
          break;
      }
      // console.log(formData.value);
      if (res.code === 0) {
        ElMessage({
          type: "success",
          message: "创建/更改成功",
        });
        closeDialog();
        getTableData();
      }
    } else {
      console.log("error submit!", fields);
    }
  });
  // return;
};

// const search = (val) => {
//   if (!!val) {
//     loading.value = true;

//   }
// };

// const searchTag = useDebounceFn(search, 500, { maxWait: 5000 });
</script>

<style lang="scss" scoped>
.centerBg {
  .el-space {
    vertical-align: middle;
  }
}
</style>
