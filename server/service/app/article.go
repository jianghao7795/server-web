package app

import (
	"server/global"
	"server/model/app"
	appReq "server/model/app/request"
	"server/model/common/request"
	"strings"
)

type ArticleService struct{}

// create
func (articleSearch *ArticleService) CreateArticle(article app.Article) (err error) {
	err = global.GVA_DB.Create(&article).Error
	return
}

// delete
func (articleSearch *ArticleService) DeleteArticle(article app.Article) (err error) {
	err = global.GVA_DB.Delete(&article).Error
	return
}

// delete by ids
func (articleSearch *ArticleService) DeleteArticleByIds(ids request.IdsReq) (err error) {
	var articleIds *[]app.Article = &[]app.Article{}
	err = global.GVA_DB.Delete(articleIds, "id in ?", ids.Ids).Error
	return
}

// update
func (articleSearch *ArticleService) UpdateArticle(article app.Article) (err error) {
	var articleDetail app.Article
	err = global.GVA_DB.Where("id = ?", article.ID).First(&articleDetail).Error
	if err != nil {
		return
	}
	err = global.GVA_DB.Save(&article).Error
	return err
}

// getDetail by id
func (articleSearch *ArticleService) GetArticle(id uint) (article app.Article, err error) {
	err = global.GVA_DB.Preload("Tag").Where("id = ?", id).First(&article).Error
	return
}

// getList

func (articleSearch *ArticleService) GetArticleInfoList(info appReq.ArticleSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&app.Article{}).Preload("Tag")
	var articles []app.Article

	if info.Title != "" {
		db = db.Where("title like ?", strings.Join([]string{"%", info.Title, "%"}, ""))
	}

	if info.TagId != 0 {
		db = db.Where("tag_id = ?", info.TagId)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("id desc").Find(&articles).Error
	return articles, total, err
}
