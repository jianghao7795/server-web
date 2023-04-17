<template>
  <n-drawer v-model:show="props.active" :width="502" placement="right">
    <n-drawer-content title="斯通纳">
      <n-form
        ref="formRef"
        :model="userPasswor"
        :rules="rules"
        label-placement="left"
        label-width="auto"
        require-mark-placement="right-hanging"
        size="large"
      >
        <n-form-item path="name">
          <n-input type="text" v-model:value="userPasswor.password" placeholder="密码" />
        </n-form-item>
        <n-form-item path="password">
          <n-input
            type="password"
            show-password-on="click"
            v-model:value="userPasswor.newPassword"
            placeholder="新密码"
          />
        </n-form-item>
        <n-form-item path="password">
          <n-input
            type="password"
            show-password-on="click"
            v-model:value="userPasswor.resetPassword"
            placeholder="重复新密码"
          />
        </n-form-item>
        <div>
          <n-button :loading="loading" type="primary" :block="true" @click="() => {}">修改密码</n-button>
        </div>
      </n-form>
    </n-drawer-content>
  </n-drawer>
</template>
<script setup lang="ts" name="Person">
import { ref } from "vue";
import type { FormInst } from "naive-ui";
const props = defineProps<{ active: boolean }>();
const emits = defineEmits<{ changeStatus: () => void }>();
const loading = ref<boolean>(false);
const formRef = ref<FormInst | null>(null);
const userPasswor = ref({
  password: "",
  newPassword: "",
  resetPassword: "",
});
const rules = {};
</script>

<style lang="less"></style>
