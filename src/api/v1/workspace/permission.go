package workspace

import (
	"github.com/kataras/iris/v12"
)

//查询当前工作空间的管理员
func SelectAdmin(ctx iris.Context) {
	//user := ctx.Values().Get("user").(userModel.SdUser)
	//global.GVA_CASBIN.add
}

//设置管理员
func SetAdmin(ctx iris.Context) {
	//user := ctx.Values().Get("user").(userModel.SdUser)
}

//删除管理员
func RemoveAdmin(ctx iris.Context) {
	//user := ctx.Values().Get("user").(userModel.SdUser)
}
