package utils

import (
	"github.com/kataras/iris/v12"
	"mayday/src/global"
	userModel "mayday/src/model/user"
	"strconv"
)

var Responser *response

type response struct{}

type Response struct {
	Code int         `json:"code" example:"200"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg" example:"操作成功"`
}

const (
	SUCCESS int = 200
	ERROR   int = 777
)

//TODO 这里控制日志输出
func Result(code int, data interface{}, msg string, ctx iris.Context) {

	ctx.Values().Get("err")
	ctx.StatusCode(iris.StatusOK)
	_, err := ctx.JSON(Response{
		code,
		data,
		msg,
	})
	if err != nil {
		global.GVA_LOG.Error("Result err :" + err.Error())
		return
	}
	user := ctx.Values().Get("user")
	if user == nil {
		user = userModel.SdUser{Id: -1}
	}
	//日志输出
	if ctx.Values().Get("err") == nil {
		global.GVA_LOG.Info("用户: " + strconv.Itoa(user.(userModel.SdUser).Id) + " " + ctx.Path() + " " + msg)
	} else {
		global.GVA_LOG.Warn("用户: " + strconv.Itoa(user.(userModel.SdUser).Id) + " " + ctx.Path() + " " + msg + " " + ctx.Values().Get("err").(error).Error())
	}
}

// Ok 成功返回
func (r *response) Ok(ctx iris.Context) {
	Result(SUCCESS, "", OptionSuccess, ctx)
}

// OkWithDetails 带详细信息成功返回
func (r *response) OkWithDetails(ctx iris.Context, msg string, data interface{}) {
	Result(SUCCESS, data, msg, ctx)
}

// OkWithMassage 带消息成功返回
func (r *response) OkWithMassage(ctx iris.Context, msg string) {
	Result(SUCCESS, "", msg, ctx)
}

// OkWithData 带数据成功返回
func (r *response) OkWithData(ctx iris.Context, data interface{}) {
	Result(SUCCESS, data, Success, ctx)
}

// FailWithDetails 带详细信息失败返回
func (r *response) FailWithDetails(ctx iris.Context, msg string, data interface{}) {
	Result(ERROR, data, msg, ctx)
}

// FailWithMsg 带信息失败返回
func (r *response) FailWithMsg(ctx iris.Context, msg string) {
	Result(ERROR, "", msg, ctx)
}
