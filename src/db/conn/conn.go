//connDB
package conn

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

/*func settings(engine *xorm.Engine, info *parse.DBConfigInfo) {
	engine.ShowSQL(info.ShowSql)
	engine.SetTZLocation(utils.SysTimeLocation)
	if info.MaxIdleConns > 0 {
		engine.SetMaxIdleConns(info.MaxIdleConns)
	}
	if info.MaxOpenConns > 0 {
		engine.SetMaxOpenConns(info.MaxOpenConns)
	}

	// 性能优化的时候才考虑，加上本机的SQL缓存
	//cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	//engine.SetDefaultCacher(cacher)
}*/

//var db *sql.DB
/*func Getconn() *sql.DB{
	if(db == nil){
		var err error
		db, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/jietong?charset=utf8")
		if err != nil {
			log.Print(err)
		}
		db.SetMaxOpenConns(50)
     	db.SetMaxIdleConns(10)
	}
	return db
}*/
