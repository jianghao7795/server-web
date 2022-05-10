package request

import (
	"server/model/app"
	"server/model/common/request"
)

type AppTabSearch struct{
    app.AppTab
    request.PageInfo
}
