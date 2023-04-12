import { defineStore } from "pinia";
import { login, getCurrentUser, registerUser } from "@/services/user";

export const useUserStore = defineStore("user", {
  state: (): { currentUser: User.CurrentUser; loading: boolean; loadingRegister: boolean } => ({
    currentUser: {
      user: { ID: 0, name: "", introduction: "", head_img: "", content: "", header: "" },
      token: "",
      exportAt: 0,
    },
    loading: false,
    loadingRegister: false,
  }),
  getters: {
    getToken: (state) => state.currentUser.token,
  },
  actions: {
    async register(data: User.Register, callback: () => void) {
      this.loadingRegister = true;
      try {
        const resp = await registerUser(data);
        this.loadingRegister = false;
        if (resp.code === 0) {
          callback();
        } else {
          window.$message.error("注册失败, 请重试");
        }
      } catch (e) {
        this.loadingRegister = false;
        window.$message.error("注册失败, 请重试");
      }
    },
    async logins(data: User.Login, callback: (imageString: string) => void) {
      this.loading = true;
      try {
        const info = await login(data);
        this.loading = false;
        this.currentUser = info.data;
        callback(info.data.user.head_img);
        localStorage.setItem("token", info.data.token);
      } catch (e) {
        this.loading = false;
        window.$message.error("登录失败");
      }
    },
    async getUser(callback: (imageUrl: string) => void) {
      const token = localStorage.getItem("token");
      if (!this.currentUser.token && token) {
        this.currentUser.token = token;
      } else {
        return;
      }
      try {
        this.loading = true;
        const info = await getCurrentUser();
        this.loading = false;
        this.currentUser.user = info.data.user;
        this.currentUser.exportAt = info.data.exportAt;
        if (info.data.user.head_img !== "") {
          callback(info.data.user.head_img);
        }
      } catch (e) {
        localStorage.removeItem("token");
        this.loading = false;
        window.$message.error("个人信息获取失败，请重新登录");
      }
    },
  },
});
