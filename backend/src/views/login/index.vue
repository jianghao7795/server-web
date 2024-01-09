<template>
  <div id="userLayout">
    <div class="login_panle">
      <div class="login_panle_form">
        <div class="login_panle_form_title" style="text-align: center">
          <img class="login_panle_form_title_logo" :src="$GIN_VUE_ADMIN.appLogo" alt />
          <p class="login_panle_form_title_p">{{ $GIN_VUE_ADMIN.appName }}</p>
        </div>
        <el-form ref="loginForm" :model="loginFormData" :rules="rules" @keyup.enter="submitForm">
          <el-form-item prop="username">
            <el-input v-model="loginFormData.username" placeholder="请输入用户名">
              <template #suffix>
                <span class="input-icon">
                  <el-icon>
                    <user />
                  </el-icon>
                </span>
              </template>
            </el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input v-model="loginFormData.password" :type="lock === 'lock' ? 'password' : 'text'" placeholder="请输入密码">
              <template #suffix>
                <span class="input-icon">
                  <el-icon>
                    <component :is="lock" @click="changeLock" />
                  </el-icon>
                </span>
              </template>
            </el-input>
          </el-form-item>
          <el-form-item prop="captcha">
            <div class="vPicBox">
              <el-input v-model="loginFormData.captcha" placeholder="请输入验证码" style="width: 60%" />
              <div class="vPic">
                <img v-if="picPath" :src="picPath" alt="请输入验证码" @click="loginVerify()" />
                <!-- <el-icon v-else @click="loginVerify()"><RefreshLeft /></el-icon> -->
                <div v-else class="centerFont" @click="loginVerify()">刷新</div>
              </div>
            </div>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" style="width: 46%" size="large" v-if="isInit" @click="checkInit">前往初始化</el-button>
            <el-button v-bind:loading="loading" type="primary" size="large" :style="isInit ? { width: '46%', marginLeft: '8%' } : { width: '100%' }" @click="submitForm">登 录</el-button>
          </el-form-item>
        </el-form>
      </div>
      <div class="login_panle_right">
        <!-- <img style="width: 100%; margin: auto 0" src="@/assets/login_left.jpg" fit="contain" /> -->
      </div>
      <div class="login_panle_foot">
        <div class="copyright">
          <bootomInfo />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Login",
};
</script>

<script setup>
import { captcha } from "@/api/user";
import { checkDB } from "@/api/initdb";
import bootomInfo from "@/views/layout/bottomInfo/bottomInfo.vue";
import { reactive, ref, onMounted } from "vue";
import { ElMessage } from "element-plus";
import { useRouter } from "vue-router";
import { useUserStore } from "@/pinia/modules/user";
import dayjs from "dayjs";
import { setLocalStorage } from "@/utils/date";
import md5 from "md5";

const router = useRouter();

// 是否需要初始化
const isInit = ref(false);
onMounted(async () => {
  const resp = await checkDB();
  if (resp.data?.needInit) {
    isInit.value = true;
  }
});

// 验证函数
const checkUsername = (rule, value, callback) => {
  if (value.length < 3) {
    return callback(new Error("请输入正确的用户名"));
  } else {
    callback();
  }
};
const checkPassword = (rule, value, callback) => {
  if (value.length < 6) {
    return callback(new Error("请输入正确的密码"));
  } else {
    callback();
  }
};

// 获取验证码
const loginVerify = () => {
  captcha({}).then((ele) => {
    if (ele.code === 0) {
      rules.captcha[1].max = ele.data.captchaLength;
      rules.captcha[1].min = ele.data.captchaLength;
      picPath.value = ele.data.picPath;
      loginFormData.captchaId = ele.data.captchaId;
    }
  });
};
onMounted(() => {
  loginVerify();
});

// 登录相关操作
const lock = ref("lock");
const changeLock = () => {
  lock.value = lock.value === "lock" ? "unlock" : "lock";
};

const loginForm = ref(null);
const picPath = ref("");
const loginFormData = reactive({
  username: "wuhao",
  password: "123456",
  captcha: "",
  captchaId: "",
});
const loading = ref(false);
const rules = reactive({
  username: [{ validator: checkUsername, trigger: "blur" }],
  password: [{ validator: checkPassword, trigger: "blur" }],
  captcha: [
    { required: true, message: "请输入验证码", trigger: "blur" },
    {
      message: "验证码格式不正确",
      trigger: "blur",
    },
  ],
});

const userStore = useUserStore();
const login = async () => {
  return await userStore.LoginIn({ ...loginFormData, password: md5(loginFormData.password) });
};
const submitForm = () => {
  loginForm.value.validate(async (v) => {
    if (v) {
      loading.value = true;
      const flag = await login();
      loading.value = false;
      // console.log(flag);
      if (!flag) {
        loginVerify();
        loginFormData.captcha = "";
      } else {
        // console.log();
        setLocalStorage("workTime", {
          workStartTime: dayjs().unix().valueOf(),
        });
      }
    } else {
      ElMessage({
        type: "error",
        message: "请正确填写登录信息",
        showClose: true,
      });
      loginVerify();
      return false;
    }
  });
};

// 跳转初始化
const checkInit = async () => {
  if (isInit) {
    userStore.NeedInit();
    router.push({ name: "Init" });
  }
  // const res = await checkDB();
  // if (res.code === 0) {
  //   if (res.data?.needInit) {
  //     userStore.NeedInit();

  //   } else {
  //     // ElMessage({
  //     //   type: "info",
  //     //   message: "已配置数据库信息，无法初始化",
  //     // });
  //   }
  // }
};
</script>

<style lang="scss" scoped>
@import "@/style/newLogin.scss";

.centerFont {
  text-align: center;
  color: #4d70ff;
  cursor: pointer;
}
</style>
