<template>
  <div>
    <div class="search-box">
      <el-form ref="searForm" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter.native="onSubmit">
        <el-form-item label="名称:">
          <el-input v-model="searchInfo.title"></el-input>
        </el-form-item>
        <el-form-item label="标签:">
          <el-select v-model="searchInfo.tag_id" clearable filterable placeholder="请选择" remote reserve-keyword>
            <el-option v-for="item in tags" :key="item.ID" :label="item.name" :value="item.ID"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="首页是否显示:">
          <el-select v-model="searchInfo.is_important" clearable>
            <el-option label="显示" value="1"></el-option>
            <el-option label="隐藏" value="2"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button icon="search" size="small" type="primary" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" size="small" @click="reset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="table-box">
      <div class="btn-list">
        <el-button icon="plus" size="small" type="primary" @click="openDialog">新增</el-button>
        <el-popconfirm :width="250" placement="top" title="确定要删除吗?" @confirm="onDelete">
          <template #reference>
            <el-button :disabled="!multipleSelection.length" icon="delete" size="small" style="margin-left: 10px">删除</el-button>
          </template>
        </el-popconfirm>
        <el-popconfirm :width="250" placement="top" title="确定不显示首页?" @confirm="OnCancelView">
          <template #reference>
            <el-button :disabled="!multipleSelection.length" icon="hide" size="small" style="margin-left: 10px">取消首页显示</el-button>
          </template>
        </el-popconfirm>
      </div>
      <el-table ref="multipleTable" v-loading="loadingInit" :data="tableData" row-key="ID" style="width: 100%" tooltip-effect="dark" @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" />
        <el-table-column label="ID" prop="ID" width="55" />
        <el-table-column label="标题" prop="title"></el-table-column>
        <el-table-column label="标签" prop="tag" width="350">
          <template #default="{ row }">
            <div class="centerBg">
              <el-space>
                <el-tag v-for="item in row.tags" :key="item.ID" :type="colorIndex(item.ID)">{{ item.name }}</el-tag>
              </el-space>
            </div>
          </template>
        </el-table-column>
        <el-table-column label="首页显示" prop="is_important">
          <template #default="{ row }">
            <div class="centerBg">
              <el-space>
                <el-tag :type="row.is_important === 1 ? 'success' : ''">{{ row.is_important === 1 ? "显示" : "隐藏" }}</el-tag>
              </el-space>
            </div>
          </template>
        </el-table-column>
        <el-table-column label="简化内容" prop="desc"></el-table-column>
        <el-table-column label="阅读量" prop="reading_quantity">
          <template #default="{ row }">
            <span>{{ row.reading_quantity }}</span>
          </template>
        </el-table-column>
        <el-table-column label="作者" prop="user">
          <template #default="{ row }">
            <span>{{ row.user?.nickName }}</span>
          </template>
        </el-table-column>
        <el-table-column label="创建时间">
          <template #default="{ row }">
            {{ formatDate(row.CreatedAt) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作">
          <template #default="scope">
            <el-button class="table-button" icon="edit" link size="small" type="primary" @click="updateArticleFunc(scope.row)">编辑</el-button>
            <el-popconfirm placement="top" title="确认删除？" width="200" v-on:confirm="deleteRow(scope.row)">
              <template #reference><el-button icon="delete" link size="small" type="primary">删除</el-button></template>
            </el-popconfirm>
          </template>
        </el-table-column>
      </el-table>
      <div class="pagination">
        <el-pagination :current-page="page" :page-size="pageSize" :page-sizes="[10, 30, 50, 100]" :total="total" background layout="total, sizes, prev, pager, next, jumper" @current-change="handleCurrentChange" @size-change="handleSizeChange" />
      </div>
    </div>
    <el-dialog :before-close="closeDialog" :model-value="dialogFormVisible" :title="type === 'update' ? '更新文章' : '新建文章'" :width="1100" draggable>
      <el-form ref="ruleFormRef" :inline-message="true" :model="formData" :rules="rules" :scroll-to-error="true" label-position="right" label-suffix=":" label-width="80px" status-icon>
        <el-form-item label="标签" prop="tags">
          <el-select v-model="formData.tags" filterable multiple placeholder="请选择" style="width: 100%" @change="changeTagsFunc" @remove-tag="removeTag">
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
          <MdEditor v-model="formData.content" @on-upload-img="onUploadImg"></MdEditor>
        </el-form-item>
        <el-form-item label="状态" prop="state">
          <el-radio-group v-model="formData.state">
            <el-radio-button :label="1">显示</el-radio-button>
            <el-radio-button :label="0">隐藏</el-radio-button>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="首页显示" prop="is_important">
          <el-radio-group v-model="formData.is_important">
            <el-radio-button :label="1">显示</el-radio-button>
            <el-radio-button :label="2">隐藏</el-radio-button>
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
import { getArticleList, deleteArticle, findArticle, createArticle, updateArticle, uploadFile, deleteArticleByIds, putArticleByIds } from "@/api/article";
import { getTagList } from "@/api/tag";
import { ref, onBeforeMount, reactive } from "vue";
import { colorIndex } from "@/utils/util";
// import { useDebounceFn } from "@vueuse/core";
import { useUserStore } from "@/pinia/modules/user";
import { MdEditor } from "md-editor-v3";
import "md-editor-v3/lib/style.css";

const userStore = useUserStore();

const formData = ref({
  title: "",
  tags: [],
  desc: "",
  content: "",
  state: 1,
  is_important: undefined,
});

const ruleFormRef = ref(null); // form ref
const rules = reactive({
  tags: [{ required: true, message: "请查询并选择", trigger: "blur" }],
  title: [{ required: true, message: "请输入", trigger: "blur" }],
});

// search tag loading
// const loading = ref(false);
const loadingInit = ref(false);

// const changeText = (e) => {
//   text.value = e;
// };

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
  getTagList({ page: 1, pageSize: 999 }).then((resp) => {
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

// 多选取消首页显示
const OnCancelView = async () => {
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: "warning",
      message: "请选择",
    });
    return;
  }

  const ids = multipleSelection.value.map((i) => i.ID);
  const resp = await putArticleByIds({ ids });
  if (resp.code === 0) {
    ElMessage({
      type: "success",
      message: "首页显示取消成功",
    });
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
    // console.log(ruleFormRef);
    ruleFormRef.value.resetFields();
  }
  dialogFormVisible.value = false;
  formData.value = {
    title: "",
    desc: "",
    content: "",
    state: 1,
  };
  // text.value = "";
};

const updateArticleFunc = async (row) => {
  const res = await findArticle({ ID: row.ID });
  type.value = "update";
  if (res.code === 0) {
    changeTags.value = res.data.rearticle.tags;
    formData.value = res.data.rearticle;
    formData.value.tags = res.data.rearticle.tags.map((i) => i.ID);
    // text.value = formData.value.content;
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
      // formData.value.content = text.value;

      switch (type.value) {
        case "create":
          res = await createArticle({ ...formData.value, tags: changeTags.value, user_id: userStore.userInfo.ID });
          break;
        case "update":
          res = await updateArticle({ ...formData.value, tags: changeTags.value });
          break;
        default:
          res = await createArticle({ ...formData.value, tags: changeTags.value, user_id: userStore.userInfo.ID });
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
