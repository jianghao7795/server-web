package mobile

import (
	v1 "server/api/v1"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

type MobileUserRouter struct{}

func (m *MobileUserRouter) InitMobileRouter(Router *gin.RouterGroup) {
	moblieUserRouter := Router.Group("mobile").Use(middleware.OperationRecord())
	moblieUserRouterWithoutRecord := Router.Group("mobile")
	var mobileUserApi = v1.ApiGroupApp.MobileApiGroup.MobileUserApi
	{
		moblieUserRouter.POST("createMobileUser", mobileUserApi.CreateMoblieUser)             // 新建MoblieUser
		moblieUserRouter.DELETE("deleteMobileUser/:id", mobileUserApi.DeleteMoblieUser)       // 删除MoblieUser
		moblieUserRouter.DELETE("deleteMobileUserByIds", mobileUserApi.DeleteMoblieUserByIds) // 批量删除MoblieUser
		moblieUserRouter.PUT("updateMobileUser/:id", mobileUserApi.UpdateMoblieUser)          // 更新MoblieUser
	}
	{
		moblieUserRouterWithoutRecord.GET("findMobileUser/:id", mobileUserApi.FindMoblieUser)   // 根据ID获取MoblieUser
		moblieUserRouterWithoutRecord.GET("getMobileUserList", mobileUserApi.GetMoblieUserList) // 获取MoblieUser列表
	}
}
