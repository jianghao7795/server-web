<template>
  <div>
    <NavBar />
    <van-form @submit="submit">
      <van-cell-group inset>
        <van-field
          v-model="phoneDetail.phone"
          name="password"
          placeholder="旧号码"
          :rules="[{ required: true, message: '旧号码' }]"
          required
          disabled
        />
        <van-field
          v-model="phoneDetail.newPhone"
          name="newPhone"
          placeholder="新手机号码"
          :rules="[
            { required: true, message: '请输入' },
            { pattern, message: '输入正确号码' },
          ]"
          required
        />
        <van-field
          v-model="phoneDetail.verification"
          name="verification"
          placeholder="请输入短信验证码"
          :rules="[{ required: true, message: '请输入' }]"
          required
          :maxlength="6"
        >
          <template #button>
            <van-count-down ref="timeRef" :time="timeDown" :auto-start="false" @finish="onFinish">
              <template #default="timeData">
                <van-button
                  size="small"
                  :type="timeBoolean ? 'default' : 'primary'"
                  :disabled="timeBoolean || phoneDetail.newPhone.length === 0"
                  @click="handleSendMessage"
                >
                  {{ timeBoolean ? `(${timeData.seconds})已发送` : '发送短信' }}
                </van-button>
              </template>
            </van-count-down>
          </template>
        </van-field>
      </van-cell-group>
      <div style="margin: 16px">
        <van-button round block type="primary" native-type="submit"> 提交 </van-button>
      </div>
    </van-form>
  </div>
</template>

<script setup lang="ts" name="ChangePhone">
  import NavBar from './components/NavBar.vue';
  import { onMounted, reactive, ref } from 'vue';
  import { useUserStore } from '@/store/modules/user';
  import type { CountDownInstance } from 'vant';

  const timeDown = ref<number>(60000);
  const timeBoolean = ref<boolean>(false);
  const timeRef = ref<CountDownInstance>();

  const pattern = /1\d{10}/;
  const phoneDetail = reactive<ResetPassword.Phone>({
    phone: '',
    newPhone: '',
    verification: '',
  });

  const handleSendMessage = () => {
    if (phoneDetail.newPhone.length !== 11) {
      return;
    }
    timeBoolean.value = true;
    timeRef.value?.start();
  };

  const onFinish = () => {
    timeBoolean.value = false;
    timeRef.value?.reset();
  };

  const userStore = useUserStore();
  onMounted(() => {
    phoneDetail.phone = userStore.getUserInfo.phone;
  });

  const submit = (value: ResetPassword.Phone) => {
    console.log('submit', value);
    console.log('user is ', userStore);
  };
</script>

<style scoped></style>
