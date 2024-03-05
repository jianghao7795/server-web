package app

import (
	v1 "server-fiber/api/v1"
	"server-fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

type ArticltRouter struct{}

// InitArticleRouter 初始化 article 路由信息
func (s *ArticltRouter) InitArticleRouter(Router *fiber.App) {
	articleRouter := Router.Use(middleware.OperationRecord())

	var articleApi = v1.ApiGroupApp.AppApiGroup.ArticleApi
	{
		articleRouter.Post("createArticle", articleApi.CreateArticle)             // 新建article
		articleRouter.Delete("DeleteArticle/:id", articleApi.DeleteArticle)       // 删除article
		articleRouter.Delete("DeleteArticleByIds", articleApi.DeleteArticleByIds) // 批量删除article
		articleRouter.Put("updateArticle/:id", articleApi.UpdateArticle)          // 更新article
		articleRouter.Put("PutArticleByIds", articleApi.PutArticleByIds)          // 批量更新 是否首页显示article
	}
	{
		Router.Get("findArticle/:id", articleApi.FindArticle)         // 根据ID获取article
		Router.Get("GetArticleList", articleApi.GetArticleList)       // 获取article列表
		Router.Get("GetArticleReading", articleApi.GetArticleReading) // 获取阅读量
	}
}
