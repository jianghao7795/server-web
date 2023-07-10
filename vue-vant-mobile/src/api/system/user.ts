import { http } from '@/utils/http/axios';
import { FileType } from '@/utils/http/axios/types';

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
export function login(params: { username: string; password: string }) {
  return http.request<BasicResponseModel<LoginResponse>>(
    {
      url: '/mobile/login',
      method: 'POST',
      data: params,
      // params,
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
  return http.request<BasicResponseModel<{}>>(
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

// 上传图片
export function uploadImage(formData: File, filename: string) {
  return http.uploadFile<BasicResponseModel<{ file: FileType }>>(
    {
      url: '/mobile/uploadImage?is_cropper=4',
      method: 'POST',
      baseURL: '/api',
    },
    {
      filename,
      file: formData,
    }
  );
}

// 更新个人信息
export function updateUser(data: { field: string; value: number | string }) {
  return http.request<BasicResponseModel<{ field: string; value: number | string }>>(
    {
      url: '/mobile/updateUser',
      method: 'PUT',
      data,
    },
    { isTransformResponse: false }
  );
}
// 更新密码
export function updatePassword(data: { id: number; password: string; newPassword: string }) {
  return http.request<BasicResponseModel<{ password: string }>>(
    {
      url: '/mobile/updatePassword',
      method: 'PUT',
      data,
    },
    { isTransformResponse: false }
  );
}
