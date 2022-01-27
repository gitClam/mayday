package utils

import (
	"github.com/kataras/iris/v12"
	"mayday/src/global"
)

const (
	CODE string = "code"
	MSG  string = "msg"
	DATA string = "data"

	SUCCESS int = 200
	ERROR   int = 777
)

func Result(code int, data interface{}, msg string, ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	_, err := ctx.JSON(iris.Map{
		CODE: code,
		MSG:  msg,
		DATA: data,
	})
	if err != nil {
		global.GVA_LOG.Error("Result err :" + err.Error())
		return
	}
}

// Ok 成功返回
func Ok(ctx iris.Context) {
	Result(SUCCESS, "", OptionSuccess, ctx)
}

// OkWithDetails 带详细信息成功返回
func OkWithDetails(ctx iris.Context, msg string, data interface{}) {
	Result(SUCCESS, data, msg, ctx)
}

// OkWithMassage 带消息成功返回
func OkWithMassage(ctx iris.Context, msg string) {
	Result(SUCCESS, "", msg, ctx)
}

// OkWithData 带数据成功返回
func OkWithData(ctx iris.Context, data interface{}) {
	Result(SUCCESS, data, Success, ctx)
}

// FailWithDetails 带详细信息失败返回
func FailWithDetails(ctx iris.Context, msg string, data interface{}) {
	Result(ERROR, data, msg, ctx)
}

// FailWithMsg 带信息失败返回
func FailWithMsg(ctx iris.Context, msg string) {
	Result(ERROR, "", msg, ctx)
}

// Fail 带信息失败返回
func Fail(ctx iris.Context) {
	Result(ERROR, "", OptionFailur, ctx)
}
