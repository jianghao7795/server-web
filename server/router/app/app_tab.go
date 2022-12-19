package app

import (
	v1 "server/api/v1"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

type TagRouter struct{}

// InitTagRouter 初始化 Tag 路由信息
func (s *TagRouter) InitTagRouter(Router *gin.RouterGroup) {
	appTabRouter := Router.Group("appTab").Use(middleware.OperationRecord())
	appTabRouterWithoutRecord := Router.Group("appTab")
	var tagApi = v1.ApiGroupApp.AppApiGroup.TagApi
	{
		appTabRouter.POST("createTag", tagApi.CreateTag)             // 新建Tag
		appTabRouter.DELETE("deleteTag", tagApi.DeleteTag)           // 删除Tag
		appTabRouter.DELETE("deleteTagByIds", tagApi.DeleteTagByIds) // 批量删除Tag
		appTabRouter.PUT("updateTag", tagApi.UpdateTag)              // 更新Tag
	}
	{
		appTabRouterWithoutRecord.GET("findTag", tagApi.FindTag)       // 根据ID获取Tag
		appTabRouterWithoutRecord.GET("getTagList", tagApi.GetTagList) // 获取Tag列表
	}
}
