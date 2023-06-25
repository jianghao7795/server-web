package system

import (
	"server/global"
	"server/model/common/request"
	"server/model/system"
)

type GithubService struct{}

func (g *GithubService) CreateApi(github []system.SysGithub) (total int, err error) {
	db := global.DB.Model(&system.SysGithub{})
	var data []system.SysGithub
	for _, item := range github {
		if item.CommitTime != "" {
			db = db.Or("commit_time = ?", item.CommitTime)
		}
	}
	db.Find(&data)
	dataInsert := []system.SysGithub{}
	var isExist = true
	for _, item := range github {
		for _, itemGithub := range data {
			if itemGithub.CommitTime == item.CommitTime {
				isExist = false
				break
			} else {
				isExist = true
			}

		}
		if isExist {
			dataInsert = append(dataInsert, item)
			isExist = true
		}
	}
	insertNumber := 0
	if len(dataInsert) != 0 {
		for _, item := range dataInsert {
			dataItem := item
			if dataItem.CommitTime != "" {
				insertNumber++
				err = db.Create(&dataItem).Error
			}
		}
	}

	return insertNumber, err
}

func (g *GithubService) GetGithubList(info request.PageInfo) (list []system.SysGithub, err error) {
	db := global.DB.Model(&system.SysGithub{})
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	err = db.Order("id desc").Limit(limit).Offset(offset).Find(&list).Error
	return
}
