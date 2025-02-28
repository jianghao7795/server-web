package request

import (
	"server/model/common/request"
	"server/model/system"
)

// api分页条件查询及排序结构体
type SearchApiParams struct {
	system.SysApi
	request.PageInfo
	OrderKey string `json:"orderKey" form:"orderKey"` // 排序
	Desc     string `json:"desc" form:"desc"`         // 排序方式:升序false(默认)|降序true
}
