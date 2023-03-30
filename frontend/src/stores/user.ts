import { defineStore } from "pinia";
import { login } from "@/services/user";

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
  },
});
