<!--
 * @Author: jianghao
 * @Date: 2022-10-17 14:50:22
 * @LastEditors: jianghao
 * @LastEditTime: 2022-10-18 16:23:56
-->

<template>
  <div>
    <el-card shadow="always" :body-style="{ padding: '20px' }">
      <template #header>
        <div class="card-header">
          <span>基本信息</span>
          <el-button :loading="loading" type="primary" @click="handleSubmit">{{ formData.ID ? "修改" : "提交" }}</el-button>
        </div>
      </template>
      <el-form :model="formData" label-position="right" label-width="160px" :rules="rules">
        <el-form-item label="网站标题:" prop="title">
          <el-input v-model="formData.title" clearable placeholder="请输入"></el-input>
        </el-form-item>
        <el-form-item label="网站简介:" prop="introduction">
          <el-input v-model.number="formData.introduction" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="网站头像:" prop="head_img">
          <el-upload :action="`${path}/base_message/upload_file`" ref="upload" :class="[fileList.length === 5 ? 'disUoloadSty' : '']" accept=".png,.jpg,.jpeg" list-type="picture-card" :disabled="fileList.length === 5" :limit="5" :headers="{ Authorization: `Bearer ${userStore.token}` }" :file-list="fileList" :on-success="uploadSuccess" :on-remove="handleRemove" :on-error="uploadError" :on-change="changeFile">
            <template #default>
              <div>
                <el-icon><Plus /></el-icon>
              </div>
            </template>

            <template #file="{ file }">
              <div>
                <img class="el-upload-list__item-thumbnail" :src="file.url" alt="" />
                <span class="el-upload-list__item-actions">
                  <span class="el-upload-list__item-preview" @click="handlePictureCardPreview(file)">
                    <el-icon><zoom-in /></el-icon>
                  </span>
                  <span class="el-upload-list__item-delete" @click="handleDownload(file)">
                    <el-icon><Download /></el-icon>
                  </span>
                  <span class="el-upload-list__item-delete" @click="handleRemove(file)">
                    <el-icon><Delete /></el-icon>
                  </span>
                </span>
              </div>
            </template>
          </el-upload>
        </el-form-item>
        <el-form-item label="版权信息:" prop="copyright">
          <el-input v-model="formData.copyright" placeholder="请输入" size="default" clearable></el-input>
        </el-form-item>
        <el-form-item label="链接:" prop="link">
          <el-input v-model.number="formData.link" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="备案信息:" prop="record_info">
          <el-input v-model.number="formData.record_info" clearable placeholder="请输入" />
        </el-form-item>
      </el-form>
      <!-- card body -->
    </el-card>
    <el-dialog v-model="dialogVisible" @closed="() => (dialogVisible = false)" title="查看图片">
      <img w-full :src="dialogImageUrl" alt="图片" style="width: 100%" />
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: "BaseMessage",
};
</script>

<script setup>
import { ref, onMounted } from "vue";
import { getBaseMessage, updateBaseMessage, createBaseMessage } from "@/api/baseMessage";
import { Delete, Download, Plus, ZoomIn } from "@element-plus/icons-vue";
import { useUserStore } from "@/pinia/modules/user";
import { ElMessage, ElNotification } from "element-plus";

const upload = ref(null);
const dialogImageUrl = ref("");
const dialogVisible = ref(false);
const fileList = ref([]);
const formData = ref({
  title: "",
  introduction: "",
  head_img: "",
  copyright: "",
  link: "",
  record_info: "",
});
const loading = ref(false);

const userStore = useUserStore();

const rules = ref({
  title: { required: true, message: "请输入", trigger: "blur" },
  introduction: { required: true, message: "请输入", trigger: "blur" },
  head_img: { required: true, message: "请输入", trigger: "blur" },
  copyright: { required: true, message: "请输入", trigger: "blur" },
});

onMounted(() => {
  getBaseMessage({ id: userStore.userInfo.ID }).then((resp) => {
    if (resp?.code === 0 && !resp?.data?.error) {
      formData.value = resp.data.baseMessage || {};
      if (resp.data.baseMessage.head_img) {
        fileList.value = resp.data.baseMessage.head_img.split(",").map((i) => ({ name: i.split("/").pop(), url: `/api/${i}`, status: "success" }));
      }
    }
  });
});

const uploadSuccess = (_response, _uploadFile, uploadFiles) => {
  formData.value.head_img = uploadFiles
    .map((i) => {
      if (i.url.includes("blob")) {
        return i.response.data.file.url.replace("/api/", "");
      }
      return i;
    })
    .join(",");
};

const uploadError = (error) => {
  const response = JSON.parse(error.message);
  ElNotification({
    title: "上传错误",
    message: response.msg,
    type: "error",
  });
};

const handlePictureCardPreview = (file) => {
  dialogImageUrl.value = file.url;
  dialogVisible.value = true;
};

const changeFile = (uploadFile, uploadFiles) => {
  // console.log(uploadFile);
  fileList.value = uploadFiles;
  // formData.value.head_img =
};

const handleDownload = (file) => {
  const a = document.createElement("a");
  a.href = file.url;
  a.download = file.name;
  a.click();
};

const handleRemove = (file) => {
  fileList.value = fileList.value.filter((i) => i.url !== file.url);
};

const handleSubmit = () => {
  loading.value = true;
  if (formData.value.ID) {
    updateBaseMessage({ ...formData.value, user_id: userStore.userInfo.ID }).then((resp) => {
      loading.value = false;
      if (resp?.code === 0) {
        ElMessage({
          message: "更新成功",
          type: "success",
        });
      }
    });
  } else {
    createBaseMessage({ ...formData.value, user_id: userStore.userInfo.ID }).then((resp) => {
      loading.value = false;
      if (resp?.code === 0) {
        ElMessage({
          message: "创建成功",
          type: "success",
        });
      }
    });
  }
};
</script>

<style scoped lang="scss">
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>
