package frontend

import (
	"server/global"
	appReq "server/model/app/request"
	"server/model/common/response"
	"server/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FrontendTagApi struct{}

var frontendService = service.ServiceGroupApp.FrontendServiceGroup

func (appTabApi *FrontendTagApi) GetTagList(c *gin.Context) {
	var pageInfo appReq.TagSearch
	_ = c.ShouldBindQuery(&pageInfo)
	// log.Println(pageInfo.Name)
	if list, err := frontendService.FrontendTag.GetTagList(c); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
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

func (appTabApi *FrontendTagApi) GetTag(c *gin.Context) {
	id := c.Param("id")
	tagId, err := strconv.Atoi(id)
	if err != nil {
		response.FailWithMessage("获取Ids失败", c)
		return
	}
	// log.Println(pageInfo.Name)
	if tagArticles, err := frontendService.FrontendTag.GetTagArticle(tagId, c); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"tag": tagArticles}, c)
	}
}
