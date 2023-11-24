package example

import (
	v1 "server/api/v1"

	"github.com/gin-gonic/gin"
)

type ExcelRouter struct{}

func (e *ExcelRouter) InitExcelRouter(Router *gin.RouterGroup) {
	excelRouter := Router.Group("excel")
	exaExcelApi := v1.ApiGroupApp.ExampleApiGroup.ExcelApi
	{
		excelRouter.POST("importExcel", exaExcelApi.ImportExcel)          // 导入文件
		excelRouter.GET("getFileInfoList", exaExcelApi.GetFileList)       // 获取上传文件成功列表
		excelRouter.GET("loadExcel", exaExcelApi.LoadExcel)               // 加载Excel数据
		excelRouter.POST("exportExcel", exaExcelApi.ExportExcel)          // 导出Excel
		excelRouter.GET("downloadTemplate", exaExcelApi.DownloadTemplate) // 下载模板文件
	}
}
