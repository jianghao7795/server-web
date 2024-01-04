/*
 * @Author: jianghao
 * @Date: 2022-10-08 07:57:03
 * @LastEditors: jianghao
 * @LastEditTime: 2022-10-19 09:08:09
 */
package main

import (
	"server/core"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /

func main() {
	core.RunServer() // 运行
}
