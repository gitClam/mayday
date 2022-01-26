package workflow_routes

import (
	"github.com/kataras/iris/v12"
	"mayday/src/middleware/jwts"

	//"strconv"
	//"time"
	"log"
	//"io"
	//"os"
	//"strconv"
	//"mayday/src/db/conn"
	"mayday/src/models"
	"mayday/src/supports/responser"

	"mayday/src/routes/workflow/order"
)

// swagger:operation POST /workflow/order/create-order workflow create_order_Workflow
// ---
// summary: 创建流程申请
// description: 创建流程申请
// parameters:
// - name: UserPhoto
//   description: 用户头像
//   type: file
//   required: true
func WorkflowOrderCreateOrder(ctx iris.Context) {
	//检查请求的用户
	user, ok := jwts.ParseToken(ctx)
	if !ok {
		log.Printf("解析TOKEN出错，请重新登录")
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.TokenParseFailur, nil)
		return
	}
	err := order.CreateOrder(ctx, user)
	if err != nil {
		log.Print(err)
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	responser.MakeSuccessRes(ctx, model.Success, nil)
}

// swagger:operation POST /workflow/order/fill-table workflow fill_table_Workflow_order_fill_table
// ---
// summary: 填写表单（会修改流程状态）
// description: 填写表单（会修改流程状态）
func WorkflowOrderFillTable(ctx iris.Context) {

}

// swagger:operation POST /workflow/order/notification workflow notification_Workflow_order
// ---
// summary: 获取待办提醒
// description: 获取待办提醒
func WorkflowOrderNotification(ctx iris.Context) {

}

// swagger:operation POST /workflow/order/order-state workflow order_state_Workflow_order
// ---
// summary: 获取流程状态
// description: 获取流程状态
// parameters:
// - name: Id
//   description: 流程ID
//   type: int
//   required: true
func WorkflowOrderOrderState(ctx iris.Context) {

}
