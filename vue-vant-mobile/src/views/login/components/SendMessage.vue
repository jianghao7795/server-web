<template>
  <van-field
    class="enter-y items-center !rounded-md"
    v-model="sms"
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
</template>

<script lang="ts" setup>
  import { ref, computed } from 'vue';
  import { useFormRules } from '../useLogin';
  import type { CountDownInstance } from 'vant';
  import { showToast } from 'vant';
  import { MessageOutlined } from '@vicons/antd';
  import { Icon } from '@vicons/utils';

  const props = defineProps<{
    data: string;
  }>();

  const emit = defineEmits<{
    (e: 'updateData', value: string): void;
  }>();

  const sms = computed({
    get: () => props.data,
    set: (val) => {
      emit('updateData', val);
    },
  });

  const { getFormRules } = useFormRules();
  const timeRef = ref<CountDownInstance>();
  const timeDown = ref<number>(60000);
  const timeBoolean = ref<boolean>(false);

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
</script>
