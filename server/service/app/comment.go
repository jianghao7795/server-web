package app

import (
	"log"
	"server/global"
	comment "server/model/app"
	commentReq "server/model/app/request"
	"server/model/common/request"
	"strings"
)

type CommentService struct {
}

// CreateComment 创建Comment记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentService *CommentService) CreateComment(comment comment.Comment) (err error) {
	err = global.GVA_DB.Create(&comment).Error
	return err
}

// DeleteComment 删除Comment记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentService *CommentService) DeleteComment(comment comment.Comment) (err error) {
	err = global.GVA_DB.Delete(&comment).Error
	return err
}

// DeleteCommentByIds 批量删除Comment记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentService *CommentService) DeleteCommentByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]comment.Comment{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateComment 更新Comment记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentService *CommentService) UpdateComment(comment comment.Comment) (err error) {
	err = global.GVA_DB.Save(&comment).Error
	return err
}

// GetComment 根据id获取Comment记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentService *CommentService) GetComment(id uint) (comment comment.Comment, err error) {
	err = global.GVA_DB.Preload("Article").Preload("SysUser").Where("id = ?", id).First(&comment).Error
	return
}

// GetCommentInfoList 分页获取Comment记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentService *CommentService) GetCommentInfoList(info commentReq.CommentSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&comment.Comment{}).Preload("Article").Preload("SysUser")
	log.Println(info.ArticleId)
	if info.ArticleId != 0 {
		db = db.Where("article_id = ", info.ArticleId)
	}
	if info.Content != "" {
		db = db.Where("content like ", strings.Join([]string{"%", info.Content, "%"}, ""))
	}
	var comments []comment.Comment
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&comments).Error
	return comments, total, err
}
