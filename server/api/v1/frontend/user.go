package frontend

import (
	"errors"
	"server/global"
	"server/model/common/response"
	"server/model/frontend"
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
		global.LOG.Error(err.Error(), zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithDetailed(userInfo, "登录成功", c)
	}
}

func (*FrontendUser) RegisterUser(c *gin.Context) {
	var userInfo loginRequest.RegisterUser
	_ = c.ShouldBindJSON(&userInfo)
	if userInfo.Password != userInfo.RePassword {
		global.LOG.Error("密码不一致!", zap.Error(errors.New("密码不一致")))
		response.FailWithMessage("密码不一致", c)
		return
	}
	if err := utils.Verify(userInfo, utils.RegisterVerifyFrontend); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err := frontendService.RegisterUser(userInfo)
	if err != nil {
		global.LOG.Error("注册失败!", zap.Error(err))
		response.FailWithDetailed(gin.H{}, err.Error(), c)
	} else {
		response.OkWithDetailed(gin.H{}, "注册成功", c)
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

func (u *FrontendUser) UpdatePassword(c *gin.Context) {
	var resetPassword frontend.ResetPassword
	_ = c.ShouldBindJSON(&resetPassword)

	if err := utils.Verify(resetPassword, utils.ResetPasswordVerifyFrontend); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if resetPassword.NewPassword != resetPassword.RepeatNewPassword {
		global.LOG.Error("密码不一致!", zap.Error(errors.New("密码不一致")))
		response.FailWithMessage("密码不一致", c)
		return
	}
	resetPassword.Password = utils.MD5V([]byte(resetPassword.Password))
	resetPassword.NewPassword = utils.MD5V([]byte(resetPassword.NewPassword))
	resetPassword.RepeatNewPassword = utils.MD5V([]byte(resetPassword.RepeatNewPassword))
	if err := frontendService.ResetPassword(resetPassword); err != nil {
		response.FailWithMessage("重置密码失败："+err.Error(), c)
		return
	}

	response.OkWithDetailed(nil, "重置密码成功", c)
}

func (u *FrontendUser) UpdateUserBackgroudImage(c *gin.Context) {
	var user frontend.User
	_ = c.ShouldBindJSON(&user)

	err := frontendService.UpdateUserBackgroudImage(user)
	if err != nil {
		response.FailWithMessage("更新失败："+err.Error(), c)
		return
	}
	response.OkWithDetailed(nil, "更新成功", c)
}

func (u *FrontendUser) UpdateUser(c *gin.Context) {
	var user frontend.User
	_ = c.ShouldBindJSON(&user)
	if err := utils.Verify(user, utils.UpdateUserVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := frontendService.UpdateUser(user); err != nil {
		response.FailWithDetailed(err.Error(), "更新失败", c)
		return
	}
	response.OkWithDetailed(nil, "更新成功", c)
}
