package workflow

import (
	"github.com/kataras/iris/v12"
	"mayday/src/global"
	UserModel "mayday/src/model/user"
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

//创建表单草稿
func CreateTableDraft(ctx iris.Context, tableDraftReq workflow.TableDraftReq) {

	user := ctx.Values().Get("user").(UserModel.SdUser)

	sdTableDraft := tableDraftReq.GetSdTableDraft()
	sdTableDraft.UserId = user.Id
	e := global.GVA_DB
	effect, err := e.Insert(sdTableDraft)
	if effect <= 0 || err != nil {
		utils.Responser.FailWithMsg(ctx, "表单草稿插入失败", err)
		return
	}
	utils.Responser.Ok(ctx)
}
