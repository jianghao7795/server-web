<template>
  <n-drawer
    v-model:show="props.active"
    :width="502"
    placement="right"
    v-bind:on-update:show="changeStatus"
  >
    <n-drawer-content title="个人信息">
      <template #footer v-if="isUpdate">
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
        <n-form-item label="Input" path="inputValue">
          <n-input v-model:value="user.name" placeholder="Input" />
        </n-form-item>
        <n-form-item label="Textarea" path="textareaValue">
          <n-input v-model:value="user.header" placeholder="Textarea" />
        </n-form-item>
      </n-form>
    </n-drawer-content>
  </n-drawer>
</template>
<script lang="ts" name="Person" setup>
import { ref } from "vue";
import type { FormInst } from "naive-ui";
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

const rules = {};
</script>

<style scoped lang="less"></style>
