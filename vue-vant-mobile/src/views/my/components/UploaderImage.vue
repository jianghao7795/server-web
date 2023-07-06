<template>
  <van-uploader
    :max-size="700 * 1024"
    :max-count="1"
    :before-read="beforeRead"
    :after-read="afterRead"
    accept="image/*"
  >
    <template #default>
      <slot name="default"></slot>
    </template>
  </van-uploader>
</template>

<script setup lang="ts">
  import { uploadImage } from '@/api/system/user';
  import { showFailToast } from 'vant';
  import type { UploaderAfterRead, UploaderBeforeRead } from 'vant/lib/uploader/types';
  type updateUserType = { field: string; value: string | number };

  const emit = defineEmits<{
    (e: 'update-img', imgUrl: string, field: string): void;
    (e: 'updateUser', data: updateUserType);
  }>();

  const props = defineProps<{ name: string }>();

  const beforeRead: UploaderBeforeRead = (file) => {
    if (Array.isArray(file)) {
      return false;
    }
    if (file.type === 'image/jpeg' || file.type === 'image/png' || file.type === 'image/jpg') {
      return true;
    }
    showFailToast('请上传正确格式的图片');
    return false;
  };

  const afterRead: UploaderAfterRead = async (file) => {
    if (Array.isArray(file)) {
      showFailToast('只能上传一张图片');
      return;
    }
    // console.log('%c [ file ]-43', 'font-size:13px; background:pink; color:#bf2c9f;', file);
    // 这里写上传逻辑
    const response = await uploadImage(file.file as File, file?.file?.name || '');
    if (response?.code === 0) {
      await emit('updateUser', {
        field: props.name,
        value: response?.data.file.url as string,
      });
      emit('update-img', response?.data.file.url as string, props.name);
    } else {
      showFailToast('上传失败，请重试');
    }
  };
</script>

<style scoped lang="less"></style>
