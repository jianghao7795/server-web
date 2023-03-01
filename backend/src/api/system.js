import service from "@/utils/request";
// @Tags systrm
// @Summary 获取配置文件内容
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /system/getSystemConfig [post]
export const getSystemConfig = () => {
  return service({
    url: "/system/getSystemConfig",
    method: "get",
  });
};

// @Tags system
// @Summary 设置配置文件内容
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body sysModel.System true
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /system/setSystemConfig [put]
export const setSystemConfig = (data) => {
  return service({
    url: "/system/setSystemConfig",
    method: "put",
    data,
  });
};

// @Tags system
// @Summary 获取服务器运行状态
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /system/getServerInfo [get]
export const getSystemState = () => {
  return service({
    url: "/system/getServerInfo",
    method: "get",
    donNotShowLoading: true,
  });
};

// @Tags start task
// @Summary 开启任务
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"开启成功"}"
// @Router /tasking/start [get]
export const startTasking = (params) => {
  return service({
    url: "/tasking/start",
    method: "get",
    params: params,
  });
};

// @Tags start task
// @Summary 重启服务
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"重启成功"}"
// @Router /system/reloadSystem [post]
export const reloadSystem = () => {
  return service({
    url: "/system/reloadSystem",
    method: "post",
    data: {},
  });
};
