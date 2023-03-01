import service from "@/utils/request";

// @Tags Tag
// @Summary 创建Tag
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Tag true "创建Tag"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tag/createTag [post]
export const createTag = (data) => {
  return service({
    url: "/tag/createTag",
    method: "post",
    data,
  });
};

// @Tags Tag
// @Summary 删除Tag
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Tag true "删除Tag"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tag/deleteTag [delete]
export const deleteTag = (params) => {
  return service({
    url: `/tag/deleteTag/${params.ID}`,
    method: "delete",
  });
};

// @Tags Tag
// @Summary 删除Tag
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Tag"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tag/deleteTag [delete]
export const deleteTagByIds = (data) => {
  return service({
    url: "/tag/deleteTagByIds",
    method: "delete",
    data,
  });
};

// @Tags Tag
// @Summary 更新Tag
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Tag true "更新Tag"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tag/updateTag [put]
export const updateTag = (data) => {
  return service({
    url: "/tag/updateTag",
    method: "put",
    data,
  });
};

// @Tags Tag
// @Summary 用id查询Tag
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Tag true "用id查询Tag"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tag/findTag [get]
export const findTag = (params) => {
  return service({
    url: `/tag/findTag/${params.ID}`,
    method: "get",
  });
};

// @Tags Tag
// @Summary 分页获取Tag列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Tag列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tag/getTagList [get]
export const getTagList = (params) => {
  return service({
    url: "/tag/getTagList",
    method: "get",
    params,
  });
};
