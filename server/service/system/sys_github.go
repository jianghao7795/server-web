package system

import (
	"server/global"
	"server/model/common/request"
	"server/model/system"

	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
)

type GithubService struct{}

func (g *GithubService) CreateApi(github []system.SysGithub) (err error) {
	db := global.DB.Model(&system.SysGithub{})

	for _, item := range github {
		var data system.SysGithub
		if ierr := db.Where("message like ?", "%"+item.Message).First(&data).Error; ierr == gorm.ErrRecordNotFound {
			iserr := db.Create(&item).Error
			if iserr != nil {
				err = multierror.Append(err, iserr)
			}
		}
	}
	return

}

func (g *GithubService) GetGithubList(info request.PageInfo) (list []system.SysGithub, err error) {
	db := global.DB.Model(&system.SysGithub{})
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	err = db.Order("id desc").Limit(limit).Offset(offset).Find(&list).Error
	return
}
