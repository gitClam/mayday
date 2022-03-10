package utils

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"mayday/src/global"
	"mayday/src/model/common/resultcode"
)

var Responser *response

type response struct{}

type Response struct {
	Code int         `json:"code" example:"200"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg" example:"操作成功"`
}

//TODO 这里控制日志输出
func Result(code int, data interface{}, ctx iris.Context) {

	ctx.Values().Get("err")
	ctx.StatusCode(iris.StatusOK)
	_, err := ctx.JSON(Response{
		code,
		data,
		resultcode.MessageMap[code],
	})
	if err != nil {
		global.GVA_LOG.Error("Result err :" + err.Error())
		return
	}
	//日志输出
	ip, err := GetIP(ctx)
	if err != nil {
		ip = "UnKnow"
	}
	if ctx.Values().Get("err") == nil {
		global.GVA_LOG.Info("ip: " + ip + " " + "用户: " + fmt.Sprint(ctx.Values().Get("user")) + " " + ctx.Path() + " " + resultcode.MessageMap[code])
	} else {
		global.GVA_LOG.Warn("ip: " + ip + " " + "用户: " + fmt.Sprint(ctx.Values().Get("user")) + " " + ctx.Path() + " " + resultcode.MessageMap[code] + " " + ctx.Values().Get("err").(error).Error())
	}
}

// Ok 成功返回
func (r *response) Ok(ctx iris.Context) {
	Result(resultcode.Success, "", ctx)
}

// OkWithDetails 带数据成功返回
func (r *response) OkWithDetails(ctx iris.Context, data interface{}) {
	Result(resultcode.Success, data, ctx)
}

// Fail失败返回
func (r *response) Fail(ctx iris.Context, code int, err ...error) {
	if len(err) != 0 {
		ctx.Values().Set("err", err[0])
	}
	Result(code, "", ctx)
}

// FailWithDetails 带详细信息失败返回
func (r *response) FailWithDetails(ctx iris.Context, code int, data interface{}, err ...error) {
	if len(err) != 0 {
		ctx.Values().Set("err", err[0])
	}
	Result(code, data, ctx)
}
