package initialize

import (
	"github.com/kataras/iris/v12"
	"mayday/src/initialize/cors"
	"mayday/src/initialize/parse"
	"mayday/src/model"
)

func Init(app *iris.Application) {
	app.Use(cors.Cors)
	parse.AppOtherParse()
	parse.DBSettingParse()
	model.RegisterLocalTimeDecoder()
}
