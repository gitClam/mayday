package user

import (
	"github.com/kataras/iris/v12/core/router"
	userSever "mayday/src/service/user"
)

func InitUserRouter(Router router.Party) {

	user := Router.Party("/user")
	{
		user.Post("/registe", userSever.UserRegister)            //用户注册
		user.Post("/login", userSever.UserLogin)                 //用户登录
		user.Delete("/cancellation", userSever.UserCancellation) //用户注销
		user.Post("/editor/message", userSever.SetUserMessage)   //修改用户信息
		user.Get("/message", userSever.UserMessage)              //获取用户信息
		user.Get("/photo/{id:int}", userSever.UserPhoto)         //获取用户头像
		user.Post("/set_photo", userSever.SetUserPhoto)          //设置用户头像头像
	}
}
