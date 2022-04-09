package workspace

import (
	"github.com/kataras/iris/v12"
	"mayday/src/global"
	"mayday/src/model/common/resultcode"
	ApplicationModel "mayday/src/model/workspace/application"
	"mayday/src/utils"
	"strconv"
	"strings"
)

// @Tags Application
// @Summary 应用Id获取流程信息
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param id path int true " 应用Id(可以多个，以 ',' 分隔开) 例：'1,2,3,4'"
// @Success 200 {object} utils.Response{data=application.SdWorkflowApplication} "错误码 （1017::数据接收失败,1023::数据不存在或查询失败)"
// @Router /workspace/application/select [Get]
func ApplicationSelect(ctx iris.Context) {
	var applicationIds []int
	for _, id := range strings.Split(ctx.URLParam("id"), ",") {
		num, err := strconv.Atoi(id)
		if err != nil {
			utils.Responser.Fail(ctx, resultcode.DataReceiveFail, err)
			return
		}
		applicationIds = append(applicationIds, num)
	}
	var allSdWorkflowApplication []ApplicationModel.SdWorkflowApplication
	for _, applicationId := range applicationIds {
		var SdWorkflowApplication []ApplicationModel.SdWorkflowApplication
		e := global.GVA_DB
		err := e.Where("ApplicationId = ?", applicationId).Find(&SdWorkflowApplication)
		if err != nil {
			utils.Responser.Fail(ctx, resultcode.DataSelectFail, err)
			return
		}
		allSdWorkflowApplication = append(allSdWorkflowApplication, SdWorkflowApplication...)
	}
	utils.Responser.OkWithDetails(ctx, allSdWorkflowApplication)
}

// @Tags Application
// @Summary 根据工作空间Id获取全部流程
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param id path int true " 工作空间Id(可以多个，以 ',' 分隔开) 例：'1,2,3,4'"
// @Success 200 {object} utils.Response{data=application.SdWorkflowApplication} "错误码 （1017::数据接收失败,1023::数据不存在或查询失败)"
// @Router /workspace/application/select/workFlowByWorkspaceId [Get]
func GetWorkflowByWorkspaceId(ctx iris.Context) {

	var workspaceIds []int
	for _, id := range strings.Split(ctx.URLParam("id"), ",") {
		num, err := strconv.Atoi(id)
		if err != nil {
			utils.Responser.Fail(ctx, resultcode.DataReceiveFail, err)
			return
		}
		workspaceIds = append(workspaceIds, num)
	}
	var allSdWorkflowApplications []ApplicationModel.SdWorkflowApplication
	for _, workspaceId := range workspaceIds {
		var SdApplications []ApplicationModel.SdApplication
		e := global.GVA_DB
		err := e.Where("workspace_id = ?", workspaceId).Find(&SdApplications)
		if err != nil {

			utils.Responser.Fail(ctx, resultcode.DataSelectFail, err)
			return

		}
		var SdWorkflowApplications []ApplicationModel.SdWorkflowApplication

		for _, SdApplication := range SdApplications {
			err := e.Where("application_id = ?", SdApplication.Id).Find(&SdWorkflowApplications)
			if err != nil {
				utils.Responser.Fail(ctx, resultcode.DataSelectFail, err)
				return
			}
			allSdWorkflowApplications = append(allSdWorkflowApplications, SdWorkflowApplications...)
			SdWorkflowApplications = SdWorkflowApplications[0:0]
		}
	}
	utils.Responser.OkWithDetails(ctx, allSdWorkflowApplications)
}

// @Tags Application
// @Summary 获取应用信息
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param id path int true " 工作空间Id(可以多个，以 ',' 分隔开) 例：'1,2,3,4'"
// @Success 200 {object} utils.Response{data=application.SdApplication} "错误码 （1017::数据接收失败,1023::数据不存在或查询失败)"
// @Router /workspace/application/select/workspace [Get]
func ApplicationSelectWorkspace(ctx iris.Context) {

	var workspaceIds []int
	for _, id := range strings.Split(ctx.URLParam("id"), ",") {
		num, err := strconv.Atoi(id)
		if err != nil {
			utils.Responser.Fail(ctx, resultcode.DataReceiveFail, err)
			return
		}
		workspaceIds = append(workspaceIds, num)
	}
	var allSdApplication []ApplicationModel.SdApplication
	for _, workspaceId := range workspaceIds {
		var SdApplications []ApplicationModel.SdApplication
		e := global.GVA_DB
		err := e.Where("workspace_id = ?", workspaceId).Find(&SdApplications)
		if err != nil {
			utils.Responser.Fail(ctx, resultcode.DataSelectFail, err)
			return
		}
		allSdApplication = append(allSdApplication, SdApplications...)
	}
	utils.Responser.OkWithDetails(ctx, allSdApplication)
}

