package workspace_routes
import (
	"github.com/kataras/iris/v12"
	"log"
	"mayday/src/db/conn"
	"mayday/src/models"
	"strconv"
	"mayday/src/supports/responser"
	"mayday/middleware/jwts"
	"mayday/src/supports/responser/vo"
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
func Job_select(ctx iris.Context) {
	var job model.SdJob
	if err := ctx.ReadForm(&job); (err != nil || job.Id == 0) {
		responser.MakeErrorRes(ctx,iris.StatusInternalServerError, model.OptionFailur , nil)
		log.Print("数据接收失败")
		return 
	}
	e := conn.MasterEngine()
	has, err := e.Where("is_deleted = 0").Get(&job)
	if (!has || err != nil) {
		log.Printf("数据库查询错误") 
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return 
	} 
	responser.MakeSuccessRes(ctx,model.Success,job)
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
	if err := ctx.ReadForm(&department); (err != nil || department.Id == 0) {
		responser.MakeErrorRes(ctx,iris.StatusInternalServerError, model.OptionFailur , nil)
		log.Print("数据接收失败")
		return 
	}
	
	var job []model.SdJob
	
	e := conn.MasterEngine()
	err := e.Where("department_id = ?",department.Id).And("is_delete = 0").Find(&job)
	if (err != nil) {
		log.Print(err)
		log.Printf("数据库查询错误") 
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return 
	} 
	responser.MakeSuccessRes(ctx,model.Success,job)
}

// swagger:operation GET /workspace/job/select/user job Job_select_user
// --- 
// summary: 获取用户职位信息
// description: 获取用户职位信息
func Job_select_user(ctx iris.Context) {
	user, ok := jwts.ParseToken(ctx)
	if !ok {
		log.Printf("解析TOKEN出错，请重新登录")
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.TokenParseFailur, nil)
		return
	}
	log.Print(user)
	var job []model.SdJob

	e := conn.MasterEngine()
	err := e.Sql("select * from sd_job where id in(select job_id from sd_user_job where user_id = ?)",user.Id).Find(&job)
	if err != nil {
		log.Printf("数据库查询错误") 
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return 
	} 
	
	responser.MakeSuccessRes(ctx,model.Success,job)
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
		responser.MakeErrorRes(ctx,iris.StatusInternalServerError, model.OptionFailur , nil)
		log.Print("数据接收失败")
		return 
	}
	job.IsDelete = 0
	e := conn.MasterEngine()
	affect, err := e.Insert(&job)
	if (affect <= 0 || err != nil) {
		log.Printf("数据库插入错误") 
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return 
	} 
	responser.MakeSuccessRes(ctx,model.Success,nil)
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
func Job_editor(ctx iris.Context) {
	var job model.SdJob
	if err := ctx.ReadForm(&job); (err != nil || job.Id == 0) {
		responser.MakeErrorRes(ctx,iris.StatusInternalServerError, model.OptionFailur , nil)
		log.Print("数据接收失败")
		return 
	}
	e := conn.MasterEngine()
	affect, err := e.Id(job.Id).Update(&job)
	if (affect <= 0 || err != nil) {
		log.Printf("数据库插入错误") 
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return 
	} 
	responser.MakeSuccessRes(ctx,model.Success,nil)
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
func Job_delete(ctx iris.Context) {
	var job model.SdJob
	if err := ctx.ReadForm(&job); (err != nil || job.Id == 0) {
		responser.MakeErrorRes(ctx,iris.StatusInternalServerError, model.OptionFailur , nil)
		log.Print("数据接收失败")
		return 
	}
	e := conn.MasterEngine().NewSession()
	defer e.Close()
	e.Begin()
	_ ,err := e.Exec("delete from sd_user_job where job_id = ?",job.Id)
	if (err != nil) {
		e.Rollback()
		log.Print(err)
		log.Printf("数据库插入错误") 
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return 
	} 
	_ ,err = e.Exec("delete from sd_job where id = ?",job.Id)
	if (err != nil) {
		e.Rollback()
		log.Print(err)
		log.Printf("数据库插入错误") 
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return 
	} 
	e.Commit()
	responser.MakeSuccessRes(ctx,model.Success,nil)
}

// swagger:operation Post /workspace/job/select-user job Job_selectUser
// --- 
// summary: 获取职位人员信息
// description: 获取职位人员信息
// parameters:
// - name: id
//   description: 职位id
//   type: string
//   required: true
func Job_selectUser(ctx iris.Context) {
	var job model.SdJob
	if err := ctx.ReadForm(&job); (err != nil || job.Id == 0) {
		responser.MakeErrorRes(ctx,iris.StatusInternalServerError, model.OptionFailur , nil)
		log.Print("数据接收失败")
		return 
	}

	var users []model.SdUser
	e := conn.MasterEngine()
	err := e.Sql("select * from sd_user where id in(select user_id from sd_user_job where job_id = ?)",job.Id).Find(&users)
	if (err != nil) {
		log.Printf("数据库查询错误") 
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return 
	} 
	responser.MakeSuccessRes(ctx,model.Success,vo.TansformUserVOList(users))
}

// swagger:operation POST /workspace/job/insert-user job job_insert
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
func Job_insert(ctx iris.Context) {
	var userJob model.SdUserJob
	var user model.SdUser
	
	user.Mail = ctx.FormValue("Mail")
	userJob.JobId ,_= strconv.Atoi(ctx.FormValue("JobId"))
	if(user.Mail == "" || userJob.JobId == 0){
		responser.MakeErrorRes(ctx,iris.StatusInternalServerError, model.OptionFailur , nil)
		log.Print("数据接收失败")
		return 
	}
	e := conn.MasterEngine()
	has,err := e.Get(&user) 
	if(!has || err != nil){
		log.Printf("数据库查询错误") 
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return 
	}
	userJob.UserId = user.Id
	affect, err := e.Insert(&userJob)
	if (affect <= 0 || err != nil) {
		log.Printf("数据库插入错误") 
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return 
	} 
	responser.MakeSuccessRes(ctx,model.Success,nil)
}

// swagger:operation POST /workspace/job/delete-user job Job_delete_user
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
func Job_delete_user(ctx iris.Context) {
	var userJob model.SdUserJob
	if err := ctx.ReadForm(&userJob); (err != nil || userJob.UserId == 0 || userJob.JobId == 0) {
		responser.MakeErrorRes(ctx,iris.StatusInternalServerError, model.OptionFailur , nil)
		log.Print("数据接收失败")
		return 
	}
	e := conn.MasterEngine()
	affect, err := e.Delete(&userJob)
	if (affect <= 0 || err != nil) {
		log.Printf("数据库插入错误") 
		responser.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return 
	} 
	responser.MakeSuccessRes(ctx,model.Success,nil)
}