package workspace

import (
	"github.com/kataras/iris/v12"
	"mayday/src/global"
	"mayday/src/model/common/resultcode"
	userModel "mayday/src/model/user"
	JobModel "mayday/src/model/workspace/job"
	"mayday/src/utils"
	"strconv"
	"strings"
)

// @Tags Job
// @Summary 获取职位信息(职位id)
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param id path int true "职位id(可以多个，以 ',' 分隔开) 例：'1,2,3,4'"
// @Success 200 {object} utils.Response{data=job.SdJob} "（1017::数据接收失败,1023::数据不存在或查询失败)"
// @Router /workspace/job/select [Get]
func JobSelect(ctx iris.Context) {
	var jobIds []int
	for _, id := range strings.Split(ctx.URLParam("id"), ",") {
		num, err := strconv.Atoi(id)
		if err != nil {
			utils.Responser.Fail(ctx, resultcode.DataReceiveFail, err)
			return
		}
		jobIds = append(jobIds, num)
	}

	var SdJob []JobModel.SdJob
	e := global.GVA_DB
	err := e.Id(jobIds).Find(&SdJob)

	if err != nil {
		utils.Responser.Fail(ctx, resultcode.DataSelectFail, err)
		return
	}
	utils.Responser.OkWithDetails(ctx, SdJob)
}

// @Tags Job
// @Summary 获取职位信息(部门id)
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param id path int true " 部门id(可以多个，以 ',' 分隔开) 例：'1,2,3,4'"
// @Success 200 {object} utils.Response{data=job.SdJob} "错误码 （1017::数据接收失败,1023::数据不存在或查询失败)"
// @Router /workspace/job/select/department [Get]
func JobSelectDepartment(ctx iris.Context) {
	var departmentId []int
	for _, id := range strings.Split(ctx.URLParam("id"), ",") {
		num, err := strconv.Atoi(id)
		if err != nil {
			utils.Responser.Fail(ctx, resultcode.DataReceiveFail, err)
			return
		}
		departmentId = append(departmentId, num)
	}

	var SdJob []JobModel.SdJob
	e := global.GVA_DB
	err := e.Where("department_id = ?", departmentId).Find(&SdJob)

	if err != nil {
		utils.Responser.Fail(ctx, resultcode.DataSelectFail, err)
		return
	}
	utils.Responser.OkWithDetails(ctx, SdJob)
}

// @Tags Job
// @Summary 获取职位信息(用户id)
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param id path int true " 用户id(可以多个，以 ',' 分隔开) 例：'1,2,3,4'"
// @Success 200 {object} utils.Response{data=job.SdJob} "错误码 （1017::数据接收失败,1023::数据不存在或查询失败)"
// @Router /workspace/job/select/user [Get]
func JobSelectUser(ctx iris.Context) {
	var UserId []int
	for _, id := range strings.Split(ctx.URLParam("id"), ",") {
		num, err := strconv.Atoi(id)
		if err != nil {
			utils.Responser.Fail(ctx, resultcode.DataReceiveFail, err)
			return
		}
		UserId = append(UserId, num)
	}
	var SdJob []JobModel.SdJob
	e := global.GVA_DB
	err := e.SQL("select * from sd_job where id in(select job_id from sd_user_job where user_id = ?)", UserId).Find(&SdJob)
	if err != nil {
		utils.Responser.Fail(ctx, resultcode.DataSelectFail, err)
		return
	}
	utils.Responser.OkWithDetails(ctx, SdJob)
}

// @Tags Job
// @Summary 创建职位
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param DepartmentId body int true "部门id"
// @Param name body string true "职位名字"
// @Success 200 {object} utils.Response{}
// @Router /workspace/job/create [POST]
func JobCreate(ctx iris.Context) {
	var SdJob JobModel.SdJob
	if err := ctx.ReadForm(&SdJob); err != nil {
		utils.Responser.Fail(ctx, resultcode.DataReceiveFail, err)
		return
	}
	e := global.GVA_DB
	affect, err := e.Insert(&SdJob)
	if affect <= 0 || err != nil {
		utils.Responser.Fail(ctx, resultcode.DataSelectFail, err)
		return
	}
	utils.Responser.Ok(ctx)
}

// @Tags Job
// @Summary 修改职位信息
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Id body int true "职位id"
// @Param DepartmentId body int true "部门id"
// @Param name body string true " 职位名字"
// @Success 200 {object} utils.Response{}
// @Router /workspace/job/editor [POST]
func JobEditor(ctx iris.Context) {
	var SdJob JobModel.SdJob
	if err := ctx.ReadForm(&SdJob); err != nil || SdJob.Id == 0 {
		utils.Responser.Fail(ctx, resultcode.DataReceiveFail, err)
		return
	}
	e := global.GVA_DB
	affect, err := e.Id(SdJob.Id).Update(&SdJob)
	if affect <= 0 || err != nil {
		utils.Responser.Fail(ctx, resultcode.DataSelectFail, err)
		return
	}
	utils.Responser.Ok(ctx)
}

