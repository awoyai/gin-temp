package service

type ServiceGroup struct {
	GreeterSrv GreeterService
	BaseSrv    BaseService
}

var ServiceGroupApp = new(ServiceGroup)
