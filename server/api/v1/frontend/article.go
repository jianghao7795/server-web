package frontend

import (
	"server/global"
	"server/model/common/response"
	"server/model/frontend/request"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FrontendArticleApi struct{}

func (s *FrontendArticleApi) GetArticleList(c *gin.Context) {
	var pageInfo request.ArticleSearch
	_ = c.ShouldBindQuery(&pageInfo)

	if list, err := frontendService.FrontendArticle.GetArticleList(pageInfo); err != nil {
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