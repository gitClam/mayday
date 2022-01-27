package user_routes

import (
	"github.com/kataras/iris/v12"
	"io"
	"log"
	"mayday/src/global"
	"mayday/src/middleware"
	"mayday/src/model/user"
	"mayday/src/utils"
	"os"
	"strconv"
	//"strconv"
	"time"
)

func UserRegister(ctx iris.Context) {
	var sdUser user.SdUser
	if err := ctx.ReadForm(&sdUser); err != nil {
		utils.Responser.FailWithMsg(ctx, "数据接收失败")
		log.Print("用户注册失败，数据接收失败")
		return
	}
	log.Print(sdUser)
	sdUser.Photo = "./data/photo/2.png"
	sdUser.IsDeleted = 0
	sdUser.CreateDate = utils.LocalTime(time.Now())
	log.Print(sdUser)
	e := global.GVA_DB
	effect, err := e.Insert(sdUser)
	if effect <= 0 || err != nil {
		log.Printf("用户注册失败")
		utils.Responser.FailWithMsg(ctx, "用户注册失败")
		return
	}

	utils.Responser.Ok(ctx)
	log.Println("ok")
}

func UserLogin(ctx iris.Context) {

	var sdUser user.SdUser
	if err := ctx.ReadForm(&sdUser); err != nil {
		utils.Responser.FailWithMsg(ctx, "用户数据接收失败")
		log.Print("用户登录失败，数据接收失败")
		return
	}

	if sdUser.Mail == "" || sdUser.Password == "" {
		utils.Responser.FailWithMsg(ctx, "用户名或密码为空")
		log.Print("用户登录失败,邮箱或密码为空")
		return
	}

	var mUser user.SdUser
	mUser.Mail = sdUser.Mail

	e := global.GVA_DB
	has, err := e.Where("is_deleted != 1").Get(&mUser)
	if !has || err != nil || mUser.IsDeleted == 1 {
		utils.Responser.FailWithMsg(ctx, "用户名不存在")
		log.Printf("数据库查询错误或用户名不存在")
		return
	}

	log.Print(mUser)

	if mUser.Password != sdUser.Password {
		utils.Responser.FailWithMsg(ctx, "密码错误")
		log.Printf("密码错误")
		return
	}

	token, err := middleware.GenerateToken(&mUser)
	log.Printf("用户[%s], 登录生成token [%s]", mUser.Name, token)
	if err != nil {
		utils.Responser.FailWithMsg(ctx, "TOKEN生成失败")
		log.Printf("数据库查询错误或用户名不存在")
		return
	}

	utils.Responser.OkWithDetails(ctx, utils.Success, user.TransformUserVOToken(token, &mUser))
}

func UserPhoto(ctx iris.Context) {

	var user user.SdUser
	Id, err := strconv.Atoi(ctx.Params().Get("id"))
	if err != nil {
		log.Printf("数据接收失败")
		utils.Responser.FailWithMsg(ctx, "")
		return
	}
	user.Id = Id
	e := global.GVA_DB
	has, err := e.Get(&user)
	if !has || err != nil {
		log.Printf("数据库查询错误或用户名不存在")
		utils.Responser.FailWithMsg(ctx, "用户名不存在")
		return
	}

	if user.Photo == "" {
		log.Printf("用户头像获取出错")
		utils.Responser.FailWithMsg(ctx, "用户头像获取出错")
		return
	}
	err1 := ctx.ServeFile(user.Photo, false)
	if err1 != nil {
		log.Printf("用户头像文件读取错误")
		utils.Responser.FailWithMsg(ctx, "头像文件读取错误")
		return
	}
}

