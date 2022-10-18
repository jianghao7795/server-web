/*
 * @Author: jianghao
 * @Date: 2022-10-17 14:22:51
 * @LastEditors: jianghao
 * @LastEditTime: 2022-10-18 17:41:39
 */
package app

import (
	"errors"
	"server/global"
	"server/model/app"
	"server/model/common/response"
	"server/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type BaseMessageApi struct{}

var baseMessageService = service.ServiceGroupApp.AppServiceGroup.BaseMessageService

// CreateBaseMessage 创建base_message
func (baseMessageApi *BaseMessageApi) CreateBaseMessage(c *gin.Context) {
	var baseMessage app.BaseMessage
	_ = c.ShouldBindJSON(&baseMessage)
	if err := baseMessageService.CreateBaseMessage(baseMessage); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

/**
 * @description: update base message
 * @return {*}
 */

func (baseMessageApi *BaseMessageApi) UpdateBaseMessage(c *gin.Context) {
	var baseMessage app.BaseMessage
	_ = c.ShouldBindJSON(&baseMessage)
	if err := baseMessageService.UpdateBaseMessage(baseMessage); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

type baseMessageNotFound struct {
	message string
}

/**
 * @description: 查找最后一条数据
 * @param {*gin.Context} c
 * @return {*}
 */
func (BaseMessageApi *BaseMessageApi) FindBaseMessage(c *gin.Context) {
	if responseBaseMessage, err := baseMessageService.FindBaseMessage(); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// respBaseMessage := baseMessageNotFound{message: "not found"}
			var respBaseMessage baseMessageNotFound
			respBaseMessage.message = "not found"
			response.OkWithData(gin.H{"baseMessage": respBaseMessage}, c)
		} else {
			global.GVA_LOG.Error("查询失败!", zap.Error(err))
			response.FailWithMessage("查询失败", c)
		}

	} else {
		response.OkWithData(gin.H{"baseMessage": responseBaseMessage}, c)
	}
}
