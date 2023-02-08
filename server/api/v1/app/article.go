package app

import (
	"server/global"
	"server/model/app"
	appReq "server/model/app/request"
	"server/model/common/request"
	"server/model/common/response"
	"server/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ArticleApi struct{}

var articleService = service.ServiceGroupApp.AppServiceGroup.ArticleService

// CreateArticle 创建Article
// @Tags Article
// @Summary 创建Article
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.Article true "创建Article"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Article/createArticle [post]
func (articleApi *ArticleApi) CreateArticle(c *gin.Context) {
	var article app.Article
	_ = c.ShouldBindJSON(&article)
	if err := articleService.CreateArticle(article); err != nil {
		global.LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteArticle 删除Article
// @Tags Article
// @Summary 删除Article
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.Article true "删除Article"
// @Success 200 {string} string "{}"
// @Router /article/deleteArticle/:id [delete]
func (*ArticleApi) DeleteArticle(c *gin.Context) {
	ids := c.Param("id")
	id, _ := strconv.Atoi(ids)
	if err := articleService.DeleteArticle(uint(id)); err != nil {
		global.LOG.Error("删除失败", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteArticleByIds 批量删除Article
// @Tags Article
// @Summary 批量删除Article
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.Article true "批量删除Article"
// @Success 200 {string} string "{}"
// @Router /article/deleteArticleByIds [delete]
func (appTabApi *ArticleApi) DeleteArticleByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := articleService.DeleteArticleByIds(IDS); err != nil {
		global.LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// CreateArticle 更新Article
// @Tags Article
// @Summary 更新Article
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.Article true "更新Article"
// @Success 200 {string} string "{"success":true, "msg":"更新成功"}"
// @Router /article/updateArticle/:id [put]
func (*ArticleApi) UpdateArticle(c *gin.Context) {
	var article app.Article
	_ = c.ShouldBindJSON(&article)
	ids := c.Param("id")
	id, _ := strconv.Atoi(ids)
	article.ID = uint(id)
	if err := articleService.UpdateArticle(article); err != nil {
		global.LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindArticle get单个Article
// @Tags Article
// @Summary get单个Article
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.Article true "get单个Article"
// @Success 200 {string} string "{"success":true, "msg":"获得成功"}"
// @Router /article/findArticle/:id [get]
func (*ArticleApi) FindArticle(c *gin.Context) {
	ids := c.Param("id")
	id, _ := strconv.Atoi(ids)
	if rearticle, err := articleService.GetArticle(uint(id)); err != nil {
		global.LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rearticle": rearticle}, c)
	}
}

// GetAppTabList 分页获取article列表
// FindArticle Get Article
// @Tags Article
// @Summary Get Article
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.Article true "Get Article"
// @Success 200 {string} string "{"success":true, "msg":"获得成功"}"
// @Router /article/getArticleList [get]
func (*ArticleApi) GetArticleList(c *gin.Context) {
	var pageInfo appReq.ArticleSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := articleService.GetArticleInfoList(pageInfo); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
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

// DeleteArticleByIds 批量更新Article
// @Tags Article
// @Summary 批量更新Article
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.Article true "批量更新Article"
// @Success 200 {string} string "{}"
// @Router /article/putArticleByIds [put]
func (*ArticleApi) PutArticleByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := articleService.PutArticleByIds(IDS); err != nil {
		global.LOG.Error("批量更新失败!", zap.Error(err))
		response.FailWithMessage("批量更新失败", c)
	} else {
		response.OkWithMessage("批量更新成功", c)
	}
}
