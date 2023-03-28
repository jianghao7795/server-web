package frontend

import (
	v1 "server/api/v1"

	"github.com/gin-gonic/gin"
)

type FrontendRouter struct{}

func (s *FrontendRouter) InitFrontendRouter(Router *gin.RouterGroup) {
	frontend := Router.Group("")
	var frontendTagApi = v1.ApiGroupApp.FrontendApiGroup.FrontendTagApi
	{
		frontend.GET("getTagList", frontendTagApi.GetTagList)
		frontend.GET("getTagArticleList/:id", frontendTagApi.GetTag)
	}
	var frontendArticleApi = v1.ApiGroupApp.FrontendApiGroup.FrontendArticleApi
	{
		frontend.GET("getArticleList", frontendArticleApi.GetArticleList)
		frontend.GET("getArticle/:id", frontendArticleApi.GetArticleDetail)
		frontend.GET("/getSearchArticle/:name/:value", frontendArticleApi.GetSearchArticle)
	}
	var frontendCommentApi = v1.ApiGroupApp.FrontendApiGroup.CommentApi
	{
		frontend.GET("/getArticleComment/:articleId", frontendCommentApi.GetCommentByArticleId)
		frontend.POST("/createdComment", frontendCommentApi.CreatedComment)
	}
	var frontendUserApi = v1.ApiGroupApp.FrontendApiGroup.FrontendUser
	{
		frontend.GET("/getImages", frontendUserApi.GetImages)
	}
}
