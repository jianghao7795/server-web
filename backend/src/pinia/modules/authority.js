import { getPolicyPathByAuthorityId } from "@/api/casbin";
import { useUserStore } from "./user";
import { defineStore } from "pinia";

export const useAuthorityStore = defineStore("authority", {
  state: () => ({
    data: [],
  }),

  getters: {
    getDate: (state) => state.data,
  },

  actions: {
    async getAuthority() {
      const userStore = useUserStore();
      //   console.log(userStore.userInfo);
      if (userStore.userInfo?.authority?.authorityId) {
        const response = await getPolicyPathByAuthorityId({ id: userStore.userInfo?.authority?.authorityId });
        if (response?.code === 0) {
          this.data = response.data.pahts;
        }
      }
    },

    isAuthority(url = "", method = "get") {
      if (url === "") return false;
      let status = false;
      const len = this.data.length;
      if (len !== 0) {
        for (let i = 0; i < len; i++) {
          if (this.data[i].method === method) {
            if (url.split("/").length === this.data[i].path.split("/").length) {
              if (this.data[i].path.includes(":")) {
                const lastIndexPath = this.data[i].path.substring(0, this.data[i].path.lastIndexOf("/"));
                const urlPath = url.substring(0, url.lastIndexOf("/"));
                if (lastIndexPath === urlPath) {
                  status = true;
                }
              }
            } else {
              status = false;
            }
          }
        }

        return status;
      }

      return false;
    },
  },
});