// @Tags Application
// @Summary 创建应用
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param WorkspaceId body string true " 工作空间id"
// @Param Name body string true " 名字"
// @Param Remark body string true " 备注"
// @Success 200 {object} utils.Response{} "错误码 （1017::数据接收失败,1023::数据不存在或查询失败)"
// @Router /workspace/application/create [POST]
func ApplicationCreate(ctx iris.Context) {
	var application ApplicationModel.SdApplication
	if err := ctx.ReadForm(&application); err != nil {
		utils.Responser.Fail(ctx, resultcode.DataReceiveFail, err)
		return
	}
	e := global.GVA_DB
	affect, err := e.Insert(&application)
	if affect <= 0 || err != nil {
		utils.Responser.Fail(ctx, resultcode.DataCreateFail, err)
		return
	}
	utils.Responser.Ok(ctx)

}

// @Tags Application
// @Summary 修改应用信息
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Id body string true " 应用id"
// @Param WorkspaceId body string false " 工作空间id"
// @Param Name body string false " 名字"
// @Param Remark body string false " 备注"
// @Success 200 {object} utils.Response{} "错误码 （1017::数据接收失败,1023::数据不存在或查询失败)"
// @Router /workspace/application/editor [POST]
func ApplicationEditor(ctx iris.Context) {
	var application ApplicationModel.SdApplication
	if err := ctx.ReadForm(&application); err != nil || application.Id == 0 {
		utils.Responser.Fail(ctx, resultcode.DataReceiveFail, err)
		return
	}
	e := global.GVA_DB
	affect, err := e.Id(application.Id).Update(&application)
	if affect <= 0 || err != nil {
		utils.Responser.Fail(ctx, resultcode.DataUpdateFail, err)
		return
	}
	utils.Responser.Ok(ctx)
}

// @Tags Application
// @Summary 删除应用
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param id path int true " 工作空间Id(可以多个，以 ',' 分隔开) 例：'1,2,3,4'"
// @Success 200 {object} utils.Response{} "错误码 （1017::数据接收失败,1024::创建失败)"
// @Router /workspace/application/delete [Delete]
func ApplicationDelete(ctx iris.Context) {
	var applicationIds []int
	for _, id := range strings.Split(ctx.URLParam("id"), ",") {
		num, err := strconv.Atoi(id)
		if err != nil {
			utils.Responser.Fail(ctx, resultcode.DataReceiveFail, err)
			return
		}
		applicationIds = append(applicationIds, num)
	}

	e := global.GVA_DB.NewSession()
	defer e.Close()
	err := e.Begin()
	if err != nil {
		utils.Responser.Fail(ctx, resultcode.DataDeleteFail, err)
		return
	}

	for _, applicationId := range applicationIds {
		var SdWorkflowApplication ApplicationModel.SdWorkflowApplication
		var SdApplication ApplicationModel.SdApplication
		effect, err := e.Where("application_id = ?", applicationId).Delete(&SdWorkflowApplication)
		if effect == 0 || err != nil {
			e.Rollback()
			utils.Responser.Fail(ctx, resultcode.DataDeleteFail, err)
			return
		}
		effect, err = e.Where("id = ?", applicationId).Delete(&SdApplication)
		if effect == 0 || err != nil {
			e.Rollback()
			utils.Responser.Fail(ctx, resultcode.DataDeleteFail, err)
			return
		}
	}
	e.Commit()
	utils.Responser.Ok(ctx)
}

// @Tags Application
// @Summary 添加流程
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param WorkflowId body string true " 流程ID"
// @Param ApplicationId body string true " 应用id"
// @Success 200 {object} utils.Response{} "错误码 （1017::数据接收失败,1022::创建失败)"
// @Router /workspace/application/insert-workflow [POST]
func ApplicationInsert(ctx iris.Context) {
	var workflowApplication ApplicationModel.SdWorkflowApplication
	if err := ctx.ReadForm(&workflowApplication); err != nil || workflowApplication.WorkflowId == 0 || workflowApplication.ApplicationId == 0 {
		utils.Responser.Fail(ctx, resultcode.DataReceiveFail, err)
		return
	}
	e := global.GVA_DB
	affect, err := e.Insert(&workflowApplication)
	if affect <= 0 || err != nil {
		utils.Responser.Fail(ctx, resultcode.DataCreateFail, err)
		return
	}
	utils.Responser.Ok(ctx)
}

// @Tags Application
// @Summary 删除流程
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param WorkflowId body string true " 流程ID"
// @Param ApplicationId body string true " 应用id"
// @Success 200 {object} utils.Response{} "错误码 （1017::数据接收失败,1024::删除失败)"
// @Router /workspace/application/delete-workflow [POST]
func ApplicationDeleteWorkflow(ctx iris.Context) {
	var workflowApplication ApplicationModel.SdWorkflowApplication
	if err := ctx.ReadForm(&workflowApplication); err != nil || workflowApplication.WorkflowId == 0 || workflowApplication.ApplicationId == 0 {
		utils.Responser.Fail(ctx, resultcode.DataReceiveFail, err)
		return
	}
	e := global.GVA_DB
	affect, err := e.Delete(&workflowApplication)
	if affect <= 0 || err != nil {
		utils.Responser.Fail(ctx, resultcode.DataDeleteFail, err)
		return
	}
	utils.Responser.Ok(ctx)
}
