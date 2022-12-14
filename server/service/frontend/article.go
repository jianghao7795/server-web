package frontend

import (
	"encoding/json"
	"errors"
	"log"
	"server/global"
	"server/model/frontend"
	frontendReq "server/model/frontend/request"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

type FrontendArticle struct{}

func (s *FrontendArticle) GetArticleList(info frontendReq.ArticleSearch, c *gin.Context) (list []frontend.Article, err error) {
	var cacheTime = global.GVA_CONFIG.Cache.Time
	var articleStr string
	articleStr, err = global.GVA_REDIS.Get(c, "article-list").Result()
	log.Println("redis cache time", cacheTime)
	// log.Println(articleStr)
	if err == redis.Nil {
		limit := info.PageSize
		offset := info.PageSize * (info.Page - 1)
		db := global.GVA_DB.Model(&frontend.Article{})
		err = db.Limit(limit).Offset(offset).Order("id desc").Preload("Tags").Preload("User").Find(&list).Error
		if err != nil {
			return list, err
		}
		strList, _ := json.Marshal(list)
		err := global.GVA_REDIS.Set(c, "article-list", strList, time.Duration(cacheTime)*time.Second).Err()
		if err != nil {
			return list, errors.New("redis 存储失败")
		}
	} else if err != nil {
		return list, err
	} else {
		// log.Println(d)
		if articleStr != "" {
			err = json.Unmarshal([]byte(articleStr), &list)
			return list, err
		}
	}
	return list, err
	// limit := info.PageSize
	// offset := info.PageSize * (info.Page - 1)
	// db := global.GVA_DB.Model(&frontend.Article{})
	// err = db.Limit(limit).Offset(offset).Order("id desc").Preload("Tags").Preload("User").Find(&list).Error
	// return list, err
}

func (s *FrontendArticle) GetAricleDetail(articleId int, c *gin.Context) (articleDetail frontend.Article, err error) {
	var cacheTime = global.GVA_CONFIG.Cache.Time
	var articleDetailStr string
	articleDetailStr, err = global.GVA_REDIS.Get(c, "article"+strconv.Itoa(articleId)).Result()
	if err == redis.Nil {
		db := global.GVA_DB.Model(&frontend.Article{})
		err = db.Where("id = ?", articleId).Preload("Tags").Preload("User").First(&articleDetail).Error
		if err != nil {
			return articleDetail, err
		}
		articleString, _ := json.Marshal(articleDetail)
		err := global.GVA_REDIS.Set(c, "article"+strconv.Itoa(articleId), articleString, time.Duration(cacheTime)*time.Second).Err()
		if err != nil {
			global.GVA_LOG.Error("Redis 存储失败!", zap.Error(nil))
		}
		return articleDetail, nil
	} else if err != nil {
		return
	} else {
		if articleDetailStr != "" {
			err = json.Unmarshal([]byte(articleDetailStr), &articleDetail)
			return articleDetail, err
		}
	}

	return
}
