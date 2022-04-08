package user_api

import (
	"github.com/kataras/iris/v12"
	"mayday/src/model/common/resultcode"
	userModel "mayday/src/model/user"
	userSever "mayday/src/service/user"
	"mayday/src/utils"
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
		utils.Responser.Fail(ctx, resultcode.DataReceiveFail, err)
		return
	}
	userSever.Register(ctx, userReq)
}

// @Tags User
// @Summary 用户登录
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param mail body string true "用户邮箱"
// @Param password body string true "用户密码"
// @Success 200 {object} utils.Response{data=user.UserDetailsRes} ”这里的token是会有信息的"
// @Router /user/login [post]
func Login(ctx iris.Context) {
	mail := ctx.FormValue("mail")
	if mail == "" {
		utils.Responser.Fail(ctx, resultcode.UsernameFail)
		return
	}
	password := ctx.FormValue("password")
	if password == "" {
		utils.Responser.Fail(ctx, resultcode.PasswordFail)
		return
	}
	userSever.Login(ctx, mail, password)
}

// @Tags User
// @Summary 获取头像
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Success 200 {string}  string  "直接返回文件的渲染视图"
// @Router /user/photo/{fileName:string}/ [get]
func GetPhoto(ctx iris.Context) {
	fileName := ctx.Params().Get("fileName")
	if fileName == "" {
		utils.Responser.Fail(ctx, resultcode.DataReceiveFail)
		return
	}
	userSever.GetUserPhoto(ctx, fileName)
}

// @Tags User
// @Summary 设置头像
// @Security ApiKeyAuth
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Param UserPhoto formData string true "头像文件"
// @Success 200 {object} utils.Response{data=user.UserPhotoFileName}
// @Router /user/set_photo [post]
func SetPhoto(ctx iris.Context) {
	userSever.SetUserPhoto(ctx)
}

// @Tags User
// @Summary 用户注销
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Success 200 {object} utils.Response
// @Router /user/cancellation [Delete]
func Cancellation(ctx iris.Context) {
	userSever.Cancellation(ctx)
}

// @Tags User
// @Summary 获取用户信息
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Success 200 {object} utils.Response{data=user.UserDetailsRes} ”这里的token是没有信息的"
// @Router /user/message [Get]
func GetUserMessage(ctx iris.Context) {
	userSever.GetUserMessage(ctx)
}

// @Tags User
// @Summary 修改用户信息
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Param userReq body user.UserReq true "用户信息"
// @Success 200 {object} utils.Response
// @Router /user/editor/message [post]
func SetUserMessage(ctx iris.Context) {
	var userReq userModel.UserReq
	if err := ctx.ReadForm(&userReq); err != nil {
		utils.Responser.Fail(ctx, resultcode.DataReceiveFail, err)
		return
	}
	userSever.SetUserMessage(ctx, userReq)
}
