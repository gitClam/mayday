package workspace

import (
	"github.com/kataras/iris/v12"
	"mayday/src/global"
	"mayday/src/model/common/resultcode"
	userModel "mayday/src/model/user"
	DepartmentModel "mayday/src/model/workspace/department"
	"mayday/src/utils"
	"strconv"
	"strings"
)

// @Tags Department
// @Summary 根据用户ID获取部门信息
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Authorization header string true "用户登录返回的TOKEN"
// @Success 200 {object} utils.Response{data=department.SdDepartment} "错误码 （1017::数据接收失败,1023::数据不存在或查询失败)"
// @Router /workspace/department/select/user [Get]
func DepartmentSelectUser(ctx iris.Context) {
	user := ctx.Values().Get("user").(userModel.SdUser)
	var department []DepartmentModel.SdDepartment

	e := global.GVA_DB
	err := e.SQL("select * from sd_department where id in (select department_id from sd_job where id in(select job_id from sd_user_job where user_id = ?))", user.Id).Find(&department)
	if err != nil {
		utils.Responser.Fail(ctx, resultcode.DataSelectFail, err)
		return
	}

	utils.Responser.OkWithDetails(ctx, department)
}

// @Tags Department
// @Summary 获取部门信息(通过部门id)
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param id path int true " 部门id(可以多个，以 ',' 分隔开) 例：'1,2,3,4'"
// @Success 200 {object} utils.Response{data=department.SdDepartment} "错误码 （1017::数据接收失败,1023::数据不存在或查询失败)"
// @Router /workspace/department/select [Get]
func DepartmentSelect(ctx iris.Context) {
	var departmentIds []int
	for _, id := range strings.Split(ctx.URLParam("id"), ",") {
		num, err := strconv.Atoi(id)
		if err != nil {
			utils.Responser.Fail(ctx, resultcode.DataReceiveFail, err)
			return
		}
		departmentIds = append(departmentIds, num)
	}

	var sdDepartment []DepartmentModel.SdDepartment
	e := global.GVA_DB
	err := e.Id(departmentIds).Find(&sdDepartment)
	if err != nil {
		utils.Responser.Fail(ctx, resultcode.DataSelectFail, err)
		return
	}
	utils.Responser.OkWithDetails(ctx, sdDepartment)
}

// @Tags Department
// @Summary 获取部门信息(通过工作空间id)
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param id path int true " 工作空间id(可以多个，以 ',' 分隔开) 例：'1,2,3,4'"
// @Success 200 {object} utils.Response{data=department.SdDepartment} "错误码 （1017::数据接收失败,1023::数据不存在或查询失败)"
// @Router /workspace/department/select/workspace [Get]
func DepartmentSelectWorkspace(ctx iris.Context) {
	var workspaceIds []int
	for _, id := range strings.Split(ctx.URLParam("id"), ",") {
		num, err := strconv.Atoi(id)
		if err != nil {
			utils.Responser.Fail(ctx, resultcode.DataReceiveFail, err)
			return
		}
		workspaceIds = append(workspaceIds, num)
	}

	var departments []DepartmentModel.SdDepartment

	e := global.GVA_DB
	err := e.Where("workspace_id = ? and is_deleted = 0", workspaceIds).Find(&departments)
	if err != nil {
		utils.Responser.Fail(ctx, resultcode.DataSelectFail, err)
		return
	}
	utils.Responser.OkWithDetails(ctx, departments)
}

// @Tags Department
// @Summary 创建部门
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param WorkspaceId body string true " 工作空间id"
// @Param Name body string true " 部门名字"
// @Param Phone body string true " 联系电话"
// @Param Remark body string true " 备注"
// @Success 200 {object} utils.Response{} "错误码 （1017::数据接收失败,1023::数据不存在或查询失败)"
// @Router /workspace/department/create [POST]
func DepartmentCreate(ctx iris.Context) {
	var sdDepartment DepartmentModel.SdDepartment
	if err := ctx.ReadForm(&sdDepartment); err != nil {
		utils.Responser.Fail(ctx, resultcode.DataReceiveFail, err)
		return
	}
	e := global.GVA_DB
	affect, err := e.Insert(&sdDepartment)
	if affect <= 0 || err != nil {
		utils.Responser.Fail(ctx, resultcode.DataSelectFail, err)
		return
	}
	utils.Responser.Ok(ctx)
}

// @Tags Department
// @Summary 修改部门信息
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Id body string true " 部门id"
// @Param WorkspaceId body string true " 工作空间id"
// @Param Name body string true " 部门名字"
// @Param Phone body string true " 联系电话"
// @Param Remark body string true " 备注"
// @Success 200 {object} utils.Response{} "错误码 （1017::数据接收失败,1023::数据不存在或查询失败)"
// @Router /workspace/department/editor [POST]
func DepartmentEditor(ctx iris.Context) {
	var sdDepartment DepartmentModel.SdDepartment
	if err := ctx.ReadForm(&sdDepartment); err != nil || sdDepartment.Id == 0 {
		utils.Responser.Fail(ctx, resultcode.DataReceiveFail, err)
		return
	}
	e := global.GVA_DB
	affect, err := e.Id(sdDepartment.Id).Update(&sdDepartment)
	if affect <= 0 || err != nil {
		utils.Responser.Fail(ctx, resultcode.DataSelectFail, err)
		return
	}
	utils.Responser.Ok(ctx)
}

// @Tags Department
// @Summary 删除部门
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param id path int true " 部门id(可以多个，以 ',' 分隔开) 例：'1,2,3,4'"
// @Success 200 {object} utils.Response{data=department.SdDepartment} "错误码 （1017::数据接收失败,1023::数据不存在或查询失败)"
// @Router /workspace/department/delete [Delete]
func DepartmentDelete(ctx iris.Context) {

	var departmentIds []int
	for _, id := range strings.Split(ctx.URLParam("id"), ",") {
		num, err := strconv.Atoi(id)
		if err != nil {
			utils.Responser.Fail(ctx, resultcode.DataReceiveFail, err)
			return
		}
		departmentIds = append(departmentIds, num)
	}

	e := global.GVA_DB.NewSession()
	defer e.Close()
	err := e.Begin()
	if err != nil {
		utils.Responser.Fail(ctx, resultcode.DataSelectFail, err)
		return
	}
	for _, departmentId := range departmentIds {

		_, err := e.Exec("delete from sd_user_job where job_id in( select id from sd_job where department_id = ?)", departmentId)
		if err != nil {
			e.Rollback()
			utils.Responser.Fail(ctx, resultcode.DataDeleteFail, err)
			return
		}
		_, err = e.Exec("delete from sd_job where department_id = ?", departmentId)
		if err != nil {
			e.Rollback()
			utils.Responser.Fail(ctx, resultcode.DataDeleteFail, err)
			return
		}
		_, err = e.Exec("delete from sd_department where id = ?", departmentId)
		if err != nil {
			e.Rollback()
			utils.Responser.Fail(ctx, resultcode.DataDeleteFail, err)
			return
		}
	}
	err = e.Commit()
	if err != nil {
		utils.Responser.Fail(ctx, resultcode.DataSelectFail, err)
		return
	}
	utils.Responser.Ok(ctx)
}
