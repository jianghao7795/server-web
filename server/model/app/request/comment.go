package request

import (
	"server/model/common/request"
)

type CommentSearch struct {
	ArticleId int    `json:"articleId" form:"articleId"`
	Content   string `json:"content" form:"content"`
	request.PageInfo
}
