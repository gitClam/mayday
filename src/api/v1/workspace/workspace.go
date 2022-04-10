package workspace

import (
	"github.com/kataras/iris/v12"
	"mayday/src/global"
	"mayday/src/model/common/resultcode"
	userModel "mayday/src/model/user"
	WorkspaceModel "mayday/src/model/workspace"
	DepartmentModel "mayday/src/model/workspace/department"
	JobModel "mayday/src/model/workspace/job"
	"mayday/src/utils"
	"strconv"
	"strings"
)

// @Tags Workspace
// @Summary 获取工作空间信息
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Success 200 {object} utils.Response{data=workspace.SdWorkspace} "（1023::数据不存在或查询失败)"
// @Router /workspace/select/user [Get]
func WorkspaceSelectWorkspaceUserid(ctx iris.Context) {

	user := ctx.Values().Get("user").(userModel.SdUser)

	var Workspace []WorkspaceModel.SdWorkspace

	e := global.GVA_DB
	err := e.SQL("select * from sd_workspace where id in (select workspace_id from sd_department where id in (select department_id from sd_job where id in(select job_id from sd_user_job where user_id = ?)))", user.Id).Find(&Workspace)
	if err != nil {
		utils.Responser.Fail(ctx, resultcode.DataSelectFail, err)
		return
	}

	utils.Responser.OkWithDetails(ctx, Workspace)
}

// @Tags Workspace
// @Summary 获取工作空间信息
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param id path int true " 工作空间id(可以多个，以 ',' 分隔开) 例：'1,2,3,4'"
// @Success 200 {object} utils.Response{data=workspace.SdWorkspace} "（1023::数据不存在或查询失败)"
// @Router /workspace/select/workspace [Get]
func WorkspaceSelectWorkspace(ctx iris.Context) {
	var workspaceIds []int
	for _, id := range strings.Split(ctx.URLParam("id"), ",") {
		num, err := strconv.Atoi(id)
		if err != nil {
			utils.Responser.Fail(ctx, resultcode.DataReceiveFail, err)
			return
		}
		workspaceIds = append(workspaceIds, num)
	}
	var sdWorkspace []WorkspaceModel.SdWorkspace
	e := global.GVA_DB
	err := e.In("id", workspaceIds).Find(&sdWorkspace)
	if err != nil {
		utils.Responser.Fail(ctx, resultcode.DataSelectFail, err)
		return
	}
	utils.Responser.OkWithDetails(ctx, sdWorkspace)
}

// @Tags Workspace
// @Summary 创建工作空间
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Name body string true " 工作空间名字"
// @Param Phone body string true " 联系电话"
// @Param Remark body string true " 备注"
// @Success 200 {object} utils.Response{} "（1017::数据接收失败,1023::数据不存在或查询失败)"
// @Router /workspace/create/workspace [Post]
func WorkspaceCreate(ctx iris.Context) {
	user := ctx.Values().Get("user").(userModel.SdUser)

	var Workspace WorkspaceModel.SdWorkspace
	if err := ctx.ReadForm(&Workspace); err != nil {
		utils.Responser.Fail(ctx, resultcode.DataReceiveFail, err)
		return
	}

	e := global.GVA_DB.NewSession()
	defer e.Close()
	e.Begin()
	affect, err := e.Insert(&Workspace)
	if affect <= 0 || err != nil {
		e.Rollback()
		utils.Responser.Fail(ctx, resultcode.DataCreateFail, err)
		return
	}
	has, err1 := e.Get(&Workspace)
	if !has || err1 != nil {
		e.Rollback()
		utils.Responser.Fail(ctx, resultcode.DataCreateFail, err)
		return
	}
	department := DepartmentModel.SdDepartment{
		WorkspaceId: Workspace.Id,
		Name:        "默认",
		Phone:       "0000000",
		Remark:      "昨木",
		IsDeleted:   0}
	affect, err = e.Insert(&department)
	if affect <= 0 || err != nil {
		e.Rollback()
		utils.Responser.Fail(ctx, resultcode.DataCreateFail, err)
		return
	}
	has, err1 = e.Get(&department)
	if !has || err1 != nil {
		e.Rollback()
		utils.Responser.Fail(ctx, resultcode.DataCreateFail, err)
		return
	}
	job := JobModel.SdJob{
		DepartmentId: department.Id,
		Name:         "默认职位",
		IsDelete:     0}
	affect, err = e.Insert(&job)
	if affect <= 0 || err != nil {
		e.Rollback()
		utils.Responser.Fail(ctx, resultcode.DataCreateFail, err)
		return
	}
	has, err1 = e.Get(&job)
	if !has || err1 != nil {
		e.Rollback()
		utils.Responser.Fail(ctx, resultcode.DataCreateFail, err)
		return
	}
	userJob := JobModel.SdUserJob{
		UserId: user.Id,
		JobId:  job.Id}
	affect, err = e.Insert(&userJob)
	if affect <= 0 || err != nil {
		e.Rollback()
		utils.Responser.Fail(ctx, resultcode.DataCreateFail, err)
		return
	}
	e.Commit()
	utils.Responser.Ok(ctx)
}

