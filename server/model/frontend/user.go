package frontend

import "server/global"

type User struct {
	global.MODEL
	Name         string `json:"name" form:"name" gorm:"column:name;comment:用户名;size:50;"`
	HeadImg      string `json:"head_img" form:"head_img" gorm:"column:head_img;comment:头像;"`
	Introduction string `json:"introduction" form:"introduction" gorm:"column:introduction;comment:简介;size:255;"`
	Content      string `json:"content" form:"content" gorm:"column:content;comment:用户信息;size:255;"`
	Header       string `json:"header" form:"header" gorm:"column:header;comment:头像;size:100"`
}

type ResetPassword struct {
	ID                uint   `json:"id" form:"id"`
	Password          string `json:"password" form:"password"`
	NewPassword       string `json:"new_password" form:"new_password"`
	RepeatNewPassword string `json:"repeat_new_password" form:"repeat_new_password"`
}

func (User) TableName() string {
	return "user"
}
