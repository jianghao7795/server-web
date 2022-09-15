package system

import (
	"server/global"
)

// Comment 结构体
// 如果含有time.Time 请自行import time包
type SysUserProblem struct {
	global.GVA_MODEL
	SysUserId int    `json:"sys_user_id" form:"sys_user_id" gorm:"column:sys_user_id;comment:user的ID"`
	Problem   string `json:"proglem" form:"proglem" gorm:"column:proglem;comment:问题"`
	Answer    string `json:"answer" form:"answer" gorm:"column:answer;comment:答案"`
}

// TableName Comment 表名
func (SysUserProblem) TableName() string {
	return "sys_user_problem"
}
