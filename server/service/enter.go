package service

import (
	"server/service/example"
	"server/service/order"
	"server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	OrderServiceGroup   order.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
