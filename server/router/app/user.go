package app

import (
	v1 "server/api/v1"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

// InitUserRouter 初始化 User 路由信息
func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("frontend-user").Use(middleware.OperationRecord())
	userRouterWithoutRecord := Router.Group("frontend-user")
	var userApi = v1.ApiGroupApp.AppApiGroup.UserApi
	{
		userRouter.POST("createUser", userApi.CreateUser)             // 新建User
		userRouter.DELETE("deleteUser/:id", userApi.DeleteUser)       // 删除User
		userRouter.DELETE("deleteUserByIds", userApi.DeleteUserByIds) // 批量删除User
		userRouter.PUT("updateUser/:id", userApi.UpdateUser)          // 更新User
	}
	{
		userRouterWithoutRecord.GET("findUser/:id", userApi.FindUser)   // 根据ID获取User
		userRouterWithoutRecord.GET("getUserList", userApi.GetUserList) // 获取User列表
	}
}
