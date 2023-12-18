<template>
  <div>
    <el-row>
      <el-col :span="6">
        <el-card style="height: 580px">
          <div slot="header">个人信息</div>

          <div class="fl-center avatar-box">
            <div class="user-card">
              <div
                class="user-headpic-update"
                :style="{
                  'background-image': `url(${userStore.userInfo.headerImg && userStore.userInfo.headerImg.slice(0, 4) !== 'http' ? path + userStore.userInfo.headerImg : userStore.userInfo.headerImg})`,
                  'background-repeat': 'no-repeat',
                  'background-size': 'cover',
                }"
              >
                <span class="update" @click="openChooseImg">
                  <el-icon>
                    <edit />
                  </el-icon>
                  重新上传
                </span>
              </div>
              <div class="user-personality">
                <p v-if="!editFlag" class="nickName">
                  {{ userStore.userInfo.nickName }}
                  <el-icon class="pointer" color="#66b1ff" @click="openEidt">
                    <edit />
                  </el-icon>
                </p>
                <p v-if="editFlag" class="nickName">
                  <el-input v-model="nickName" />
                  <el-icon class="pointer" color="#67c23a" @click="enterEdit">
                    <check />
                  </el-icon>
                  <el-icon class="pointer" color="#f23c3c" @click="closeEdit">
                    <close />
                  </el-icon>
                </p>
                <p class="person-info">这个家伙很懒，什么都没有留下</p>
              </div>
              <div class="user-information">
                <ul>
                  <li>
                    <el-icon>
                      <user />
                    </el-icon>
                    {{ userStore.userInfo.nickName }}
                  </li>
                  <el-tooltip class="item" effect="light" content="前端事业群" placement="top">
                    <li>
                      <el-icon>
                        <data-analysis />
                      </el-icon>
                      前端事业群
                    </li>
                  </el-tooltip>
                  <li>
                    <el-icon>
                      <video-camera />
                    </el-icon>
                    中国
                  </li>
                  <el-tooltip class="item" effect="light" content="Vue" placement="top">
                    <li>
                      <el-icon>
                        <medal />
                      </el-icon>
                      Vue
                    </li>
                  </el-tooltip>
                </ul>
              </div>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="18">
        <el-card style="height: 580px">
          <div slot="header"></div>
          <div>
            <div class="user-addcount">
              <el-tabs v-model="activeName" @tab-click="handleClick">
                <el-tab-pane label="账号绑定" name="second">
                  <ul>
                    <li>
                      <p class="title">密保手机</p>
                      <p class="desc">
                        已绑定手机:{{ userStore.userInfo.phone }}
                        <a href="javascript:void(0)" @click="changePhoneFlag = true">立即修改</a>
                      </p>
                    </li>
                    <li>
                      <p class="title">密保邮箱</p>
                      <p class="desc">
                        已绑定邮箱：{{ userStore.userInfo.email }}
                        <a href="javascript:void(0)" @click="changeEmailFlag = true">立即修改</a>
                      </p>
                    </li>
                    <li>
                      <p class="title">密保问题</p>
                      <p class="desc">
                        {{ !hasSetting ? "未设置密保问题" : "修改密保" }}
                        <a href="javascript:void(0)" @click="changeSecurity(true)">
                          {{ !hasSetting ? "去设置" : "立即修改" }}
                        </a>
                      </p>
                    </li>
                    <li>
                      <p class="title">修改密码</p>
                      <p class="desc">
                        修改个人密码
                        <a href="javascript:void(0)" @click="showPassword = true">修改密码</a>
                      </p>
                    </li>
                  </ul>
                </el-tab-pane>
              </el-tabs>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <ChooseImg ref="chooseImgRef" @enter-img="enterImg" />

    <el-dialog v-model="showPassword" title="修改密码" width="360px" @close="clearPassword">
      <el-form ref="modifyPwdForm" :model="pwdModify" :rules="rules" label-width="80px">
        <el-form-item :minlength="6" label="原密码" prop="password">
          <el-input v-model="pwdModify.password" show-password />
        </el-form-item>
        <el-form-item :minlength="6" label="新密码" prop="newPassword">
          <el-input v-model="pwdModify.newPassword" show-password />
        </el-form-item>
        <el-form-item :minlength="6" label="确认密码" prop="confirmPassword">
          <el-input v-model="pwdModify.confirmPassword" show-password />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="showPassword = false">取 消</el-button>
          <el-button size="small" type="primary" @click="savePassword">确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog v-model="changePhoneFlag" title="绑定手机" width="600px">
      <el-form :model="phoneForm">
        <el-form-item label="手机号" label-width="120px">
          <el-input v-model="phoneForm.phone" placeholder="请输入手机号" autocomplete="off" />
        </el-form-item>
        <el-form-item label="验证码" label-width="120px">
          <div class="code-box">
            <el-input v-model="phoneForm.code" autocomplete="off" placeholder="请自行设计短信服务，此处为模拟随便写" style="width: 300px" />
            <el-button size="small" type="primary" :disabled="time > 0" @click="getCode">
              {{ time > 0 ? `(${time}s)后重新获取` : "获取验证码" }}
            </el-button>
          </div>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button size="small" @click="closeChangePhone">取消</el-button>
          <el-button type="primary" size="small" @click="changePhone">更改</el-button>
        </span>
      </template>
    </el-dialog>

    <el-dialog v-model="changeEmailFlag" title="绑定邮箱" width="600px">
      <el-form :model="emailForm">
        <el-form-item label="邮箱" label-width="120px">
          <el-input v-model="emailForm.email" placeholder="请输入邮箱" autocomplete="off" />
        </el-form-item>
        <el-form-item label="验证码" label-width="120px">
          <div class="code-box">
            <el-input v-model="emailForm.code" placeholder="请自行设计邮件服务，此处为模拟随便写" autocomplete="off" style="width: 300px" />
            <el-button size="small" type="primary" :disabled="emailTime > 0" @click="getEmailCode">
              {{ emailTime > 0 ? `(${emailTime}s)后重新获取` : "获取验证码" }}
            </el-button>
          </div>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button size="small" @click="closeChangeEmail">取消</el-button>
          <el-button type="primary" size="small" @click="changeEmail">更改</el-button>
        </span>
      </template>
    </el-dialog>

    <el-dialog title="设置密保问题" v-model="securityQuestion" width="600px" @close="changeSecurity(false)">
      <div v-if="!!hasSetting">
        <el-steps :active="active" simple :space="200" class="botto">
          <el-step title="回答问题" :icon="Edit" />
          <el-step title="设置问题" :icon="Upload" />
          <el-step title="成功" :icon="Picture" />
        </el-steps>
        <div v-if="active === 0">
          <el-form :model="random">
            <el-form-item label="问题" label-width="100px">
              <el-input v-model="random.problem" :disabled="true" />
            </el-form-item>
            <el-form-item label="答案" label-width="100px">
              <el-input v-model="random.answer" />
            </el-form-item>
          </el-form>
        </div>
        <div v-if="active === 1">
          <div v-for="item in securityQuestionList" :key="item.id">
            <el-form :model="item">
              <el-form-item label="问题" label-width="100px">
                <el-input v-model="item.problem" />
              </el-form-item>
              <el-form-item label="答案" label-width="100px">
                <el-input v-model="item.answer" />
              </el-form-item>
            </el-form>
          </div>
          <div>
            <el-form-item label="" label-width="100px">
              <el-button size="small" @click="addChallenge" :disabled="securityQuestionList.length >= 4" :plain="true" style="width: 100%" :icon="Plus"></el-button>
            </el-form-item>
          </div>
        </div>
        <div style="text-align: center">
          <el-button v-show="active < 3" v-if="active !== 0" @click="forwardStep">上一步</el-button>
          <el-button v-else-if="active === 0" @click="forwardStep">换个问题</el-button>
          <el-button type="primary" @click="nextStep">{{ active === 3 ? "完成" : "下一步" }}</el-button>
        </div>
      </div>
      <div v-else>
        <div v-for="item in securityQuestionList" :key="item.id">
          <el-form :model="item">
            <el-form-item label="问题" label-width="100px">
              <el-input v-model="item.problem" />
            </el-form-item>
            <el-form-item label="答案" label-width="100px">
              <el-input v-model="item.answer" />
            </el-form-item>
          </el-form>
        </div>
        <div>
          <el-form-item label="" label-width="100px">
            <el-button size="small" @click="addChallenge" :disabled="securityQuestionList.length >= 4" :plain="true" style="width: 100%" :icon="Plus"></el-button>
          </el-form-item>
        </div>
        <div style="text-align: center">
          <el-button type="primary" @click="setProblem">设置</el-button>
        </div>
      </div>
    </el-dialog>

    <ImageCropModal :img-url="cropUrl" :dialog-form-visible="cropStatus" :crop-data="cropData" @changeImageStatus="changeImageStatus" @changeAvatar="changeAvatar" />
  </div>
