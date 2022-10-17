/*
 * @Author: jianghao
 * @Date: 2022-10-17 11:13:10
 * @LastEditors: jianghao
 * @LastEditTime: 2022-10-17 11:19:46
 */

package app

import (
	v1 "server/api/v1"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

type BaseMessageRouter struct{}

// InitArticleRouter 初始化 base_message 路由信息
func (r *BaseMessageRouter) InitBaseMessageRouter(c *gin.RouterGroup) {
	baseMessageRouter := c.Group("base_message").Use(middleware.OperationRecord())
	baseMessageRouterWithoutRecord := c.Group("base_message")
	var baseMessageApi = v1.ApiGroupApp.AppApiGroup.BaseMessageApi
	{
		baseMessageRouter.POST("createBaseMessage", baseMessageApi.CreateBaseMessage)
		baseMessageRouter.PUT("updateBaseMessage", baseMessageApi.UpdateBaseMessage)
	}
	{
		baseMessageRouterWithoutRecord.GET("getBaseMessage", baseMessageApi.FindBaseMessage)
	}
}
