package system

type ServiceGroup struct {
	BaseSrv BaseService
	UserSrv UserService
	RoleSrv RoleService
	MenuSrv MenuService
}
