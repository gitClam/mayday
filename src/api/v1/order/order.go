package order

import (
	"github.com/kataras/iris/v12"
	"mayday/src/model/common/resultcode"
	userModel "mayday/src/model/user"
	"mayday/src/service/order"
	"mayday/src/utils"
)

func CreateOrder(ctx iris.Context) {
	user := ctx.Values().Get("user").(userModel.SdUser)
	err := order.CreateOrderService(ctx, &user)
	if err != nil {
		utils.Responser.Fail(ctx, resultcode.DataCreateFail, err)
		return
	}
	utils.Responser.Ok(ctx)
}

func GetOrderNotification(ctx iris.Context) {

}

func GetOrderState(ctx iris.Context) {

}
