import service from '@/utils/request';

export const createArticle = (data) => {
  return service({
    url: '/article/createArticle',
    method: 'post',
    data,
  });
};

export const deleteArticle = (data) => {
  return service({
    url: '/article/deleteArticle',
    method: 'delete',
    data,
  });
};

export const deleteArticleByIds = (data) => {
  return service({
    url: '/article/deleteArticleByIds',
    method: 'delete',
    data,
  });
};

export const updateArticle = (data) => {
  return service({
    url: '/article/updateArticle',
    method: 'put',
    data,
  });
};

export const findArticle = (params) => {
  return service({
    url: '/article/findArticle',
    method: 'get',
    params,
  });
};

export const getArticleList = (params) => {
  return service({
    url: '/article/getArticleList',
    method: 'get',
    params,
  });
};
