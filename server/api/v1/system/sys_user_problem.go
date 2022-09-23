package system

import (
	"server/global"
	"server/model/common/response"
	problemReq "server/model/system"
	"server/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type UserProblem struct{}

var userProblem = service.ServiceGroupApp.SystemServiceGroup.Problem

func (*UserProblem) GetProblemSetting(c *gin.Context) {
	var search problemReq.SysUserProblem
	SysUserId := c.Param("id")
	SysUserProblemId, err := strconv.Atoi(SysUserId)
	search.SysUserId = SysUserProblemId
	if err != nil {
		global.GVA_LOG.Error("传参错误!", zap.Error(err))
		response.FailWithMessage("传错参数,请传user id", c)
		return
	}
	list, err := userProblem.GetUserProblemSettingList(&search)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List: list,
	}, "获取成功", c)
}

type UpdateProblemSettingData struct {
	Data []problemReq.SysUserProblem `json:"data" binding:"required,dive,required"`
}

func (*UserProblem) UpdateProblemSetting(c *gin.Context) {
	var dataProblem UpdateProblemSettingData
	err := c.ShouldBindJSON(&dataProblem)
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			global.GVA_LOG.Error("传参错误!", zap.Error(err))
			response.FailWithDetailed(err.Error(), "传错参数", c)
			return
		}
		global.GVA_LOG.Error("传参错误!", zap.Error(err))

		response.FailWithDetailed(errs.Translate(global.Validate), "传错参数", c)
		return
	}
	message, err := userProblem.SetUserProblemSetting(dataProblem.Data)
	if err != nil {
		global.GVA_LOG.Error("传参错误!", zap.Error(err))
		response.FailWithDetailed(err, "传错参数", c)
		return
	}
	response.OkWithMessage(message, c)
}

func (*UserProblem) HasSetting(c *gin.Context) {
	SysUserId := c.Param("uid")
	SysUserProblemId, _ := strconv.Atoi(SysUserId)
	isSetting, err := userProblem.HasSetting(SysUserProblemId)
	if err != nil {
		global.GVA_LOG.Error("传参错误!", zap.Error(err))
		response.FailWithDetailed(err, "传错参数", c)
		return
	}
	response.OkWithDetailed(isSetting, "获取成功", c)
}

type VerifyProblemSettingData struct {
	Data problemReq.SysUserProblem `json:"data" binding:"required,dive,required"`
}

func (*UserProblem) VerifyAnswer(c *gin.Context) {
	var dataProblem VerifyProblemSettingData
	err := c.ShouldBindJSON(&dataProblem)
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			global.GVA_LOG.Error("传参错误!", zap.Error(err))
			response.FailWithDetailed(err.Error(), "传错参数", c)
			return
		}
		global.GVA_LOG.Error("传参错误!", zap.Error(err))
		response.FailWithDetailed(errs.Translate(global.Validate), "传错参数", c)
		return
	}
	ispassed, err := userProblem.VerifyAnswer(&dataProblem.Data)
	if err != nil {
		global.GVA_LOG.Error("未查到此问题!", zap.Error(err))
		response.FailWithDetailed(err, "未查到此问题", c)
		return
	}
	response.OkWithDetailed(ispassed, "已验证", c)
}
