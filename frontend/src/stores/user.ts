import { defineStore } from "pinia";
import { login, getCurrentUser } from "@/services/user";

export const useUserStore = defineStore("user", {
  state: (): { currentUser: User.CurrentUser; loading: boolean } => ({
    currentUser: {
      user: { ID: 0, name: "", introduction: "", head_img: "", content: "", password: "" },
      token: "",
      exportAt: 0,
    },
    loading: false,
  }),
  getters: {
    getToken: (state) => state.currentUser.token,
  },
  actions: {
    async logins(data: User.Login, callback: () => void) {
      this.loading = true;
      try {
        const info = await login(data);
        this.loading = false;
        this.currentUser = info.data;
        callback();
        localStorage.setItem("token", info.data.token);
      } catch (e) {
        this.loading = false;
        window.$message.error("登录失败");
      }
    },
    async getUser() {
      const token = localStorage.getItem("token");
      if (!this.currentUser.token && token) {
        this.currentUser.token = token;
      } else {
        return;
      }
      try {
        const info = await getCurrentUser();
        this.currentUser.user = info.data.user;
        this.currentUser.exportAt = info.data.exportAt;
      } catch (e) {
        window.$message.error("个人信息获取失败，请重新登录");
      }
    },
  },
});
