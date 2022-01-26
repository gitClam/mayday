//connDB
package initialize

import (
	"fmt"
	"mayday/src/initialize/parse"
	//"go-iris/utils"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	//"github.com/xorm.io/core"
	"log"
)

var (
	masterEngine *xorm.Engine
	//slaveEngine  *xorm.Engine
	lock sync.Mutex
)

func MasterEngine() *xorm.Engine {
	if masterEngine != nil {
		return masterEngine
	}

	lock.Lock()
	defer lock.Unlock()

	if masterEngine != nil {
		return masterEngine
	}

	master := parse.DBConfig.Master

	engine, err := xorm.NewEngine(master.Dialect, GetConnURL(&master))
	if err != nil {
		log.Printf("@@@ Instance Master DB error!! %s", err)
		return nil
	}
	//settings(engine, &master)
	//engine.SetMapper(core.GonicMapper{})

	masterEngine = engine

	return masterEngine
}

// GetConnURL 获取数据库连接的url
// true：master主库
func GetConnURL(info *parse.DBConfigInfo) (url string) {
	//db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	url = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		info.User,
		info.Password,
		info.Host,
		info.Port,
		info.Database,
		info.Charset)
	//golog.Infof("@@@ DB conn==>> %s", url)
	return
}
