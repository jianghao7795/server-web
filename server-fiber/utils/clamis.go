package utils

import (
	"errors"
	"server-fiber/global"
	systemReq "server-fiber/model/system/request"
	"strings"

	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
)

func GetClaims(c *fiber.Ctx) (*systemReq.CustomClaims, error) {
	tokenString := c.Request().Header.Peek("Authorization")
	tokenValue := strings.Split(string(tokenString), " ")
	if tokenValue[0] != "Bearer" {
		return nil, errors.New("token 错误")
	}
	token := tokenValue[1]
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		global.LOG.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在token且claims是否为规定结构")
	}
	return claims, err
}

// 从Gin的Context中获取从jwt解析出来的用户ID
func GetUserID(c *fiber.Ctx) uint {
	cl, err := GetClaims(c)
	if err != nil {
		return 0
	}
	return uint(cl.BaseClaims.ID)

}

// 从Gin的Context中获取从jwt解析出来的用户UUID
func GetUserUuid(c *fiber.Ctx) uuid.UUID {
	cl, err := GetClaims(c)
	if err != nil {
		return uuid.UUID{}
	}
	return cl.UUID

}

// 从Gin的Context中获取从jwt解析出来的用户角色id
func GetUserAuthorityId(c *fiber.Ctx) string {
	cl, err := GetClaims(c)
	if err != nil {
		return ""
	}
	return cl.AuthorityId

}

// 从Gin的Context中获取从jwt解析出来的用户角色id
func GetUserInfo(c *fiber.Ctx) *systemReq.CustomClaims {
	cl, err := GetClaims(c)
	if err != nil {
		return nil
	}
	return cl

}