</template>

<script>
export default {
  name: "Person",
};
</script>

<script setup>
import ChooseImg from "@/components/chooseImg/index.vue";
import { setSelfInfo, changePassword, getProblemList, isSettingProblem, VerifyAnswer, putProblem } from "@/api/user.js";
import { reactive, ref, onMounted } from "vue";
import { ElMessage } from "element-plus";
import { useUserStore } from "@/pinia/modules/user";
import { Edit, Picture, Plus, Upload } from "@element-plus/icons-vue";
import ImageCropModal from "./ImageCropModal.vue";

const path = ref(import.meta.env.VITE_BASE_API + "/");
const activeName = ref("second");
const rules = reactive({
  password: [
    { required: true, message: "请输入密码", trigger: "blur" },
    { min: 6, message: "最少6个字符", trigger: "blur" },
  ],
  newPassword: [
    { required: true, message: "请输入新密码", trigger: "blur" },
    { min: 6, message: "最少6个字符", trigger: "blur" },
  ],
  confirmPassword: [
    { required: true, message: "请输入确认密码", trigger: "blur" },
    { min: 6, message: "最少6个字符", trigger: "blur" },
    {
      validator: (rule, value, callback) => {
        if (value !== pwdModify.value.newPassword) {
          callback(new Error("两次密码不一致"));
        } else {
          callback();
        }
      },
      trigger: "blur",
    },
  ],
});

