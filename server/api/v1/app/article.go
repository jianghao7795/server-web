package app

import (
	"server/global"
	"server/model/app"
	appReq "server/model/app/request"
	"server/model/common/request"
	"server/model/common/response"
	"server/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ArticleApi struct{}

var articleService = service.ServiceGroupApp.AppServiceGroup.ArticleService

// CreateArticle 创建article
func (articleApi *ArticleApi) CreateArticle(c *gin.Context) {
	var article app.Article
	_ = c.ShouldBindJSON(&article)
	if err := articleService.CreateArticle(article); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// delete id
func (*ArticleApi) DeleteArticle(c *gin.Context) {
	var article app.Article
	_ = c.ShouldBindJSON(&article)

	if err := articleService.DeleteArticle(article); err != nil {
		global.GVA_LOG.Error("删除失败", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// delete ids
func (appTabApi *ArticleApi) DeleteArticleByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := articleService.DeleteArticleByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// update
func (*ArticleApi) UpdateArticle(c *gin.Context) {
	var article app.Article
	_ = c.ShouldBindJSON(&article)
	if err := articleService.UpdateArticle(article); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// get info
func (*ArticleApi) FindArticle(c *gin.Context) {
	var article app.Article
	_ = c.ShouldBindQuery(&article)
	if reappTab, err := articleService.GetArticle(article.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reappTab": reappTab}, c)
	}
}

// GetAppTabList 分页获取article列表
func (*ArticleApi) GetArticleList(c *gin.Context) {
	var pageInfo appReq.ArticleSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := articleService.GetArticleInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}

}
