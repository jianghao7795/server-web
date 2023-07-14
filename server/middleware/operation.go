package middleware

import (
	"bytes"
	// "encoding/json"
	ioutil "io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	json "github.com/bytedance/sonic"

	"server/global"
	"server/model/common/response"
	"server/model/system"
	"server/service"
	"server/service/frontend"
	"server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var operationRecordService = service.ServiceGroupApp.SystemServiceGroup.OperationRecordService

// 写入操作历史

func OperationRecord() gin.HandlerFunc {
	// var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return func(c *gin.Context) {
		var body []byte
		var userId int
		if c.Request.Method != http.MethodGet {
			var err error
			body, err = ioutil.ReadAll(c.Request.Body)
			if err != nil {
				global.LOG.Error("read body from request error:", zap.Error(err))
			} else {
				c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			}
		} else {
			query := c.Request.URL.RawQuery
			query, _ = url.QueryUnescape(query)
			split := strings.Split(query, "&")
			m := make(map[string]string)
			for _, v := range split {
				kv := strings.Split(v, "=")
				if len(kv) == 2 {
					m[kv[0]] = kv[1]
				}
			}
			body, _ = json.Marshal(&m)
		}
		claims, _ := utils.GetClaims(c)
		if claims.BaseClaims.ID != 0 {
			userId = int(claims.BaseClaims.ID)
		} else {
			id, err := strconv.Atoi(c.Request.Header.Get("x-user-id"))
			if err != nil {
				userId = 0
			}
			userId = id
		}
		record := system.SysOperationRecord{
			Ip:       c.ClientIP(),
			Method:   c.Request.Method,
			Path:     c.Request.URL.Path,
			Agent:    c.Request.UserAgent(),
			Body:     string(body),
			UserID:   userId,
			TypePort: system.Backend,
		}
		// 上传文件时候 中间件日志进行裁断操作
		if strings.Contains(c.GetHeader("Content-Type"), "multipart/form-data") {
			if len(record.Body) > 512 {
				record.Body = "File or Length out of limit"
			}
		}
		writer := responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer
		now := time.Now()

		c.Next()

		latency := time.Since(now)
		record.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		record.Status = c.Writer.Status()
		record.Latency = latency
		record.Resp = writer.body.String()

		if err := operationRecordService.CreateSysOperationRecord(record); err != nil {
			global.LOG.Error("create operation record error:", zap.Error(err))
		}
	}
}

func OperationRecordFrontend() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body []byte
		if c.Request.Method != http.MethodGet {
			var err error
			body, err = ioutil.ReadAll(c.Request.Body)
			if err != nil {
				global.LOG.Error("read body from request error:", zap.Error(err))
			} else {
				c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			}
		} else {
			query := c.Request.URL.RawQuery
			query, _ = url.QueryUnescape(query)
			split := strings.Split(query, "&")
			m := make(map[string]string)
			for _, v := range split {
				kv := strings.Split(v, "=")
				if len(kv) == 2 {
					m[kv[0]] = kv[1]
				}
			}
			body, _ = json.Marshal(&m)
		}
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			response.FailWithMessage("token 失效", c)
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		myClaims, _ := frontend.ParseToken(parts[1])
		record := system.SysOperationRecord{
			Ip:       c.ClientIP(),
			Method:   c.Request.Method,
			Path:     c.Request.URL.Path,
			Agent:    c.Request.UserAgent(),
			Body:     string(body),
			UserID:   int(myClaims.ID),
			TypePort: system.Frontend,
		}
		if strings.Contains(c.GetHeader("Content-Type"), "multipart/form-data") {
			if len(record.Body) > 512 {
				record.Body = "File or Length out of limit"
			}
		}
		writer := responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer
		now := time.Now()

		c.Next()

		latency := time.Since(now)
		record.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		record.Status = c.Writer.Status()
		record.Latency = latency
		record.Resp = writer.body.String()
		if err := operationRecordService.CreateSysOperationRecord(record); err != nil {
			global.LOG.Error("create operation record error:", zap.Error(err))
		}
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