const random = ref({});
const kiomRabdom = ref(0);
const active = ref(0);
const securityQuestion = ref(false);
const securityQuestionList = ref([]);
const hasSetting = ref(false);
//被裁剪图片的url
const cropUrl = ref("");
// 裁剪图片的modal status
const cropStatus = ref(false);
// 裁剪图片中的某条数据
const cropData = ref({});

const userStore = useUserStore();

const addChallenge = () => {
  securityQuestionList.value = [...securityQuestionList.value, { problem: undefined, answer: undefined, sys_user_id: userStore.userInfo.ID }];
};

const setProblem = async () => {
  securityQuestionList.value = securityQuestionList.value.map((i) => ({ ...i, sys_user_id: userStore.userInfo.ID }));
  const resp = await putProblem({ data: securityQuestionList.value });
  if (resp?.code === 0) {
    if (resp.data) {
      ElMessage.success({
        message: "设置成功",
      });
      securityQuestion.value = false;
      hasSetting.value = true;
    } else {
      ElMessage.warning({
        message: "设置失败",
      });
    }
  }
};

const nextStep = async () => {
  const activeValue = active.value;
  if (activeValue === 0) {
    const resp = await VerifyAnswer({ data: random.value });
    if (resp?.code === 0) {
      if (resp.data) {
        ElMessage.success({
          message: "回答成功，请设置问题并提交",
        });
        active.value = active.value + 1;
      } else {
        ElMessage.warning({
          message: "回答错误，请更改答案",
        });
      }
      // console.log(resp);
      // active.value = active.value + 1;
    }
  }

  if (activeValue === 1) {
    console.log(securityQuestionList.value);
    const resp = await putProblem({ data: securityQuestionList.value });
    if (resp?.code === 0) {
      if (resp.data) {
        ElMessage.success({
          message: "设置成功",
        });
        active.value = active.value + 2;
      } else {
        ElMessage.warning({
          message: "设置失败",
        });
      }
      // console.log(resp);
      // active.value = active.value + 1;
    }
  }

  if (activeValue === 3) {
    securityQuestion.value = false;
  }
};

