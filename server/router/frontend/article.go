package frontend

import (
	v1 "server/api/v1"

	"github.com/gin-gonic/gin"
)

type FrontendRouter struct{}

func (s *FrontendRouter) InitFrontendRouter(Router *gin.RouterGroup) {
	frontend := Router.Group("frontend")
	var frontendTagApi = v1.ApiGroupApp.FrontendApiGroup.FrontendTagApi
	{
		frontend.GET("getTagList", frontendTagApi.GetAppTabList)
	}
	var frontendArticleApi = v1.ApiGroupApp.FrontendApiGroup
	{
		frontend.GET("getArticleList", frontendArticleApi.GetArticleList)
	}
}
