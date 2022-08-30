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
