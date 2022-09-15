package request

import "server/global"

type Problem struct {
	global.GVA_MODEL
	SysUserId uint   `json:"sys_user_id"`
	Problem   string `json:"Problem"`
	Answer    string `json:"answer"`
}
