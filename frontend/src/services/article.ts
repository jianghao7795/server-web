import { http } from "@/utils/request";

export const getArticleList = (params?: API.SearchArticle) => {
  // return request({
  //   url: "/frontend/getArticleList",
  //   method: "get",
  // });
  return http.get<API.Response<API.Article>>("/getArticleList", {
    method: "get",
    params: params,
  });
};

export const getArticleDetail = (id: string) => {
  return http.get<API.Response<API.Article | undefined>>(`/getArticle/${id}`, {
    method: "get",
  });
};

// 搜索文章的 tag 或 title
export const getArticleSearch = (params: API.SearchArticle) => {
  return http.get<API.Response<API.Article>>(`/getSearchArticle/${params.name}/${params.value}`, {
    method: "get",
  });
};
