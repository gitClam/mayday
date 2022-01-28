package user_routes

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
	"mayday/src/global"
	"mayday/src/middleware"
	"mayday/src/model/user"
	"mayday/src/utils"
	"strconv"
	"time"
)

// @Tags User
// @Summary 用户注册
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param userReq body user.UserReq true "用户信息"
// @Success 200 {object} utils.Response
// @Router /user/registe [post]
func UserRegister(ctx iris.Context) {
	var sdUser user.SdUser
	if err := ctx.ReadForm(&sdUser); err != nil {
		utils.Responser.FailWithMsg(ctx, "数据接收失败")
		return
	}

	sdUser.Photo = global.GVA_CONFIG.System.DefaultHeadPortrait
	sdUser.CreateDate = utils.LocalTime(time.Now())

	e := global.GVA_DB
	effect, err := e.Insert(sdUser)
	if effect <= 0 || err != nil {
		utils.Responser.FailWithMsg(ctx, "用户注册失败")
		return
	}

	utils.Responser.Ok(ctx)
	global.GVA_LOG.Info("用户: " + sdUser.Mail + " 注册成功")
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
func UserLogin(ctx iris.Context) {

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

	var mUser user.SdUser
	mUser.Mail = mail

	e := global.GVA_DB
	has, err := e.Get(&mUser)
	if !has || err != nil {
		utils.Responser.FailWithMsg(ctx, "用户名不存在")
		return
	}

	if mUser.Password != password {
		utils.Responser.FailWithMsg(ctx, "密码错误")
		return
	}

	token, err := middleware.GenerateToken(&mUser)
	global.GVA_LOG.Info(fmt.Sprintf("用户[%s], 登录生成token [%s]", mUser.Name, token))
	if err != nil {
		utils.Responser.FailWithMsg(ctx, "TOKEN生成失败")
		return
	}

	utils.Responser.OkWithDetails(ctx, utils.Success, user.GetUserDetailsResWithToken(token, &mUser))
	global.GVA_LOG.Info("用户: " + mail + " 登录成功")
}

// @Tags User
// @Summary 获取头像
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param id path int true "用户id"
// @Success 200 {string}  string  "直接返回文件的渲染视图"
// @Router /user/photo/{id:int} [get]
func UserPhoto(ctx iris.Context) {

	Id, err := strconv.Atoi(ctx.Params().Get("id"))
	if err != nil {
		utils.Responser.FailWithMsg(ctx, "数据接收失败")
		return
	}

	var sdUser user.SdUser
	sdUser.Id = Id

	e := global.GVA_DB
	has, err := e.Get(&sdUser)
	if !has || err != nil {
		utils.Responser.FailWithMsg(ctx, "用户名不存在")
		return
	}

	if sdUser.Photo == "" {
		utils.Responser.FailWithMsg(ctx, "用户头像未设置")
		return
	}

	err1 := ctx.ServeFile(sdUser.Photo, false)
	if err1 != nil {
		utils.Responser.FailWithMsg(ctx, "头像文件读取错误")
		return
	}
}

// @Tags User
// @Summary 设置头像
// @Security ApiKeyAuth
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Param UserPhoto formData string true "头像文件"
// @Success 200 {object} utils.Response
// @Router /user/set_photo [post]
func SetUserPhoto(ctx iris.Context) {

	token, ok := middleware.ParseToken(ctx)
	if !ok {
		utils.Responser.FailWithMsg(ctx, "解析TOKEN出错，请重新登录")
		return
	}

	var mUser user.SdUser
	mUser.Id = token.Id
	mUser.Name = token.Name

	has, err := global.GVA_DB.Get(&mUser)
	if !has || err != nil {
		utils.Responser.FailWithMsg(ctx, "用户名不存在")
		return
	}

	file, _, err := ctx.FormFile("UserPhoto")
	if err != nil {
		utils.Responser.FailWithMsg(ctx, "图片接收失败")
		return
	}

	photoPath := global.GVA_CONFIG.System.PhotoPath + mUser.Mail

	err = utils.IO.Save(photoPath, file)
	if err != nil {
		global.GVA_LOG.Error("头像文件保存出错：", zap.Error(err))
		utils.Responser.FailWithMsg(ctx, "图片文件保存失败")
		return
	}

	mUser.Photo = photoPath
	affected, err := global.GVA_DB.Id(mUser.Id).Update(mUser)
	if affected <= 0 || err != nil {
		utils.Responser.FailWithMsg(ctx, "图片更新失败")
		return
	}

	utils.Responser.Ok(ctx)
	global.GVA_LOG.Info("用户: " + mUser.Mail + " 头像保存成功")
}

// @Tags User
// @Summary 用户注销
// @Security ApiKeyAuth
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Success 200 {object} utils.Response
// @Router /user/cancellation [Delete]
func UserCancellation(ctx iris.Context) {

	token, ok := middleware.ParseToken(ctx)
	if !ok {
		utils.Responser.FailWithMsg(ctx, "解析TOKEN出错，请重新登录")
		return
	}

	var mUser user.SdUser
	mUser.Id = token.Id

	effect, err := global.GVA_DB.Id(mUser.Id).Delete(&mUser)
	if effect <= 0 || err != nil {
		utils.Responser.FailWithMsg(ctx, "用户注销失败")
		return
	}

	utils.Responser.Ok(ctx)
}

// @Tags User
// @Summary 获取用户信息
// @Security ApiKeyAuth
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Success 200 {object} utils.Response{data=user.UserDetailsRes} ”这里的token是没有信息的"
// @Router /user/message [Get]
func UserMessage(ctx iris.Context) {

	token, ok := middleware.ParseToken(ctx)
	if !ok {
		utils.Responser.FailWithMsg(ctx, "解析TOKEN出错，请重新登录")
		return
	}

	var mUser user.SdUser
	mUser.Id = token.Id
	mUser.Name = token.Name

	has, err := global.GVA_DB.Get(&mUser)
	if !has || err != nil {
		utils.Responser.FailWithMsg(ctx, "用户名不存在")
		return
	}

	utils.Responser.OkWithDetails(ctx, utils.Success, user.GetUserDetailsResWithOutToken(&mUser))

}

// @Tags User
// @Summary 修改用户信息
// @Security ApiKeyAuth
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Param userReq body user.UserReq true "用户信息"
// @Success 200 {object} utils.Response
// @Router /user/editor/message [post]
func SetUserMessage(ctx iris.Context) {

	token, ok := middleware.ParseToken(ctx)
	if !ok {
		utils.Responser.FailWithMsg(ctx, "解析TOKEN出错，请重新登录")
		return
	}

	var mUser user.SdUser
	if err := ctx.ReadForm(&mUser); err != nil {
		utils.Responser.FailWithMsg(ctx, "数据接收失败")
		return
	}

	affected, err := global.GVA_DB.Id(token.Id).Update(mUser)
	if affected <= 0 || err != nil {
		utils.Responser.FailWithMsg(ctx, "数据更新失败")
		return
	}

	utils.Responser.Ok(ctx)

}
