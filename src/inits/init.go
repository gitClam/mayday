package inits

import (
	//"github.com/kataras/iris/v12/v12/middleware/logger"
	"mayday/src/models"
	"mayday/src/inits/parse"
	//"github.com/kataras/iris/v12/v12/middleware/recover"
)

func Init(){
	parse.AppOtherParse()
	parse.DBSettingParse()
	model.RegisterLocalTimeDecoder()
}