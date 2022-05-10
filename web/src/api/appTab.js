import service from '@/utils/request'

// @Tags AppTab
// @Summary 创建AppTab
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppTab true "创建AppTab"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appTab/createAppTab [post]
export const createAppTab = (data) => {
  return service({
    url: '/appTab/createAppTab',
    method: 'post',
    data
  })
}

// @Tags AppTab
// @Summary 删除AppTab
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppTab true "删除AppTab"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appTab/deleteAppTab [delete]
export const deleteAppTab = (data) => {
  return service({
    url: '/appTab/deleteAppTab',
    method: 'delete',
    data
  })
}

// @Tags AppTab
// @Summary 删除AppTab
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppTab"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appTab/deleteAppTab [delete]
export const deleteAppTabByIds = (data) => {
  return service({
    url: '/appTab/deleteAppTabByIds',
    method: 'delete',
    data
  })
}

// @Tags AppTab
// @Summary 更新AppTab
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppTab true "更新AppTab"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appTab/updateAppTab [put]
export const updateAppTab = (data) => {
  return service({
    url: '/appTab/updateAppTab',
    method: 'put',
    data
  })
}

// @Tags AppTab
// @Summary 用id查询AppTab
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AppTab true "用id查询AppTab"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appTab/findAppTab [get]
export const findAppTab = (params) => {
  return service({
    url: '/appTab/findAppTab',
    method: 'get',
    params
  })
}

// @Tags AppTab
// @Summary 分页获取AppTab列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取AppTab列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appTab/getAppTabList [get]
export const getAppTabList = (params) => {
  return service({
    url: '/appTab/getAppTabList',
    method: 'get',
    params
  })
}
