package app

type ArticleTag struct {
	ArticleId uint `json:"article_id" form:"article_id" gorm:"column:article_id;comment:文章id;size:10;"`
	TagId     uint `json:"tag_id" form:"tag_id" gorm:"column:tag_id;comment:标签id;size:10;"`
}

func (ArticleTag) TableName() string {
	return "article_tag"
}
