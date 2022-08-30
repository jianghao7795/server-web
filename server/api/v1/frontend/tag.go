package frontend

import (
	"server/global"
	appReq "server/model/app/request"
	"server/model/common/response"
	"server/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FrontendTagApi struct{}

var frontendService = service.ServiceGroupApp.FrontendServiceGroup

func (appTabApi *FrontendTagApi) GetAppTabList(c *gin.Context) {
	var pageInfo appReq.AppTabSearch
	_ = c.ShouldBindQuery(&pageInfo)
	// log.Println(pageInfo.Name)
	if list, err := frontendService.FrontendTag.GetTagList(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List: list,
			// Total:    total,
			// Page:     pageInfo.Page,
			// PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
