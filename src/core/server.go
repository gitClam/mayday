package core

import (
	"github.com/kataras/iris/v12"
	"log"
	"mayday/src/global"
	"mayday/src/initialize"
)

func RunServer() {
	app := iris.New()

	initialize.Init(app)
	initialize.Routers(app)

	err := app.Run(iris.Addr(global.GVA_CONFIG.System.Port))

	if err != nil {
		log.Print("服务器启动失败 " + err.Error())
		return
	}

}
