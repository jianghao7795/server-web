import { http } from "@/../utils/request";
import type { API } from "@/type/article";

export const getTagList = (params?: API.SearchTag) => {
  return http.get<API.Response<API.Tag>>("/frontend/getTagList", {
    method: "get",
    params: params,
  });
};
