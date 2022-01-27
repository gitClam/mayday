package workspace_routes

import (
	"github.com/kataras/iris/v12"
	"log"
	"mayday/src/global"
	"mayday/src/middleware"
	"mayday/src/model"
	userModel "mayday/src/model/user"
	"mayday/src/utils"
	"strconv"
)

// swagger:operation Post /workspace/job/select job job_select
// ---
// summary: 获取职位信息
// description: 获取职位信息
// parameters:
// - name: id
//   description: 职位id
//   type: string
//   required: true
func JobSelect(ctx iris.Context) {
	var job model.SdJob
	if err := ctx.ReadForm(&job); err != nil || job.Id == 0 {
		utils.Responser.FailWithMsg(ctx, utils.OptionFailur)
		log.Print("数据接收失败")
		return
	}
	e := global.GVA_DB
	has, err := e.Where("is_deleted = 0").Get(&job)
	if !has || err != nil {
		log.Printf("数据库查询错误")
		utils.Responser.FailWithMsg(ctx, utils.OptionFailur)
		return
	}
	utils.Responser.OkWithDetails(ctx, utils.Success, job)
}

// swagger:operation Post /workspace/job/select/department job Job_select_department
// ---
// summary: 获取职位信息
// description: 获取职位信息
// parameters:
// - name: id
//   description: 部门id
//   type: string
//   required: true
func Job_select_department(ctx iris.Context) {
	var department model.SdDepartment
	if err := ctx.ReadForm(&department); err != nil || department.Id == 0 {
		utils.Responser.FailWithMsg(ctx, utils.OptionFailur)
		log.Print("数据接收失败")
		return
	}

	var job []model.SdJob

	e := global.GVA_DB
	err := e.Where("department_id = ?", department.Id).And("is_delete = 0").Find(&job)
	if err != nil {
		log.Print(err)
		log.Printf("数据库查询错误")
		utils.Responser.FailWithMsg(ctx, utils.OptionFailur)
		return
	}
	utils.Responser.OkWithDetails(ctx, utils.Success, job)
}

// swagger:operation GET /workspace/job/select/user job Job_select_user
// ---
// summary: 获取用户职位信息
// description: 获取用户职位信息
func Job_select_user(ctx iris.Context) {
	user, ok := middleware.ParseToken(ctx)
	if !ok {
		log.Printf("解析TOKEN出错，请重新登录")
		utils.Responser.FailWithMsg(ctx, "解析TOKEN出错，请重新登录")
		return
	}
	log.Print(user)
	var job []model.SdJob

	e := global.GVA_DB
	err := e.Sql("select * from sd_job where id in(select job_id from sd_user_job where user_id = ?)", user.Id).Find(&job)
	if err != nil {
		log.Printf("数据库查询错误")
		utils.Responser.FailWithMsg(ctx, utils.OptionFailur)
		return
	}

	utils.Responser.OkWithDetails(ctx, utils.Success, job)
}

// swagger:operation POST /workspace/job/create job job_create
// ---
// summary: 创建职位
// description: 创建职位
// parameters:
// - name: DepartmentId
//   description: 部门id
//   type: string
//   required: true
// - name: name
//   description: 名字
//   type: string
//   required: true
func Job_create(ctx iris.Context) {
	var job model.SdJob
	if err := ctx.ReadForm(&job); err != nil {
		utils.Responser.FailWithMsg(ctx, utils.OptionFailur)
		log.Print("数据接收失败")
		return
	}
	job.IsDelete = 0
	e := global.GVA_DB
	affect, err := e.Insert(&job)
	if affect <= 0 || err != nil {
		log.Printf("数据库插入错误")
		utils.Responser.FailWithMsg(ctx, utils.OptionFailur)
		return
	}
	utils.Responser.Ok(ctx)
}

// swagger:operation POST /workspace/job/editor job job_editor
// ---
// summary: 修改职位信息
// description: 修改职位信息
// parameters:
// - name: DepartmentId
//   description: 部门id
//   type: string
//   required: false
// - name: name
//   description: 名字
//   type: string
//   required: false
func JobEditor(ctx iris.Context) {
	var job model.SdJob
	if err := ctx.ReadForm(&job); err != nil || job.Id == 0 {
		utils.Responser.FailWithMsg(ctx, utils.OptionFailur)
		log.Print("数据接收失败")
		return
	}
	e := global.GVA_DB
	affect, err := e.Id(job.Id).Update(&job)
	if affect <= 0 || err != nil {
		log.Printf("数据库插入错误")
		utils.Responser.FailWithMsg(ctx, utils.OptionFailur)
		return
	}
	utils.Responser.Ok(ctx)
}

