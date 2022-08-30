package frontend

import "server/global"

type Article struct {
	global.GVA_MODEL
	TagId      int    `json:"tag_id" form:"tag_id" gorm:"column:tag_id;comment:标签ID"`
	Title      string `json:"title" form:"title" gorm:"column:title;comment:文章标题;size:191;"`
	Desc       string `json:"desc" form:"desc" gorm:"column:desc;comment:文章简述;"`
	Content    string `json:"content" form:"content" gorm:"column:content;comment:文章内容;"`
	State      int    `json:"state" form:"state" gorm:"column:state;comment:文章状态;"`
	ModifiedBy string `json:"modified_by" form:"modified_by" gorm:"column:modified_by;comment:修改人;"`
	CreatedBy  string `json:"created_by" form:"created_by" gorm:"column:created_by;comment:创建人;"`
	Tag        AppTab `json:"tag" form:"tag" gorm:"foreignKey:TagId"`
}

// 表名
func (Article) TableName() string {
	return "article"
}
