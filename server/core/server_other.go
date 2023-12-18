//go:build !windows
// +build !windows

package core

import (
	"time"

	"github.com/fvbock/endless" // 重启服务的包
	"github.com/gin-gonic/gin"
)

// interface server 被endless.endlessServer 实现了 ListenAndServe
func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 20 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}
