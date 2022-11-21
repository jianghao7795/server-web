package frontend

import (
	"encoding/json"
	"errors"
	"log"
	"server/global"
	"server/model/common/response"
	"server/model/frontend"
	"server/model/frontend/request"
	"strconv"
	"time"

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

	if list, err := frontendService.FrontendArticle.GetArticleList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		// global.GVA_REDIS.Set(c, "list", list, time.Hour)
		// err = global.GVA_REDIS.Do("set", "name", "abc")
		global.GVA_REDIS.Set(c, "list", "list", 1*time.Hour)
		val, _ := global.GVA_REDIS.Get(c, "list").Result()
		log.Println(val)
		b, _ := json.Marshal(list)

		global.GVA_REDIS.Set(c, "key4", string(b), 1*time.Hour)
		d, _ := global.GVA_REDIS.Get(c, "key4").Result()
		var f []frontend.Article
		err = json.Unmarshal([]byte(d), &f)
		if err != nil {
			panic(err)
		}
		log.Println(f)

		response.OkWithDetailed(response.PageResult{
			List: list,
			// Total:    total,
			// Page:     pageInfo.Page,
			// PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func (s *FrontendArticleApi) GetAricleDetail(c *gin.Context) {
	id := c.Param("id")
	aritcleId, err := strconv.Atoi(id)
	if err != nil {
		response.FailWithMessage("获取Id失败", c)
	}

	articleDetail, err := frontendService.FrontendArticle.GetAricleDetail(aritcleId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.FailWithMessage("文章没有，请重新查询", c)
		return
	}
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)

	} else {
		response.OkWithData(gin.H{"article": articleDetail}, c)
	}
}
