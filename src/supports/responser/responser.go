package responser

import (
	"github.com/kataras/iris/v12"
	"log"
	"mayday/src/models"
)

// MakeSuccessRes 成功返回
func MakeSuccessRes(ctx iris.Context, msg string, data interface{}) {
	ctx.StatusCode(iris.StatusOK) //200
	_, err := ctx.JSON(iris.Map{
		model.CODE: iris.StatusOK,
		model.MSG:  msg,
		model.DATA: data,
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
		model.CODE: code,
		model.MSG:  msg,
		model.DATA: data,
	})
	if err != nil {
		log.Println("MakeErrorRes err :" + err.Error())
		return
	}
}
