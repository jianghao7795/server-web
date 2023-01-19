import type { MessageApiInjection, NotificationApiInjection } from "naive-ui";

declare namespace API {
  export type Article = {
    ID: number;
    UpdatedAt: string;
    CreatedAt: string;
    content: string;
    desc: string;
    state: 1 | 0;
    user?: User;
    tags?: Tag[];
    title: string;
    user_id: number;
  };

  export type User = {
    ID: number;
    CreatedAt: string;
    UpdatedAt: string;
    username: string;
    nick_name: string;
    header_img: string;
    phone: string;
    email: string;
  };

  export type Tag = {
    ID: 12;
    CreatedAt: string;
    UpdatedAt: string;
    name: string;
    status: 1;
    aritcles?: Article[];
  };

  export type Response<T> = {
    code: 0 | 7;
    data?: {
      [i: string]: T[] | T;
    };
    msg: string;
  };

  export type SearchArticle = {
    page?: number;
    pageSize?: number;
    title?: string;
    is_important?: number;
    name?: "tag" | "article";
    value?: string;
  };

  export type SearchTag = {
    name?: string;
  };
}
