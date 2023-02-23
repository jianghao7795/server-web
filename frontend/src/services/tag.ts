import { http } from "@/utils/request";

export const getTagList = (params?: API.SearchTag) => {
  return http.get<API.Response<API.Tag>>("/getTagList", {
    method: "get",
    params: params,
  });
};
