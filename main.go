package main

import (
	"github.com/kataras/iris/v12"
	"log"
	"mayday/src/initialize"
	"mayday/src/initialize/parse"
	"mayday/src/routes"
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

func main() {

	app := iris.New()

	initialize.Init(app)
	route_Controller.Hub(app)

	err := app.Run(iris.Addr(parse.O.Port))
	if err != nil {
		log.Print("服务器启动失败 " + err.Error())
		return
	}
}
