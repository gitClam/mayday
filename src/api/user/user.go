package user_api

import (
	"github.com/kataras/iris/v12"
	"mayday/src/middleware"
	userModel "mayday/src/model/user"
	userSever "mayday/src/service/user"
	"mayday/src/utils"
	"strconv"
)

// @Tags User
// @Summary 用户注册
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param userReq body user.UserReq true "用户信息"
// @Success 200 {object} utils.Response
// @Router /user/registe [post]
func Register(ctx iris.Context) {

	var userReq userModel.UserReq
	if err := ctx.ReadForm(&userReq); err != nil {
		utils.Responser.FailWithMsg(ctx, "数据接收失败")
		return
	}
	userSever.Register(ctx, userReq)
}

// @Tags User
// @Summary 用户登录
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Mail body string true "用户邮箱"
// @Param Password body string true "用户密码"
// @Success 200 {object} utils.Response{data=user.UserDetailsRes} ”这里的token是会有信息的"
// @Router /user/login [post]
func Login(ctx iris.Context) {
	mail := ctx.FormValue("Mail")
	if mail == "" {
		utils.Responser.FailWithMsg(ctx, "用户邮箱为空")
		return
	}
	password := ctx.FormValue("Password")
	if ctx.FormValue("Password") == "" {
		utils.Responser.FailWithMsg(ctx, "用户密码为空")
		return
	}
	userSever.Login(ctx, mail, password)
}

// @Tags User
// @Summary 获取头像
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param id path int true "用户id"
// @Success 200 {string}  string  "直接返回文件的渲染视图"
// @Router /user/photo/{id:int} [get]
func GetPhoto(ctx iris.Context) {
	id, err := strconv.Atoi(ctx.Params().Get("id"))
	if err != nil {
		utils.Responser.FailWithMsg(ctx, "数据接收失败")
		return
	}
	userSever.GetUserPhoto(ctx, id)
}

// @Tags User
// @Summary 设置头像
// @Security ApiKeyAuth
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Param GetPhoto formData string true "头像文件"
// @Success 200 {object} utils.Response
// @Router /user/set_photo [post]
func SetPhoto(ctx iris.Context) {
	user, ok := middleware.ParseToken(ctx)
	if !ok {
		utils.Responser.FailWithMsg(ctx, "解析TOKEN出错，请重新登录")
		return
	}
	userSever.SetUserPhoto(ctx, *user)
}

// @Tags User
// @Summary 用户注销
// @Security ApiKeyAuth
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Success 200 {object} utils.Response
// @Router /user/cancellation [Delete]
func Cancellation(ctx iris.Context) {
	user, ok := middleware.ParseToken(ctx)
	if !ok {
		utils.Responser.FailWithMsg(ctx, "解析TOKEN出错，请重新登录")
		return
	}
	userSever.Cancellation(ctx, *user)
}

// @Tags User
// @Summary 获取用户信息
// @Security ApiKeyAuth
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Success 200 {object} utils.Response{data=user.UserDetailsRes} ”这里的token是没有信息的"
// @Router /user/message [Get]
func GetUserMessage(ctx iris.Context) {
	user, ok := middleware.ParseToken(ctx)
	if !ok {
		utils.Responser.FailWithMsg(ctx, "解析TOKEN出错，请重新登录")
		return
	}
	userSever.GetUserMessage(ctx, *user)
}

// @Tags User
// @Summary 修改用户信息
// @Security ApiKeyAuth
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Param userReq body user.UserReq true "用户信息"
// @Success 200 {object} utils.Response
// @Router /user/editor/message [post]
func SetUserMessage(ctx iris.Context) {
	user, ok := middleware.ParseToken(ctx)
	if !ok {
		utils.Responser.FailWithMsg(ctx, "解析TOKEN出错，请重新登录")
		return
	}

	var userReq userModel.UserReq
	if err := ctx.ReadForm(&userReq); err != nil {
		utils.Responser.FailWithMsg(ctx, "数据接收失败")
		return
	}
	userSever.SetUserMessage(ctx, *user, userReq)
}
