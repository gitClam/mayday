package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/go-xorm/xorm"
	"mayday/src/config"
	"sync"
)

var (
	GVA_DB     *xorm.Engine
	GVA_REDIS  *redis.Client
	GVA_CONFIG config.Server
	//GVA_VP     *viper.Viper
	//GVA_LOG    *oplogging.Logger
	//GVA_LOG                 *zap.Logger
	//GVA_Concurrency_Control = &singleflight.Group{}

	//BlackCache local_cache.Cache
	lock sync.RWMutex
)
