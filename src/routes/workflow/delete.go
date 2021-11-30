package workflow_routes

import (
	_ "encoding/json"
	//"mayday/middleware/jwts"
	"log"
	"mayday/src/db/conn"

	//"time"
	"mayday/src/models"
	"mayday/src/supports/responser"

	"github.com/kataras/iris/v12"
)

// swagger:operation POST /workflow/delete/workflow workflow delete_workflow
// ---
// summary: 删除流程（已发布）
// description: 删除流程（已发布）
// parameters:
// - name: id
//   description: 流程ID
//   type: int
//   required: true
func Workflow_delete_workflow(ctx iris.Context) {

	var workflow model.SdWorkflow
	if err := ctx.ReadJSON(&workflow); err != nil {
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		log.Print("数据接收失败")
		return
	}
	workflow.IsDeleted = 1
	e := conn.MasterEngine()
	affected, err := e.Id(workflow.Id).Cols("is_deleted").Update(workflow)
	if affected <= 0 || err != nil {
		log.Print(err)
		log.Printf("流程删除失败")
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	responser.MakeSuccessRes(ctx, model.Success, nil)
}

// swagger:operation POST /workflow/delete/table table delete_table
// ---
// summary: 删除表单（已发布）
// description: 删除表单（已发布）
// parameters:
// - name: id
//   description: 表单ID
//   type: int
//   required: true
func Workflow_delete_table(ctx iris.Context) {
	log.Print("删除流程表单1")
	var table model.SdTable
	if err := ctx.ReadJSON(&table); err != nil {
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		log.Print("数据接收失败")
		return
	}
	e := conn.MasterEngine()
	effect, err := e.Id(table.Id).Delete(table)
	if effect <= 0 || err != nil {
		log.Print(err)
		log.Printf("数据库操作失败")
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	responser.MakeSuccessRes(ctx, model.Success, nil)

}

// swagger:operation POST /workflow/delete/workflow-draft workflow delete_workflow-draft
// ---
// summary: 删除流程（草稿）
// description: 删除流程（草稿）
// parameters:
// - name: id
//   description: 流程ID
//   type: int
//   required: true
func Workflow_delete_workflow_draft(ctx iris.Context) {

	var workflowDraft model.SdWorkflowDraft
	if err := ctx.ReadJSON(&workflowDraft); err != nil {
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		log.Print("数据接收失败")
		return
	}
	workflowDraft.IsDeleted = 1

	e := conn.MasterEngine()
	affected, err := e.Id(workflowDraft.Id).Cols("is_deleted").Update(workflowDraft)
	if affected <= 0 || err != nil {
		log.Print(err)
		log.Printf("流程删除失败")
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	responser.MakeSuccessRes(ctx, model.Success, nil)
}

// swagger:operation POST /workflow/delete/table-draft table delete_table-draft
// ---
// summary: 删除表单（草稿）
// description: 删除表单（草稿）
// parameters:
// - name: id
//   description: 表单草稿ID
//   type: int
//   required: true
func Workflow_delete_table_draft(ctx iris.Context) {
	log.Print("删除流程表单草稿")
	var table model.SdTableDraft
	if err := ctx.ReadJSON(&table); err != nil {
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		log.Print("数据接收失败")
		return
	}
	e := conn.MasterEngine()
	effect, err := e.Id(table.Id).Delete(table)
	if effect <= 0 || err != nil {
		log.Print(err)
		log.Printf("数据库操作失败")
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	responser.MakeSuccessRes(ctx, model.Success, nil)
}
