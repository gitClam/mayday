package middleware

import (
	"github.com/casbin/casbin/v2"
	xormadapter "github.com/casbin/xorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"mayday/src/global"
)

func Casbin() (*casbin.Enforcer, error) {
	Adapter, err := xormadapter.NewAdapter("mysql", global.GVA_CONFIG.Mysql.GetConnURL())
	if err != nil {
		global.GVA_LOG.Error("casbin rbac_model or policy init error, message: %v \r\n", zap.Error(err))
		return nil, err
	}
	e, err := casbin.NewEnforcer("config/rbac_model.conf", Adapter)
	if err != nil {
		global.GVA_LOG.Error("casbin rbac_model or policy init error, message: %v \r\n", zap.Error(err))
		return nil, err
	}
	return e, err
}
