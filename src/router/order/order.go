package order

import (
	"github.com/kataras/iris/v12/core/router"
	orderApi "mayday/src/api/v1/order"
)

func InitOrderRouter(Router router.Party) {
	user := Router.Party("/order")
	{
		user.Get("/state", orderApi.GetOrderState)               //获取事件状态
		user.Post("/create", orderApi.CreateOrder)               //创建事件
		user.Get("/notification", orderApi.GetOrderNotification) //获取代办
		user.Post("/handle", orderApi.Handle)                    //事件处理
	}
}
