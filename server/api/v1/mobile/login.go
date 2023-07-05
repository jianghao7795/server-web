package mobile

import (
	"log"
	"server/global"
	"server/model/common/response"
	"server/model/mobile"
	"server/service"
	"server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MobileLoginApi struct{}

var mobileService = service.ServiceGroupApp.MobileServiceGroup.MobileLoginService

func (*MobileLoginApi) Login(c *gin.Context) {
	var l mobile.Login
	_ = c.ShouldBindJSON(&l)
	log.Println(l)
	if err := utils.Verify(l, utils.MobileLoginVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if user, err := mobileService.Login(l); err != nil {
		global.LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
		response.FailWithMessage("用户名不存在或者密码错误", c)
	} else {
		response.OkWithDetailed(user, "登录成功", c)
	}

}
