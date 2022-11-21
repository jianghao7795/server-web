package system

import (
	v1 "server/api/v1"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

type SysRouter struct{}

func (s *SysRouter) InitSystemRouter(Router *gin.RouterGroup) {
	sysRouter := Router.Group("system").Use(middleware.OperationRecord())
	sysGetRouter := Router.Group("system")
	systemApi := v1.ApiGroupApp.SystemApiGroup.SystemApi
	{
		sysRouter.PUT("setSystemConfig", systemApi.SetSystemConfig) // 设置配置文件内容
		sysRouter.POST("reloadSystem", systemApi.ReloadSystem)      // 重启服务
	}
	{
		sysRouter.GET("getSystemConfig", systemApi.GetSystemConfig) // 获取配置文件内容
		sysGetRouter.GET("getServerInfo", systemApi.GetServerInfo)  // 获取服务器信息
	}
}
