package app

import (
	v1 "server/api/v1"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

type CommentRouter struct {
}

// InitCommentRouter 初始化 Comment 路由信息
func (s *CommentRouter) InitCommentRouter(Router *gin.RouterGroup) {
	commentRouter := Router.Group("comment").Use(middleware.OperationRecord())
	commentRouterWithoutRecord := Router.Group("comment")
	var commentApi = v1.ApiGroupApp.AppApiGroup.CommentApi
	{
		commentRouter.POST("createComment", commentApi.CreateComment)             // 新建Comment
		commentRouter.DELETE("deleteComment", commentApi.DeleteComment)           // 删除Comment
		commentRouter.DELETE("deleteCommentByIds", commentApi.DeleteCommentByIds) // 批量删除Comment
		commentRouter.PUT("updateComment", commentApi.UpdateComment)              // 更新Comment
		commentRouter.PUT("pariseComment", commentApi.PutLikeItOrDislike)         //点赞
	}
	{
		commentRouterWithoutRecord.GET("findComment", commentApi.FindComment)       // 根据ID获取Comment
		commentRouterWithoutRecord.GET("getCommentList", commentApi.GetCommentList) // 获取Comment列表
	}
}
