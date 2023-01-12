package frontend

import "server/global"

type User struct {
	global.MODEL
	Username  string `json:"username" form:"username"`
	NickName  string `json:"nick_name" form:"nick_name"`
	HeaderImg string `json:"header_img" form:"header_img"`
	Phone     string `json:"phone" form:"phone"`
	Email     string `json:"email" form:"email"`
}

func (User) TableName() string {
	return "sys_users"
}
