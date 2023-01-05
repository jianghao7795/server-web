/*
 * @Author: jianghao
 * @Date: 2022-10-17 15:13:04
 * @LastEditors: jianghao
 * @LastEditTime: 2022-10-17 15:20:32
 */
import service from "@/utils/request";
/**
 * @description: 获取basemessage
 * @return {*}
 */
export const getBaseMessage = (params) => {
  return service({
    url: `/base_message/getBaseMessage/${params.id}`,
    method: "get",
  });
};
/**
 * @description: 创建baseMessage
 * @param {*} data
 * @return {*}
 */
export const createBaseMessage = (data) => {
  return service({
    url: "/base_message/createBaseMessage",
    method: "post",
    data,
  });
};
export const updateBaseMessage = (data) => {
  return service({
    url: "/base_message/updateBaseMessage",
    method: "put",
    data,
  });
};
