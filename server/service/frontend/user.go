package frontend

import (
	"errors"
	"server/global"
	"server/model/frontend"
	frontendRequest "server/model/frontend/request"
	frontendResponse "server/model/frontend/response"
	"server/utils"

	"gorm.io/gorm"
)

type User struct{}

func (u *User) Login(data frontendRequest.LoginForm) (userInter frontendResponse.LoginResponse, err error) {
	var user frontend.User
	data.Password = utils.MD5V([]byte(data.Password))
	err = global.DB.Where("name = ? and password = ?", data.Name, data.Password).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return userInter, errors.New("用户名已注册")
	}
	userInter.User = user
	return
}
