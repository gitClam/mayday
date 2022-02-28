package workflow

import (
	"github.com/kataras/iris/v12"
	workflowModel "mayday/src/model/workflow"
	workflowSever "mayday/src/service/workflow"
	"mayday/src/utils"
	"strconv"
	"strings"
)

// @Tags Workflow
// @Summary 获取流程详细信息
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Param id path int true "流程id(可以多个，以 ',' 分隔开) 例：'1,2,3,4'"
// @Success 200 {object} utils.Response{data=[]workflow.SdWorkflow} "返回流程的详细信息"
// @Router /workflow/get/workflow [get]
func GetWorkflowById(ctx iris.Context) {
	var workflowsId []int
	for _, id := range strings.Split(ctx.URLParam("id"), ",") {
		num, err := strconv.Atoi(id)
		if err != nil {
			utils.Responser.FailWithMsg(ctx, "数据接收失败", err)
			return
		}
		workflowsId = append(workflowsId, num)
	}
	workflowSever.GetWorkflowById(ctx, workflowsId)

}

// @Tags Workflow
// @Summary 获取当前用户的流程草稿列表
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Success 200 {object} utils.Response{data=user.UserDetailsRes} "返回流程草稿的详细信息"
// @Router /workflow/get/workflow-draft/user [get]
func GetWorkflowDraftByUser(ctx iris.Context) {
	workflowSever.GetWorkflowDraftByUser(ctx)
}

// @Tags Workflow
// @Summary 获取流程草稿详细信息
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Param id path int true "流程草稿id(可以多个，以 ',' 分隔开) 例：'1,2,3,4'"
// @Success 200 {object} utils.Response{data=user.UserDetailsRes} "返回流程的详细信息"
// @Router /workflow/get/workflow-draft/id [get]
func GetWorkflowDraftById(ctx iris.Context) {
	var workflowDraftsId []int
	for _, id := range strings.Split(ctx.URLParam("id"), ",") {
		num, err := strconv.Atoi(id)
		if err != nil {
			utils.Responser.FailWithMsg(ctx, "数据接收失败", err)
			return
		}
		workflowDraftsId = append(workflowDraftsId, num)
	}
	workflowSever.GetWorkflowDraftById(ctx, workflowDraftsId)
}

// @Tags Workflow
// @Summary 创建流程
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Param userReq body workflowModel.WorkflowReq true "流程信息"
// @Success 200 {object} utils.Response
// @Router /workflow/create/workflow [post]
func CreateWorkflow(ctx iris.Context) {

	var workflowReq workflowModel.WorkflowReq
	if err := ctx.ReadJSON(&workflowReq); err != nil {
		utils.Responser.FailWithMsg(ctx, "数据接收失败", err)
		return
	}
	workflowSever.CreateWorkflow(ctx, workflowReq)
}

// @Tags Workflow
// @Summary 创建流程草稿
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Param userReq body workflowModel.WorkflowDraftReq true "流程信息"
// @Success 200 {object} utils.Response
// @Router /workflow/create/workflow-draft [post]
func CreateWorkflowDraft(ctx iris.Context) {

	var workflowDraftReq workflowModel.WorkflowDraftReq
	if err := ctx.ReadJSON(&workflowDraftReq); err != nil {
		utils.Responser.FailWithMsg(ctx, "数据接收失败", err)
		return
	}
	workflowSever.CreateWorkflowDraft(ctx, workflowDraftReq)
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
// @Param id path string true "流程id(可以多个，以 ',' 分隔开) 例：'1,2,3,4'"
// @Success 200 {object} utils.Response
// @Router /workflow/delete/workflow [delete]
func DeleteWorkflow(ctx iris.Context) {

	var workflowsId []int
	for _, id := range strings.Split(ctx.URLParam("id"), ",") {
		num, err := strconv.Atoi(id)
		if err != nil {
			utils.Responser.FailWithMsg(ctx, "数据接收失败", err)
			return
		}
		workflowsId = append(workflowsId, num)
	}
	workflowSever.DeleteWorkflow(ctx, workflowsId)
}

// @Tags Workflow
// @Summary 删除流程草稿
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Param id path int true "流程id(可以多个，以 ',' 分隔开) 例：'1,2,3,4'"
// @Success 200 {object} utils.Response
// @Router /workflow/delete/workflow-draft [delete]
func DeleteWorkflowDraft(ctx iris.Context) {
	var workflowsId []int
	for _, id := range strings.Split(ctx.URLParam("id"), ",") {
		num, err := strconv.Atoi(id)
		if err != nil {
			utils.Responser.FailWithMsg(ctx, "数据接收失败", err)
			return
		}
		workflowsId = append(workflowsId, num)
	}
	workflowSever.DeleteWorkflowDraft(ctx, workflowsId)
}
