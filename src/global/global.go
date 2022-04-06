package global

import (
	"github.com/casbin/casbin/v2"
	"github.com/go-redis/redis/v8"
	"github.com/go-xorm/xorm"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"mayday/src/config"
	"sync"
)

var (
	GVA_DB     *xorm.Engine
	GVA_REDIS  *redis.Client
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
	GVA_LOG    *zap.Logger
	GVA_CASBIN *casbin.Enforcer
	//GVA_Concurrency_Control = &singleflight.Group{}

	//BlackCache local_cache.Cache
	lock sync.RWMutex
)
