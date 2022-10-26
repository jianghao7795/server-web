package app

import (
	"errors"
	"server/global"
	comment "server/model/app"
	commentReq "server/model/app/request"
	"server/model/common/request"
	"strings"

	"gorm.io/gorm"
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
	err = global.GVA_DB.Preload("Article").Where("id = ?", id).First(&comment).Error
	return
}

// GetCommentInfoList 分页获取Comment记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentService *CommentService) GetCommentInfoList(info commentReq.CommentSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&comment.Comment{}).Preload("Article").Preload("Praise")
	if info.ArticleId != 0 {
		db = db.Where("article_id = ?", info.ArticleId)
	}
	if info.Content != "" {
		db = db.Where("content like ?", strings.Join([]string{"%", info.Content, "%"}, ""))
		// db = db.Where("MATCH(content) AGAINST('+" + info.Content + "')")
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

// LikeIt 点赞一条记录
func (*CommentService) PutLikeItOrDislike(info comment.Praise) (praise comment.Praise, err error) {
	db := global.GVA_DB.Model(&comment.Praise{})

	if info.ID == 0 {
		var praise comment.Praise
		err = db.Raw("Select id, comment_id, user_id, created_at, updated_at from praise where user_id = ? and comment_id = ? limit 1", info.UserId, info.CommentId).Scan(&praise).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				err = db.Create(&info).Error
			} else {
				return praise, err
			}
		}
		if praise.ID != 0 {
			info.ID = praise.ID
			info.CreatedAt = praise.CreatedAt
			info.UpdatedAt = praise.UpdatedAt
			err = db.Exec("UPDATE `praise` SET `created_at`=?,`updated_at`=?,`deleted_at`=NULL,`comment_id`=?,`user_id`=? where id = ? ORDER BY `praise`.`id` LIMIT 1", info.CreatedAt, info.UpdatedAt, info.CommentId, info.UserId, info.ID).Error

		}
	} else {
		err = db.Where("id = ?", info.ID).Delete(&info).Error
	}

	// if err != nil {
	// 	return err
	// }
	return info, err
}
