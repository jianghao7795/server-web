/*
 * @Author: jianghao
 * @Date: 2022-10-08 07:57:03
 * @LastEditors: jianghao
 * @LastEditTime: 2022-10-19 09:08:09
 */
package main

import (
	"server/core"
	"server/global"
	"server/initialize"

	utilsInit "server/utils"

	"go.uber.org/zap"
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
	global.GVA_VP = core.Viper()       // 初始化Viper 配置
	global.GVA_LOG = core.Zap()        // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG) // 部署到全局
	global.GVA_DB = initialize.Gorm()  // gorm连接数据库
	initialize.Timer()
	// conn, err := global.GVA_Timer.FindCron("ClearDB")
	// log.Println(conn, err)
	initialize.DBList()
	utilsInit.TransInit("zh")
	if global.GVA_DB != nil {
		// initialize.RegisterTables(global.GVA_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}
