package app

import (
	"server/api/v1"
	"server/middleware"
	"github.com/gin-gonic/gin"
)

type AppTabRouter struct {
}

// InitAppTabRouter 初始化 AppTab 路由信息
func (s *AppTabRouter) InitAppTabRouter(Router *gin.RouterGroup) {
	appTabRouter := Router.Group("appTab").Use(middleware.OperationRecord())
	appTabRouterWithoutRecord := Router.Group("appTab")
	var appTabApi = v1.ApiGroupApp.AppApiGroup.AppTabApi
	{
		appTabRouter.POST("createAppTab", appTabApi.CreateAppTab)   // 新建AppTab
		appTabRouter.DELETE("deleteAppTab", appTabApi.DeleteAppTab) // 删除AppTab
		appTabRouter.DELETE("deleteAppTabByIds", appTabApi.DeleteAppTabByIds) // 批量删除AppTab
		appTabRouter.PUT("updateAppTab", appTabApi.UpdateAppTab)    // 更新AppTab
	}
	{
		appTabRouterWithoutRecord.GET("findAppTab", appTabApi.FindAppTab)        // 根据ID获取AppTab
		appTabRouterWithoutRecord.GET("getAppTabList", appTabApi.GetAppTabList)  // 获取AppTab列表
	}
}
