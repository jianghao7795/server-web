package system

import (
	v1 "server/api/v1"

	"github.com/gin-gonic/gin"
)

type InitRouter struct{}

func (s *InitRouter) InitInitRouter(Router *gin.RouterGroup) {
	initRouter := Router.Group("init")
	dbApi := v1.ApiGroupApp.SystemApiGroup.DBApi
	{
		initRouter.POST("initdb", dbApi.InitDB)  // 创建Api
		initRouter.GET("checkdb", dbApi.CheckDB) // 检查是否初始化
	}
}
