package service

type ServiceGroup struct {
	GameSrv      GameService
	WhitelistSrv WhitelistService
	BaseSrv      BaseService
}

var ServiceGroupApp = new(ServiceGroup)
