package example

import (
	"os"
	"server/global"
	"server/model/common/request"
	"server/model/common/response"
	"server/model/example"
	"server/utils"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ExcelApi struct{}

// /excel/importExcel 接口，与upload接口作用类似，只是把文件存到resource/excel目录下，用于导入Excel时存放Excel文件(ExcelImport.xlsx)
// /excel/loadExcel接口，用于读取resource/excel目录下的文件((ExcelImport.xlsx)并加载为[]model.SysBaseMenu类型的示例数据
// /excel/exportExcel 接口，用于读取前端传来的tableData，生成Excel文件并返回
// /excel/downloadTemplate 接口，用于下载resource/excel目录下的 ExcelTemplate.xlsx 文件，作为导入的模板

// @Tags excel
// @Summary 导出Excel
// @Security ApiKeyAuth
// @accept application/json
// @Produce  application/octet-stream
// @Param data body example.ExcelInfo true "导出Excel文件信息"
// @Success 200
// @Router /excel/exportExcel [post]
func (e *ExcelApi) ExportExcel(c *gin.Context) {
	// time.Sleep(3 * time.Second)
	var excelInfo example.ExcelInfo
	_ = c.ShouldBindJSON(&excelInfo)
	if strings.Contains(excelInfo.FileName, "..") {
		response.FailWithMessage("包含非法字符", c)
		return
	}
	err := excelService.GetMenuData(&excelInfo.InfoList)
	if err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	filePath := global.CONFIG.Excel.Dir + excelInfo.FileName
	err = excelService.ParseInfoList2Excel(excelInfo.InfoList, filePath)
	if err != nil {
		global.LOG.Error("转换Excel失败!", zap.Error(err))
		response.FailWithMessage("转换Excel失败", c)
		return
	}
	c.Writer.Header().Add("success", "true")
	c.File(filePath)
}

// @Tags excel
// @Summary 导入Excel文件
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "导入Excel文件"
// @Success 200 {object} response.Response{msg=string} "导入Excel文件"
// @Router /excel/importExcel [post]
func (e *ExcelApi) ImportExcel(c *gin.Context) {
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		global.LOG.Error("接收文件失败!", zap.Error(err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	var filepath_time = time.Now().Format("2006/01/02")
	var filenameMd5 = utils.MD5V([]byte(header.Filename)) + "_" + time.Now().Format("20060102150405") + "." + strings.Split(header.Filename, ".")[len(strings.Split(header.Filename, "."))-1]
	var fileTypeName = strings.Split(header.Filename, ".")[len(strings.Split(header.Filename, "."))-1]
	importExcel := example.FielUploadImport{
		FileName:    header.Filename,
		FileNameMd5: filenameMd5,
		State:       1,
		FileSize:    header.Size,
		FilePath:    global.CONFIG.Excel.Dir + filepath_time + "/" + filenameMd5,
		FileType:    fileTypeName,
	}
	err = excelService.ImportExcel(&importExcel)
	if err != nil {
		global.LOG.Error("导入失败!", zap.Error(err))
		response.FailWithMessage("导入失败", c)
		return
	}

	mkdirErr := os.MkdirAll(global.CONFIG.Excel.Dir+filepath_time, os.ModePerm)
	if mkdirErr != nil {
		global.LOG.Error("创建目录失败：", zap.Any("err", mkdirErr.Error()))
		return
	}
	err = c.SaveUploadedFile(header, global.CONFIG.Excel.Dir+filepath_time+"/"+filenameMd5)
	if err != nil {
		global.LOG.Error("保存文件失败：", zap.Any("err", err.Error()))
		return
	}
	response.OkWithMessage("导入成功", c)
}

// @Tags excel
// @Summary 加载Excel数据
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "加载Excel数据,返回包括列表,总数,页码,每页数量"
// @Router /excel/loadExcel [get]
func (e *ExcelApi) LoadExcel(c *gin.Context) {
	menus, err := excelService.ParseExcel2InfoList()
	if err != nil {
		global.LOG.Error("加载数据失败!", zap.Error(err))
		response.FailWithMessage("加载数据失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     menus,
		Total:    int64(len(menus)),
		Page:     1,
		PageSize: 999,
	}, "加载数据成功", c)
}

// @Tags excel
// @Summary 下载模板
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param fileName query string true "模板名称"
// @Success 200
// @Router /excel/downloadTemplate [get]
func (e *ExcelApi) DownloadTemplate(c *gin.Context) {
	fileName := c.Query("fileName")
	filePath := global.CONFIG.Excel.Dir + fileName

	fi, err := os.Stat(filePath)
	if err != nil {
		global.LOG.Error("文件不存在!", zap.Error(err))
		response.FailWithMessage("文件不存在", c)
		return
	}
	if fi.IsDir() {
		global.LOG.Error("不支持下载文件夹!", zap.Error(err))
		response.FailWithMessage("不支持下载文件夹", c)
		return
	}
	c.Writer.Header().Add("success", "true") // 增加返回头 信息
	c.File(filePath)
}

func (e *ExcelApi) GetFileList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	list, total, err := excelService.GetFileList(pageInfo)
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

func (e *ExcelApi) DeleteFile(c *gin.Context) {
	id := c.Param("id")
	fileId, _ := strconv.Atoi(id)
	if err := excelService.DeleteFile(int64(fileId)); err != nil {
		global.LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
