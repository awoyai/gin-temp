package system

import (
	"sync"

	"github.com/awoyai/gin-temp/global"
	"github.com/awoyai/gin-temp/model/system"
	"github.com/casbin/casbin/v2"
	casbinModel "github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

var CasbinRepo = new(casbinRepo)

type casbinRepo struct{}

var (
	syncedCachedEnforcer *casbin.Enforcer
	once                 sync.Once
)

func (casbinRepo) Casbin() *casbin.Enforcer {
	once.Do(func() {
		gormadapter.TurnOffAutoMigrate(global.DB)
		a, err := gormadapter.NewAdapterByDBWithCustomTable(global.DB, &system.CasbinRule{}, "tb_casbin")
		if err != nil {
			panic("NewAdapter err: " + err.Error())
		}
		m, err := casbinModel.NewModelFromString(`
		[request_definition]
		r = sub, obj, act
	
		[policy_definition]
		p = sub, obj
	
		[role_definition]
		g = _, _
	
		[policy_effect]
		e = some(where (p.eft == allow))
	
		[matchers]
		m = g(r.sub, p.sub) && r.obj == p.obj|| r.sub != "root" && g(r.sub, "root")
		`)
		if err != nil {
			panic("NewModelFromString err: " + err.Error())
		}

		syncedCachedEnforcer, err = casbin.NewEnforcer(m, a)
		if err != nil {
			panic(err)
		}
		_ = syncedCachedEnforcer.LoadPolicy()
	})
	return syncedCachedEnforcer
}

func (r casbinRepo) CheckPermission(sub, obj, act string) (bool, error) {
	return r.Casbin().Enforce(sub, obj, act)
}
