package initialize

import (
	"github.com/kataras/iris/v12"
	"mayday/src/middleware"
	"mayday/src/model"
)

func Init(app *iris.Application) {
	app.Use(middleware.Cors)
	model.RegisterLocalTimeDecoder()
}
