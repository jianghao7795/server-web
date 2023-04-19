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
    reading_quantity: number;
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
    ID: number;
    CreatedAt: string;
    UpdatedAt: string;
    name: string;
    status: 1;
    aritcles?: Article[];
  };

  export type SearchArticle = {
    page?: number;
    pageSize?: number;
    title?: string;
    is_important?: number;
    name?: "tag" | "article";
    value?: string;
    sort?: string;
  };

  export type SearchTag = {
    name?: string;
  };

  export type AboutMe = {
    aboutMe: string;
  };
}
