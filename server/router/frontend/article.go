package frontend

import (
	v1 "server/api/v1"

	"github.com/gin-gonic/gin"
)

type FrontendRouter struct{}

func (s *FrontendRouter) InitFrontendRouter(Router *gin.RouterGroup) {
	frontend := Router.Group("frontend")
	var frontendApi = v1.ApiGroupApp.FrontendApiGroup.FrontendApi
	{
		frontend.GET("getTagList", frontendApi.GetAppTabList)
	}
}
