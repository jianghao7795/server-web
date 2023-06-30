import service from "@/utils/request";

// @Tags MobileUser
// @Summary 创建MobileUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MobileUser true "创建MobileUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mobile/createMobileUser [post]
export const createMobileUser = (data) => {
  return service({
    url: "/mobile/createMobileUser",
    method: "post",
    data,
  });
};

// @Tags MobileUser
// @Summary 删除MobileUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MobileUser true "删除MobileUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mobile/deleteMobileUser [delete]
export const deleteMobileUser = (data) => {
  return service({
    url: `/mobile/deleteMobileUser/${data.ID}`,
    method: "delete",
  });
};

// @Tags MobileUser
// @Summary 删除MobileUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MobileUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mobile/deleteMobileUser [delete]
export const deleteMobileUserByIds = (data) => {
  return service({
    url: "/mobile/deleteMobileUserByIds",
    method: "delete",
    data,
  });
};

// @Tags MobileUser
// @Summary 更新MobileUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MobileUser true "更新MobileUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /mobile/updateMobileUser [put]
export const updateMobileUser = (data) => {
  return service({
    url: `/mobile/updateMobileUser/${data.ID}`,
    method: "put",
    data,
  });
};

// @Tags MobileUser
// @Summary 用id查询MobileUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.MobileUser true "用id查询MobileUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /mobile/findMobileUser [get]
export const findMobileUser = (params) => {
  return service({
    url: `/mobile/findMobileUser/${params.ID}`,
    method: "get",
  });
};

// @Tags MobileUser
// @Summary 分页获取MobileUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取MobileUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mobile/getMobileUserList [get]
export const getMobileUserList = (params) => {
  return service({
    url: "/mobile/getMobileUserList",
    method: "get",
    params,
  });
};
