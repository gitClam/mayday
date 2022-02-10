package user_Server

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
	"mayday/src/global"
	"mayday/src/middleware"
	userModel "mayday/src/model/user"
	"mayday/src/utils"
	"time"
)

//用户注册
func Register(ctx iris.Context, userReq userModel.UserReq) {

	sdUser := userReq.GetSdUser()
	sdUser.Photo = global.GVA_CONFIG.System.DefaultHeadPortrait
	sdUser.CreateDate = utils.LocalTime(time.Now())

	e := global.GVA_DB
	effect, err := e.Insert(sdUser)
	if effect <= 0 || err != nil {
		utils.Responser.FailWithMsg(ctx, "用户注册失败")
		return
	}

	utils.Responser.Ok(ctx)
}

//用户登录
func Login(ctx iris.Context, mail string, password string) {

	var mUser userModel.SdUser
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
	if err != nil {
		utils.Responser.FailWithMsg(ctx, "TOKEN生成失败")
		global.GVA_LOG.Warn("用户: "+mUser.Name+" TOKEN生成失败", zap.Error(err))
		return
	}

	utils.Responser.OkWithDetails(ctx, utils.Success, userModel.GetUserDetailsResWithToken(token, &mUser))
	global.GVA_LOG.Info(fmt.Sprintf("用户[%s], 登录生成token [%s]", mUser.Name, token))
}

//获取用户头像
func GetUserPhoto(ctx iris.Context, id int) {

	var sdUser userModel.SdUser
	sdUser.Id = id

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

//设置头像
func SetUserPhoto(ctx iris.Context, user userModel.SdUser) {

	has, err := global.GVA_DB.Get(&user)
	if !has || err != nil {
		utils.Responser.FailWithMsg(ctx, "用户名不存在")
		return
	}

	file, _, err := ctx.FormFile("GetPhoto")
	if err != nil {
		utils.Responser.FailWithMsg(ctx, "图片接收失败")
		return
	}

	photoPath := global.GVA_CONFIG.System.PhotoPath + user.Mail

	err = utils.IO.Save(photoPath, file)
	if err != nil {
		utils.Responser.FailWithMsg(ctx, "图片文件保存失败")
		return
	}

	if user.Photo != photoPath {
		affected, err := global.GVA_DB.Id(user.Id).Update(user)
		if affected <= 0 || err != nil {
			utils.Responser.FailWithMsg(ctx, "图片更新失败")
			return
		}
	}

	utils.Responser.Ok(ctx)
	global.GVA_LOG.Info("用户: " + user.Mail + " 头像保存成功")
}

//用户注销
func Cancellation(ctx iris.Context, user userModel.SdUser) {
	effect, err := global.GVA_DB.Id(user.Id).Delete(&user)
	if effect <= 0 || err != nil {
		utils.Responser.FailWithMsg(ctx, "用户注销失败")
		return
	}

	utils.Responser.Ok(ctx)
	global.GVA_LOG.Info("用户: " + user.Mail + " 已注销")
}

//获取用户信息
func GetUserMessage(ctx iris.Context, user userModel.SdUser) {

	has, err := global.GVA_DB.Get(&user)
	if !has || err != nil {
		utils.Responser.FailWithMsg(ctx, "用户名不存在")
		return
	}

	utils.Responser.OkWithDetails(ctx, utils.Success, userModel.GetUserDetailsResWithOutToken(&user))
	global.GVA_LOG.Info("用户: " + user.Mail + " 获取用户信息")
}

//修改用户信息
func SetUserMessage(ctx iris.Context, user userModel.SdUser, msg userModel.UserReq) {

	affected, err := global.GVA_DB.Id(user.Id).Update(msg)
	if affected <= 0 || err != nil {
		utils.Responser.FailWithMsg(ctx, "数据更新失败")
		return
	}

	utils.Responser.Ok(ctx)
	global.GVA_LOG.Info("用户: " + user.Mail + " 修改用户信息")
}