const forwardStep = () => {
  if (active.value === 0) {
    if (kiomRabdom.value === securityQuestionList.value.length - 1) {
      kiomRabdom.value = 0;
    } else {
      kiomRabdom.value = kiomRabdom.value + 1;
    }
    random.value = securityQuestionList.value[kiomRabdom.value];

    return;
  }
  active.value = active.value - 1;
};

// console.log(userStore.userInfo);

onMounted(() => {
  isSettingProblem({ uid: userStore.userInfo.ID }).then((resp) => {
    if (resp?.code === 0) {
      hasSetting.value = resp?.data || false;
    }
  });
  // getProblemList({ id: userStore.userInfo.ID }).then((resp) => {
  //   if (resp.code === 0) {
  //     securityQuestionList.value = resp.data.list || [];
  //     random.value = resp.data.list[0];
  //     random.value;
  //   }
  // });
});
// watch([], (newUserStore, prevUserStore) => {
//   console.log(newUserStore, prevUserStore);
// });

const changeSecurity = (status = false) => {
  securityQuestion.value = status;
  if (status && hasSetting.value) {
    getProblemList({ id: userStore.userInfo.ID }).then((resp) => {
      if (resp.code === 0) {
        securityQuestionList.value = resp.data.list || [];
        random.value = resp.data.list[0];
      }
    });
  } else {
    securityQuestionList.value = [{ problem: undefined, answer: undefined }];
    // console.log(securityQuestionList.value);
  }
};

const modifyPwdForm = ref(null);
const showPassword = ref(false);
const pwdModify = ref({});
const nickName = ref("");
const editFlag = ref(false);
const savePassword = () => {
  modifyPwdForm.value.validate(async (valid) => {
    if (valid) {
      const res = await changePassword({
        username: userStore.userInfo.userName,
        password: pwdModify.value.password,
        newPassword: pwdModify.value.newPassword,
      });
      if (res.code === 0) {
        ElMessage.success("修改密码成功！");
      }
      showPassword.value = false;
    } else {
      return false;
    }
  });
};

const clearPassword = () => {
  pwdModify.value = {
    password: "",
    newPassword: "",
    confirmPassword: "",
  };
  modifyPwdForm.value.clearValidate();
};

const chooseImgRef = ref(null);
const openChooseImg = () => {
  chooseImgRef.value.open();
};

const changeImageStatus = (status = false, url = "", data) => {
  if (status) {
    cropUrl.value = url;
    cropStatus.value = status;
    cropData.value = data;
  } else {
    cropUrl.value = "";
    cropStatus.value = false;
    cropData.value = {};
  }
};

const enterImg = async (url, record) => {
  changeImageStatus(true, `/backend/${url}`, record);
};

const changeAvatar = async (url) => {
  const res = await setSelfInfo({ headerImg: url });
  if (res.code === 0) {
    userStore.ResetUserInfo({ headerImg: url });
    ElMessage({
      type: "success",
      message: "设置成功",
    });
  }
};

const openEidt = () => {
  nickName.value = userStore.userInfo.nickName;
  editFlag.value = true;
};

const closeEdit = () => {
  nickName.value = "";
  editFlag.value = false;
};

const enterEdit = async () => {
  const res = await setSelfInfo({
    nickName: nickName.value,
  });
  if (res.code === 0) {
    userStore.ResetUserInfo({ nickName: nickName.value });
    ElMessage({
      type: "success",
      message: "设置成功",
    });
  }
  nickName.value = "";
  editFlag.value = false;
};

const handleClick = (tab, event) => {
  console.log(tab, event);
};

const changePhoneFlag = ref(false);
const time = ref(0);
const phoneForm = reactive({
  phone: "",
  code: "",
});

