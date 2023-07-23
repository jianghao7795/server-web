package mobile

import (
	"server/global"
	"server/model/common/response"
	"server/model/mobile"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RegisterMobile struct{}

func (*RegisterMobile) Register(c *gin.Context) {
	var data mobile.Register
	_ = c.ShouldBindJSON(&data)
	data.Sign = "It is SB"
	if err := mobileRegisterService.Register(data); err != nil {
		global.LOG.Error("注册失败!", zap.Error(err))
		response.FailWithMessage400("注册失败，请重试", c)
	} else {
		response.OkWithDetailed("", "注册成功", c)
	}
}
