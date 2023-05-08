package app

import (
	"log"
	"server/global"
	"server/model/common/response"
	"time"

	"github.com/gin-gonic/gin"
)

type TaskNameApi struct{}

// StartTasking start Tasking
// @Tags start Tasking
// @Summary start Tasking
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"开启成功"}"
// @Router /tasking/start [get]
func (*TaskNameApi) StartTasking(c *gin.Context) {
	tasking := c.Query("task")
	if tasking == "" {
		global.LOG.Error("请传入任务名!")
		response.FailWithMessage("请传入任务名", c)
		return
	}
	task, status := global.Timer.FindCron(tasking)

	log.Println(task)

	if !status {
		global.LOG.Error("开启失败!")
		response.FailWithMessage("开启失败，没有这个任务", c)
	} else {
		global.Timer.StartTask(tasking)
		time.Sleep(5 * time.Second)
		global.Timer.StopTask(tasking)
		response.OkWithMessage("开启成功", c)
	}
}
