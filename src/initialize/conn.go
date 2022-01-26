//connDB
package initialize

import (
	"mayday/src/global"
	//"go-iris/utils"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	//"github.com/xorm.io/core"
	"log"
)

var (
	lock sync.Mutex
)

func XormMysql() *xorm.Engine {
	if global.GVA_DB != nil {
		return global.GVA_DB
	}

	lock.Lock()

	defer lock.Unlock()

	if global.GVA_DB != nil {
		return global.GVA_DB
	}

	master := global.GVA_CONFIG.Mysql
	engine, err := xorm.NewEngine(master.Dialect, master.GetConnURL())
	if err != nil {
		log.Printf("@@@ Instance Master DB error!! %s", err)
		return nil
	}

	global.GVA_DB = engine

	return global.GVA_DB
}
