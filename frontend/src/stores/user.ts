import { defineStore } from "pinia";
import { login } from "@/services/user";

export const useUserStore = defineStore("user", {
  state: (): { user: User.UserInfo; token: string; exportAt: number } => ({
    user: { ID: 0, name: "", introduction: "", head_img: "", content: "", password: "" },
    token: "",
    exportAt: 0,
  }),
  getters: {
    getToken: (state) => state.token,
  },
  actions: {
    async logins(data: User.Login) {
      console.log(data);
      try {
        const info = await login(data);
        console.log(info);
      } catch (e) {
        console.log(e);
      }
    },
  },
});
