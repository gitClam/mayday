package core

import (
	"github.com/kataras/iris/v12"
	"log"
	"mayday/src/initialize"
	"mayday/src/initialize/parse"
	"mayday/src/router"
)

func RunServer() {
	app := iris.New()

	initialize.Init(app)
	router.Hub(app)

	err := app.Run(iris.Addr(parse.O.Port))
	if err != nil {
		log.Print("服务器启动失败 " + err.Error())
		return
	}
}
