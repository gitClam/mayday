package workflow

import (
	"github.com/kataras/iris/v12"
	"mayday/src/middleware"
	workflowModel "mayday/src/model/workflow"
	workflowSever "mayday/src/service/workflow"
	"mayday/src/utils"
)

// @Tags Workflow
// @Summary 获取流程详细信息
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Param id path int true "流程id"
// @Success 200 {object} utils.Response{data=user.UserDetailsRes} "返回流程的详细信息"
// @Router /workflow/get/workflow/{id:int} [get]
func GetWorkflowById(ctx iris.Context) {

}

// @Tags Workflow
// @Summary 获取当前用户的流程草稿列表
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Success 200 {object} utils.Response{data=user.UserDetailsRes} "返回流程草稿的详细信息"
// @Router /workflow/get/workflow-draft [get]
func GetWorkflowDraftByUser(ctx iris.Context) {

}

// @Tags Workflow
// @Summary 获取流程草稿详细信息
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Param id path int true "流程草稿id"
// @Success 200 {object} utils.Response{data=user.UserDetailsRes} "返回流程的详细信息"
// @Router /workflow/get/workflow-draft/{id:int} [get]
func GetWorkflowDraftById(ctx iris.Context) {

}

// @Tags Workflow
// @Summary 创建流程
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Param userReq body user.UserReq true "流程信息"
// @Success 200 {object} utils.Response
// @Router /workflow/create/workflow [post]
func CreateWorkflow(ctx iris.Context) {

	user, ok := middleware.ParseToken(ctx)
	if !ok {
		utils.Responser.FailWithMsg(ctx, "解析TOKEN出错，请重新登录")
		return
	}

	var workflow workflowModel.SdWorkflow
	if err := ctx.ReadJSON(&workflow); err != nil {
		utils.Responser.FailWithMsg(ctx, "数据接收失败")
		return
	}
	workflowSever.CreateWorkflow(ctx, *user, workflow)

}

// @Tags Workflow
// @Summary 创建流程草稿
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Param userReq body user.UserReq true "流程信息"
// @Success 200 {object} utils.Response
// @Router /workflow/create/workflow-draft [post]
func CreateWorkflowDraft(ctx iris.Context) {

}

// @Tags Workflow
// @Summary 修改流程信息
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Param userReq body user.UserReq true "流程信息"
// @Success 200 {object} utils.Response
// @Router /workflow/update/workflow [post]
func UpdateWorkflow(ctx iris.Context) {

}

// @Tags Workflow
// @Summary 修改流程状态
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Param id body int true "流程id"
// @Success 200 {object} utils.Response
// @Router /workflow/update/workflow-state [post]
func UpdateWorkflowState(ctx iris.Context) {

}

// @Tags Workflow
// @Summary 修改流程草稿
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Param userReq body user.UserReq true "流程信息"
// @Success 200 {object} utils.Response
// @Router /workflow/update/workflow-draft [post]
func UpdateWorkflowDraft(ctx iris.Context) {

}

// @Tags Workflow
// @Summary 删除流程
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Param id path int true "流程id"
// @Success 200 {object} utils.Response
// @Router /workflow/delete/workflow/{id:int} [delete]
func DeleteWorkflow(ctx iris.Context) {

}

// @Tags Workflow
// @Summary 删除流程草稿
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Param id path int true "流程id"
// @Success 200 {object} utils.Response
// @Router /workflow/delete/workflow-draft/{id:int} [delete]
func DeleteWorkflowDraft(ctx iris.Context) {

}