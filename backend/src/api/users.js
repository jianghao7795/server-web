import service from "@/utils/request";

// @Tags User
// @Summary 创建User
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.User true "创建User"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/createUser [post]
export const createUser = (data) => {
  return service({
    url: "/frontend-user/createUser",
    method: "post",
    data,
  });
};

// @Tags User
// @Summary 删除User
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.User true "删除User"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /user/deleteUser [delete]
export const deleteUser = (data) => {
  return service({
    url: `/frontend-user/deleteUser/${data.ID}`,
    method: "delete",
  });
};

// @Tags User
// @Summary 删除User
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除User"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /user/deleteUser [delete]
export const deleteUserByIds = (data) => {
  return service({
    url: "/frontend-user/deleteUserByIds",
    method: "delete",
    data,
  });
};

// @Tags User
// @Summary 更新User
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.User true "更新User"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /user/updateUser [put]
export const updateUser = (data) => {
  return service({
    url: `/frontend-user/updateUser/${data.ID}`,
    method: "put",
    data,
  });
};

// @Tags User
// @Summary 用id查询User
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.User true "用id查询User"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /user/findUser [get]
export const findUser = (params) => {
  return service({
    url: `/frontend-user/findUser/${params.ID}`,
    method: "get",
  });
};

// @Tags User
// @Summary 分页获取User列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取User列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/getUserList [get]
export const getUserList = (params) => {
  return service({
    url: "/frontend-user/getUserList",
    method: "get",
    params,
  });
};
