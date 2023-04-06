declare namespace User {
  export type UserInfo = {
    ID: number;
    name: string;
    introduction: string;
    head_img: string;
    content: string;
    // password: string;
    header: string;
  };

  export type Login = {
    name: string;
    password: string;
  };

  export type Images = {
    url: string;
    name: string;
    ID: number;
  };

  export type UpdateImages = {
    head_img: string;
    ID: number;
  };

  export type CurrentUser = {
    user: User.UserInfo;
    token: string;
    exportAt: number;
  };
}
