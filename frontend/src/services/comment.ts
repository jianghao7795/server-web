import { http } from "@/utils/request";

export const getArticleComment = (params: { articleId: string }) => {
  // return request({
  //   url: "/frontend/getArticleList",
  //   method: "get",
  // });
  return http.get<API.Response<Comment.comment>>(`/getArticleComment/${params.articleId}`, {
    method: "get",
  });
};
