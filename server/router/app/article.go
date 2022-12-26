package app

import (
	v1 "server/api/v1"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

type ArticltRouter struct{}

// InitArticleRouter 初始化 article 路由信息
func (s *ArticltRouter) InitArticleRouter(Router *gin.RouterGroup) {
	articleRouter := Router.Group("article").Use(middleware.OperationRecord())
	articleRouterWithoutRecord := Router.Group("article")
	var articleApi = v1.ApiGroupApp.AppApiGroup.ArticleApi
	{
		articleRouter.POST("createArticle", articleApi.CreateArticle)             // 新建article
		articleRouter.DELETE("deleteArticle/:id", articleApi.DeleteArticle)       // 删除article
		articleRouter.DELETE("deleteArticleByIds", articleApi.DeleteArticleByIds) // 批量删除article
		articleRouter.PUT("updateArticle/:id", articleApi.UpdateArticle)          // 更新article
	}
	{
		articleRouterWithoutRecord.GET("findArticle/:id", articleApi.FindArticle)   // 根据ID获取article
		articleRouterWithoutRecord.GET("getArticleList", articleApi.GetArticleList) // 获取article列表
	}
}
