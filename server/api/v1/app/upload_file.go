/*
 * @Author: jianghao
 * @Date: 2022-10-17 15:45:55
 * @LastEditors: jianghao
 * @LastEditTime: 2022-10-17 17:04:47
 */

package app

import (
	"server/global"
	"server/model/app"
	"server/model/common/response"
	"server/service"

	responseUplodFile "server/model/app/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FileUploadAndDownloadApi struct{}

var fileUploadService = service.ServiceGroupApp.AppServiceGroup.FileUploadService

// @Tags ExaFileUploadAndDownload
// @Summary 上传文件示例
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "上传文件示例"
// @Success 200 {object} response.Response{data=string,msg=string} "上传文件示例,返回包括文件详情"
// @Router /fileUploadAndDownload/upload [post]
func (u *FileUploadAndDownloadApi) UploadFile(c *gin.Context) {
	var file app.FileUploadAndDownload
	noSave := c.DefaultQuery("noSave", "0")
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		global.GVA_LOG.Error("接收文件失败!", zap.Error(err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	file, err = fileUploadService.UploadFile(header, noSave) // 文件上传后拿到文件路径
	if err != nil {
		global.GVA_LOG.Error("修改数据库链接失败!", zap.Error(err))
		response.FailWithMessage("修改数据库链接失败", c)
		return
	}
	response.OkWithDetailed(responseUplodFile.ResponseUploadFile{File: file}, "上传成功", c)
}