// @Tags Job
// @Summary 删除职位
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param id path int true " 职位id(可以多个，以 ',' 分隔开) 例：'1,2,3,4'"
// @Success 200 {object} utils.Response{}
// @Router /workspace/job/delete [DELETE]
func JobDelete(ctx iris.Context) {
	var JobIds []int
	for _, id := range strings.Split(ctx.URLParam("id"), ",") {
		num, err := strconv.Atoi(id)
		if err != nil {
			utils.Responser.Fail(ctx, resultcode.DataReceiveFail, err)
			return
		}
		JobIds = append(JobIds, num)
	}

	e := global.GVA_DB.NewSession()
	defer e.Close()
	err := e.Begin()
	if err != nil {
		utils.Responser.Fail(ctx, resultcode.DataSelectFail, err)
		return
	}
	for jobId := range JobIds {
		_, err := e.Exec("delete from sd_user_job where job_id = ?", jobId)
		if err != nil {
			e.Rollback()
			utils.Responser.Fail(ctx, resultcode.DataSelectFail, err)
			return
		}
		_, err = e.Exec("delete from sd_job where id = ?", jobId)
		if err != nil {
			e.Rollback()
			utils.Responser.Fail(ctx, resultcode.DataSelectFail, err)
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

// @Tags Job
// @Summary 获取职位人员信息
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param id path int true " 职位id(可以多个，以 ',' 分隔开) 例：'1,2,3,4'"
// @Success 200 {object} utils.Response{user.UserAbstractRes}
// @Router /workspace/job/select-user [Get]
func SelectUserByJobId(ctx iris.Context) {
	var JobIds []int
	for _, id := range strings.Split(ctx.URLParam("id"), ",") {
		num, err := strconv.Atoi(id)
		if err != nil {
			utils.Responser.Fail(ctx, resultcode.DataReceiveFail, err)
			return
		}
		JobIds = append(JobIds, num)
	}
	var AllUsers []userModel.SdUser
	for jobId := range JobIds {
		var users []userModel.SdUser
		e := global.GVA_DB
		err := e.SQL("select * from sd_user where id in(select user_id from sd_user_job where job_id = ?)", jobId).Find(&users)
		if err != nil {
			utils.Responser.Fail(ctx, resultcode.DataSelectFail, err)
			return
		}
		AllUsers = append(AllUsers, users...)
	}
	utils.Responser.OkWithDetails(ctx, userModel.GetUserAbstractResList(AllUsers))
}

// @Tags Job
// @Summary 添加员工
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Mail body string true "用户邮箱"
// @Param JobId body string true "职位id"
// @Success 200 {object} utils.Response{}
// @Router /workspace/job/insert-user [POST]
func JobInsert(ctx iris.Context) {
	var userJob JobModel.SdUserJob
	var user userModel.SdUser

	user.Mail = ctx.FormValue("Mail")
	userJob.JobId, _ = strconv.Atoi(ctx.FormValue("JobId"))
	if user.Mail == "" || userJob.JobId == 0 {
		utils.Responser.Fail(ctx, resultcode.DataReceiveFail)
		return
	}
	e := global.GVA_DB
	has, err := e.Get(&user)
	if !has || err != nil {
		utils.Responser.Fail(ctx, resultcode.DataSelectFail, err)
		return
	}
	userJob.UserId = user.Id
	affect, err := e.Insert(&userJob)
	if affect <= 0 || err != nil {
		utils.Responser.Fail(ctx, resultcode.DataCreateFail, err)
		return
	}
	utils.Responser.Ok(ctx)
}

// @Tags Job
// @Summary 删除员工
// @Security ApiKeyAuth
// @accept application/x-www-form-urlencoded
// @Produce application/json
// @Param UserId body int true "用户id"
// @Param JobId body int true "职位id"
// @Success 200 {object} utils.Response{}
// @Router /workspace/job/delete-user [POST]
func JobDeleteUser(ctx iris.Context) {
	var userJob JobModel.SdUserJob
	if err := ctx.ReadForm(&userJob); err != nil || userJob.UserId == 0 || userJob.JobId == 0 {
		utils.Responser.Fail(ctx, resultcode.DataReceiveFail)
		return
	}
	e := global.GVA_DB
	affect, err := e.Delete(&userJob)
	if affect <= 0 || err != nil {
		utils.Responser.Fail(ctx, resultcode.DataDeleteFail, err)
		return
	}
	utils.Responser.Ok(ctx)
}
