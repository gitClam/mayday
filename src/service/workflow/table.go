package workflowService

import (
	"github.com/kataras/iris/v12"
	"mayday/src/global"
	"mayday/src/model/common/resultcode"
	UserModel "mayday/src/model/user"
	WorkflowModel "mayday/src/model/workflow"
	"mayday/src/utils"
)

//创建表单
func CreateTable(ctx iris.Context, tableReq WorkflowModel.CreateTableReq) {
	//TODO 权限检查
	sdTable := tableReq.GetSdTable()

	e := global.GVA_DB
	effect, err := e.Insert(sdTable)
	if effect <= 0 || err != nil {
		utils.Responser.Fail(ctx, resultcode.DataCreateFail, err)
		return
	}
	utils.Responser.Ok(ctx)
}

//创建表单草稿
func CreateTableDraft(ctx iris.Context, tableDraftReq WorkflowModel.CreateTableDraftReq) {

	user := ctx.Values().Get("user").(UserModel.SdUser)

	sdTableDraft := tableDraftReq.GetSdTableDraft()
	sdTableDraft.UserId = user.Id
	e := global.GVA_DB
	effect, err := e.Insert(sdTableDraft)
	if effect <= 0 || err != nil {
		utils.Responser.Fail(ctx, resultcode.DataCreateFail, err)
		return
	}
	utils.Responser.Ok(ctx)
}

//删除表单
func DeleteTable(ctx iris.Context, id []int) {
	//TODO 权限检查
	e := global.GVA_DB
	effect, err := e.Id(id).Delete(new(WorkflowModel.SdTable))
	if effect <= 0 || err != nil {
		utils.Responser.Fail(ctx, resultcode.DataDeleteFail, err)
		return
	}
	utils.Responser.Ok(ctx)
}

//删除表单草稿
func DeleteTableDraft(ctx iris.Context, id []int) {
	user := ctx.Values().Get("user").(UserModel.SdUser)
	var sdTableDrafts []WorkflowModel.SdTableDraft
	e := global.GVA_DB

	err := e.Id(id).Find(&sdTableDrafts)
	if err != nil {
		utils.Responser.Fail(ctx, resultcode.DataSelectFail, err)
		return
	}

	// 只能删除自己的草稿
	for _, sdTableDraft := range sdTableDrafts {
		if sdTableDraft.UserId != user.Id {
			utils.Responser.Fail(ctx, resultcode.PermissionsLess, err)
			return
		}
	}

	effect, err := e.Id(id).Delete(new(WorkflowModel.SdTableDraft))
	if effect <= 0 || err != nil {
		utils.Responser.Fail(ctx, resultcode.DataDeleteFail, err)
		return
	}
	utils.Responser.Ok(ctx)
}

//获取表单信息
func GetTableById(ctx iris.Context, id []int) {
	//TODO 验证权限
	var sdTables []WorkflowModel.SdTable
	e := global.GVA_DB
	err := e.Id(id).Find(&sdTables)
	if err != nil {
		utils.Responser.Fail(ctx, resultcode.DataSelectFail, err)
		return
	}
	utils.Responser.OkWithDetails(ctx, sdTables)
}

//获取表单信息
func GetTableByWorkspaceId(ctx iris.Context, id []int) {
	//TODO 验证权限
	var sdTables []WorkflowModel.SdTable
	e := global.GVA_DB
	err := e.In("workspace_id", id).Find(&sdTables)
	if err != nil {
		utils.Responser.Fail(ctx, resultcode.DataSelectFail, err)
		return
	}
	utils.Responser.OkWithDetails(ctx, sdTables)
}

//获取用户的表单草稿信息
func GetTableDraftByUser(ctx iris.Context) {
	user := ctx.Values().Get("user").(UserModel.SdUser)
	var sdTableDrafts WorkflowModel.SdTableDraft

	e := global.GVA_DB
	err := e.Where("user_id = ?", user.Id).Find(&sdTableDrafts)
	if err != nil {
		utils.Responser.Fail(ctx, resultcode.DataSelectFail, err)
		return
	}

	utils.Responser.OkWithDetails(ctx, sdTableDrafts)
}

//获取表单草稿信息
func GetTableDraftById(ctx iris.Context, id []int) {
	//TODO 验证权限
	var sdTableDraft []WorkflowModel.SdTableDraft
	e := global.GVA_DB
	err := e.Id(id).Find(&sdTableDraft)
	if err != nil {
		utils.Responser.Fail(ctx, resultcode.DataSelectFail, err)
		return
	}
	utils.Responser.OkWithDetails(ctx, sdTableDraft)
}

//修改表单
func UpdateTable(ctx iris.Context, tableReq WorkflowModel.UpdateTableReq) {
	//TODO 验证权限
	e := global.GVA_DB
	sdTable := tableReq.GetSdTable()
	effect, err := e.Id(sdTable.Id).Update(sdTable)
	if effect <= 0 || err != nil {
		utils.Responser.Fail(ctx, resultcode.DataUpdateFail, err)
		return
	}
	utils.Responser.Ok(ctx)
}

func UpdateTableDraft(ctx iris.Context, tableDraftReq WorkflowModel.UpdateTableDraftReq) {
	//TODO 验证权限
	e := global.GVA_DB
	sdTableDraft := tableDraftReq.GetSdTableDraft()
	effect, err := e.Id(sdTableDraft.Id).Update(sdTableDraft)
	if effect <= 0 || err != nil {
		utils.Responser.Fail(ctx, resultcode.DataUpdateFail, err)
		return
	}
	utils.Responser.Ok(ctx)
}
