package frontend

import (
	"server/global"
	"server/model/common/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (u *FrontendUser) GetImages(c *gin.Context) {
	imageList, err := frontendService.GetImagesList()
	if err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(imageList, "获取成功", c)
	}
}
