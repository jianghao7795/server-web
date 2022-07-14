package app

import (
	"server/global"
	"server/model/app"
	appReq "server/model/app/request"
	"server/model/common/request"
)

type AppTabService struct {
}

// CreateAppTab 创建AppTab记录
// Author [piexlmax](https://github.com/piexlmax)
func (appTabService *AppTabService) CreateAppTab(appTab app.AppTab) (err error) {
	err = global.GVA_DB.Create(&appTab).Error
	return err
}

// DeleteAppTab 删除AppTab记录
// Author [piexlmax](https://github.com/piexlmax)
func (appTabService *AppTabService) DeleteAppTab(appTab app.AppTab) (err error) {
	err = global.GVA_DB.Delete(&appTab).Error
	return err
}

// DeleteAppTabByIds 批量删除AppTab记录
// Author [piexlmax](https://github.com/piexlmax)
func (appTabService *AppTabService) DeleteAppTabByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]app.AppTab{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateAppTab 更新AppTab记录
// Author [piexlmax](https://github.com/piexlmax)
func (appTabService *AppTabService) UpdateAppTab(appTab app.AppTab) (err error) {
	var appTabNew app.AppTab
	// var count int64
	err = global.GVA_DB.Where("id = ?", appTab.ID).First(&appTabNew).Error
	if err != nil {
		// log.Println(err)
		return err
	}
	err = global.GVA_DB.Save(&appTab).Error
	return err
}

// GetAppTab 根据id获取AppTab记录
// Author [piexlmax](https://github.com/piexlmax)
func (appTabService *AppTabService) GetAppTab(id uint) (appTab app.AppTab, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&appTab).Error
	return
}

// GetAppTabInfoList 分页获取AppTab记录
// Author [piexlmax](https://github.com/piexlmax)
func (appTabService *AppTabService) GetAppTabInfoList(info appReq.AppTabSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&app.AppTab{})
	var appTabs []app.AppTab
	if info.Name != "" {
		db = db.Where("name = ?", info.Name)
	}
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("id desc").Find(&appTabs).Error
	return appTabs, total, err
}
