package middleware

import (
	"server/model/common/response"
	"server/service"
	"server/utils"

	"github.com/gin-gonic/gin"
)

var casbinService = service.ServiceGroupApp.SystemServiceGroup.CasbinService

// 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		waitUse, _ := utils.GetClaims(c)
		// 获取请求的PATH
		obj := c.Request.URL.Path
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		sub := waitUse.AuthorityId
		e := casbinService.Casbin()
		// 判断策略中是否存在
		success, _ := e.Enforce(sub, obj, act)
		// log.Println(obj, act, sub, e)
		// if global.GVA_CONFIG.System.Env == "develop" || success {
		if success {
			c.Next()
		} else {
			// 上传文件 由于是ajxs 必须返回400 错误 才能展示错误信息
			if obj == "/base_message/upload_file" {
				response.FailWithDetailed400(gin.H{}, "权限不足", c)
				c.Abort()
				return
			}
			response.FailWithDetailed(gin.H{}, "权限不足", c)
			c.Abort()
			return
		}
	}
}
