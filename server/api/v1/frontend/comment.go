package frontend

import (
	"server/global"
	"server/model/common/response"
	"server/model/frontend"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CommentApi struct{}

var commentService = frontendService.Comment

func (s *CommentApi) GetCommentByArticleId(c *gin.Context) {
	articleId := c.Param("articleId")
	id, err := strconv.Atoi(articleId)
	if err != nil {
		global.LOG.Error("获取articleId 失败!", zap.Error(err))
		response.FailWithMessage("获取articleId 失败", c)
	}
	if articleComment, err := commentService.GetCommentByArticleId(id); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(articleComment, "获取成功", c)
	}
}

func (s *CommentApi) CreatedComment(c *gin.Context) {
	var comment frontend.Comment
	_ = c.ShouldBindJSON(&comment)
	if id, err := commentService.CreatedComment(comment); err != nil {
		global.LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithId("创建成功", id, c)
	}

}
