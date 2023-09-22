package frontend

import (
	"errors"
	"server/global"
	"server/model/common/response"
	"server/model/frontend/request"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type FrontendArticleApi struct{}

func (s *FrontendArticleApi) GetArticleList(c *gin.Context) {
	var pageInfo request.ArticleSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if pageInfo.Page == 0 {
		pageInfo.Page = 1
	}

	if pageInfo.PageSize == 0 {
		pageInfo.PageSize = 10
	}

	if list, total, err := frontendService.FrontendArticle.GetArticleList(pageInfo, c); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		// log.Println("total is ", total)
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func (s *FrontendArticleApi) GetArticleDetail(c *gin.Context) {
	id := c.Param("id")
	aritcleId, err := strconv.Atoi(id)
	if err != nil {
		global.LOG.Error("获取Id失败!", zap.Error(err))
		response.FailWithMessage("获取Id失败", c)
		return
	}

	articleDetail, err := frontendService.FrontendArticle.GetAricleDetail(aritcleId, c)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.FailWithMessage("文章没有，请重新查询", c)
		return
	}
	if err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)

	} else {
		response.OkWithData(gin.H{"article": articleDetail}, c)
	}
}

func (s *FrontendArticleApi) GetSearchArticle(c *gin.Context) {
	var searchValue request.ArticleSearch
	searchValue.Name = c.Param("name")
	searchValue.Value = c.Param("value")
	searchValue.Sort = c.Query("sort")
	if list, err := frontendService.GetSearchArticle(searchValue); err != nil {
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
