declare namespace User {
  export type UserInfo = {
    ID: number;
    name: string;
    introduction: string;
    head_img: string;
    content: string;
    password: string;
  };

  export type Login = {
    name: string;
    password: string;
  };

  export type Images = {
    name: string;
    url: string;
    ID: number;
  };

  export type CurrentUser = {
    user: User.UserInfo;
    token: string;
    exportAt: number;
  };
}
