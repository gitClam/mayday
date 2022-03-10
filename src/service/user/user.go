package user_Server

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"mayday/src/global"
	"mayday/src/middleware"
	"mayday/src/model/common/resultcode"
	"mayday/src/model/common/timedecoder"
	userModel "mayday/src/model/user"
	"mayday/src/utils"
	"strconv"
	"time"
)

//用户注册
func Register(ctx iris.Context, userReq userModel.UserReq) {

	if userReq.Password == "" || userReq.Mail == "" {
		utils.Responser.Fail(ctx, resultcode.EmptyMaliOrPassWord)
		return
	}

	sdUser := userReq.GetSdUser()
	sdUser.Photo = global.GVA_CONFIG.System.DefaultHeadPortrait
	sdUser.CreateDate = timedecoder.LocalTime(time.Now())

	e := global.GVA_DB
	effect, err := e.Insert(sdUser)
	if effect <= 0 || err != nil {
		utils.Responser.Fail(ctx, resultcode.RegisterFail, err)
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
		utils.Responser.Fail(ctx, resultcode.UsernameFail, err)
		return
	}

	if mUser.Password != password {
		utils.Responser.Fail(ctx, resultcode.PasswordFail)
		return
	}

	token, err := middleware.GenerateToken(&mUser)
	if err != nil {
		utils.Responser.Fail(ctx, resultcode.TokenCreateFail, err)
		return
	}

	ctx.Values().Set("user", fmt.Sprint(strconv.Itoa(mUser.Id)+" "+mUser.Name))
	utils.Responser.OkWithDetails(ctx, userModel.GetUserDetailsResWithToken(token, &mUser))
}

//获取用户头像
func GetUserPhoto(ctx iris.Context, fileName string) {

	err := ctx.ServeFile(global.GVA_CONFIG.System.PhotoPath+"/"+fileName, false)
	if err != nil {
		utils.Responser.Fail(ctx, resultcode.PhotoReadFail, err)
		return
	}
}

//设置头像
func SetUserPhoto(ctx iris.Context) {
	user := ctx.Values().Get("user").(userModel.SdUser)
	has, err := global.GVA_DB.Get(&user)
	if !has || err != nil {
		utils.Responser.Fail(ctx, resultcode.UsernameFail, err)
		return
	}

	file, _, err := ctx.FormFile("UserPhoto")
	if err != nil {
		utils.Responser.Fail(ctx, resultcode.DataReceiveFail, err)
		return
	}

	fileName := user.Mail + "_" + strconv.FormatInt(time.Now().Unix(), 10)

	err = utils.IO.Save(global.GVA_CONFIG.System.PhotoPath+"/"+fileName, file)
	if err != nil {
		utils.Responser.Fail(ctx, resultcode.PhotoSaveFail, err)
		return
	}

	if user.Photo != fileName {
		user.Photo = fileName
		affected, err := global.GVA_DB.Id(user.Id).Update(user)
		if affected <= 0 || err != nil {
			utils.Responser.Fail(ctx, resultcode.PhotoUpdateFail, err)
			return
		}
	}

	utils.Responser.OkWithDetails(ctx, userModel.UserPhotoFileName{FileName: fileName})
}

//用户注销
func Cancellation(ctx iris.Context) {
	user := ctx.Values().Get("user").(userModel.SdUser)
	effect, err := global.GVA_DB.Id(user.Id).Delete(&user)
	if effect <= 0 || err != nil {
		utils.Responser.Fail(ctx, resultcode.CancellationFail, err)
		return
	}
	utils.Responser.Ok(ctx)
}

//获取用户信息
func GetUserMessage(ctx iris.Context) {
	user := ctx.Values().Get("user").(userModel.SdUser)
	has, err := global.GVA_DB.Get(&user)
	if !has || err != nil {
		utils.Responser.Fail(ctx, resultcode.UsernameFail, err)
		return
	}
	utils.Responser.OkWithDetails(ctx, userModel.GetUserDetailsResWithOutToken(&user))
}

//修改用户信息
func SetUserMessage(ctx iris.Context, msg userModel.UserReq) {
	user := ctx.Values().Get("user").(userModel.SdUser)
	sdUser := msg.GetSdUser()
	affected, err := global.GVA_DB.Id(user.Id).Update(sdUser)
	if affected <= 0 || err != nil {
		utils.Responser.Fail(ctx, resultcode.DataUpdateFail, err)
		return
	}

	utils.Responser.Ok(ctx)
}
