package service

import (
	"github.com/awoyai/gin-temp/service/greeter"
	"github.com/awoyai/gin-temp/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	GreeterServiceGroup greeter.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
