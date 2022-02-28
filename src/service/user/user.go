package user_Server

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"mayday/src/global"
	"mayday/src/middleware"
	userModel "mayday/src/model/user"
	"mayday/src/utils"
	"strconv"
	"time"
)

//用户注册
func Register(ctx iris.Context, userReq userModel.UserReq) {

	if userReq.Password == "" || userReq.Mail == "" {
		utils.Responser.FailWithMsg(ctx, "邮箱或密码为空")
		return
	}

	sdUser := userReq.GetSdUser()
	sdUser.Photo = global.GVA_CONFIG.System.DefaultHeadPortrait
	sdUser.CreateDate = utils.LocalTime(time.Now())

	e := global.GVA_DB
	effect, err := e.Insert(sdUser)
	if effect <= 0 || err != nil {
		utils.Responser.FailWithMsg(ctx, "用户注册失败", err)
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
		utils.Responser.FailWithMsg(ctx, "用户名不存在", err)
		return
	}

	if mUser.Password != password {
		utils.Responser.FailWithMsg(ctx, "密码错误")
		return
	}

	token, err := middleware.GenerateToken(&mUser)
	if err != nil {
		utils.Responser.FailWithMsg(ctx, "TOKEN生成失败", err)
		return
	}

	ctx.Values().Set("user", fmt.Sprint(strconv.Itoa(mUser.Id)+" "+mUser.Name))
	utils.Responser.OkWithDetails(ctx, utils.Success, userModel.GetUserDetailsResWithToken(token, &mUser))
}

//获取用户头像
func GetUserPhoto(ctx iris.Context, fileName string) {

	err := ctx.ServeFile(global.GVA_CONFIG.System.PhotoPath+"/"+fileName, false)
	if err != nil {
		utils.Responser.FailWithMsg(ctx, "头像文件读取错误", err)
		return
	}
}

//设置头像
func SetUserPhoto(ctx iris.Context) {
	user := ctx.Values().Get("user").(userModel.SdUser)
	has, err := global.GVA_DB.Get(&user)
	if !has || err != nil {
		utils.Responser.FailWithMsg(ctx, "用户名不存在", err)
		return
	}

	file, _, err := ctx.FormFile("UserPhoto")
	if err != nil {
		utils.Responser.FailWithMsg(ctx, "图片接收失败", err)
		return
	}

	fileName := user.Mail + "_" + strconv.FormatInt(time.Now().Unix(), 10)

	err = utils.IO.Save(global.GVA_CONFIG.System.PhotoPath+"/"+fileName, file)
	if err != nil {
		utils.Responser.FailWithMsg(ctx, "图片文件保存失败", err)
		return
	}

	if user.Photo != fileName {
		user.Photo = fileName
		affected, err := global.GVA_DB.Id(user.Id).Update(user)
		if affected <= 0 || err != nil {
			utils.Responser.FailWithMsg(ctx, "图片更新失败", err)
			return
		}
	}

	utils.Responser.OkWithDetails(ctx, "成功", iris.Map{"fileName": fileName})
}

//用户注销
func Cancellation(ctx iris.Context) {
	user := ctx.Values().Get("user").(userModel.SdUser)
	effect, err := global.GVA_DB.Id(user.Id).Delete(&user)
	if effect <= 0 || err != nil {
		utils.Responser.FailWithMsg(ctx, "用户注销失败", err)
		return
	}
	utils.Responser.Ok(ctx)
}

//获取用户信息
func GetUserMessage(ctx iris.Context) {
	user := ctx.Values().Get("user").(userModel.SdUser)
	has, err := global.GVA_DB.Get(&user)
	if !has || err != nil {
		utils.Responser.FailWithMsg(ctx, "用户名不存在", err)
		return
	}
	utils.Responser.OkWithDetails(ctx, utils.Success, userModel.GetUserDetailsResWithOutToken(&user))
}

//修改用户信息
func SetUserMessage(ctx iris.Context, msg userModel.UserReq) {
	user := ctx.Values().Get("user").(userModel.SdUser)
	affected, err := global.GVA_DB.Id(user.Id).Update(msg)
	if affected <= 0 || err != nil {
		utils.Responser.FailWithMsg(ctx, "数据更新失败", err)
		return
	}

	utils.Responser.Ok(ctx)
}
