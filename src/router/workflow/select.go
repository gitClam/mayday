package workflow_routes

import (
	_ "encoding/json"
	"log"
	"mayday/src/global"
	"mayday/src/middleware"
	workflowModel "mayday/src/model/workflow"
	"mayday/src/utils"

	"github.com/kataras/iris/v12"
)

// swagger:operation GET /workflow/select/workflow workflow select_workflow
// ---
// summary: 查询流程（已发布）
// description: 查询流程（已发布）
func WorkflowSelectWorkflow(ctx iris.Context) {

	var workflows []workflowModel.SdWorkflow

	e := global.GVA_DB
	err := e.Sql("select * from sd_workflow where is_deleted != 1").Find(&workflows)
	if err != nil {
		log.Print(err)
		log.Printf("流程查询失败")
		utils.Responser.FailWithMsg(ctx, "")
		return
	}
	utils.Responser.OkWithDetails(ctx, utils.Success, workflows)
}

// swagger:operation GET /workflow/select/workflow-byId workflow select_workflow_byId
// ---
// summary: 查询流程（已发布）
// description: 查询流程（已发布）
// parameters:
// - name: id
//   description: 流程ID
//   type: int
//   required: true
func WorkflowSelectWorkflowById(ctx iris.Context) {

	var workflow workflowModel.SdWorkflow
	if err := ctx.ReadForm(&workflow); err != nil {
		log.Print(err)
		utils.Responser.FailWithMsg(ctx, "")
		log.Print("数据接收失败")
		return
	}
	e := global.GVA_DB
	has, err := e.Id(workflow.Id).Get(&workflow)
	if !has || err != nil {
		log.Print(err)
		log.Printf("流程查询失败")
		utils.Responser.FailWithMsg(ctx, "")
		return
	}
	utils.Responser.OkWithDetails(ctx, utils.Success, workflow)
}

// swagger:operation Post /workflow/select/table table select_table
// ---
// summary: 查询表单（已发布）
// description: 查询表单（已发布）
// parameters:
// - name: id
//   description: 表单ID
//   type: int
//   required: true
func WorkflowSelectTable(ctx iris.Context) {
	var table workflowModel.SdTable
	if err := ctx.ReadJSON(&table); err != nil {
		log.Print(err)
		utils.Responser.FailWithMsg(ctx, "")
		log.Print("数据接收失败")
		return
	}
	e := global.GVA_DB
	has, err := e.Id(table.Id).Get(&table)
	if !has || err != nil {
		log.Print(err)
		log.Printf("流程查询失败")
		utils.Responser.FailWithMsg(ctx, "")
		return
	}
	utils.Responser.OkWithDetails(ctx, utils.Success, table)
}

// swagger:operation GET /workflow/select/workflow-draft workflow select_workflow-draft
// ---
// summary: 查询流程（草稿）
// description: 查询流程（草稿）
// parameters:
// - name: id
//   description: 流程ID
//   type: int
//   required: true
func WorkflowSelectWorkflowDraft(ctx iris.Context) {
	user, ok := middleware.ParseToken(ctx)
	if !ok {
		log.Printf("解析TOKEN出错，请重新登录")
		utils.Responser.FailWithMsg(ctx, "解析TOKEN出错，请重新登录")
		return
	}

	var workflowDrafts []workflowModel.SdWorkflowDraft

	e := global.GVA_DB
	err := e.Sql("select * from sd_workflow_draft where owner_id = ? and is_deleted != 1", user.Id).Find(&workflowDrafts)
	if err != nil {
		log.Print(err)
		log.Printf("流程草稿查询失败")
		utils.Responser.FailWithMsg(ctx, "")
		return
	}
	utils.Responser.OkWithDetails(ctx, utils.Success, workflowDrafts)
}

// swagger:operation Post /workflow/select/table-draft table select_table-draft
// ---
// summary: 查询表单（草稿）
// description: 查询表单（草稿）
// parameters:
// - name: id
//   description: 表单草稿ID
//   type: int
//   required: true
func WorkflowSelectTableDraft(ctx iris.Context) {
	var tableDraft workflowModel.SdTableDraft
	if err := ctx.ReadJSON(&tableDraft); err != nil {
		log.Print(err)
		utils.Responser.FailWithMsg(ctx, "")
		log.Print("数据接收失败")
		return
	}
	e := global.GVA_DB
	has, err := e.Id(tableDraft.Id).Get(&tableDraft)
	if !has || err != nil {
		log.Print(err)
		log.Printf("流程查询失败")
		utils.Responser.FailWithMsg(ctx, "")
		return
	}
	utils.Responser.OkWithDetails(ctx, utils.Success, tableDraft)
}

// swagger:operation Post /workflow/select/table-workSpace table select_table_draft_workSpace_Workflow
// ---
// summary: 查询工作空间拥有的表单
// description: 查询工作空间拥有的表单
// parameters:
// - name: workspace_id
//   description: 工作空间ID
//   type: int
//   required: true
func WorkflowSelectTableDraftWorkspace(ctx iris.Context) {
	var workflow workflowModel.SdWorkflow
	if err := ctx.ReadJSON(&workflow); err != nil {
		log.Print(err)
		utils.Responser.FailWithMsg(ctx, "")
		log.Print("数据接收失败")
		return
	}
	var table []workflowModel.SdTable
	e := global.GVA_DB
	err := e.Where("workspace_id = ?", workflow.Id).Find(&table)
	if err != nil {
		log.Print(err)
		log.Printf("工作空间流程草稿查询失败")
		utils.Responser.FailWithMsg(ctx, "")
		return
	}
	utils.Responser.OkWithDetails(ctx, utils.Success, table)
}

// swagger:operation Post /workflow/select/table-draft-user table select_table_draft_user_Workflow_select
// ---
// summary: 查询用户拥有的表单（草稿）
// description: 查询用户拥有查询表单（草稿）
func WorkflowSelectTableDraftUser(ctx iris.Context) {
	user, ok := middleware.ParseToken(ctx)
	if !ok {
		log.Printf("解析TOKEN出错，请重新登录")
		utils.Responser.FailWithMsg(ctx, "解析TOKEN出错，请重新登录")
		return
	}
	var tableDraft []workflowModel.SdTableDraft
	e := global.GVA_DB
	err := e.Where("user_id = ?", user.Id).Find(&tableDraft)
	if err != nil {
		log.Print(err)
		log.Printf("个人流程草稿查询失败")
		utils.Responser.FailWithMsg(ctx, "")
		return
	}
	utils.Responser.OkWithDetails(ctx, utils.Success, tableDraft)
}
