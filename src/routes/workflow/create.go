package workflow_routes

import (
	_ "encoding/json"
	"log"
	"mayday/middleware/jwts"
	"mayday/src/db/conn"
	"mayday/src/models"
	"mayday/src/supports/responser"
	"time"

	"github.com/kataras/iris/v12"
)

// swagger:operation POST /workflow/create/workflow workflow create_workflow
// ---
// summary: 创建流程（发布）
// description: 创建流程（发布）
// parameters:
// - name: Name
//   description: 流程名字
//   type: json
//   required: true
// - name: Structure
//   description: 流程结构
//   type: json
//   required: true
// - name: Tables
//   description: 表单样式
//   type: json
//   required: true
// - name: Remarks
//   description: 备注
//   type: string
//   required: false
// - name: IsStart
//   description: 是否开启(默认为开启)
//   type: bool
//   required: false
func Workflow_create_workflow(ctx iris.Context) {
	log.Print("创建流程（发布）")
	user, ok := jwts.ParseToken(ctx)
	if !ok {
		log.Printf("解析TOKEN出错，请重新登录")
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.TokenParseFailur, nil)
		return
	}
	log.Print(user)
	var workflow model.SdWorkflow
	if err := ctx.ReadJSON(&workflow); err != nil {
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		log.Print("数据接收失败")
		return
	}

	workflow.IsDeleted = 0
	workflow.CreateTime = model.LocalTime(time.Now())
	workflow.CreateUser = user.Id
	e := conn.MasterEngine()
	effect, err := e.Insert(workflow)
	if effect <= 0 || err != nil {
		log.Print(err)
		log.Printf("流程创建失败")
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	responser.MakeSuccessRes(ctx, model.Success, nil)
}

// swagger:operation POST /workflow/create/table table create_table
// ---
// summary: 创建表单（发布）
// description: 创建表单（发布）
// parameters:
// - name: WorkspaceId
//   description: 要发布的工作空间的id
//   type: int
//   required: true
// - name: Data
//   description: 表单的具体数据
//   type: json(string)
//   required: true
// - name: Name
//   description: 表单的名字
//   type: string
//   required: true
func Workflow_create_table(ctx iris.Context) {
	log.Print("创建流程表单")
	var table model.SdTable
	if err := ctx.ReadJSON(&table); err != nil {
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		log.Print("数据接收失败")
		return
	}
	e := conn.MasterEngine()
	effect, err := e.Insert(table)
	if effect <= 0 || err != nil {
		log.Print(err)
		log.Printf("数据库操作失败")
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	responser.MakeSuccessRes(ctx, model.Success, nil)

}

// swagger:operation POST /workflow/create/workflow-draft workflow create_workflow-draft
// ---
// summary: 创建流程（草稿）
// description: 创建流程（草稿）
// parameters:
// - name: Structure
//   description: 流程结构
//   type: json(string)
//   required: true
func Workflow_create_workflow_draft(ctx iris.Context) {
	log.Print("创建流程（草稿）")
	user, ok := jwts.ParseToken(ctx)
	if !ok {
		log.Printf("解析TOKEN出错，请重新登录")
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.TokenParseFailur, nil)
		return
	}
	log.Print(user.Name)
	var workflowDraft model.SdWorkflowDraft
	if err := ctx.ReadJSON(&workflowDraft); err != nil {
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		log.Print("数据接收失败")
		return
	}

	workflowDraft.IsDeleted = 0
	workflowDraft.OwnerId = user.Id

	e := conn.MasterEngine()
	effect, err := e.Insert(workflowDraft)
	if effect <= 0 || err != nil {
		log.Print(err)
		log.Printf("流程创建失败")
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	responser.MakeSuccessRes(ctx, model.Success, nil)

}

// swagger:operation POST /workflow/create/table-draft table Workflow_create_table_draft-draft
// ---
// summary: 创建表单（草稿）
// description: 创建表单（草稿）
// parameters:
// - name: Data
//   description: 表单的具体数据
//   type: json(string)
//   required: true
// - name: Name
//   description: 表单的名字
//   type: string
//   required: true
func Workflow_create_table_draft(ctx iris.Context) {
	log.Print("创建表单（草稿）")
	user, ok := jwts.ParseToken(ctx)
	if !ok {
		log.Printf("解析TOKEN出错，请重新登录")
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.TokenParseFailur, nil)
		return
	}
	var tableDraft model.SdTableDraft
	if err := ctx.ReadJSON(&tableDraft); err != nil {
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		log.Print("数据接收失败")
		return
	}
	tableDraft.UserId = user.Id
	e := conn.MasterEngine()
	effect, err := e.Insert(tableDraft)
	if effect <= 0 || err != nil {
		log.Print(err)
		log.Printf("数据库操作失败")
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	responser.MakeSuccessRes(ctx, model.Success, nil)
}
