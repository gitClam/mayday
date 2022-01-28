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

// @title        JieTong API
// @version      1.0
// @description  除了用户登录和注册以及头像获取三个接口
// @description  其他的都需要用户携带TOKEN进行用户验证，否则无法访问接口
// @description
// @description  TOKEN 格式 ： KEY：Authorization VALUE： "JWT " + 登录时返回的对应TOKEN   （放在请求的header中）
// @description
// @license.name  ha 1.0
// @host          47.107.108.127:80
// @BasePath      /
// @accept        json
// @produce       json

func main() {

	//配置文件初始化
	global.GVA_VP = core.Viper()

	//日志工具初始化
	global.GVA_LOG = core.Zap()

	//启服
	core.RunServer()
}
