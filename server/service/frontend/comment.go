package frontend

import (
	"server/global"
	"server/model/frontend"
)

type Comment struct{}

func (commentService *Comment) GetCommentByArticleId(articleId int) (list []frontend.Comment, err error) {
	db := global.DB.Model(&frontend.Comment{})
	db = db.Where("article_id = ?", articleId)

	if err != nil {
		return
	}

	var commentList []frontend.Comment
	err = db.Where("parent_id = ?", 0).Preload("Article").Find(&commentList).Error
	// err = db.Limit(limit).Offset(offset).Where("parent_id = ?", 0).Find(&commentList).Error
	if len(commentList) > 0 {
		for comment := range commentList {
			err = commentService.findChildrenComment(&commentList[comment])
		}
	}

	return commentList, err
}

func (commentService *Comment) findChildrenComment(comment *frontend.Comment) (err error) {
	err = global.DB.Where("parent_id = ?", comment.ID).Find(&comment.Children).Error
	if len(comment.Children) > 0 {
		for k := range comment.Children {
			err = commentService.findChildrenComment(&comment.Children[k])
		}
	}
	return err
}

func (c *Comment) CreatedComment(info frontend.Comment) (err error) {
	err = global.DB.Create(&info).Error
	return
}
