package main

import (
	"mayday/src/core"
	"mayday/src/global"
)

//
//                       _oo0oo_
//                      o8888888o
//                      88" . "88
//                      (| -_- |)
//                      0\  =  /0
//                    ___/`---'\___
//                  .' \\|     |// '.
//                 / \\|||  :  |||// \
//                / _||||| -:- |||||- \
//               |   | \\\  -  /// |   |
//               | \_|  ''\---/''  |_/ |
//               \  .-\__  '-'  ___/-. /
//             ___'. .'  /--.--\  `. .'___
//          ."" '<  `.___\_<|>_/___.' >' "".
//         | | :  `- \`.;`\ _ /`;.`/ - ` : | |
//         \  \ `_.   \_ __\ /__ _/   .-` /  /
//     =====`-.____`.___ \_____/___.-`___.`-======
//                       ==---==
//
//
//     ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//
//               佛祖保佑         永无BUG
//

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /v2

func main() {

	//app := iris.New()
	//
	//initialize.Init(app)
	//route_Controller.Hub(app)
	//
	//err := app.Run(iris.Addr(global.GVA_CONFIG.System.Port))
	//if err != nil {
	//	log.Print("服务器启动失败 " + err.Error())
	//	return
	//}

	//配置文件初始化
	global.GVA_VP = core.Viper()

	//日志工具初始化
	global.GVA_LOG = core.Zap()

	//启服
	core.RunServer()
}
