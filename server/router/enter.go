package router

import (
	"server/router/example"
	"server/router/order"
	"server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Order   order.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
