import { http } from "@/../utils/request";
import type { API } from "@/type/article";

export const getArticleList = (params?: API.Article) => {
  // return request({
  //   url: "/frontend/getArticleList",
  //   method: "get",
  // });
  return http.get<API.Article>("/frontend/getArticleList", {
    method: "get",
    params: params,
  });
};
