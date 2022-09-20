package system

import (
	"server/global"
	"server/model/system"
	"server/model/system/request"
	"server/model/system/response"
)

type Problem struct{}

func (*Problem) GetUserProblemList(info *request.Problem) (list interface{}, err error) {
	db := global.GVA_DB.Model(&system.SysUserProblem{})
	var problemList []response.SysUserProblem
	err = db.Where("sys_user_id = ?", info.SysUserId).Find(&problemList).Error
	return problemList, err
}

func (*Problem) GetUserProblemSettingList(info *request.Problem) (list interface{}, err error) {
	db := global.GVA_DB.Model(&system.SysUserProblem{})
	var problemList []response.SysUserProblemSetting
	err = db.Where("sys_user_id = ?", info.SysUserId).Find(&problemList).Error
	return problemList, err
}

// func (*Problem) SetUserProblemSetting(problem *[]request.Problem) error {

// }
