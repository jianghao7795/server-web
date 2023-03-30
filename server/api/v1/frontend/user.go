package frontend

import (
	"server/global"
	"server/model/common/response"
	loginRequest "server/model/frontend/request"
	"server/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FrontendUser struct{}

func (u *FrontendUser) Login(c *gin.Context) {
	var user loginRequest.LoginForm
	_ = c.ShouldBindJSON(&user)
	if err := utils.Verify(user, utils.LoginVerifyFrontend); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userInfo, err := frontendService.Login(user)
	if err != nil {
		global.LOG.Error("登录失败!", zap.Error(err))
		response.FailWithMessage("登录失败", c)
	} else {
		response.OkWithDetailed(userInfo, "获取成功", c)
	}
}

func (u *FrontendUser) GetCurrent(c *gin.Context) {
	token := c.Request.Header.Get("authorization")
	if token == "" {
		response.FailWithMessage("请登录", c)
		return
	}
	var tokenSplit = strings.Split(token, " ")
	userInfo, err := frontendService.GetUser(tokenSplit[1])
	if err != nil {
		response.FailWithMessage("token失效，请重新登录", c)
		return
	}
	response.OkWithDetailed(userInfo, "获取成功", c)
}
