package mobile

import (
	v1 "server/api/v1"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

type MobileLoginRouter struct{}

func (m *MobileUserRouter) InitMobileLoginRouter(Router *gin.RouterGroup) {
	moblieLoginRouter := Router.Group("") //.Use(middleware.JWTAuthMobileMiddleware())
	var mobileLoginApi = v1.ApiGroupApp.MobileApiGroup.MobileLoginApi
	{
		moblieLoginRouter.POST("login", mobileLoginApi.Login)
	}
	mobileGetUserApi := Router.Group("").Use(middleware.JWTAuthMobileMiddleware())
	{
		mobileGetUserApi.GET("getUserInfo", mobileLoginApi.GetUserInfo)
		mobileGetUserApi.PUT("updateUser", mobileLoginApi.UpdateMobileUser)
		mobileGetUserApi.PUT("updatePassword", mobileLoginApi.UpdatePassword)
	}
	exaFileUploadAndDownloadApi := v1.ApiGroupApp.ExampleApiGroup.FileUploadAndDownloadApi
	{
		mobileGetUserApi.POST("uploadImage", exaFileUploadAndDownloadApi.UploadFile)
	}

}
