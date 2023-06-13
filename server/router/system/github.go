package system

import (
	v1 "server/api/v1"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

type GithubRouter struct{}

func (g *GithubRouter) InitGithubRouter(Router *gin.RouterGroup) {
	githubRouter := Router.Group("github").Use(middleware.OperationRecord())
	githubRouterRecord := Router.Group("github")

	githubRouterApi := v1.ApiGroupApp.SystemApiGroup.SystemGithubApi
	{
		githubRouter.POST("createGithub", githubRouterApi.CreateGithub) // 创建github
	}
	{
		githubRouterRecord.GET("getGithubList", githubRouterApi.GetGithubList)
	}
}
