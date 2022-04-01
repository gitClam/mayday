package core

import (
	"github.com/kataras/iris/v12"
	"mayday/src/global"
	"mayday/src/initialize"
)

func RunServer() {
	//初始化数据库连接
	initialize.Mysql()
	//初始化时间格式解析器
	initialize.RegisterLocalTimeDecoder()

	app := iris.New()
	//路由分配
	initialize.Routers(app)

	err := app.Run(iris.Addr(global.GVA_CONFIG.System.Port))

	if err != nil {
		global.GVA_LOG.Error("服务器启动失败 " + err.Error())
		return
	}

}
