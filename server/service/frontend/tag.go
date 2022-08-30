package frontend

import (
	"server/global"
	"server/model/app"
)

type FrontendTag struct{}

func (s *FrontendTag) GetTagList() (list []app.AppTab, err error) {
	db := global.GVA_DB.Model(&app.AppTab{})
	var tags []app.AppTab
	err = db.Order("id desc").Find(&tags).Error
	return tags, err
}
