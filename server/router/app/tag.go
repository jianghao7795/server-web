package app

import (
	v1 "server/api/v1"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

type TagRouter struct{}

// InitTagRouter 初始化 Tag 路由信息
func (s *TagRouter) InitTagRouter(Router *gin.RouterGroup) {
	tagRouter := Router.Group("tag").Use(middleware.OperationRecord())
	tagRouterWithoutRecord := Router.Group("tag")
	var tagApi = v1.ApiGroupApp.AppApiGroup.TagApi
	{
		tagRouter.POST("createTag", tagApi.CreateTag)             // 新建Tag
		tagRouter.DELETE("deleteTag", tagApi.DeleteTag)           // 删除Tag
		tagRouter.DELETE("deleteTagByIds", tagApi.DeleteTagByIds) // 批量删除Tag
		tagRouter.PUT("updateTag", tagApi.UpdateTag)              // 更新Tag
	}
	{
		tagRouterWithoutRecord.GET("findTag", tagApi.FindTag)       // 根据ID获取Tag
		tagRouterWithoutRecord.GET("getTagList", tagApi.GetTagList) // 获取Tag列表
	}
}
