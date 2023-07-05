package mobile

import (
	v1 "server/api/v1"

	"github.com/gin-gonic/gin"
)

type MobileLoginRouter struct{}

func (m *MobileUserRouter) InitMobileLoginRouter(Router *gin.RouterGroup) {
	moblieLoginRouter := Router.Group("") //.Use(middleware.JWTAuthMobileMiddleware())
	var mobileLoginApi = v1.ApiGroupApp.MobileApiGroup.MobileLoginApi
	{
		moblieLoginRouter.POST("login", mobileLoginApi.Login)
	}

}
