package frontend

import (
	"server/global"
	"server/model/frontend"
)

type FrontendTag struct{}

func (s *FrontendTag) GetTagList() (list []frontend.AppTab, err error) {
	db := global.GVA_DB.Model(&frontend.AppTab{})
	var tags []frontend.AppTab
	err = db.Order("id desc").Find(&tags).Error
	return tags, err
}

func (s *FrontendTag) GetTagArticle(tagId int) (tagArticle frontend.AppTab, err error) {
	db := global.GVA_DB.Model(&frontend.AppTab{})
	err = db.Where("id = ?", tagId).Order("id desc").Preload("Articles").First(&tagArticle).Error
	return tagArticle, err
}
