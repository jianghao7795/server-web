<template>
  <van-form ref="formRef" v-if="getShow" class="flex flex-col" @submit="handleRegister">
    <van-cell-group inset class="enter-y !mx-0 !mb-60px">
      <van-field
        class="enter-y items-center !rounded-md"
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
        class="enter-y items-center !rounded-md"
        v-model="formData.realname"
        name="realname"
        placeholder="真实姓名"
        :rules="getFormRules.username"
      >
        <template #left-icon>
          <Icon>
            <UserOutlined />
          </Icon>
        </template>
      </van-field>

      <van-field
        class="enter-y items-center !rounded-md"
        v-model="formData.phone"
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

      <!-- <van-field
        class="enter-y items-center !rounded-md"
        v-model="formData.sms"
        center
        clearable
        placeholder="请输入短信验证码"
        :rules="getFormRules.sms"
      >
        <template #left-icon>
          <Icon>
            <EditOutlined />
          </Icon>
        </template>
        <template #button>
          <van-button size="small" type="primary">发送验证码</van-button>
        </template>
      </van-field> -->
      <SendMessage :data="formData.sms" @updateData="(v:string) => formData.sms = v" />
      <van-field
        class="enter-y items-center !rounded-md"
        v-model="formData.password"
        :type="switchPassType ? 'password' : 'text'"
        name="password"
        placeholder="密码"
        :rules="getFormRules.password"
        @click-right-icon="switchPassType = !switchPassType"
      >
        <template #left-icon>
          <Icon>
            <LockOutlined />
          </Icon>
        </template>
        <template #right-icon>
          <Icon v-if="switchPassType">
            <EyeInvisibleOutlined />
          </Icon>
          <Icon v-else>
            <EyeOutlined />
          </Icon>
        </template>
      </van-field>

      <van-field
        class="enter-y items-center !rounded-md"
        v-model="formData.confirmPassword"
        :type="switchConfirmPassType ? 'password' : 'text'"
        name="confirmPassword"
        placeholder="确认密码"
        :rules="getFormRules.confirmPassword"
        @click-right-icon="switchConfirmPassType = !switchConfirmPassType"
      >
        <template #left-icon>
          <Icon>
            <LockOutlined />
          </Icon>
        </template>
        <template #right-icon>
          <Icon v-if="switchConfirmPassType">
            <EyeInvisibleOutlined />
          </Icon>
          <Icon v-else>
            <EyeOutlined />
          </Icon>
        </template>
      </van-field>

      <van-field
        name="policy"
        class="enter-y items-center px-1 !rounded-md"
        :rules="getFormRules.policy"
      >
        <template #input>
          <van-checkbox v-model="formData.policy" icon-size="14px" shape="square">
            我同意 xxx <a @click="viewPrivacyPolicy">隐私政策</a>
          </van-checkbox>
        </template>
      </van-field>
    </van-cell-group>

    <van-action-sheet v-model:show="show" title="隐私政策">
      <div class="content">内容</div>
      <div class="content">内容</div>
      <div class="content">内容</div>
      <div class="content">内容</div>
      <div class="content">内容</div>
      <div class="content">内容</div>
      <div class="content">内容</div>
      <div class="content">内容</div>
      <div class="content">内容</div>
      <div class="content">内容</div>
      <div class="content">内容</div>
      <div class="content">内容</div>
      <div class="content">内容</div>
      <div class="content">内容</div>
      <div class="content">内容</div>
      <div class="content">内容</div>
      <div class="content">内容</div>
      <div class="content">内容</div>
      <div class="content">内容</div>
    </van-action-sheet>

    <van-button
      class="enter-y !mb-25px !rounded-md"
      type="primary"
      block
      native-type="submit"
      :loading="loading"
    >
      注 册
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
  import { FormInstance, showToast } from 'vant';
  import { Icon } from '@vicons/utils';
  import {
    UserOutlined,
    MobileOutlined,
    LockOutlined,
    EyeOutlined,
    EyeInvisibleOutlined,
  } from '@vicons/antd';
  import { LoginStateEnum, useLoginState, useFormRules } from './useLogin';
  import SendMessage from './components/SendMessage.vue';
  import { register } from '@/api/system/user';

  const { handleBackLogin, getLoginState } = useLoginState();
  const getShow = computed(() => unref(getLoginState) === LoginStateEnum.REGISTER);

  const loading = ref(false);
  const formRef = ref<FormInstance>();

  const show = ref(false);

  const viewPrivacyPolicy = () => {
    show.value = true;
  };

  const formData = reactive({
    username: '',
    phone: '',
    realname: '',
    sms: '',
    password: '',
    confirmPassword: '',
    policy: false,
  });

  const { getFormRules } = useFormRules(formData);

  const switchPassType = ref(true);
  const switchConfirmPassType = ref(true);

  function handleRegister() {
    formRef.value
      ?.validate()
      .then(async () => {
        try {
          loading.value = true;
          // do something
          const resp = await register(formData);
          showToast('注册成功，请登录');
          handleBackLogin();
          // console.log('%c [  ]-167', 'font-size:13px; background:pink; color:#bf2c9f;');
        } finally {
          loading.value = false;

          // console.log('%c [  ]-171', 'font-size:13px; background:pink; color:#bf2c9f;');
        }
      })
      .catch(() => {
        console.log(formData);
        console.error('验证失败');
      });
  }
</script>

<style scoped lang="less"></style>
