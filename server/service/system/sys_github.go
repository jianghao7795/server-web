package system

import (
	"errors"
	"log"
	"server/global"
	"server/model/common/request"
	"server/model/system"

	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
)

type GithubService struct{}

func (g *GithubService) CreateApi(github []system.SysGithub) (err error) {
	db := global.DB.Model(&system.SysGithub{})
	var data system.SysGithub
	var ierr error
	for _, item := range github {
		items := item
		ierr = db.Where("commit_time = ?", item.CommitTime).First(&data).Error
		if errors.Is(ierr, gorm.ErrRecordNotFound) {
			iserr := db.Create(&items).Error
			log.Println("iserr -------", iserr)
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
