package system

import (
	v1 "server/api/v1"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

type MenuRouter struct{}

func (s *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	menuRouter := Router.Group("menu").Use(middleware.OperationRecord())
	menuRouterWithoutRecord := Router.Group("menu")
	authorityMenuApi := v1.ApiGroupApp.SystemApiGroup.AuthorityMenuApi
	{
		menuRouter.POST("addBaseMenu", authorityMenuApi.AddBaseMenu)             // 新增菜单
		menuRouter.POST("addMenuAuthority", authorityMenuApi.AddMenuAuthority)   //	增加menu和角色关联关系
		menuRouter.DELETE("deleteBaseMenu/:id", authorityMenuApi.DeleteBaseMenu) // 删除菜单
		menuRouter.PUT("updateBaseMenu", authorityMenuApi.UpdateBaseMenu)        // 更新菜单
	}
	{
		menuRouterWithoutRecord.GET("getMenu", authorityMenuApi.GetMenu)                     // 获取菜单树
		menuRouterWithoutRecord.GET("getMenuList", authorityMenuApi.GetMenuList)             // 分页获取基础menu列表
		menuRouterWithoutRecord.GET("getBaseMenuTree", authorityMenuApi.GetBaseMenuTree)     // 获取用户动态路由
		menuRouterWithoutRecord.GET("getMenuAuthority", authorityMenuApi.GetMenuAuthority)   // 获取指定角色menu
		menuRouterWithoutRecord.GET("getBaseMenuById/:id", authorityMenuApi.GetBaseMenuById) // 根据id获取菜单
	}
	return menuRouter
}
