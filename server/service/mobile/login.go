package mobile

import (
	"errors"
	"server/global"
	"server/model/mobile"
	"server/model/mobile/response"
	"server/utils"

	"gorm.io/gorm"
)

type MobileLoginService struct{}

func (*MobileLoginService) Login(data mobile.Login) (m response.LoginResponse, err error) {
	var user mobile.MobileUser
	data.Password = utils.MD5V([]byte(data.Password))
	err = global.DB.Where("username = ? and password = ?", data.Username, data.Password).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return m, errors.New("密码错误")
	}

	tokenString, expiresAt, err := utils.MakeToken(data, user.ID)
	if err != nil {
		return
	}

	m.User = user
	m.Token = tokenString
	m.ExpiresAt = expiresAt
	return
}
