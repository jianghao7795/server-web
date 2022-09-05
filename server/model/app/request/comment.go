package request

import (
	comment "server/model/app"
	"server/model/common/request"
)

type CommentSearch struct {
	comment.Comment
	request.PageInfo
}
