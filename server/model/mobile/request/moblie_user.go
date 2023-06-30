package request

import (
	"server/model/common/request"
	"server/model/mobile"
)

type MoblieUserSearch struct {
	mobile.MobileUser
	request.PageInfo
}
