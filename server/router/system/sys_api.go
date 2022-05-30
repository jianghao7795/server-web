package system

import (
	v1 "server/api/v1"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

type ApiRouter struct{}

func (s *ApiRouter) InitApiRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("api").Use(middleware.OperationRecord())
	apiRouterWithoutRecord := Router.Group("api")
	apiRouterApi := v1.ApiGroupApp.SystemApiGroup.SystemApiApi
	{
		apiRouter.POST("createApi", apiRouterApi.CreateApi)               // 创建Api
		apiRouter.DELETE("deleteApi/:id", apiRouterApi.DeleteApi)         // 删除Api
		apiRouter.GET("getApiById/:id", apiRouterApi.GetApiById)          // 获取单条Api消息
		apiRouter.PUT("updateApi/:id", apiRouterApi.UpdateApi)            // 更新api
		apiRouter.DELETE("deleteApisByIds", apiRouterApi.DeleteApisByIds) // 删除选中api
	}
	{
		apiRouterWithoutRecord.GET("getAllApis", apiRouterApi.GetAllApis) // 获取所有api
		apiRouterWithoutRecord.GET("getApiList", apiRouterApi.GetApiList) // 获取Api列表
	}
}
