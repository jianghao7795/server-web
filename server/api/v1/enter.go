package v1

import (
	"server/api/v1/app"
	"server/api/v1/example"
	"server/api/v1/order"
	"server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
	OrderApiGroup   order.ApiGroup
	AppApiGroup     app.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
