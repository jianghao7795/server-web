<template>
  <div>
    <NavBar />
    <van-divider>基本信息</van-divider>
    <van-field
      label="头像"
      label-class="font-bold"
      input-align="right"
      :center="true"
      :border="false"
      is-link
      readonly
    >
      <template #input>
        <UploaderImage @update-img="onChangeImage" name="avatar" @updateUser="changeUser">
          <van-image class="avatar" round fit="cover" :src="`/api/mobile/${state.avatar}`" />
        </UploaderImage>
      </template>
    </van-field>

    <van-field
      label="昵称"
      readonly
      label-class="font-bold"
      input-align="right"
      :center="true"
      :border="false"
      v-model="state.nickname"
      is-link
      to="/editNickname"
    />

    <van-field
      label="性别"
      readonly
      label-class="font-bold"
      input-align="right"
      :center="true"
      :border="false"
      v-model="state.genderText"
      is-link
      @click="showGenderPicker = true"
    />

    <van-field
      label="签名"
      readonly
      label-class="font-bold"
      input-align="right"
      :center="true"
      :border="false"
      v-model="state.sign"
      is-link
      to="/editSign"
    />

    <van-field
      label="主页封面"
      label-class="font-bold"
      input-align="right"
      :center="true"
      :border="false"
      is-link
      readonly
    >
      <template #input>
        <UploaderImage @update-img="onChangeImage" name="cover" @updateUser="changeUser">
          <van-image
            class="cover"
            fit="cover"
            :src="state.cover ? `/api/mobile/${state.cover}` : `/api/mobile/${state.avatar}`"
          />
        </UploaderImage>
      </template>
    </van-field>

    <van-field
      label="行业"
      readonly
      label-class="font-bold"
      input-align="right"
      :center="true"
      :border="false"
      v-model="state.industryText"
      is-link
      @click="showIndustryPicker = true"
    />

    <van-popup v-model:show="showGenderPicker" position="bottom" round>
      <van-picker
        visible-option-num="3"
        v-model="state.genderValues"
        :columns="genderColumns"
        @confirm="handleGender"
        @cancel="showGenderPicker = false"
      />
    </van-popup>

    <van-popup v-model:show="showIndustryPicker" position="bottom" round>
      <van-picker
        v-model="state.industryValues"
        :columns="industryColumns"
        @confirm="handleIndustry"
        @cancel="showIndustryPicker = false"
      />
    </van-popup>
  </div>
</template>

<script setup lang="ts">
  import { onMounted, reactive, ref } from 'vue';
  import NavBar from './components/NavBar.vue';
  import UploaderImage from './components/UploaderImage.vue';
  import { useUserStore } from '@/store/modules/user';
  import { FormColumns, genderColumns, industryColumns } from './pickColumns';
  import { showFailToast, showToast } from 'vant';
  import { updateUser } from '@/api/system/user';

  const userStore = useUserStore();
  // const { avatar, gender, industry, cover } = userStore.getUserInfo;

  const showGenderPicker = ref(false);
  const showIndustryPicker = ref(false);

  const changeUser = async (data: { field: string; value: string | number }) => {
    const response = await updateUser(data);
    if (response?.code === 0) {
      showToast('更新成功');
      return true;
    } else {
      showFailToast('更新失败');
      return false;
    }
  };

  const state = reactive({
    nickname: '',
    sign: '',
    genderText: '',
    industryText: '',
    industryValues: [0],
    genderValues: [0],
    gender: 0,
    industry: 0,
    cover: '',
    avatar: '',
  });

  const handleGender = async ({ selectedOptions, selectedValues }) => {
    const response = await updateUser({ field: 'gender', value: selectedValues[0] });
    if (response?.code === 0) {
      state.genderText = selectedOptions[0].text;
      showToast('更新成功');
      showGenderPicker.value = false;
    }
  };

  const handleIndustry = async ({ selectedValues, selectedOptions }) => {
    const response = await updateUser({ field: 'industry', value: selectedValues[0] });
    if (response?.code === 0) {
      state.industryText = selectedOptions[0].text;
      showToast('更新成功');
      // do something
      showIndustryPicker.value = false;
    }
  };

  const getFromText = (columns: FormColumns[], value = 0) =>
    columns.find((item) => item.value === value)?.text;

  function initState() {
    Object.keys(state).forEach((key) => {
      state[key] = userStore.getUserInfo[key];
    });
    // set field text value.
    state.genderText = getFromText(genderColumns, state.gender) ?? '';
    state.industryText = getFromText(industryColumns, state.industry) ?? '';
    // set the pick selected value.
    state.industryValues = [state.industry ?? 0];
    state.genderValues = [state.gender];
  }

  const onChangeImage = (imgUrl: string, field: string) => {
    const setUser = { ...userStore.getUserInfo, [field]: imgUrl };
    userStore.setUserInfo(setUser);
    initState();
  };

  onMounted(() => {
    initState();
  });
</script>

<style scoped lang="less">
  .avatar {
    width: 140px;
    height: 140px;
  }
  .cover {
    width: 170px;
    height: 100px;
    :deep(.van-image__img) {
      border-radius: 10px !important;
    }
  }

  :deep(.van-field__control) {
    color: var(--van-text-color-2);
  }
</style>
