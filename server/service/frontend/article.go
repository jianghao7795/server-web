package frontend

import (
	"encoding/json"
	"errors"
	"server/global"
	"server/model/frontend"
	frontendReq "server/model/frontend/request"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

type FrontendArticle struct{}

func (s *FrontendArticle) GetArticleList(info frontendReq.ArticleSearch, c *gin.Context) (list []frontend.Article, err error) {
	var cacheTime = global.CONFIG.Cache.Time
	var articleStr string
	if info.IsImportant != 0 {
		articleStr, err = global.REDIS.Get(c, "article-list-home").Result()
	} else {
		articleStr, err = global.REDIS.Get(c, "article-list").Result()
	}

	if err == redis.Nil {
		limit := info.PageSize
		offset := info.PageSize * (info.Page - 1)
		db := global.DB.Model(&frontend.Article{})
		if info.IsImportant != 0 {
			db = db.Where("is_important = ?", info.IsImportant)
		}
		if info.Title != "" {
			db = db.Where("title like ?", strings.Join([]string{"%", info.Title, "%"}, ""))
		}
		err = db.Limit(limit).Offset(offset).Order("id desc").Preload("Tags").Find(&list).Error
		if err != nil {
			return list, err
		}
		strList, _ := json.Marshal(list)
		if info.IsImportant != 0 {
			// articleStr, err = global.REDIS.Get(c, "article-list-home").Result()
			err = global.REDIS.Set(c, "article-list-home", strList, time.Duration(cacheTime)*time.Second).Err()
		} else {
			err = global.REDIS.Set(c, "article-list", strList, time.Duration(cacheTime)*time.Second).Err()
		}

		if err != nil {
			return list, errors.New("redis 存储失败")
		}
	} else if err != nil {
		return list, err
	} else {
		if articleStr != "" {
			err = json.Unmarshal([]byte(articleStr), &list)
			return list, err
		}
	}
	return list, err
}

func (s *FrontendArticle) GetAricleDetail(articleId int, c *gin.Context) (articleDetail frontend.Article, err error) {
	var cacheTime = global.CONFIG.Cache.Time
	var articleDetailStr string
	articleDetailStr, err = global.REDIS.Get(c, "article"+strconv.Itoa(articleId)).Result()
	if err == redis.Nil {
		db := global.DB.Model(&frontend.Article{})
		err = db.Where("id = ?", articleId).Preload("Tags").First(&articleDetail).Error
		if err != nil {
			return articleDetail, err
		}
		articleString, _ := json.Marshal(articleDetail)
		err := global.REDIS.Set(c, "article"+strconv.Itoa(articleId), articleString, time.Duration(cacheTime)*time.Second).Err()
		if err != nil {
			global.LOG.Error("Redis 存储失败!", zap.Error(err))
		}
		return articleDetail, err
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

func (s *FrontendArticle) GetSearchArticle(info frontendReq.ArticleSearch) (list []frontend.Article, err error) {
	db := global.DB.Model(&frontend.Article{})
	// Preload("Tags", func(dbg *gorm.DB) *gorm.DB {
	// 	return dbg.Where("name = ?", info.Value)
	// })
	// 	type User struct {
	//     gorm.Model
	//     Name      string
	//     Phone     string
	//     Languages []*Language `gorm:"many2many:user_languages;"`
	// }

	// // Language 一种语言属于多个用户，使用 `user_languages` 作为连接表
	// type Language struct {
	//     gorm.Model
	//     Name  string
	//     Users []*User `gorm:"many2many:user_languages;"`
	// }

	if info.Name == "tag" {
		// 多对多关联 Association
		var id uint
		global.DB.Model(&frontend.Tag{}).Select("id").Where("name = ?", info.Value).First(&id)
		dbTag := &frontend.Tag{MODEL: global.MODEL{ID: id}}
		err = global.DB.Model(dbTag).Preload("Tags").Association("Articles").Find(&list)
	}

	if info.Name == "article" {
		err = db.Where("title like ?", strings.Join([]string{"%", info.Value, "%"}, "")).Preload("Tags").Order("id desc").Find(&list).Error
	}

	return
}
