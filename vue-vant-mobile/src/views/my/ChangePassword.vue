<template>
  <div>
    <NavBar />
    <van-form @submit="submit">
      <van-cell-group inset>
        <van-field
          v-model="pwdDetail.password"
          name="password"
          label="密码"
          placeholder="密码"
          :rules="[{ required: true, message: '密码' }]"
          required
        />
        <van-field
          v-model="pwdDetail.newPassword"
          type="password"
          name="newPassword"
          label="新密码"
          placeholder="新密码"
          :rules="[
            { required: true, message: '请输入' },
            { pattern, message: '密码不少于6位' },
          ]"
          required
        />
        <van-field
          v-model="pwdDetail.repeatPassword"
          type="password"
          name="repeatPassword"
          label="重复密码"
          placeholder="重复密码"
          :rules="[
            { required: true, message: '请输入' },
            {
              validator,
              message: '与新密码不一致',
            },
          ]"
          required
        />
      </van-cell-group>
      <div style="margin: 16px">
        <van-button round block type="primary" native-type="submit"> 提交 </van-button>
      </div>
    </van-form>
  </div>
</template>

<script setup lang="ts">
  import NavBar from './components/NavBar.vue';
  import { reactive } from 'vue';
  import { useUserStore } from '@/store/modules/user';

  const pattern = /\S{6}/g;
  const pwdDetail = reactive<ResetPassword.Password>({
    password: '',
    newPassword: '',
    repeatPassword: '',
  });

  const userStore = useUserStore();

  const validator = (val: string) => {
    if (val === pwdDetail.newPassword) {
      return true;
    }

    return false;
  };

  const submit = (value: ResetPassword.Password) => {
    console.log('submit', value);
    console.log('user is ', userStore);
  };
</script>

<style scoped></style>
