package workflow

import (
	"github.com/kataras/iris/v12"
	"mayday/src/global"
	"mayday/src/model/workflow"
	"mayday/src/utils"
)

//创建表单
func CreateTable(ctx iris.Context, tableReq workflow.TableReq) {
	sdTable := tableReq.GetSdTable()
	e := global.GVA_DB
	effect, err := e.Insert(sdTable)
	if effect <= 0 || err != nil {
		utils.Responser.FailWithMsg(ctx, "表单插入失败", err)
		return
	}
	utils.Responser.Ok(ctx)
}
