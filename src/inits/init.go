package inits

import (
	"github.com/kataras/iris/v12"
	"mayday/src/models"
	"mayday/src/inits/parse"
	"mayday/src/inits/cors"
)

func Init(app *iris.Application){
	app.Use(cors.Cors)
	parse.AppOtherParse()
	parse.DBSettingParse()
	model.RegisterLocalTimeDecoder()
}