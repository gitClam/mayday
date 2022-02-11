package workflow_routes

import (
	_ "encoding/json"
	"log"
	"mayday/src/global"
	"mayday/src/middleware"
	"mayday/src/model/workflow"
	"mayday/src/utils"
	"time"

	"github.com/kataras/iris/v12"
)

func WorkflowCreateWorkflow(ctx iris.Context) {
	log.Print("创建流程（发布）")
	user, ok := middleware.ParseToken(ctx)
	if !ok {
		log.Printf("解析TOKEN出错，请重新登录")
		utils.Responser.FailWithMsg(ctx, "解析TOKEN出错，请重新登录")
		return
	}
	log.Print(user)
	var workflow workflow.SdWorkflow
	if err := ctx.ReadJSON(&workflow); err != nil {
		utils.Responser.FailWithMsg(ctx, "")
		log.Print("数据接收失败")
		return
	}

	workflow.IsDeleted = 0
	workflow.CreateTime = utils.LocalTime(time.Now())
	workflow.CreateUser = user.Id
	e := global.GVA_DB
	effect, err := e.Insert(workflow)
	if effect <= 0 || err != nil {
		log.Print(err)
		log.Printf("流程创建失败")
		utils.Responser.FailWithMsg(ctx, "")
		return
	}
	utils.Responser.Ok(ctx)
}

func WorkflowCreateTable(ctx iris.Context) {
	log.Print("创建流程表单")
	var table workflow.SdTable
	if err := ctx.ReadJSON(&table); err != nil {
		utils.Responser.FailWithMsg(ctx, "")
		log.Print("数据接收失败")
		return
	}
	e := global.GVA_DB
	effect, err := e.Insert(table)
	if effect <= 0 || err != nil {
		log.Print(err)
		log.Printf("数据库操作失败")
		utils.Responser.FailWithMsg(ctx, "")
		return
	}
	utils.Responser.Ok(ctx)
}

func WorkflowCreateWorkflowDraft(ctx iris.Context) {
	log.Print("创建流程（草稿）")
	user, ok := middleware.ParseToken(ctx)
	if !ok {
		log.Printf("解析TOKEN出错，请重新登录")
		utils.Responser.FailWithMsg(ctx, "解析TOKEN出错，请重新登录")
		return
	}
	log.Print(user.Name)
	var workflowDraft workflow.SdWorkflowDraft
	if err := ctx.ReadJSON(&workflowDraft); err != nil {
		utils.Responser.FailWithMsg(ctx, "")
		log.Print("数据接收失败")
		return
	}

	workflowDraft.IsDeleted = 0
	workflowDraft.OwnerId = user.Id

	e := global.GVA_DB
	effect, err := e.Insert(workflowDraft)
	if effect <= 0 || err != nil {
		log.Print(err)
		log.Printf("流程创建失败")
		utils.Responser.FailWithMsg(ctx, "")
		return
	}
	utils.Responser.Ok(ctx)

}

func WorkflowCreateTableDraft(ctx iris.Context) {
	log.Print("创建表单（草稿）")
	user, ok := middleware.ParseToken(ctx)
	if !ok {
		log.Printf("解析TOKEN出错，请重新登录")
		utils.Responser.FailWithMsg(ctx, "解析TOKEN出错，请重新登录")
		return
	}
	var tableDraft workflow.SdTableDraft
	if err := ctx.ReadJSON(&tableDraft); err != nil {
		utils.Responser.FailWithMsg(ctx, "")
		log.Print("数据接收失败")
		return
	}
	tableDraft.UserId = user.Id
	e := global.GVA_DB
	effect, err := e.Insert(tableDraft)
	if effect <= 0 || err != nil {
		log.Print(err)
		log.Printf("数据库操作失败")
		utils.Responser.FailWithMsg(ctx, "")
		return
	}
	utils.Responser.Ok(ctx)
}
