package example

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"server/global"
	"server/model/common/request"
	"server/model/common/response"
	"server/model/example"
	fileDimensionReq "server/model/example/request"
	exampleRes "server/model/example/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FileUploadAndDownloadApi struct{}

// @Tags ExaFileUploadAndDownload
// @Summary 上传文件示例
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "上传文件示例"
// @Success 200 {object} response.Response{data=exampleRes.ExaFileResponse,msg=string} "上传文件示例,返回包括文件详情"
// @Router /fileUploadAndDownload/upload [post]
func (u *FileUploadAndDownloadApi) UploadFile(c *gin.Context) {
	var file example.ExaFileUploadAndDownload
	noSave := c.DefaultQuery("noSave", "0")
	isCropper, _ := strconv.Atoi(c.DefaultQuery("is_cropper", "1"))
	fileImages, header, err := c.Request.FormFile("file")
	if err != nil {
		global.LOG.Error("接收文件失败!", zap.Error(err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	if header.Header.Get("content-type") == "image/svg+xml" {
		var fileDimension fileDimensionReq.FileDimension
		fileDimension.Height = 2
		fileDimension.Width = 1
		fileDimension.Proportion = 2.00
		file, err = fileUploadAndDownloadService.UploadFile(header, noSave, fileDimension, isCropper) // 文件上传后拿到文件路径
		if err != nil {
			global.LOG.Error("修改数据库链接失败!", zap.Error(err))
			response.FailWithMessage("修改数据库链接失败", c)
			return
		}
	} else {
		ct, _, err := image.Decode(fileImages)
		if err != nil {
			global.LOG.Error("获取文件失败!", zap.Error(err))
			response.FailWithMessage("获取文件失败", c)
			return
		}
		fileCtx := ct.Bounds()
		var fileDimension fileDimensionReq.FileDimension
		fileDimension.Height = fileCtx.Dy()
		fileDimension.Width = fileCtx.Dx()
		fileDimension.Proportion = float64(fileCtx.Dx()) / float64(fileCtx.Dy())

		file, err = fileUploadAndDownloadService.UploadFile(header, noSave, fileDimension, isCropper) // 文件上传后拿到文件路径
		if err != nil {
			global.LOG.Error("修改数据库链接失败!", zap.Error(err))
			response.FailWithMessage("修改数据库链接失败", c)
			return
		}
	}

	response.OkWithDetailed(exampleRes.ExaFileResponse{File: file}, "上传成功", c)
}

// EditFileName 编辑文件名或者备注
func (u *FileUploadAndDownloadApi) EditFileName(c *gin.Context) {
	var file example.ExaFileUploadAndDownload
	_ = c.ShouldBindJSON(&file)
	if err := fileUploadAndDownloadService.EditFileName(file); err != nil {
		global.LOG.Error("编辑失败!", zap.Error(err))
		response.FailWithMessage("编辑失败", c)
		return
	}
	response.OkWithMessage("编辑成功", c)
}

// @Tags ExaFileUploadAndDownload
// @Summary 删除文件
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body example.ExaFileUploadAndDownload true "传入文件里面id即可"
// @Success 200 {object} response.Response{msg=string} "删除文件"
// @Router /fileUploadAndDownload/deleteFile [post]
func (u *FileUploadAndDownloadApi) DeleteFile(c *gin.Context) {
	var file example.ExaFileUploadAndDownload
	id := c.Param("id")
	ids, _ := strconv.Atoi(id)
	file.ID = uint(ids)
	if err := fileUploadAndDownloadService.DeleteFile(file); err != nil {
		global.LOG.Error("删除失败!", zap.Error(err))
		response.FailWithDetailed(gin.H{"msg": err.Error()}, "删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// @Tags ExaFileUploadAndDownload
// @Summary 分页文件列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "分页文件列表,返回包括列表,总数,页码,每页数量"
// @Router /fileUploadAndDownload/getFileList [post]
func (u *FileUploadAndDownloadApi) GetFileList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	list, total, err := fileUploadAndDownloadService.GetFileRecordInfoList(pageInfo)
	if err != nil {
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
