package app

import (
	v1 "server/api/v1"

	"github.com/gin-gonic/gin"
)

type TaskRouter struct{}

func (t *TaskRouter) InitTaskRouter(Router *gin.RouterGroup) {
	// articleRouter := Router.Group("article").Use(middleware.OperationRecord())
	taskRouterWithoutRecord := Router.Group("tasking")
	var taskApi = v1.ApiGroupApp.AppApiGroup.TaskNameApi
	{
		taskRouterWithoutRecord.GET("start", taskApi.StartTasking) // 开启任务
	}
}