// swagger:operation DELETE /workspace/job/delete job job_delete
// ---
// summary: 删除职位
// description: 删除职位
// parameters:
// - name: Id
//   description: 职位id
//   type: string
//   required: false
func JobDelete(ctx iris.Context) {
	var job model.SdJob
	if err := ctx.ReadForm(&job); err != nil || job.Id == 0 {
		utils.Responser.FailWithMsg(ctx, utils.OptionFailur)
		log.Print("数据接收失败")
		return
	}
	e := global.GVA_DB.NewSession()
	defer e.Close()
	e.Begin()
	_, err := e.Exec("delete from sd_user_job where job_id = ?", job.Id)
	if err != nil {
		e.Rollback()
		log.Print(err)
		log.Printf("数据库插入错误")
		utils.Responser.FailWithMsg(ctx, "")
		return
	}
	_, err = e.Exec("delete from sd_job where id = ?", job.Id)
	if err != nil {
		err := e.Rollback()
		if err != nil {
			return
		}
		log.Print(err)
		log.Printf("数据库插入错误")
		utils.Responser.FailWithMsg(ctx, "")
		return
	}
	e.Commit()
	utils.Responser.Ok(ctx)
}

// swagger:operation Post /workspace/job/select-user job selectUser_Job
// ---
// summary: 获取职位人员信息
// description: 获取职位人员信息
// parameters:
// - name: id
//   description: 职位id
//   type: string
//   required: true
func JobSelectUser(ctx iris.Context) {
	var job model.SdJob
	if err := ctx.ReadForm(&job); err != nil || job.Id == 0 {
		utils.Responser.FailWithMsg(ctx, "")
		log.Print("数据接收失败")
		return
	}

	var users []userModel.SdUser
	e := global.GVA_DB
	err := e.Sql("select * from sd_user where id in(select user_id from sd_user_job where job_id = ?)", job.Id).Find(&users)
	if err != nil {
		log.Printf("数据库查询错误")
		utils.Responser.FailWithMsg(ctx, "")
		return
	}
	utils.Responser.OkWithDetails(ctx, utils.Success, userModel.TransformUserVOList(users))
}

// swagger:operation POST /workspace/job/insert-user job insert_job
// ---
// summary: 添加员工
// description: 添加员工
// parameters:
// - name: Mail
//   description: 用户邮箱
//   type: string
//   required: true
// - name: JobId
//   description: 职位id
//   type: string
//   required: true
func JobInsert(ctx iris.Context) {
	var userJob model.SdUserJob
	var user userModel.SdUser

	user.Mail = ctx.FormValue("Mail")
	userJob.JobId, _ = strconv.Atoi(ctx.FormValue("JobId"))
	if user.Mail == "" || userJob.JobId == 0 {
		utils.Responser.FailWithMsg(ctx, utils.OptionFailur)
		log.Print("数据接收失败")
		return
	}
	e := global.GVA_DB
	has, err := e.Get(&user)
	if !has || err != nil {
		log.Printf("数据库查询错误")
		utils.Responser.FailWithMsg(ctx, utils.OptionFailur)
		return
	}
	userJob.UserId = user.Id
	affect, err := e.Insert(&userJob)
	if affect <= 0 || err != nil {
		log.Printf("数据库插入错误")
		utils.Responser.FailWithMsg(ctx, "")
		return
	}
	utils.Responser.Ok(ctx)
}

// swagger:operation POST /workspace/job/delete-user job delete_user_Job
// ---
// summary: 删除员工
// description: 删除员工
// parameters:
// - name: UserId
//   description: 用户ID
//   type: string
//   required: true
// - name: JobId
//   description: 职位id
//   type: string
//   required: true
func JobDeleteUser(ctx iris.Context) {
	var userJob model.SdUserJob
	if err := ctx.ReadForm(&userJob); err != nil || userJob.UserId == 0 || userJob.JobId == 0 {
		utils.Responser.FailWithMsg(ctx, "")
		log.Print("数据接收失败")
		return
	}
	e := global.GVA_DB
	affect, err := e.Delete(&userJob)
	if affect <= 0 || err != nil {
		log.Printf("数据库插入错误")
		utils.Responser.FailWithMsg(ctx, "")
		return
	}
	utils.Responser.Ok(ctx)
}
