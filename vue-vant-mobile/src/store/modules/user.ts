import { defineStore } from 'pinia';
import { createStorage } from '@/utils/Storage';
import { store } from '@/store';
import { ACCESS_TOKEN, CURRENT_USER } from '@/store/mutation-types';
import { ResultEnum } from '@/enums/httpEnum';
import { getUserInfo, login } from '@/api/system/user';
import { PageEnum } from '@/enums/pageEnum';
import router from '@/router';
import { useRoute } from 'vue-router';

const route = useRoute();
interface UserInfo {
  ID: string | number;
  username: string;
  realname: string;
  nickname: string;
  avatar: string;
  cover: string;
  gender: number;
  phone: string;
  sign?: string;
  industry?: number;
}

interface IUserState {
  token?: string;
  userInfo: Nullable<UserInfo>;
  lastUpdateTime: number;
}

interface LoginParams {
  username: string;
  password: string;
}

const Storage = createStorage({ storage: localStorage });

export const useUserStore = defineStore({
  id: 'app-user',
  state: (): IUserState => ({
    userInfo: null,
    token: undefined,
    lastUpdateTime: 0,
  }),
  getters: {
    getUserInfo(): UserInfo {
      return this.userInfo || Storage.get(CURRENT_USER, '') || {};
    },
    getToken(): string {
      const token: string = Storage.get(ACCESS_TOKEN, '');
      return this.token || token;
    },
    getLastUpdateTime(): number {
      return this.lastUpdateTime;
    },
  },
  actions: {
    setToken(token: string | undefined, expire: number) {
      this.token = token ? token : '';
      Storage.set(ACCESS_TOKEN, token, expire);
    },
    setUserInfo(info: UserInfo | null) {
      this.userInfo = info;
      this.lastUpdateTime = new Date().getTime();
      Storage.set(CURRENT_USER, info);
    },

    async Login(params: LoginParams) {
      try {
        const response = await login(params);
        const { data, code } = response;
        if (code === ResultEnum.SUCCESS) {
          // save token
          this.setToken(data.token, data.expiresAt);
        }
        return Promise.resolve(response);
      } catch (error) {
        return Promise.reject(error);
      }
    },

    async GetUserInfo() {
      getUserInfo()
        .then((res) => {
          // console.log(res);
          this.setUserInfo(res);
        })
        .catch((error) => {
          throw error;
        });
    },

    async Logout(url: string) {
      // if (this.getToken) {
      //   try {
      //     await doLogout();
      //   } catch {
      //     console.error('注销Token失败');
      //   }
      // }
      this.setToken(undefined, 0);
      this.setUserInfo(null);
      Storage.remove(ACCESS_TOKEN);
      Storage.remove(CURRENT_USER);
      router.push({
        path: PageEnum.BASE_LOGIN,
        query: {
          redirect: url,
        },
      });
      // location.reload();
    },
  },
});

// Need to be used outside the setup
export function useUserStoreWidthOut() {
  return useUserStore(store);
}
