// 自动生成模板AppTab
package frontend

import (
	"server/global"
)

// AppTab 结构体
// 如果含有time.Time 请自行import time包
type AppTab struct {
	global.MODEL
	Name     string    `json:"name" form:"name" gorm:"column:name;comment:标签名称;size:191;"`
	Status   int       `json:"status" form:"status" gorm:"column:status;comment:状态;"`
	Articles []Article `json:"aritcles" form:"aritcles" gorm:"many2many:article_tag"`
}

// TableName AppTab 表名
func (AppTab) TableName() string {
	return "app_tab"
}
