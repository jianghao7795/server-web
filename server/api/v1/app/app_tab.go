package app

import (
	"log"
	"server/global"
	"server/model/app"
	appReq "server/model/app/request"
	"server/model/common/request"
	"server/model/common/response"
	"server/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AppTabApi struct {
}

var appTabService = service.ServiceGroupApp.AppServiceGroup.AppTabService

// CreateAppTab 创建AppTab
// @Tags AppTab
// @Summary 创建AppTab
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppTab true "创建AppTab"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appTab/createAppTab [post]
func (appTabApi *AppTabApi) CreateAppTab(c *gin.Context) {
	var appTab app.AppTab
	_ = c.ShouldBindJSON(&appTab)
	if err := appTabService.CreateAppTab(appTab); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAppTab 删除AppTab
// @Tags AppTab
// @Summary 删除AppTab
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppTab true "删除AppTab"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appTab/deleteAppTab [delete]
func (appTabApi *AppTabApi) DeleteAppTab(c *gin.Context) {
	var appTab app.AppTab
	_ = c.ShouldBindJSON(&appTab)
	if err := appTabService.DeleteAppTab(appTab); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAppTabByIds 批量删除AppTab
// @Tags AppTab
// @Summary 批量删除AppTab
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppTab"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /appTab/deleteAppTabByIds [delete]
func (appTabApi *AppTabApi) DeleteAppTabByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := appTabService.DeleteAppTabByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAppTab 更新AppTab
// @Tags AppTab
// @Summary 更新AppTab
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppTab true "更新AppTab"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appTab/updateAppTab [put]
func (appTabApi *AppTabApi) UpdateAppTab(c *gin.Context) {
	var appTab app.AppTab
	_ = c.ShouldBindJSON(&appTab)
	log.Println(appTab)
	if err := appTabService.UpdateAppTab(appTab); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAppTab 用id查询AppTab
// @Tags AppTab
// @Summary 用id查询AppTab
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query app.AppTab true "用id查询AppTab"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appTab/findAppTab [get]
func (appTabApi *AppTabApi) FindAppTab(c *gin.Context) {
	var appTab app.AppTab
	_ = c.ShouldBindQuery(&appTab)
	if err, reappTab := appTabService.GetAppTab(appTab.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reappTab": reappTab}, c)
	}
}

// GetAppTabList 分页获取AppTab列表
// @Tags AppTab
// @Summary 分页获取AppTab列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query appReq.AppTabSearch true "分页获取AppTab列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appTab/getAppTabList [get]
func (appTabApi *AppTabApi) GetAppTabList(c *gin.Context) {
	var pageInfo appReq.AppTabSearch
	_ = c.ShouldBindQuery(&pageInfo)
	// log.Println(pageInfo.Name)
	if err, list, total := appTabService.GetAppTabInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
