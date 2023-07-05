import { http } from '@/utils/http/axios';

export interface BasicResponseModel<T = any> {
  code: number;
  msg: string;
  data: T;
}

export type UserType = {
  ID: number;
  username: string;
  password: string;
  nickname: string;
  realname: string;
  avatar: string;
  cover: string;
  sign: string;
  industry: number;
  gender: 0 | 1;
  phone: string;
};

export type LoginResponse = {
  user: UserType;
  expiresAt: string;
  token: string;
};

/**
 * @description: 用户登录
 */
export function login(params: any) {
  return http.request<BasicResponseModel<LoginResponse>>(
    {
      url: '/mobile/login',
      method: 'POST',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 获取用户信息
 */
export function getUserInfo() {
  return http.request<UserType>({
    url: '/mobile/getUserInfo',
    method: 'get',
  });
}

/**
 * @description: 用户登出
 */
export function doLogout() {
  return http.request<{}>({
    url: '/logout',
    method: 'POST',
  });
}

/**
 * @description: 用户修改密码
 */
export function changePassword(params: any, uid: any) {
  return http.request(
    {
      url: `/user/u${uid}/changepw`,
      method: 'POST',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}
