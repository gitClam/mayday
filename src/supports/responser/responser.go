package responser

import (
	"github.com/kataras/iris/v12"
	"log"
)

const (
	CODE string = "code"
	MSG  string = "msg"
	DATA string = "data"
)

// MakeSuccessRes 成功返回
func MakeSuccessRes(ctx iris.Context, msg string, data interface{}) {
	ctx.StatusCode(iris.StatusOK) //200
	_, err := ctx.JSON(iris.Map{
		CODE: iris.StatusOK,
		MSG:  msg,
		DATA: data,
	})
	if err != nil {
		log.Println("MakeSuccessRes err :" + err.Error())
		return
	}
}

// MakeErrorRes 错误返回
func MakeErrorRes(ctx iris.Context, code int, msg string, data interface{}) {
	ctx.StatusCode(iris.StatusOK)
	//ctx.StatusCode(code)
	_, err := ctx.JSON(iris.Map{
		CODE: code,
		MSG:  msg,
		DATA: data,
	})
	if err != nil {
		log.Println("MakeErrorRes err :" + err.Error())
		return
	}
}
