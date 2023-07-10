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
  import { showSuccessToast, showFailToast } from 'vant';
  import { updatePassword } from '@/api/system/user';

  const pattern = /\S{6}/;
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

  const submit = async (value: ResetPassword.Password) => {
    const resp = await updatePassword({
      id: userStore.getUserInfo.ID as number,
      password: value.password,
      newPassword: value.newPassword,
    });
    if (resp?.code === 0) {
      showSuccessToast('更新成功，请重新登录');
      await userStore.Logout();
    } else {
      showFailToast(resp.msg);
    }
  };
</script>

<style scoped></style>
