<template>
  <div>
    <NavBar>
      <template #right><span class="text-32px" @click="handleNickname">保存</span></template>
    </NavBar>
    <van-form ref="formRef">
      <van-field
        class="mt-20px"
        name="sign"
        clearable
        v-model="formValue.sign"
        rows="4"
        autosize
        label="签名"
        type="textarea"
        maxlength="70"
        placeholder="介绍一下你自己"
        show-word-limit
      />
    </van-form>
  </div>
</template>

<script setup lang="ts">
  import NavBar from './components/NavBar.vue';
  import { useUserStore } from '@/store/modules/user';
  import { onMounted, reactive, ref } from 'vue';
  import { closeToast, FormInstance, showLoadingToast, showToast } from 'vant';
  import { updateUser } from '@/api/system/user';

  const userStore = useUserStore();

  const { sign } = userStore.getUserInfo;
  const formRef = ref<FormInstance>();

  const formValue = reactive({
    sign: '',
  });

  function handleNickname() {
    formRef.value
      ?.validate()
      .then(async () => {
        try {
          const formValue = formRef.value?.getValues();
          showLoadingToast({
            duration: 0,
            forbidClick: true,
            message: '更新中',
          });
          const response = await updateUser({ field: 'sign', value: formValue?.sign as string });
          if (response?.code === 0) {
            const changeUser = userStore.getUserInfo;
            userStore.setUserInfo({ ...changeUser, sign: response.data.value as string });
            closeToast();
            showToast('更新成功');
          }
          // do something
        } finally {
          // after successful
        }
      })
      .catch(() => {
        console.error('验证失败');
      });
  }

  onMounted(() => {
    formValue.sign = sign ?? '';
  });
</script>

<style scoped lang="less">
  .note {
    margin-top: 15px;
    font-size: 25px;
    color: var(--van-text-color-2);
  }
</style>
