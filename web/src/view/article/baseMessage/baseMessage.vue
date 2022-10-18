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
          <el-upload
            action="/api/base_message/upload_file"
            ref="upload"
            accept=".png,.jpg,.jpeg"
            list-type="picture-card"
            :headers="{ 'x-token': userStore.token }"
            v-model:file-list="fileList"
            :limit="1"
            :on-change="changeFile"
            :on-success="uploadSuccess"
            :disabled="fileList.length === 1"
          >
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
    <el-dialog v-model="dialogVisible" @closed="() => (dialogVisible = false)">
      <img w-full :src="dialogImageUrl" alt="图片" />
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { getBaseMessage, updateBaseMessage, createBaseMessage } from "@/api/baseMessage";
import { Delete, Download, Plus, ZoomIn } from "@element-plus/icons-vue";
import { useUserStore } from "@/pinia/modules/user";
import { ElMessage } from "element-plus";

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
  getBaseMessage().then((resp) => {
    if (resp?.code === 0 && resp?.data?.baseMessage?.ID) {
      formData.value = resp.data.baseMessage || {};
      fileList.value = [{ name: resp.data.baseMessage.head_img.split("/").pop(), url: `/api/${resp.data.baseMessage.head_img}` }];
    }
  });
});

const uploadSuccess = (response) => {
  // console.log(response);
  if (response?.data?.file?.url) {
    formData.value.head_img = response?.data?.file?.url;
  } else {
    formData.value.head_img = "";
  }
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
  fileList.value = [];
};

const handleSubmit = () => {
  loading.value = true;
  if (formData.value.ID) {
    updateBaseMessage(formData.value).then((resp) => {
      loading.value = false;
      if (resp?.code === 0) {
        ElMessage({
          message: "更新成功",
          type: "success",
        });
      }
    });
  } else {
    createBaseMessage(formData.value).then((resp) => {
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
.avatar-uploader .el-upload {
  border: 1px dashed var(--el-border-color);
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
  transition: var(--el-transition-duration-fast);
}

.avatar-uploader .el-upload:hover {
  border-color: var(--el-color-primary);
}

.el-icon.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  text-align: center;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.text {
  font-size: 14px;
}

.item {
  margin-bottom: 18px;
}
</style>
