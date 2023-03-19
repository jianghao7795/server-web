package frontend

import (
	"log"
	"server/global"
	"server/model/common/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CommentApi struct{}

func (s *CommentApi) GetCommentByArticleId(c *gin.Context) {
	articleId := c.Param("articleId")
	id, _ := strconv.Atoi(articleId)
	log.Println("artilceId", id)
	if articleComment, err := frontendService.Comment.GetCommentByArticleId(id); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(articleComment, "获取成功", c)
	}
}
