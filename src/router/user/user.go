package user

import (
	"github.com/kataras/iris/v12/core/router"
	userApi "mayday/src/api/v1/user"
)

func InitUserRouter(Router router.Party) {

	user := Router.Party("/user")
	{
		user.Post("/registe", userApi.Register)                  //用户注册
		user.Post("/login", userApi.Login)                       //用户登录
		user.Delete("/cancellation", userApi.Cancellation)       //用户注销
		user.Post("/editor/message", userApi.SetUserMessage)     //修改用户信息
		user.Get("/message", userApi.GetUserMessage)             //获取用户信息
		user.Get("/photo/{fileName:string}", userApi.GetPhoto)   //获取用户头像
		user.Post("/set_photo", userApi.SetPhoto)                //设置用户头像头像
		user.Get("/messageByUserId", userApi.GetUserMessageById) //
	}
}
