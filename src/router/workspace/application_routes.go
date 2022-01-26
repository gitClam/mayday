package workspace_routes

import (
	"github.com/kataras/iris/v12"
	"log"
	"mayday/src/initialize"
	"mayday/src/model"
	"mayday/src/utils"
	//"mayday/middleware/jwts"
)

// swagger:operation Post /workspace/application/select application select_application
// ---
// summary: 获取应用信息
// description: 根据ID获取应用信息
// parameters:
// - name: id
//   description: 应用id
//   type: string
//   required: true
func ApplicationSelect(ctx iris.Context) {
	var application model.SdApplication
	if err := ctx.ReadForm(&application); err != nil || application.Id == 0 {
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		log.Print("数据接收失败")
		return
	}
	e := initialize.MasterEngine()
	has, err := e.Where("is_deleted = 0").Get(&application)
	if !has || err != nil {
		log.Printf("数据库查询错误")
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	utils.MakeSuccessRes(ctx, model.Success, application)
}

// swagger:operation Post /workspace/application/select/workspace application select_workspace_Application
// ---
// summary: 获取应用信息
// description: 获取应用信息
// parameters:
// - name: WorkspaceId
//   description: 工作空间id
//   type: string
//   required: true
func ApplicationSelectWorkspace(ctx iris.Context) {

	var application model.SdApplication
	if err := ctx.ReadForm(&application); err != nil || application.WorkspaceId == 0 {
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		log.Print("数据接收失败")
		return
	}

	var applications []model.SdApplication

	e := initialize.MasterEngine()
	has, err := e.Where("is_deleted = 0 and workspace_id = ?", application.WorkspaceId).Get(&applications)
	if !has || err != nil {
		log.Printf("数据库查询错误")
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	utils.MakeSuccessRes(ctx, model.Success, applications)
}

// swagger:operation POST /workspace/application/create application create_application
// ---
// summary: 创建应用
// description: 创建应用
// parameters:
// - name: WorkspaceId
//   description: 工作空间id
//   type: string
//   required: true
// - name: name
//   description: 名字
//   type: string
//   required: true
// - name: Remark
//   description: 备注
//   type: string
//   required: true
func ApplicationCreate(ctx iris.Context) {
	var application model.SdApplication
	if err := ctx.ReadForm(&application); err != nil {
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		log.Print("数据接收失败")
		return
	}
	application.IsDeleted = 0
	e := initialize.MasterEngine()
	affect, err := e.Insert(&application)
	if affect <= 0 || err != nil {
		log.Printf("数据库插入错误")
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	utils.MakeSuccessRes(ctx, model.Success, nil)

}

// swagger:operation POST /workspace/application/editor application editor_application
// ---
// summary: 修改应用信息
// description: 修改应用信息
// parameters:
// - name: WorkspaceId
//   description: 工作空间id
//   type: string
//   required: true
// - name: name
//   description: 名字
//   type: string
//   required: true
// - name: Remark
//   description: 备注
//   type: string
//   required: true
func ApplicationEditor(ctx iris.Context) {
	var application model.SdApplication
	if err := ctx.ReadForm(&application); err != nil || application.Id == 0 {
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		log.Print("数据接收失败")
		return
	}
	e := initialize.MasterEngine()
	affect, err := e.Id(application.Id).Update(&application)
	if affect <= 0 || err != nil {
		log.Printf("数据库插入错误")
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	utils.MakeSuccessRes(ctx, model.Success, nil)
}

// swagger:operation DELETE /workspace/application/delete application delete_application
// ---
// summary: 删除应用
// description: 删除应用
func ApplicationDelete(ctx iris.Context) {
	var application model.SdApplication
	if err := ctx.ReadForm(&application); err != nil || application.Id == 0 {
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		log.Print("数据接收失败")
		return
	}
	application.IsDeleted = 1
	e := initialize.MasterEngine()
	affect, err := e.Id(application.Id).Update(&application)
	if affect <= 0 || err != nil {
		log.Printf("数据库插入错误")
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	utils.MakeSuccessRes(ctx, model.Success, nil)
}

// swagger:operation POST /workspace/application/insert-workflow application insert_application
// ---
// summary: 添加流程
// description: 添加流程
// parameters:
// - name: WorkflowId
//   description: 流程ID
//   type: string
//   required: true
// - name: ApplicationId
//   description: 应用id
//   type: string
//   required: true
func ApplicationInsert(ctx iris.Context) {
	var workflowApplication model.SdWorkflowApplication
	if err := ctx.ReadForm(&workflowApplication); err != nil || workflowApplication.WorkflowId == 0 || workflowApplication.ApplicationId == 0 {
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		log.Print("数据接收失败")
		return
	}
	e := initialize.MasterEngine()
	affect, err := e.Insert(&workflowApplication)
	if affect <= 0 || err != nil {
		log.Printf("数据库插入错误")
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	utils.MakeSuccessRes(ctx, model.Success, nil)
}

// swagger:operation DELETE /workspace/application/delete-workflow application delete_workflow_Application
// ---
// summary: 删除流程
// description: 删除流程
// parameters:
// - name: WorkflowId
//   description: 流程ID
//   type: string
//   required: true
// - name: ApplicationId
//   description: 应用id
//   type: string
//   required: true
func ApplicationDeleteWorkflow(ctx iris.Context) {
	var workflowApplication model.SdWorkflowApplication
	if err := ctx.ReadForm(&workflowApplication); err != nil || workflowApplication.WorkflowId == 0 || workflowApplication.ApplicationId == 0 {
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		log.Print("数据接收失败")
		return
	}
	e := initialize.MasterEngine()
	affect, err := e.Delete(&workflowApplication)
	if affect <= 0 || err != nil {
		log.Printf("数据库插入错误")
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	utils.MakeSuccessRes(ctx, model.Success, nil)
}
