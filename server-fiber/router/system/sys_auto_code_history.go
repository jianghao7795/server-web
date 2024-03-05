package system

import (
	v1 "server-fiber/api/v1"

	"github.com/gofiber/fiber/v2"
)

type AutoCodeHistoryRouter struct{}

func (s *AutoCodeRouter) InitAutoCodeHistoryRouter(Router *fiber.App) {
	autoCodeHistoryRouter := Router.Group("autoCode")
	autoCodeHistoryApi := v1.ApiGroupApp.SystemApiGroup.AutoCodeHistoryApi
	{
		autoCodeHistoryRouter.Post("GetMeta", autoCodeHistoryApi.First)         // 根据id获取meta信息
		autoCodeHistoryRouter.Post("rollback", autoCodeHistoryApi.RollBack)     // 回滚
		autoCodeHistoryRouter.Post("delSysHistory", autoCodeHistoryApi.Delete)  // 删除回滚记录
		autoCodeHistoryRouter.Post("GetSysHistory", autoCodeHistoryApi.GetList) // 获取回滚记录分页
	}
}
