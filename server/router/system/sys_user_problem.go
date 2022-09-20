package system

import (
	v1 "server/api/v1"

	"github.com/gin-gonic/gin"
)

type ProblemRouter struct{}

func (*ProblemRouter) InitProblemRouter(Router *gin.RouterGroup) {
	problemRouter := Router.Group("problem")
	var problemApi = v1.ApiGroupApp.SystemApiGroup.UserProblem
	{
		problemRouter.GET("getProblemList/:id", problemApi.GetProblemSetting)
	}
}
