package frontend

import (
	"log"
	"server/global"
	"server/model/common/response"
	"server/model/frontend"
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

func (s *CommentApi) CreatedComment(c *gin.Context) {
	var comment frontend.Comment
	_ = c.ShouldBindJSON(&comment)
	if id, err := frontendService.Comment.CreatedComment(comment); err != nil {
		global.LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithId("创建成功", id, c)
	}

}
