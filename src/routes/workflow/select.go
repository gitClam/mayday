package workflow_routes

import (
	_ "encoding/json"
	"log"
	"mayday/middleware/jwts"
	"mayday/src/db/conn"

	//"time"
	"mayday/src/models"
	"mayday/src/supports/responser"

	"github.com/kataras/iris/v12"
)

// swagger:operation GET /workflow/select/workflow workflow select_workflow
// ---
// summary: 查询流程（已发布）
// description: 查询流程（已发布）
func Workflow_select_workflow(ctx iris.Context) {

	var workflows []model.SdWorkflow

	e := conn.MasterEngine()
	err := e.Sql("select * from sd_workflow where is_deleted != 1").Find(&workflows)
	if err != nil {
		log.Print(err)
		log.Printf("流程查询失败")
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	responser.MakeSuccessRes(ctx, model.Success, workflows)
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
func Workflow_select_workflow_byId(ctx iris.Context) {

	var workflow model.SdWorkflow
	if err := ctx.ReadForm(&workflow); err != nil {
		log.Print(err)
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		log.Print("数据接收失败")
		return
	}
	e := conn.MasterEngine()
	has, err := e.Id(workflow.Id).Get(&workflow)
	if !has || err != nil {
		log.Print(err)
		log.Printf("流程查询失败")
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	responser.MakeSuccessRes(ctx, model.Success, workflow)
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
func Workflow_select_table(ctx iris.Context) {
	var table model.SdTable
	if err := ctx.ReadJSON(&table); err != nil {
		log.Print(err)
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		log.Print("数据接收失败")
		return
	}
	e := conn.MasterEngine()
	has, err := e.Id(table.Id).Get(&table)
	if !has || err != nil {
		log.Print(err)
		log.Printf("流程查询失败")
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	responser.MakeSuccessRes(ctx, model.Success, table)
}

// swagger:operation GET /workflow/select/workflow-draft workflow select_workflow-draft
// ---
// summary: 查询流程（草稿）
// description: 查询流程（草稿）

func Workflow_select_workflow_draft(ctx iris.Context) {
	user, ok := jwts.ParseToken(ctx)
	if !ok {
		log.Printf("解析TOKEN出错，请重新登录")
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.TokenParseFailur, nil)
		return
	}

	var workflowDrafts []model.SdWorkflowDraft

	e := conn.MasterEngine()
	err := e.Sql("select * from sd_workflow_draft where owner_id = ? and is_deleted != 1", user.Id).Find(&workflowDrafts)
	if err != nil {
		log.Print(err)
		log.Printf("流程草稿查询失败")
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	responser.MakeSuccessRes(ctx, model.Success, workflowDrafts)
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
func Workflow_select_table_draft(ctx iris.Context) {
	var tableDraft model.SdTableDraft
	if err := ctx.ReadJSON(&tableDraft); err != nil {
		log.Print(err)
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		log.Print("数据接收失败")
		return
	}
	e := conn.MasterEngine()
	has, err := e.Id(tableDraft.Id).Get(&tableDraft)
	if !has || err != nil {
		log.Print(err)
		log.Printf("流程查询失败")
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	responser.MakeSuccessRes(ctx, model.Success, tableDraft)
}

// swagger:operation Post /workflow/select/table-workSpace table Workflow_select_table_draft_workSpace
// ---
// summary: 查询工作空间拥有的表单
// description: 查询工作空间拥有的表单
// parameters:
// - name: workspace_id
//   description: 工作空间ID
//   type: int
//   required: true
func Workflow_select_table_draft_workSpace(ctx iris.Context) {
	var workflow model.SdWorkflow
	if err := ctx.ReadJSON(&workflow); err != nil {
		log.Print(err)
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		log.Print("数据接收失败")
		return
	}
	var table []model.SdTable
	e := conn.MasterEngine()
	has, err := e.Where("workspace_id = ?", workflow.Id).Find(&table)
	if !has || err != nil {
		log.Print(err)
		log.Printf("工作空间流程草稿查询失败")
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	responser.MakeSuccessRes(ctx, model.Success, table)
}

// swagger:operation Post /workflow/select/table-draft-user table Workflow_select_table_draft_user
// ---
// summary: 查询用户拥有的表单（草稿）
// description: 查询用户拥有查询表单（草稿）
func Workflow_select_table_draft_user(ctx iris.Context) {
	user, ok := jwts.ParseToken(ctx)
	if !ok {
		log.Printf("解析TOKEN出错，请重新登录")
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.TokenParseFailur, nil)
		return
	}
	var tableDraft []model.SdTableDraft
	e := conn.MasterEngine()
	has, err := e.Where("user_id = ?", user.Id).Find(&tableDraft)
	if !has || err != nil {
		log.Print(err)
		log.Printf("个人流程草稿查询失败")
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	responser.MakeSuccessRes(ctx, model.Success, tableDraft)
}
