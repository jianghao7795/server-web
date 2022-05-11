// 自动生成模板AppTab
package app

import (
	"server/global"
)

// AppTab 结构体
// 如果含有time.Time 请自行import time包
type AppTab struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:客户名;size:191;"`
      SysUserId  *bool `json:"sysUserId" form:"sysUserId" gorm:"column:sys_user_id;comment:管理ID;"`
}


// TableName AppTab 表名
func (AppTab) TableName() string {
  return "app_tab"
}

