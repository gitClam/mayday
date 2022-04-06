//connDB
package initialize

import (
	"go.uber.org/zap"
	"mayday/src/global"
	"os"

	//"go-iris/utils"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	lock sync.Mutex
)

func Mysql() {
	if global.GVA_DB != nil {
		return
	}

	lock.Lock()

	defer lock.Unlock()

	if global.GVA_DB != nil {
		return
	}

	master := global.GVA_CONFIG.Mysql
	engine, err := xorm.NewEngine(master.Dialect, master.GetConnURL())
	if err != nil {
		global.GVA_LOG.Error("数据库连接创建失败", zap.Error(err))
		os.Exit(0)
		return
	}

	global.GVA_DB = engine
}
