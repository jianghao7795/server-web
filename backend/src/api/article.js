import service from "@/utils/request";

export const createArticle = (data) => {
  return service({
    url: "/article/createArticle",
    method: "post",
    data,
  });
};

export const deleteArticle = (params) => {
  return service({
    url: `/article/deleteArticle/${params.ID}`,
    method: "delete",
  });
};

export const deleteArticleByIds = (data) => {
  return service({
    url: "/article/deleteArticleByIds",
    method: "delete",
    data,
  });
};

export const updateArticle = (data) => {
  return service({
    url: `/article/updateArticle/${data.ID}`,
    method: "put",
    data,
  });
};

export const findArticle = (params) => {
  return service({
    url: `/article/findArticle/${params.ID}`,
    method: "get",
  });
};

export const getArticleList = (params) => {
  return service({
    url: "/article/getArticleList",
    method: "get",
    params,
  });
};
// 上传文件

export const uploadFile = (params) => {
  return service({
    url: "/fileUploadAndDownload/upload",
    method: "post",
    data: params.form,
    headers: {
      "Content-Type": "multipart/form-data",
    },
  });
};

//批量更新 是否首页显示
export const putArticleByIds = (data) => {
  return service({
    url: "/article/putArticleByIds",
    method: "put",
    data: data,
  });
};

// 文章阅读量
export const getReading = () => {
  return service({
    url: "/article/getArticleReading",
    method: "get",
  });
};
