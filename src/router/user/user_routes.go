package user_routes

import (
	"github.com/kataras/iris/v12"
	"io"
	"log"
	"mayday/src/global"
	"mayday/src/middleware/jwts"
	"mayday/src/model"
	"mayday/src/utils"
	"os"
	"strconv"
	//"strconv"
	"time"
)

// swagger:operation POST /user/registe user registe
// ---
// summary: 用户注册
// description: 用户注册
// parameters:
// - name: name
//   description: 用户昵称
//   type: string
//   required: true
// - name: password
//   description: 用户 密码
//   type: string
//   required: true
// - name: realname
//   description: 真实姓名
//   type: string
//   required: false
// - name: age
//   description: 用户年龄
//   type: int
//   required: false
// - name: birthday
//   description: 用户生日
//   type: datetime
//   required: false
// - name: sex
//   in: 男/女
//   description: 用户性别
//   type: string
//   required: true
// - name: Wechat
//   description: 微信
//   type: string
//   required: false
// - name: Qqnumber
//   description: QQ
//   type: string
//   required: false
// - name: Info
//   description: 备注
//   type: string
//   required: false
// - name: mail
//   description: 邮箱
//   type: string
//   required: true
// - name: company
//   description: 公司
//   type: string
//   required: false
// - name: vocation
//   description: 职业
//   type: string
//   required: false
// - name: department
//   description: 部门
//   type: string
//   required: false
// Responses:
//       '200':
//         schema:
//           $ref: '#/responses/forbidden'

func UserRegister(ctx iris.Context) {
	var user model.SdUser
	if err := ctx.ReadForm(&user); err != nil {
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.RegisteFailur, "数据接收失败")
		log.Print("用户注册失败，数据接收失败")
		return
	}
	log.Print(user)
	user.Photo = "./data/photo/2.png"
	user.IsDeleted = 0
	user.CreateDate = model.LocalTime(time.Now())
	log.Print(user)
	e := global.GVA_DB
	effect, err := e.Insert(user)
	if effect <= 0 || err != nil {
		log.Printf("用户注册失败")
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.RegisteFailur, "用户注册失败")
		return
	}

	utils.MakeSuccessRes(ctx, model.Success, nil)
	log.Println("ok")
}

// swagger:operation POST /user/login user login
// ---
// summary: 用户登录
// description: 用户登录
// parameters:
// - name: mail
//   description: 用户邮箱
//   type: string
//   required: true
// - name: password
//   description: 用户密码
//   type: string
//   required: true
func UserLogin(ctx iris.Context) {

	var user model.SdUser
	if err := ctx.ReadForm(&user); err != nil {
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.LoginFailur, nil)
		log.Print("用户登录失败，数据接收失败")
		return
	}

	if user.Mail == "" || user.Password == "" {
		utils.MakeErrorRes(ctx, 3333, model.LoginFailur, nil)
		log.Print("用户登录失败,邮箱或密码为空")
		return
	}

	var mUser model.SdUser
	mUser.Mail = user.Mail

	e := global.GVA_DB
	has, err := e.Where("is_deleted != 1").Get(&mUser)
	if !has || err != nil || mUser.IsDeleted == 1 {
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.LoginFailur, nil)
		log.Printf("数据库查询错误或用户名不存在")
		return
	}

	log.Print(mUser)

	if mUser.Password != user.Password {
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.LoginFailur, nil)
		log.Printf("密码错误")
		return
	}

	token, err := jwts.GenerateToken(&mUser)
	log.Printf("用户[%s], 登录生成token [%s]", mUser.Name, token)
	if err != nil {
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.TokenCreateFailur, nil)
		log.Printf("数据库查询错误或用户名不存在")
		return
	}

	utils.MakeSuccessRes(ctx, model.Success, utils.TransformUserVOToken(token, &mUser))
}

// swagger:operation GET /user/photo/{id:int} user get_photo
// ---
// summary: 获取用户头像
// description: 获取用户头像

func UserPhoto(ctx iris.Context) {

	var user model.SdUser
	Id, err := strconv.Atoi(ctx.Params().Get("id"))
	if err != nil {
		log.Printf("数据接收失败")
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	user.Id = Id
	e := global.GVA_DB
	has, err := e.Get(&user)
	if !has || err != nil {
		log.Printf("数据库查询错误或用户名不存在")
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, "用户名不存在")
		return
	}

	if user.Photo == "" {
		log.Printf("用户头像获取出错")
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, "用户头像获取出错")
		return
	}
	err1 := ctx.ServeFile(user.Photo, false)
	if err1 != nil {
		log.Printf("用户头像文件读取错误")
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, "用户头像获取出错")
		return
	}
}

