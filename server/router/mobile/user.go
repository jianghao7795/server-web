package mobile

import (
	v1 "server/api/v1"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

type MobileUserRouter struct{}

func (m *MobileUserRouter) InitMobileRouter(Router *gin.RouterGroup) {
	moblieUserRouter := Router.Group("moblieUser").Use(middleware.OperationRecord())
	moblieUserRouterWithoutRecord := Router.Group("moblieUser")
	var moblieUserApi = v1.ApiGroupApp.MobileApiGroup.MobileUserApi
	{
		moblieUserRouter.POST("createMoblieUser", moblieUserApi.CreateMoblieUser)             // 新建MoblieUser
		moblieUserRouter.DELETE("deleteMoblieUser/:id", moblieUserApi.DeleteMoblieUser)       // 删除MoblieUser
		moblieUserRouter.DELETE("deleteMoblieUserByIds", moblieUserApi.DeleteMoblieUserByIds) // 批量删除MoblieUser
		moblieUserRouter.PUT("updateMoblieUser/:id", moblieUserApi.UpdateMoblieUser)          // 更新MoblieUser
	}
	{
		moblieUserRouterWithoutRecord.GET("findMoblieUser/:id", moblieUserApi.FindMoblieUser)   // 根据ID获取MoblieUser
		moblieUserRouterWithoutRecord.GET("getMoblieUserList", moblieUserApi.GetMoblieUserList) // 获取MoblieUser列表
	}
}
