package order

import (
	"github.com/kataras/iris/v12/core/router"
	orderApi "mayday/src/api/v1/order"
)

func InitOrderRouter(Router router.Party) {
	order := Router.Party("/order")
	{
		//order.Get("/state", orderApi.GetOrderState) //获取事件状态
		//order.Post("/create", orderApi.CreateOrder) //创建事件
		//order.Get("/notification", orderApi.GetOrderNotification) //获取代办
		order.Post("/process-structure", orderApi.ProcessStructure)
		order.Post("/create", orderApi.CreateOrder)      //创建流程实例
		order.Post("/list", orderApi.WorkOrderList)      //获取代办
		order.Post("/handle", orderApi.ProcessWorkOrder) //事件处理
		order.Get("/unity", orderApi.UnityWorkOrder)
		order.Post("/inversion", orderApi.InversionWorkOrder) // 转交工单
		//order.Get("/urge", orderApi.UrgeWorkOrder)
		//order.Post("/active-order/:id", orderApi.ActiveOrder)
		order.Delete("/delete/:id", orderApi.DeleteWorkOrder) // 删除工单
		//order.Post("/reopen/:id", orderApi.ReopenWorkOrder)
	}
}
