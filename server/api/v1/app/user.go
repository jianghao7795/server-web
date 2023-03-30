package app

import (
	"errors"
	"server/global"
	"server/model/app"
	appReq "server/model/app/request"
	"server/model/common/request"
	"server/model/common/response"
	"server/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserApi struct {
}

var userService = service.ServiceGroupApp.AppServiceGroup.UserService

// CreateUser 创建User
// @Tags User
// @Summary 创建User
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.User true "创建User"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/createUser [post]
func (userApi *UserApi) CreateUser(c *gin.Context) {
	var user app.User
	_ = c.ShouldBindJSON(&user)
	if err := userService.CreateUser(user); err != nil {
		global.LOG.Error(err.Error(), zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteUser 删除User
// @Tags User
// @Summary 删除User
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.User true "删除User"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /user/deleteUser [delete]
func (userApi *UserApi) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	IDS, _ := strconv.Atoi(id)
	if err := userService.DeleteUser(IDS); err != nil {
		global.LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteUserByIds 批量删除User
// @Tags User
// @Summary 批量删除User
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除User"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /user/deleteUserByIds [delete]
func (userApi *UserApi) DeleteUserByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := userService.DeleteUserByIds(IDS); err != nil {
		global.LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateUser 更新User
// @Tags User
// @Summary 更新User
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.User true "更新User"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /user/updateUser [put]
func (userApi *UserApi) UpdateUser(c *gin.Context) {
	var user app.User
	ID := c.Param("id")
	IDS, _ := strconv.Atoi(ID)
	notFound := userService.FindIsUser(uint(IDS))
	if notFound {
		global.LOG.Error("未找到，该用户!", zap.Error(errors.New("未找到，该用户")))
		response.FailWithMessage("未找到，该用户", c)
	}
	_ = c.ShouldBindJSON(&user)
	user.ID = uint(IDS)
	if err := userService.UpdateUser(user); err != nil {
		global.LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindUser 用id查询User
// @Tags User
// @Summary 用id查询User
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query app.User true "用id查询User"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /user/findUser [get]
func (userApi *UserApi) FindUser(c *gin.Context) {
	ID := c.Param("id")
	IDS, _ := strconv.Atoi(ID)
	if reuser, err := userService.GetUser(uint(IDS)); err != nil {
		global.LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reuser": reuser}, c)
	}
}

// GetUserList 分页获取User列表
// @Tags User
// @Summary 分页获取User列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query appReq.UserSearch true "分页获取User列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/getUserList [get]
func (userApi *UserApi) GetUserList(c *gin.Context) {
	var pageInfo appReq.UserSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := userService.GetUserInfoList(pageInfo); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
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
