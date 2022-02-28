package workflow

import (
	"github.com/kataras/iris/v12"
	"mayday/src/global"
	UserModel "mayday/src/model/user"
	WorkflowModel "mayday/src/model/workflow"
	"mayday/src/utils"
)

//创建表单
func CreateTable(ctx iris.Context, tableReq WorkflowModel.TableReq) {
	//TODO 权限检查
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
func CreateTableDraft(ctx iris.Context, tableDraftReq WorkflowModel.TableDraftReq) {

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

//删除表单
func DeleteTable(ctx iris.Context, id []int) {
	//TODO 权限检查
	e := global.GVA_DB
	effect, err := e.Id(id).Delete(new(WorkflowModel.SdTable))
	if effect <= 0 || err != nil {
		utils.Responser.FailWithMsg(ctx, "表单删除失败", err)
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
		utils.Responser.FailWithMsg(ctx, "表单草稿不存在", err)
		return
	}

	// 只能删除自己的草稿
	for _, sdTableDraft := range sdTableDrafts {
		if sdTableDraft.UserId != user.Id {
			utils.Responser.FailWithMsg(ctx, "非法请求", err)
			return
		}
	}

	effect, err := e.Id(id).Delete(new(WorkflowModel.SdTableDraft))
	if effect <= 0 || err != nil {
		utils.Responser.FailWithMsg(ctx, "表单草稿删除失败", err)
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
		utils.Responser.FailWithMsg(ctx, "表单查询失败", err)
		return
	}
	utils.Responser.OkWithData(ctx, sdTables)
}

//获取用户的表单草稿信息
func GetTableDraftByUser(ctx iris.Context) {
	user := ctx.Values().Get("user").(UserModel.SdUser)
	var sdTableDrafts WorkflowModel.SdTableDraft

	e := global.GVA_DB
	err := e.Where("user_id = ?", user.Id).Find(&sdTableDrafts)
	if err != nil {
		utils.Responser.FailWithMsg(ctx, "表单草稿查询失败")
		return
	}

	utils.Responser.OkWithDetails(ctx, utils.Success, sdTableDrafts)
}

//获取表单草稿信息
func GetTableDraftById(ctx iris.Context, id []int) {
	//TODO 验证权限
	var sdTableDraft []WorkflowModel.SdTableDraft
	e := global.GVA_DB
	err := e.Id(id).Find(&sdTableDraft)
	if err != nil {
		utils.Responser.FailWithMsg(ctx, "表单草稿查询失败", err)
		return
	}
	utils.Responser.OkWithData(ctx, sdTableDraft)
}
