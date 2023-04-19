<template>
  <n-drawer
    v-model:show="props.active"
    :width="502"
    placement="right"
    v-bind:on-update:show="changeStatus"
  >
    <n-drawer-content title="个人信息">
      <template #footer>
        <div>
          <NButton>提交</NButton>
        </div>
      </template>

      <n-form
        ref="formRef"
        :model="user"
        :rules="rules"
        label-placement="left"
        label-width="auto"
        require-mark-placement="right-hanging"
      >
        <n-form-item label="头像" path="header">
          <n-upload
            list-type="image-card"
            :file-list="fileList"
            :max="1"
          ></n-upload>
        </n-form-item>
        <n-form-item label="账号" path="name">
          <n-input v-model:value="userInfo.name" placeholder="Textarea" />
        </n-form-item>
        <n-form-item label="简介" path="introduction">
          <n-input
            v-model:value="userInfo.introduction"
            placeholder="Textarea"
          />
        </n-form-item>
        <n-form-item label="介绍" path="content">
          <n-input v-model:value="userInfo.content" placeholder="Textarea" />
        </n-form-item>
      </n-form>
    </n-drawer-content>
  </n-drawer>
</template>
<script lang="ts" name="Person" setup>
import { ref } from "vue";
import type { FormInst, UploadFileInfo } from "naive-ui";
import { useUserStore } from "@/stores/user";

const props = defineProps<{ active: boolean }>();
const emits = defineEmits<{ (e: "changeStatus", status: boolean): void }>();
const changeStatus = (show: boolean) => {
  emits("changeStatus", show);
};

const isUpdate = ref<boolean>(false);
const loading = ref<boolean>(false);
const formRef = ref<FormInst | null>(null);
const userStore = useUserStore();
const user = userStore.currentUser.user;

const userInfo = ref<User.UserInfo>(user);
const fileList = ref<UploadFileInfo[]>([
  {
    id: "a",
    name: "头像.png",
    status: "finished",
    url: "/api/" + user.header,
  },
]);
console.log(userInfo.value);
const rules = {};
</script>

<style scoped lang="less"></style>