// @Tags Workspace
// @Summary 修改工作空间信息
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Id body string true "工作空间id"
// @Param Name body string true "工作空间名字"
// @Param Phone body string true "联系电话"
// @Param Remark body string true "备注"
// @Success 200 {object} utils.Response{} "（1017::数据接收失败,1023::数据不存在或查询失败)"
// @Router /workspace/editor/workspace [Post]
func WorkspaceEditor(ctx iris.Context) {
	var Workspace WorkspaceModel.SdWorkspace
	if err := ctx.ReadForm(&Workspace); err != nil || Workspace.Id == 0 {
		utils.Responser.Fail(ctx, resultcode.DataReceiveFail, err)
		return
	}
	e := global.GVA_DB
	affect, err := e.Id(Workspace.Id).Update(&Workspace)
	if affect <= 0 || err != nil {
		utils.Responser.Fail(ctx, resultcode.DataUpdateFail, err)
		return
	}
	utils.Responser.Ok(ctx)
}

// @Tags Workspace
// @Summary 删除工作空间
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param id path int true " 部门id(可以多个，以 ',' 分隔开) 例：'1,2,3,4'"
// @Success 200 {object} utils.Response{} "（1017::数据接收失败,1023::数据不存在或查询失败)"
// @Router /workspace/delete/workspace [Delete]
func WorkspaceDelete(ctx iris.Context) {

	var workspaceIds []int
	for _, id := range strings.Split(ctx.URLParam("id"), ",") {
		num, err := strconv.Atoi(id)
		if err != nil {
			utils.Responser.Fail(ctx, resultcode.DataReceiveFail, err)
			return
		}
		workspaceIds = append(workspaceIds, num)
	}

	e := global.GVA_DB.NewSession()
	defer e.Close()
	e.Begin()
	for _, workspaceId := range workspaceIds {
		_, err := e.Exec("delete from sd_user_job where job_id in( select id from sd_job where department_id in (select id from sd_department where workspace_id = ?))", workspaceId)
		if err != nil {
			e.Rollback()
			utils.Responser.Fail(ctx, resultcode.DataDeleteFail, err)
			return
		}
		_, err = e.Exec("delete from sd_job where department_id in (select id from sd_department where workspace_id = ?)", workspaceId)
		if err != nil {
			e.Rollback()
			utils.Responser.Fail(ctx, resultcode.DataDeleteFail, err)
			return
		}
		_, err = e.Exec("delete from sd_department where workspace_id = ?", workspaceId)
		if err != nil {
			e.Rollback()
			utils.Responser.Fail(ctx, resultcode.DataDeleteFail, err)
			return
		}
		_, err = e.Exec("delete from sd_workspace where id = ?", workspaceId)
		if err != nil {
			e.Rollback()
			utils.Responser.Fail(ctx, resultcode.DataDeleteFail, err)
			return
		}
	}
	e.Commit()
	utils.Responser.Ok(ctx)
}

func WorkspaceSelectUseByWorkspaceId(ctx iris.Context) {
	var workspaceIds []int
	for _, id := range strings.Split(ctx.URLParam("id"), ",") {
		num, err := strconv.Atoi(id)
		if err != nil {
			utils.Responser.Fail(ctx, resultcode.DataReceiveFail, err)
			return
		}
		workspaceIds = append(workspaceIds, num)
	}
	var allUser []userModel.SdUser
	for _, workspaceId := range workspaceIds {
		var users []userModel.SdUser
		err := global.GVA_DB.SQL("select * from sd_user where id IN (SELECT user_id FROM sd_user_job WHERE job_id IN (SELECT id FROM sd_job  WHERE department_id IN (SELECT id FROM sd_department WHERE workspace_id = ?)))", workspaceId).Find(&users)
		if err != nil {
			utils.Responser.Fail(ctx, resultcode.DataSelectFail, err)
			return
		}
		allUser = append(allUser, users...)
	}
	utils.Responser.OkWithDetails(ctx, allUser)
}

//删除员工
func WorkspaceUserDelete(ctx iris.Context) {
	userId := ctx.FormValue("userId")
	workspaceId := ctx.FormValue("workspaceId")
	if userId == "" || workspaceId == "" {
		utils.Responser.Fail(ctx, resultcode.DataReceiveFail)
		return
	}

	_, err := global.GVA_DB.Exec("DELETE FROM sd_user_job WHERE user_id = ? AND job_id IN (SELECT id FROM sd_job WHERE department_id IN (SELECT id FROM sd_department WHERE workspace_id = ?))", userId, workspaceId)
	if err != nil {
		utils.Responser.Fail(ctx, resultcode.DataDeleteFail)
		return
	}
	utils.Responser.Ok(ctx)
}
