<template>
  <div>
    <NavBar>
      <template #right><span class="text-32px" @click="handleNickname">保存</span></template>
    </NavBar>
    <van-form ref="formRef">
      <van-field
        class="mt-20px"
        name="nickname"
        v-model="formValue.nickname"
        placeholder="请输入昵称（2-12字）"
        :rules="[
          {
            validator: validateNickname(),
            trigger: 'onChange',
          },
        ]"
      />
    </van-form>

    <div class="note px-30px">
      <p>昵称支持2-12个中文字符或3-24个英文字符，</p>
      <p>符号仅支持”-“和”_“和”.“以及“·”</p>
    </div>
  </div>
</template>

<script setup lang="ts">
  import NavBar from './components/NavBar.vue';
  import { useUserStore } from '@/store/modules/user';
  import { onMounted, reactive, ref } from 'vue';
  import type { FormInstance } from 'vant';
  import { closeToast, showLoadingToast, showToast } from 'vant';
  import { updateUser } from '@/api/system/user';

  const userStore = useUserStore();

  const { nickname } = userStore.getUserInfo;
  const formRef = ref<FormInstance>();

  const formValue = reactive({
    nickname: '',
  });

  const validateNickname = () => {
    return async (value: string) => {
      const pattern = /^[\u4E00-\u9FA5A-Za-z0-9-_.·]+$/;
      if (!pattern.test(value)) {
        return Promise.resolve('请输入正确内容');
      }
      if (value.length < 2 || value.length > 12) {
        return Promise.resolve('长度不符合');
      }
      return Promise.resolve(true);
    };
  };

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
          const response = await updateUser({
            field: 'nickname',
            value: formValue?.nickname as string,
          });
          if (response?.code === 0) {
            const changeUser = userStore.getUserInfo;
            userStore.setUserInfo({ ...changeUser, nickname: response.data.value as string });
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
    formValue.nickname = nickname;
  });
</script>

<style scoped lang="less">
  .note {
    margin-top: 15px;
    font-size: 25px;
    color: var(--van-text-color-2);
  }
</style>
