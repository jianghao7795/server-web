package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseMobile struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR401 = 401
)

func ResultMobile(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

// 返回400 错误信息
func FailWithDetailed401(data interface{}, message string, c *gin.Context) {
	Result400(ERROR401, data, message, c)
}

func FailWithMessage401(message string, c *gin.Context) {
	Result400(ERROR401, map[string]interface{}{}, message, c)
}

func Result4001(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusBadRequest, ResponseMobile{
		code,
		data,
		msg,
	})
}
