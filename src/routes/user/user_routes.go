package user_routes

import (
	"github.com/kataras/iris/v12"
	//"strconv"
	"time"
	"log"
	"io"
	"os"
	"strconv"
	"mayday/src/db/conn"
	"mayday/src/models"
	"mayday/src/supports/responser"
	"mayday/src/supports/responser/vo"
	"mayday/middleware/jwts"
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

func User_registe(ctx iris.Context) {
	var user model.SdUser
	if err := ctx.ReadForm(&user); err != nil {
		responser.MakeErrorRes(ctx,iris.StatusInternalServerError, model.RegisteFailur , nil)
		log.Print("用户注册失败，数据接收失败")
		return 
	}
	log.Print(user)
	user.Photo = "./data/photo/2.png"
	user.IsDeleted = 0;
	user.CreateDate = model.LocalTime(time.Now())
	log.Print(user)
	e := conn.MasterEngine()	
	effect, err := e.Insert(user)	
	if effect <= 0 || err != nil {
		log.Printf("用户注册失败。")
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.RegisteFailur, nil)
		return 
	} 
	
	responser.MakeSuccessRes(ctx,model.Success,nil)
	log.Print("ok")
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
func User_login(ctx iris.Context) {
	log.Print("13526")
	
	var user model.SdUser
	if err := ctx.ReadForm(&user); err != nil {
		responser.MakeErrorRes(ctx,iris.StatusInternalServerError, model.LoginFailur , nil)
		log.Print("用户登录失败，数据接收失败")
		return 
	}
	
	if(user.Mail == "" || user.Password == ""){
		responser.MakeErrorRes(ctx,3333, model.LoginFailur , nil)
		log.Print("用户登录失败,邮箱或密码为空")
		return 
	}	
	
	var mUser model.SdUser
	mUser.Mail = user.Mail
	
	e := conn.MasterEngine()
	has, err := e.Where("is_deleted != 1").Get(&mUser)
	if ( !has || err != nil || mUser.IsDeleted == 1) {
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.LoginFailur, nil)
		log.Printf("数据库查询错误或用户名不存在")
		return 
	} 
	
	log.Print(mUser)
	
	if mUser.Password != user.Password{
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.LoginFailur, nil)
		log.Printf("用户密码错误" )
		return
	}
	
	token, err := jwts.GenerateToken(&mUser);
	log.Printf("用户[%s], 登录生成token [%s]", mUser.Name, token)
	if err != nil {
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.TokenCreateFailur, nil)
		log.Printf("数据库查询错误或用户名不存在")
		return
	}
	
	responser.MakeSuccessRes(ctx,model.Success,vo.TansformUserVOToken(token,&mUser))
	log.Print("ok")
}

// swagger:operation GET /user/photo/{id:int} user get_photo
// --- 
// summary: 获取用户头像
// description: 获取用户头像

func User_photo(ctx iris.Context) {

	var user model.SdUser
	Id , err := strconv.Atoi(ctx.Params().Get("id"))
	if(err != nil){
		log.Printf("ID获取失败")
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return 
	}
	user.Id = Id
	e := conn.MasterEngine()
	has, err := e.Get(&user)
	if ( !has || err != nil) {
		log.Printf("数据库查询错误或用户名不存在")
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return 
	} 
	
	if user.Photo == ""{
		log.Printf("用户头像获取出错")
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return 
	}
	err1 := ctx.ServeFile(user.Photo, false)
	if err1 != nil {
		log.Printf("用户头像文件读取错误")
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
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
func Set_user_photo(ctx iris.Context){
	log.Print("修改用户头像")
	user, ok := jwts.ParseToken(ctx)
	if !ok {
		log.Printf("解析TOKEN出错，请重新登录")
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.TokenParseFailur, nil)
		return
	}
	log.Print(user)
	var mUser model.SdUser
	mUser.Id = user.Id
	mUser.Name = user.Name
	e := conn.MasterEngine()
	has, err := e.Get(&mUser)
	if ( !has || err != nil) {
		
		log.Printf("数据库查询错误或用户名不存在") 
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return 
	}  
	
	log.Print("接收图片中")
	file, _, err := ctx.FormFile("UserPhoto")
	if err != nil {
		log.Print("图片文件不存在")
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	defer file.Close()
	//fname := info.Filename
	if(mUser.Photo == "./data/photo/2.png" || mUser.Photo == ""){
		mUser.Photo = "./data/photo/" +  mUser.Mail
		affected, err := e.Id(mUser.Id).Update(mUser)
		if ( affected <= 0 || err != nil) {
		log.Printf("数据库更新失败") 
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return 
		}  
	}
	
	out, err := os.OpenFile(mUser.Photo, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Print("文件打开失败")
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	defer out.Close()
	
	io.Copy(out, file)
	
	responser.MakeSuccessRes(ctx,model.Success,nil)	
	
	log.Print("图片已保存")
}
// swagger:operation Delete /user/cancellation user cancellation
// --- 
// summary: 用户注销
// description: 用户注销

func User_cancellation(ctx iris.Context) {
	user, ok := jwts.ParseToken(ctx)
	if !ok {
		log.Printf("解析TOKEN出错，请重新登录")
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.TokenParseFailur, nil)
		return
	}
	var mUser model.SdUser
	mUser.Id = user.Id
	e := conn.MasterEngine()
	log.Print(mUser.Id)
	has, err := e.Id(mUser.Id).Get(&mUser)
	if ( !has || err != nil) {
		log.Print(err)
		log.Printf("数据库查询错误或用户名不存在") 
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return 
	} 
	mUser.IsDeleted = 1
	affected, err1 := e.Id(mUser.Id).Update(&mUser)
	if affected <= 0 || err1 != nil{
		log.Print(err)
		log.Printf("数据库修改失败")
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return 
	}
	responser.MakeSuccessRes(ctx,model.Success,nil)
}

// swagger:operation GET /user/message user message
// --- 
// summary: 获取用户信息
// description: 获取用户信息


func User_message(ctx iris.Context) {

	user, ok := jwts.ParseToken(ctx)
	if !ok {
		log.Printf("解析TOKEN出错，请重新登录")
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.TokenParseFailur, nil)
		return
	}
	var mUser model.SdUser
	mUser.Id = user.Id
	mUser.Name = user.Name

	e := conn.MasterEngine()
	has, err := e.Where(" is_deleted != 1 ").Get(&mUser)
	if ( !has || err != nil) {
		log.Printf("数据库查询错误或用户名不存在") 
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return 
	} 
	
	responser.MakeSuccessRes(ctx,model.Success,vo.TansformUserVO(&mUser))
	
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

func Set_User_message(ctx iris.Context) {

	user, ok := jwts.ParseToken(ctx)
	if !ok {
		log.Printf("解析TOKEN出错，请重新登录")
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.TokenParseFailur, nil)
		return
	}
	
	var mUser model.SdUser
	if err := ctx.ReadForm(&mUser); err != nil {
		log.Print(err)
		responser.MakeErrorRes(ctx,model.OtherErrorCode, model.OptionFailur , nil)
		log.Print("数据接收失败")
		return 
	}
	
	e := conn.MasterEngine()
	affected, err := e.Id(user.Id).Update(mUser)
	if ( affected <= 0 || err != nil) {
		log.Print(err)
		log.Printf("数据库更新失败") 
		responser.MakeErrorRes(ctx, model.OtherErrorCode, model.OptionFailur, nil)
		return 
	} 
	
	responser.MakeSuccessRes(ctx,model.Success,nil)
	
}