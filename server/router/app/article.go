package app

import (
	v1 "server/api/v1"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

type ArticltRouter struct{}

// InitAppTabRouter 初始化 AppTab 路由信息
func (s *ArticltRouter) InitArticleRouter(Router *gin.RouterGroup) {
	articleRouter := Router.Group("article").Use(middleware.OperationRecord())
	articleRouterWithoutRecord := Router.Group("article")
	var articleApi = v1.ApiGroupApp.AppApiGroup.ArticleApi
	{
		articleRouter.POST("createArticle", articleApi.CreateArticle)             // 新建AppTab
		articleRouter.DELETE("deleteArticle", articleApi.DeleteArticle)           // 删除AppTab
		articleRouter.DELETE("deleteArticleByIds", articleApi.DeleteArticleByIds) // 批量删除AppTab
		articleRouter.PUT("updateArticle", articleApi.UpdateArticle)              // 更新AppTab
	}
	{
		articleRouterWithoutRecord.GET("findArticle", articleApi.FindArticle)       // 根据ID获取AppTab
		articleRouterWithoutRecord.GET("getArticleList", articleApi.GetArticleList) // 获取AppTab列表
	}
}
