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

// swagger:operation POST /workflow/editor/workflow workflow editor_workflow
// ---
// summary: 修改流程（已发布）
// description: 修改流程（已发布）
// parameters:
// - name: id
//   description: 流程ID
//   type: int
//   required: true
func Workflow_editor_workflow(ctx iris.Context) {

	var workflow model.SdWorkflow
	if err := ctx.ReadJSON(&workflow); err != nil || workflow.Id == 0 {
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		log.Print("数据接收失败")
		return
	}

	e := conn.MasterEngine()
	affected, err := e.Id(workflow.Id).Update(workflow)
	if affected <= 0 || err != nil {
		log.Print(err)
		log.Printf("流程更新失败")
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	responser.MakeSuccessRes(ctx, model.Success, nil)
}

// swagger:operation POST /workflow/editor/workflow-state workflow editor_workflow_state
// ---
// summary: 修改流程状态（已发布）
// description: 修改流程状态（已发布）
// parameters:
// - name: id
//   description: 流程ID
//   type: int
//   required: true
func Workflow_editor_workflow_state(ctx iris.Context) {

	var workflow model.SdWorkflow
	if err := ctx.ReadForm(&workflow); err != nil || workflow.Id == 0 {
		log.Print(workflow.Id)
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		log.Print("数据接收失败")
		return
	}

	e := conn.MasterEngine()
	has, err := e.Id(workflow.Id).Get(&workflow)
	if !has || err != nil {
		log.Print(err)
		log.Printf("流程更新失败")
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}

	if workflow.IsStart == 0 {
		workflow.IsStart = 1
	} else {
		workflow.IsStart = 0
	}

	affected, err := e.Id(workflow.Id).Cols("is_start").Update(workflow)
	if affected <= 0 || err != nil {
		log.Print(err)
		log.Printf("流程更新失败")
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	responser.MakeSuccessRes(ctx, model.Success, nil)
}

// swagger:operation POST /workflow/editor/table table editor_table
// ---
// summary: 修改表单（已发布）
// description: 修改表单（已发布）
// parameters:
// - name: id
//   description: 表单ID
//   type: int
//   required: true
func Workflow_editor_table(ctx iris.Context) {
	log.Print("修改流程表单")
	var table model.SdTable
	if err := ctx.ReadJSON(&table); err != nil {
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		log.Print("数据接收失败")
		return
	}
	e := conn.MasterEngine()
	effect, err := e.Id(table.Id).Update(table)
	if effect <= 0 || err != nil {
		log.Print(err)
		log.Printf("数据库操作失败")
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	responser.MakeSuccessRes(ctx, model.Success, nil)
}

// swagger:operation POST /workflow/editor/workflow-draft workflow editor_workflow_draft
// ---
// summary: 修改流程（草稿）
// description: 修改流程（草稿）
// parameters:
// - name: id
//   description: 流程ID
//   type: int
//   required: true
func Workflow_editor_workflow_draft(ctx iris.Context) {

}

// swagger:operation POST /workflow/editor/table-draft table editor_table_draft
// ---
// summary: 修改表单（草稿）
// description: 修改表单（草稿）
// parameters:
// - name: id
//   description: 表单ID
//   type: int
//   required: true
func Workflow_editor_table_draft(ctx iris.Context) {
	log.Print("修改流程表单草稿")
	var table model.SdTableDraft
	if err := ctx.ReadJSON(&table); err != nil {
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		log.Print("数据接收失败")
		return
	}
	e := conn.MasterEngine()
	effect, err := e.Id(table.Id).Update(table)
	if effect <= 0 || err != nil {
		log.Print(err)
		log.Printf("数据库操作失败")
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	responser.MakeSuccessRes(ctx, model.Success, nil)
}
