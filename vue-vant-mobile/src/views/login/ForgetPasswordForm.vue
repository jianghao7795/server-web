<template>
  <van-form ref="formRef" v-if="getShow" class="flex flex-col items-center" @submit="handleReset">
    <van-field
      class="enter-y items-center mb-25px !rounded-md"
      v-model="formData.username"
      name="username"
      placeholder="用户名"
      :rules="getFormRules.username"
    >
      <template #left-icon>
        <Icon>
          <UserOutlined />
        </Icon>
      </template>
    </van-field>

    <van-field
      class="enter-y items-center mb-25px !rounded-md"
      v-model="formData.mobile"
      name="password"
      placeholder="手机号码"
      :rules="getFormRules.mobile"
    >
      <template #left-icon>
        <Icon>
          <MobileOutlined />
        </Icon>
      </template>
    </van-field>

    <van-field
      class="enter-y items-center mb-70px !rounded-md"
      v-model="formData.sms"
      center
      clearable
      placeholder="请输入短信验证码"
      :rules="getFormRules.sms"
    >
      <template #left-icon>
        <Icon>
          <MessageOutlined />
        </Icon>
      </template>
      <template #button>
        <van-count-down ref="timeRef" :time="timeDown" :auto-start="false" @finish="onFinish">
          <template #default="timeData">
            <van-button
              size="small"
              :type="timeBoolean ? 'default' : 'primary'"
              :disabled="timeBoolean"
              @click="handleSendMessage"
            >
              {{ timeBoolean ? `(${timeData.seconds})已发送` : '发送短信' }}
            </van-button>
          </template>
        </van-count-down>
      </template>
    </van-field>

    <van-button
      class="enter-y !mb-25px !rounded-md"
      type="primary"
      block
      native-type="submit"
      :loading="loading"
    >
      重 置
    </van-button>

    <van-button
      class="enter-y !mb-150px !rounded-md"
      plain
      type="primary"
      block
      @click="handleBackLogin"
    >
      返 回
    </van-button>
  </van-form>
</template>

<script setup lang="ts">
  import { computed, reactive, ref, unref } from 'vue';
  import type { CountDownInstance, FormInstance } from 'vant';
  import { showToast } from 'vant';
  import { Icon } from '@vicons/utils';
  import { UserOutlined, MobileOutlined, MessageOutlined } from '@vicons/antd';
  import { LoginStateEnum, useLoginState, useFormRules } from './useLogin';

  const { handleBackLogin, getLoginState } = useLoginState();
  const { getFormRules } = useFormRules();
  const getShow = computed(() => unref(getLoginState) === LoginStateEnum.RESET_PASSWORD);

  const timeRef = ref<CountDownInstance>();
  const timeDown = ref<number>(60000);
  const timeBoolean = ref<boolean>(false);

  const loading = ref<boolean>(false);
  const formRef = ref<FormInstance>();
  const formData = reactive({
    username: '',
    mobile: '',
    sms: '',
  });

  const handleSendMessage = () => {
    timeBoolean.value = true;
    timeRef.value?.start();
  };

  const onFinish = () => {
    timeBoolean.value = false;
    // console.log(timeBoolean.value);
    showToast('倒计时结束');
    timeRef.value?.reset();
  };

  function handleReset() {
    formRef.value
      ?.validate()
      .then(async () => {
        try {
          loading.value = true;
          throw '密码错误';
        } finally {
          loading.value = false;
        }
      })
      .catch((e: Error) => {
        console.error('验证失败');
        showToast(e);
      });
  }
</script>

<style scoped lang="less"></style>
