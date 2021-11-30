package workflow_routes
import (
	"github.com/kataras/iris/v12"
	//"strconv"
	//"time"
	"log"
	//"io"
	//"os"
	//"strconv"
	//"mayday/src/db/conn"
	"mayday/src/models"
	"mayday/src/supports/responser"
	//"mayday/src/supports/responser/vo"
	"mayday/middleware/jwts"
	"mayday/src/routes/workflow/order"
)
// swagger:operation POST /workflow/order/create-order workflow Workflow_order_create_order
// --- 
// summary: 创建流程申请
// description: 创建流程申请
// parameters:
// - name: UserPhoto 
//   description: 用户头像
//   type: file
//   required: true
func Workflow_order_create_order(ctx iris.Context) {
	//检查请求的用户（检测TOKEN类的东西）            
	user, ok := jwts.ParseToken(ctx)
	if !ok {
		log.Printf("解析TOKEN出错，请重新登录")
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.TokenParseFailur, nil)
		return
	}
	err := order.Create_order(ctx , user)
	if(err != nil){
		log.Print(err)
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	responser.MakeSuccessRes(ctx,model.Success,nil)
}
// swagger:operation POST /workflow/order/fill-table workflow Workflow_order_fill_table
// --- 
// summary: 填写表单（会修改流程状态）
// description: 填写表单（会修改流程状态）
// parameters:
// - name: UserPhoto
//   description: 用户头像
//   type: file
//   required: true
func Workflow_order_fill_table(ctx iris.Context) {
	
}
// swagger:operation POST /workflow/order/notification workflow Workflow_order_notification
// --- 
// summary: 获取消息提醒
// description: 获取消息提醒
// parameters:
// - name: UserPhoto
//   description: 用户头像
//   type: file
//   required: true
func Workflow_order_notification(ctx iris.Context) {
	
}
// swagger:operation POST /workflow/order/order-state workflow Workflow_order_order_state
// --- 
// summary: 获取流程状态		
// description: 获取流程状态		
// parameters:
// - name: Id
//   description: 流程ID
//   type: int
//   required: true
func Workflow_order_order_state(ctx iris.Context) {
	
}