package app

import (
	"server/global"
	"server/model/system"
)

type Article struct {
	global.GVA_MODEL
	TagId   int    `json:"tag_id" form:"tag_id" gorm:"column:tag_id;comment:标签ID;"`
	Title   string `json:"title" form:"title" gorm:"column:title;comment:文章标题;size:191;"`
	Desc    string `json:"desc" form:"desc" gorm:"column:desc;comment:文章简述;"`
	Content string `json:"content" form:"content" gorm:"column:content;comment:文章内容;"`
	State   int    `json:"state" form:"state" gorm:"column:state;comment:文章状态;"`
	UserId  int    `json:"user_id" form:"user_id" gorm:"column:user_id;comment:article 作者id;"`
	// Username string `json:"username" form:"username" gorm:"column:username;comment:文章内容;"`
	Tag  AppTab         `json:"tag" form:"tag" gorm:"foreignKey:TagId"`
	User system.SysUser `json:"sys_user" form:"sys_user" gorm:"foreignKey:user_id"`
}

// 表名
func (Article) TableName() string {
	return "article"
}
