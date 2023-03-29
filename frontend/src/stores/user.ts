import { defineStore } from "pinia";
import { login } from "@/services/user";

export const useUserStore = defineStore("user", {
  state: (): { user: User.UserInfo; token: string } => ({
    user: { ID: 0, name: "", introduction: "", head_img: "", content: "", password: "" },
    token: "",
  }),
  getters: {
    getToken: (state) => state.token,
  },
  actions: {
    async logins(data: User.Login) {
      console.log(data);
      const info = await login(data);
    },
  },
});
