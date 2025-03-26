package initialize

import (
	"sync"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

var (
	syncedEnforcer *casbin.SyncedEnforcer
	once           sync.Once
)

// InitializeRBAC khởi tạo RBAC với Casbin
func InitializeRBAC(db *gorm.DB) (*casbin.SyncedEnforcer, error) {
	var err error

	once.Do(func() {
		// Khởi tạo adapter cho GORM
		adapter, err := gormadapter.NewAdapterByDB(db)
		if err != nil {
			return
		}

		// Định nghĩa model RBAC
		m := `
		[request_definition]
		r = sub, obj, act

		[policy_definition]
		p = sub, obj, act

		[role_definition]
		g = _, _

		[policy_effect]
		e = some(where (p.eft == allow))

		[matchers]
		m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
		`

		// Tạo model từ string
		casbinModel, err := model.NewModelFromString(m)
		if err != nil {
			return
		}

		// Khởi tạo enforcer
		syncedEnforcer, err = casbin.NewSyncedEnforcer(casbinModel, adapter)
		if err != nil {
			return
		}

		// Load policies từ DB
		err = syncedEnforcer.LoadPolicy()
		if err != nil {
			return
		}
	})

	return syncedEnforcer, err
}
