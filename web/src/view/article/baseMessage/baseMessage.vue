<!--
 * @Author: jianghao
 * @Date: 2022-10-17 14:50:22
 * @LastEditors: jianghao
 * @LastEditTime: 2022-10-17 15:23:04
-->

<template>
  <el-card shadow="always" :body-style="{ padding: '20px' }">
    <template #header>
      <div>
        <span>基本信息</span>
      </div>
    </template>
    <el-form :model="formData" label-position="right" label-width="160px">
      <el-form-item label="网站标题:">
        <el-input v-model="formData.title" clearable placeholder="请输入"></el-input>
      </el-form-item>
      <el-form-item label="网站简介:">
        <el-input v-model.number="formData.introduction" clearable placeholder="请输入" />
      </el-form-item>
      <el-form-item label="网站头像:">
        <el-input v-model="formData.head_img" clearable placeholder="请输入" />
      </el-form-item>
      <el-form-item label="版权信息:">
        <el-input v-model="formData.copyright" placeholder="请输入" size="default" clearable></el-input>
      </el-form-item>
      <el-form-item label="链接:">
        <el-input v-model.number="formData.link" clearable placeholder="请输入" />
      </el-form-item>
      <el-form-item label="备案信息:">
        <el-input v-model.number="formData.record_info" clearable placeholder="请输入" />
      </el-form-item>
    </el-form>
    <!-- card body -->
  </el-card>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { getBaseMessage } from "@/api/baseMessage";
const formData = ref({
  title: "",
  introduction: "",
  head_img: "",
  copyright: "",
  link: "",
  record_info: "",
});

onMounted(() => {
  getBaseMessage().then((resp) => {
    if (resp?.code === 0 && resp?.data?.baseMessage?.ID) {
      formData.value = resp.data.baseMessage || {};
    }
  });
});
</script>

<style scoped lang="scss"></style>
