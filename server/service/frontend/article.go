package frontend

import (
	"server/global"
	"server/model/frontend"
	frontendReq "server/model/frontend/request"
	"strings"
)

type FrontendArticle struct{}

func (s *FrontendArticle) GetArticleList(info frontendReq.ArticleSearch) (list []frontend.Article, err error) {
	db := global.GVA_DB.Model(&frontend.Article{}).Preload("Tag")
	if info.Title != "" {
		db = db.Where("title like ?", strings.Join([]string{"%", info.Title, "%"}, ""))
	}
	err = db.Order("id desc").Find(&list).Error
	return list, err
}