func SetUserPhoto(ctx iris.Context) {
	token, ok := middleware.ParseToken(ctx)
	if !ok {
		log.Printf("解析TOKEN出错，请重新登录")
		utils.Responser.FailWithMsg(ctx, "解析TOKEN出错，请重新登录")
		return
	}
	var mUser user.SdUser
	mUser.Id = token.Id
	mUser.Name = token.Name
	e := global.GVA_DB
	has, err := e.Get(&mUser)
	if !has || err != nil {

		log.Printf("数据库查询错误或用户名不存在")
		utils.Responser.FailWithMsg(ctx, "用户名不存在")
		return
	}

	log.Print("接收图片中")
	file, _, err := ctx.FormFile("UserPhoto")
	if err != nil {
		log.Print("图片文件不存在")
		utils.Responser.FailWithMsg(ctx, "图片接收失败")
		return
	}
	defer file.Close()
	//fname := info.Filename
	if mUser.Photo == "./data/photo/2.png" || mUser.Photo == "" {
		mUser.Photo = "./data/photo/" + mUser.Mail
		affected, err := e.Id(mUser.Id).Update(mUser)
		if affected <= 0 || err != nil {
			log.Printf("数据库更新失败")
			utils.Responser.FailWithMsg(ctx, "图片更新失败")
			return
		}
	}

	out, err := os.OpenFile(mUser.Photo, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Print("文件打开失败")
		utils.Responser.FailWithMsg(ctx, "图片文件保存失败")
		return
	}
	defer out.Close()

	io.Copy(out, file)

	utils.Responser.Ok(ctx)

	log.Print("图片已保存")
}

func UserCancellation(ctx iris.Context) {
	token, ok := middleware.ParseToken(ctx)
	if !ok {
		log.Printf("解析TOKEN出错，请重新登录")
		utils.Responser.FailWithMsg(ctx, "解析TOKEN出错，请重新登录")
		return
	}
	var mUser user.SdUser
	mUser.Id = token.Id
	e := global.GVA_DB
	log.Print(mUser.Id)
	has, err := e.Id(mUser.Id).Get(&mUser)
	if !has || err != nil {
		log.Print(err)
		log.Printf("数据库查询错误或用户名不存在")
		utils.Responser.FailWithMsg(ctx, "用户名不存在")
		return
	}
	mUser.IsDeleted = 1
	affected, err1 := e.Id(mUser.Id).Update(&mUser)
	if affected <= 0 || err1 != nil {
		log.Print(err)
		log.Printf("数据库修改失败")
		utils.Responser.FailWithMsg(ctx, "注销失败")
		return
	}
	utils.Responser.Ok(ctx)
}

func UserMessage(ctx iris.Context) {

	token, ok := middleware.ParseToken(ctx)
	if !ok {
		log.Printf("解析TOKEN出错，请重新登录")
		utils.Responser.FailWithMsg(ctx, "解析TOKEN出错，请重新登录")
		return
	}
	var mUser user.SdUser
	mUser.Id = token.Id
	mUser.Name = token.Name

	e := global.GVA_DB
	has, err := e.Where(" is_deleted != 1 ").Get(&mUser)
	if !has || err != nil {
		log.Printf("数据库查询错误或用户名不存在")
		utils.Responser.FailWithMsg(ctx, "用户名不存在")
		return
	}

	utils.Responser.OkWithDetails(ctx, utils.Success, user.TransformUserVO(&mUser))

}

func SetUserMessage(ctx iris.Context) {

	token, ok := middleware.ParseToken(ctx)
	if !ok {
		log.Printf("解析TOKEN出错，请重新登录")
		utils.Responser.FailWithMsg(ctx, "解析TOKEN出错，请重新登录")
		return
	}

	var mUser user.SdUser
	if err := ctx.ReadForm(&mUser); err != nil {
		log.Print(err)
		utils.Responser.FailWithMsg(ctx, "数据接收失败")
		log.Print("数据接收失败")
		return
	}

	e := global.GVA_DB
	affected, err := e.Id(token.Id).Update(mUser)
	if affected <= 0 || err != nil {
		log.Print(err)
		log.Printf("数据库更新失败")
		utils.Responser.FailWithMsg(ctx, "数据更新失败")
		return
	}

	utils.Responser.Ok(ctx)

}