const getCode = async () => {
  time.value = 60;
  let timer = setInterval(() => {
    time.value--;
    if (time.value <= 0) {
      clearInterval(timer);
      timer = null;
    }
  }, 1000);
};

const closeChangePhone = () => {
  changePhoneFlag.value = false;
  phoneForm.phone = "";
  phoneForm.code = "";
};

const changePhone = async () => {
  const res = await setSelfInfo({ phone: phoneForm.phone });
  if (res.code === 0) {
    ElMessage.success("修改成功");
    userStore.ResetUserInfo({ phone: phoneForm.phone });
    closeChangePhone();
  }
};

const changeEmailFlag = ref(false);
const emailTime = ref(0);
const emailForm = reactive({
  email: "",
  code: "",
});

const getEmailCode = async () => {
  emailTime.value = 60;
  let timer = setInterval(() => {
    emailTime.value--;
    if (emailTime.value <= 0) {
      clearInterval(timer);
      timer = null;
    }
  }, 1000);
};

const closeChangeEmail = () => {
  changeEmailFlag.value = false;
  emailForm.email = "";
  emailForm.code = "";
};

const changeEmail = async () => {
  const res = await setSelfInfo({ email: emailForm.email });
  if (res.code === 0) {
    ElMessage.success("修改成功");
    userStore.ResetUserInfo({ email: emailForm.email });
    closeChangeEmail();
  }
};
</script>

<style lang="scss">
.botto {
  margin-bottom: 15px;
}
.avatar-uploader .el-upload {
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}
.avatar-uploader .el-upload:hover {
  border-color: #409eff;
}
.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  line-height: 178px;
  text-align: center;
}
.avatar {
  width: 178px;
  height: 178px;
  display: block;
}
.avatar-box {
  // box-shadow: -2px 0 20px -16px;
  width: 100%;
  height: 100%;
  .user-card {
    min-height: calc(90vh - 200px);
    padding: 30px 20px;
    text-align: center;
    .el-avatar {
      border-radius: 50%;
    }
    .user-personality {
      padding: 24px 0;
      text-align: center;
      p {
        font-size: 16px;
      }
      .nickName {
        display: flex;
        justify-content: center;
        align-items: center;
        font-size: 26px;
      }
      .person-info {
        margin-top: 6px;
        font-size: 14px;
        color: #999;
      }
    }
    .user-information {
      width: 100%;
      height: 100%;
      text-align: left;
      ul {
        display: inline-block;
        height: 100%;
        width: 100%;
        li {
          width: 100%;
          white-space: nowrap;
          overflow: hidden;
          text-overflow: ellipsis;
          i {
            margin-right: 8px;
          }
          padding: 20px 0;
          font-size: 16px;
          font-weight: 400;
          color: #606266;
        }
      }
    }
  }
}
.user-addcount {
  ul {
    li {
      .title {
        padding: 10px;
        font-size: 18px;
        color: #696969;
      }
      .desc {
        font-size: 16px;
        padding: 0 10px 20px 10px;
        color: #a9a9a9;
        a {
          color: rgb(64, 158, 255);
          float: right;
        }
      }
      border-bottom: 2px solid #f0f2f5;
    }
  }
}
.user-headpic-update {
  width: 120px;
  height: 120px;
  line-height: 120px;
  margin: 0 auto;
  display: flex;
  justify-content: center;
  border-radius: 20px;
  &:hover {
    color: #fff;
    background: linear-gradient(to bottom, rgba(255, 255, 255, 0.15) 0%, rgba(0, 0, 0, 0.15) 100%), radial-gradient(at top center, rgba(255, 255, 255, 0.4) 0%, rgba(0, 0, 0, 0.4) 120%) #989898;
    background-blend-mode: multiply, multiply;
    .update {
      color: #fff;
      cursor: pointer;
    }
  }
  .update {
    height: 120px;
    width: 120px;
    text-align: center;
    color: transparent;
  }
}
.pointer {
  cursor: pointer;
}
.code-box {
  display: flex;
  justify-content: space-between;
}
</style>