// swagger:operation POST /user/set_photo user set_photo
// ---
// summary: 设置用户头像
// description: 设置用户头像
// parameters:
// - name: UserPhoto
//   description: 用户头像
//   type: file
//   required: true
func SetUserPhoto(ctx iris.Context) {
	user, ok := jwts.ParseToken(ctx)
	if !ok {
		log.Printf("解析TOKEN出错，请重新登录")
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.TokenParseFailur, "解析TOKEN出错，请重新登录")
		return
	}
	log.Print(user)
	var mUser model.SdUser
	mUser.Id = user.Id
	mUser.Name = user.Name
	e := global.GVA_DB
	has, err := e.Get(&mUser)
	if !has || err != nil {

		log.Printf("数据库查询错误或用户名不存在")
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, "用户名不存在")
		return
	}

	log.Print("接收图片中")
	file, _, err := ctx.FormFile("UserPhoto")
	if err != nil {
		log.Print("图片文件不存在")
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, "图片文件不存在")
		return
	}
	defer file.Close()
	//fname := info.Filename
	if mUser.Photo == "./data/photo/2.png" || mUser.Photo == "" {
		mUser.Photo = "./data/photo/" + mUser.Mail
		affected, err := e.Id(mUser.Id).Update(mUser)
		if affected <= 0 || err != nil {
			log.Printf("数据库更新失败")
			utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, "更新失败")
			return
		}
	}

	out, err := os.OpenFile(mUser.Photo, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Print("文件打开失败")
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, "文件打开失败")
		return
	}
	defer out.Close()

	io.Copy(out, file)

	utils.MakeSuccessRes(ctx, model.Success, nil)

	log.Print("图片已保存")
}

// swagger:operation Delete /user/cancellation user cancellation
// ---
// summary: 用户注销
// description: 用户注销

func UserCancellation(ctx iris.Context) {
	user, ok := jwts.ParseToken(ctx)
	if !ok {
		log.Printf("解析TOKEN出错，请重新登录")
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.TokenParseFailur, "解析TOKEN出错，请重新登录")
		return
	}
	var mUser model.SdUser
	mUser.Id = user.Id
	e := global.GVA_DB
	log.Print(mUser.Id)
	has, err := e.Id(mUser.Id).Get(&mUser)
	if !has || err != nil {
		log.Print(err)
		log.Printf("数据库查询错误或用户名不存在")
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, "用户名不存在")
		return
	}
	mUser.IsDeleted = 1
	affected, err1 := e.Id(mUser.Id).Update(&mUser)
	if affected <= 0 || err1 != nil {
		log.Print(err)
		log.Printf("数据库修改失败")
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, "注销失败")
		return
	}
	utils.MakeSuccessRes(ctx, model.Success, nil)
}

// swagger:operation GET /user/message user message
// ---
// summary: 获取用户信息
// description: 获取用户信息

func UserMessage(ctx iris.Context) {

	user, ok := jwts.ParseToken(ctx)
	if !ok {
		log.Printf("解析TOKEN出错，请重新登录")
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.TokenParseFailur, "解析TOKEN出错，请重新登录")
		return
	}
	var mUser model.SdUser
	mUser.Id = user.Id
	mUser.Name = user.Name

	e := global.GVA_DB
	has, err := e.Where(" is_deleted != 1 ").Get(&mUser)
	if !has || err != nil {
		log.Printf("数据库查询错误或用户名不存在")
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, "用户名不存在")
		return
	}

	utils.MakeSuccessRes(ctx, model.Success, utils.TransformUserVO(&mUser))

}

// swagger:operation POST /user/editor/message user editor_message
// ---
// summary: 修改用户信息
// description: 修改用户信息
// parameters:
// - name: name
//   description: 用户昵称
//   type: string
//   required: false
// - name: password
//   description: 用户 密码
//   type: string
//   required: false
// - name: realname
//   description: 真实姓名
//   type: string
//   required: false
// - name: age
//   description: 用户年龄
//   type: int
//   required: false
// - name: birthday
//   description: 用户生日
//   type: datetime
//   required: false
// - name: sex
//   in: 男/女
//   description: 用户性别
//   type: string
//   required: false
// - name: Wechat
//   description: 微信
//   type: string
//   required: false
// - name: Qqnumber
//   description: QQ
//   type: string
//   required: false
// - name: Info
//   description: 备注
//   type: string
//   required: false
// - name: mail
//   description: 邮箱
//   type: string
//   required: false
// - name: company
//   description: 公司
//   type: string
//   required: false
// - name: vocation
//   description: 职业
//   type: string
//   required: false
// - name: phone
//   description: 联系电话
//   type: string
//   required: false
// - name: department
//   description: 部门
//   type: string
//   required: false
// Responses:
//       '200':
//         schema:
//           $ref: '#/responses/forbidden'

func SetUserMessage(ctx iris.Context) {

	user, ok := jwts.ParseToken(ctx)
	if !ok {
		log.Printf("解析TOKEN出错，请重新登录")
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.TokenParseFailur, "解析TOKEN出错，请重新登录")
		return
	}

	var mUser model.SdUser
	if err := ctx.ReadForm(&mUser); err != nil {
		log.Print(err)
		utils.MakeErrorRes(ctx, model.OtherErrorCode, model.OptionFailur, "数据接收失败")
		log.Print("数据接收失败")
		return
	}

	e := global.GVA_DB
	affected, err := e.Id(user.Id).Update(mUser)
	if affected <= 0 || err != nil {
		log.Print(err)
		log.Printf("数据库更新失败")
		utils.MakeErrorRes(ctx, model.OtherErrorCode, model.OptionFailur, "数据库更新失败")
		return
	}

	utils.MakeSuccessRes(ctx, model.Success, nil)

}
