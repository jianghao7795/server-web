package request

import (
	"server/model/common/request"
	"server/model/mobile"
)

type MoblieUserSearch struct {
	mobile.MobileUser
	request.PageInfo
}

type MobileUpdate struct {
	Field string      `json:"field" form:"field"`
	Value interface{} `json:"value" form:"value"`
}
