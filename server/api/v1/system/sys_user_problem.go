package system

import (
	"server/global"
	"server/model/common/response"
	problemReq "server/model/system/request"
	"server/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserProblem struct{}

var userProblem = service.ServiceGroupApp.SystemServiceGroup.Problem

func (*UserProblem) GetProblemSetting(c *gin.Context) {
	var search problemReq.Problem
	SysUserId := c.Param("id")
	SysUserProblemId, err := strconv.ParseUint(SysUserId, 10, 64)
	search.SysUserId = uint(SysUserProblemId)
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
