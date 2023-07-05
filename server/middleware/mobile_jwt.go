package middleware

import (
	"server/model/common/response"
	"server/utils"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuthMobileMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			response.FailWithMessage("token 失效", c)
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.FailWithMessage("token 不正确", c)
			c.Abort()
			return
		}

		user, err := utils.ParseToken(parts[1])
		if err != nil {

			response.FailWithMessage("token 失效， 请重新登录", c)
			c.Abort()
			return
		}
		c.Request.Header.Set("user_id", strconv.FormatUint(uint64(user.ID), 10))
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
