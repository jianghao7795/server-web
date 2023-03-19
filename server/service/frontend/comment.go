package frontend

import (
	"server/global"
	"server/model/frontend"
)

type Comment struct{}

func (c *Comment) GetCommentByArticleId(articleId int) (comment frontend.Comment, err error) {
	db := global.DB.Model(&frontend.Comment{})
	err = db.Where("article_id = ?", articleId).First(&comment).Error
	if err != nil {
		return
	}
	dbChirdren := global.DB.Model(&frontend.Comment{})
	err = dbChirdren.Where("parent_id = ?", comment.ID).Find(&comment.Children).Error
	if err != nil {
		return
	}
	return
}
