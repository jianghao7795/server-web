<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="文章：">
          <el-select
            v-model="searchInfo.articleId"
            :multiple="false"
            filterable
            remote
            clearable
            reserve-keyword
            placeholder="请输入"
            :remote-method="searchArticle"
            :loading="loading"
          >
            <el-option v-for="item in options" :key="item.ID" :label="item.title" :value="item.ID" />
          </el-select>
        </el-form-item>
        <el-form-item label="内容：">
          <el-input v-model="searchInfo.content" placeholder="请输入"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="small" type="primary" icon="plus" @click="openDialog">新增</el-button>
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px">
            <el-button size="small" link @click="deleteVisible = false">取消</el-button>
            <el-button size="small" type="primary" @click="onDelete">确定</el-button>
          </div>
          <template #reference>
            <el-button icon="delete" size="small" style="margin-left: 10px" :disabled="!multipleSelection.length" @click="deleteVisible = true">
              删除
            </el-button>
          </template>
        </el-popover>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column header-align="center" align="center" prop="ID" label="ID" width="55"></el-table-column>
        <el-table-column align="left" label="文章" prop="articleId" width="120">
          <template #default="scope">{{ scope.row?.article?.title }}</template>
        </el-table-column>
        <el-table-column align="left" label="上级" prop="parentId" width="120" />
        <el-table-column align="left" label="内容" prop="content" width="120" />
        <el-table-column align="left" label="用户" prop="userId" width="120">
          <template #default="scope">
            {{ scope.row?.SysUser?.nickName }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="赞数" prop="praise" width="120">
          <template #header="scope">
            <span class="iconfont icon-aixin"></span>
            /
            <span class="iconfont icon-aixin1"></span>
          </template>
          <template #default="scope">
            <div class="like-num-wrapper">
              <transition :name="scope.row.praise ? 'plus' : 'minus'">
                <div class="like-num" :style="{ color: scope.row.praise ? 'red' : '#333' }" :key="scope.row.praise">
                  <span class="iconfont icon-aixin"></span>
                  {{ scope.row.praise }}
                </div>
              </transition>
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="评论时间" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="操作">
          <template #default="scope">
            <el-button link icon="edit" type="primary" size="small" class="table-button" @click="updateCommentFunc(scope.row)">编辑</el-button>
            <el-button link icon="delete" type="primary" size="small" @click="deleteRow(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
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
    <el-dialog :model-value="dialogFormVisible" :before-close="closeDialog" :title="!formData.articleId ? '创建评论' : '更新评论'">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="文章:">
          <el-select
            v-model="formData.articleId"
            :multiple="false"
            filterable
            remote
            clearable
            reserve-keyword
            placeholder="请输入"
            :remote-method="searchArticle"
          >
            <el-option v-for="item in options" :key="item.ID" :label="item.title" :value="item.ID"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="上级:">
          <el-input v-model.number="formData.parentId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="内容:">
          <el-input v-model="formData.content" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户:">
          <el-select v-model="formData.userId" placeholder="请输入">
            <el-option v-for="item in userInfo" :key="item.ID" :label="item.nickName" :value="item.ID"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="赞数:">
          <el-input v-model.number="formData.praise" clearable placeholder="请输入" />
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
  name: "Comment",
};
</script>

<script setup>
import { createComment, deleteComment, deleteCommentByIds, updateComment, findComment, getCommentList } from "@/api/comment";

// 全量引入格式化工具 请按需保留
import { formatDate } from "@/utils/format";
import { ElMessage, ElMessageBox } from "element-plus";
import { ref } from "vue";
import { useDebounceFn } from "@vueuse/core";
import { getArticleList } from "@/api/article";
// import { ThumbsDownFilled, ThumbsUpFilled } from "@vicons/carbon";

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  articleId: undefined,
  parentId: undefined,
  content: "",
  userId: undefined,
  praise: 0,
});

// =========== 表格控制部分 ===========
const page = ref(1);
const total = ref(0);
const pageSize = ref(10);
const tableData = ref([]);
const searchInfo = ref({});
const loading = ref(false);
const options = ref([]);
const userInfo = ref([]);

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
  // console.log(searchInfo.value);
  const table = await getCommentList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value });
  if (table.code === 0) {
    tableData.value = table.data.list;
    total.value = table.data.total;
    page.value = table.data.page;
    pageSize.value = table.data.pageSize;
  }
};

getTableData();

// 搜索文章
const search = (query) => {
  if (query) {
    getArticleList({ title: query }).then((resp) => {
      if (resp.code === 0) {
        options.value = resp.data.list;
      }
    });
  }
};

const searchArticle = useDebounceFn(search, 800);

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
    deleteCommentFunc(row);
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
  const res = await deleteCommentByIds({ ids });
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
const updateCommentFunc = async (row) => {
  const res = await findComment({ ID: row.ID });
  type.value = "update";
  if (res.code === 0) {
    formData.value = res.data.recomment;
    // formData.value.SysUser = undefined;
    // formData.value.article = undefined;
    dialogFormVisible.value = true;
    userInfo.value = [res.data.recomment.SysUser];
    options.value = [res.data.recomment.article];
  }
};

// 删除行
const deleteCommentFunc = async (row) => {
  const res = await deleteComment({ ID: row.ID });
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
  console.log(dialogFormVisible.value);
  setTimeout(() => {
    formData.value = {
      articleId: undefined,
      parentId: undefined,
      content: "",
      userId: undefined,
      praise: 0,
    };
  }, 300);
};
// 弹窗确定
const enterDialog = async () => {
  let res;
  // console.log(formData.value);
  switch (type.value) {
    case "create":
      res = await createComment(formData.value);
      break;
    case "update":
      res = await updateComment({ ...formData.value, SysUser: undefined, article: undefined });
      break;
    default:
      res = await createComment(formData.value);
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

<style scoped lang="scss">
.like-num-wrapper {
  position: relative;
  font-size: 13px;
  height: 17px;
  overflow-y: hidden;

  .like-num {
    top: 0;
    left: 0;
    position: relative;
    line-height: 17px;
  }
}
.plus-enter-active,
.plus-leave-active {
  transition: all 0.3s ease-in;
}

.plus-enter,
.plus-leave {
  transform: translateY(0);
}

.plus-enter-to,
.plus-leave-to {
  transform: translateY(-17px);
}

// 点赞数字-1动画
.minus-enter-active,
.minus-leave-active {
  transition: all 0.3s ease-in;
}

.minus-enter {
  transform: translateY(-34px);
}

.minus-enter-to {
  transform: translateY(-17px);
}

.minus-leave {
  transform: translateY(0);
}

.minus-leave-to {
  transform: translateY(17px);
}

/** 动画进行时的class **/
.zoom-enter-active,
.zoom-leave-active {
  transition: all 0.15s cubic-bezier(0.42, 0, 0.34, 1.55);
}

/** 设置进场开始的状态和离场结束的状态，都是缩放到0 **/
.zoom-enter,
.zoom-leave-to {
  transform: scale(0);
}

/** 设置进场结束的状态和离场开始的状态, 都是缩放到1 **/
.zoom-enter-to,
.zoom-leave {
  transform: scale(1);
}
.minus-enter {
  transform: translateY(-34px);
}
</style>
