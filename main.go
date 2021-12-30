package main

import (
	"github.com/kataras/iris/v12"
	"mayday/src/inits"
	"mayday/src/routes"
	"mayday/src/inits/parse"
	"log"
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
type test struct {
	ProjectId    string
	SeriesId string    
	TestOwner  string 
}
func main() {
	
	app := iris.New()
	
	inits.Init(app)
    	route_Controller.Hub(app)
    	
    	app.Post("/hello", func(ctx iris.Context) {
    	   var abc test
    	   ctx.ReadJSON(&abc)
    	   log.Print(abc)
    	   if(abc.ProjectId == "1" && abc.SeriesId == "2" && abc.TestOwner == "3"){
    	   	log.Print(abc)
    	   	ctx.JSON(iris.Map{"message": "ok"})
    	   	return
    	   }
        ctx.JSON(iris.Map{"message": "no ok"})
    })
    
	app.Run(iris.Addr(parse.O.Port))
}