package frontend

import (
	"errors"
	"server/global"
	"server/model/app"
	"server/model/frontend"
	frontendRequest "server/model/frontend/request"
	frontendResponse "server/model/frontend/response"
	"server/utils"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type MyClaims struct {
	ID                   uint   `json:"id"`
	Name                 string `json:"name"`
	Password             string `json:"password"`
	jwt.RegisteredClaims        // 注意!这是jwt-go的v4版本新增的，原先是jwt.StandardClaims
}

type UpdateImage struct {
	ID        int
	HeapImage string
}

var MySecret = []byte(global.CONFIG.JWT.SigningKey)

type FrontendUser struct{}

func (u *FrontendUser) Login(data frontendRequest.LoginForm) (userInter frontendResponse.LoginResponse, err error) {
	var user frontend.User
	data.Password = utils.MD5V([]byte(data.Password))
	err = global.DB.Where("username = ? and password = ?", data.Name, data.Password).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return userInter, errors.New("密码错误")
	}
	tokenString, expiresAt, err := MakeToken(data, user.ID)
	if err != nil {
		return
	}

	userInter.User = user
	userInter.Token = tokenString
	userInter.ExpiresAt = expiresAt
	return
}

func (u *FrontendUser) RegisterUser(data frontendRequest.RegisterUser) (err error) {
	user := app.User{
		Name:         data.Name,
		Introduction: data.Introduction,
		Content:      data.Content,
		Header:       data.Header,
	}
	user.Password = utils.MD5V([]byte(data.Password))
	if user.Header == "" {
		user.Header = "uploads/file/2023/04/10/b53b3a3d6ab90ce0268229151c9bde11_162839.jpeg"
	}
	var userLog frontend.User
	err = global.DB.Where("name = ?", data.Name).First(&userLog).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return global.DB.Create(&user).Error
	}
	return errors.New("账号已注册，请重新注册")
}

func (u *FrontendUser) GetUser(tokenString string) (userInter frontendResponse.LoginResponse, err error) {
	myClaims, err := ParseToken(tokenString)
	if err != nil {
		return
	}
	var userInfo frontend.User
	err = global.DB.Where("name = ? and password = ?", myClaims.Name, myClaims.Password).First(&userInfo).Error
	userInter.User = userInfo
	userInter.ExpiresAt = myClaims.RegisteredClaims.ExpiresAt.Unix()
	return
}

func (u *FrontendUser) GetUserInfo(id uint) (frontend.User, error) {
	var userInfo frontend.User
	err := global.DB.Where("id = ?", id).First(&userInfo).Error

	return userInfo, err
}

func (u *FrontendUser) UpdateUserBackgroudImage(data frontend.User) (err error) {
	err = global.DB.Model(&frontend.User{}).Where("id = ?", data.ID).Update("head_img", data.HeadImg).Error
	return
}

func (u *FrontendUser) UpdateUser(data frontend.User) (err error) {
	err = global.DB.Save(&data).Error
	return
}

func (u *FrontendUser) ResetPassword(data frontend.ResetPassword) (err error) {
	var userInfo frontend.User
	err = global.DB.Where("id = ? and password = ?", data.ID, data.Password).First(&userInfo).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return global.DB.Model(&frontend.User{}).Where("id = ?", data.ID).Update("password", data.NewPassword).Error
}

func MakeToken(data frontendRequest.LoginForm, id uint) (tokenString string, expiresAt int64, err error) {
	claim := MyClaims{
		ID:       id,
		Name:     data.Name,
		Password: data.Password,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour * time.Duration(1))), // 过期时间3小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                        // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                                        // 生效时间
		}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim) // 使用HS256算法
	tokenString, err = token.SignedString(MySecret)
	return tokenString, claim.ExpiresAt.Unix(), err
}

func Secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil // 这是我的secret
	}
}

func ParseToken(tokenss string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenss, &MyClaims{}, Secret())
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("that's not even a token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token is expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token not active yet")
			} else {
				return nil, errors.New("couldn't handle this token")
			}
		}
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("couldn't handle this token")
}
