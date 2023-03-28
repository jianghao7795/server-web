import { defineStore } from "pinia";

export const useUserStore = defineStore("user", {
  state: (): { user: User.UserInfo; token: string } => ({
    user: { ID: 0, name: "", introduction: "", head_img: "", content: "" },
    token: "",
  }),
  getters: {
    getToken: (state) => state.token,
  },
  actions: {
    login(data: User.Login) {},
  },
});
