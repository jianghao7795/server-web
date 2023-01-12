package system

import (
	v1 "server/api/v1"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

type AuthorityBtnRouter struct{}

func (s *AuthorityBtnRouter) InitAuthorityBtnRouterRouter(Router *gin.RouterGroup) {
	authorityRouter := Router.Group("authorityBtn").Use(middleware.OperationRecord())
	authorityRouterWithoutRecord := Router.Group("authorityBtn")
	authorityBtnApi := v1.ApiGroupApp.SystemApiGroup.AuthorityBtnApi
	{
		authorityRouter.POST("setAuthorityBtn", authorityBtnApi.SetAuthorityBtn)
		authorityRouter.DELETE("canRemoveAuthorityBtn/:id", authorityBtnApi.CanRemoveAuthorityBtn)
	}
	{
		authorityRouterWithoutRecord.GET("getAuthorityBtn", authorityBtnApi.GetAuthorityBtn)

	}
}
