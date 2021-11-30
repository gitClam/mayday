package main

import (
	"github.com/kataras/iris/v12"
	"fmt"
	"mayday/src/inits"
	"mayday/src/routes"
	"mayday/src/inits/parse"
	//"github.com/kataras/iris/v12/v12/middleware/logger"
	//"github.com/kataras/iris/v12/v12/middleware/recover"
)
// Package classification testProject API.
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http, https
//     Host: localhost
//     BasePath: /v1
//     Version: 0.0.1
//     Contact: Haojie.zhao<haojie.zhao@changhong.com>
//
//     Consumes:
//     - application/json
//     - application/xml
//
//     Produces:
//     - application/json1
//     - application/xml
//
// swagger:meta
func main() {
	
	app := iris.New()
	app.Use(Cors)
	
	inits.Init()
    	route_Controller.Hub(app)


	/*app.Get("/hello", func(ctx iris.Context) {	
	ctx.WriteString("hello")
	})*/

	
	app.Handle("GET", "/test", func(ctx iris.Context){
	//log.Print("ok")
	err := ctx.SendFile("./data/photo/1", "1")
		if err != nil {
			fmt.Println(err)
		}
	})
	
	app.Handle("GET", "/test1", func(ctx iris.Context){
	//log.Print("ok")
	err := ctx.ServeFile("./data/photo/2.png", false)
		if err != nil {
			fmt.Println(err)
		}
	})
	
	/*app.Post("/hello", func(ctx iris.Context) {	
	ctx.WriteString("hello")
	err := ctx.SendFile("./data/photo/1", "1")
		if err != nil {
			fmt.Println(err)
		}
	})*/
	
	app.Run(iris.Addr(parse.O.Port))
}
func Cors(ctx iris.Context) {
    ctx.Header("Access-Control-Allow-Origin", "*")
    if ctx.Request().Method == "OPTIONS" {
        ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
        ctx.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
        ctx.StatusCode(204)
        return
    }
    ctx.Next()
}