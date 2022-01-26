package workspace_routes

import (
	"github.com/kataras/iris/v12"
	"log"
	"mayday/src/db/conn"
	"mayday/src/middleware/jwts"
	"mayday/src/model"
	"mayday/src/utils"
)

const (
	SelectSql string = "select * from sd_department where id in (select department_id from sd_job where id in(select job_id from sd_user_job where user_id = ?))"
)

// swagger:operation GET /workspace/department/select/user department select_user_Department
// ---
// summary: 获取部门信息
// description: 根据用户ID获取部门信息
func DepartmentSelectUser(ctx iris.Context) {

	user, ok := jwts.ParseToken(ctx)
	if !ok {
		log.Printf("解析TOKEN出错，请重新登录")
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.TokenParseFailur, nil)
		return
	}

	var department []model.SdDepartment

	e := conn.MasterEngine()
	err := e.SQL(SelectSql, user.Id).Find(&department)
	if err != nil {
		log.Printf("数据库查询错误")
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}

	utils.MakeSuccessRes(ctx, model.Success, department)
}

// swagger:operation Post /workspace/department/select department select_department
// ---
// summary: 获取部门信息
// description: 获取部门信息
// parameters:
// - name: id
//   description: 部门id
//   type: string
//   required: true
func DepartmentSelect(ctx iris.Context) {
	var department model.SdDepartment
	if err := ctx.ReadForm(&department); err != nil || department.Id == 0 {
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		log.Print("数据接收失败")
		return
	}
	e := conn.MasterEngine()
	has, err := e.Where("is_deleted = 0").Get(&department)
	if !has || err != nil {
		log.Printf("数据库查询错误")
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	utils.MakeSuccessRes(ctx, model.Success, department)
}

// swagger:operation Post /workspace/department/select/workspace department select_workspace_Department
// ---
// summary: 获取部门信息
// description: 获取部门信息
// parameters:
// - name: id
//   description: 工作空间id
//   type: string
//   required: true
func DepartmentSelectWorkspace(ctx iris.Context) {
	var Workspace model.SdWorkspace
	if err := ctx.ReadForm(&Workspace); err != nil || Workspace.Id == 0 {
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		log.Print("数据接收失败")
		return
	}

	var departments []model.SdDepartment

	e := conn.MasterEngine()
	err := e.Where("workspace_id = ? and is_deleted = 0", Workspace.Id).Find(&departments)
	if err != nil {
		log.Print(err)
		log.Printf("数据库查询错误")
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	utils.MakeSuccessRes(ctx, model.Success, departments)
}

// swagger:operation POST /workspace/department/create department create_department
// ---
// summary: 创建部门
// description: 创建部门
// parameters:
// - name: WorkspaceId
//   description: 工作空间id
//   type: string
//   required: true
// - name: Name
//   description: 部门名字
//   type: string
//   required: true
// - name: Phone
//   description: 联系电话
//   type: string
//   required: false
// - name: Remark
//   description: 备注
//   type: string
//   required: false
func DepartmentCreate(ctx iris.Context) {
	var department model.SdDepartment
	if err := ctx.ReadForm(&department); err != nil {
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		log.Print("数据接收失败")
		return
	}
	department.IsDeleted = 0
	e := conn.MasterEngine()
	affect, err := e.Insert(&department)
	if affect <= 0 || err != nil {
		log.Printf("数据库插入错误")
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	utils.MakeSuccessRes(ctx, model.Success, nil)
}

// swagger:operation POST /workspace/department/editor department editor_department
// ---
// summary: 修改部门信息
// description: 修改部门信息
// parameters:
// - name: Id
//   description: 部门id
//   type: string
//   required: true
// - name: WorkspaceId
//   description: 工作空间id
//   type: string
//   required: true
// - name: Name
//   description: 部门名字
//   type: string
//   required: true
// - name: Phone
//   description: 联系电话
//   type: string
//   required: false
// - name: Remark
//   description: 备注
//   type: string
//   required: false
func DepartmentEditor(ctx iris.Context) {
	var department model.SdDepartment
	if err := ctx.ReadForm(&department); err != nil || department.Id == 0 {
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		log.Print("数据接收失败")
		return
	}
	e := conn.MasterEngine()
	affect, err := e.Id(department.Id).Update(&department)
	if affect <= 0 || err != nil {
		log.Printf("数据库插入错误")
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	utils.MakeSuccessRes(ctx, model.Success, nil)
}

// swagger:operation DELETE /workspace/department/delete department delete_department
// ---
// summary: 删除部门
// description: 删除部门
// parameters:
// - name: Id
//   description: 部门id
//   type: string
//   required: true
// - name: WorkspaceId
func DepartmentDelete(ctx iris.Context) {
	var department model.SdDepartment
	if err := ctx.ReadForm(&department); err != nil || department.Id == 0 {
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		log.Print("数据接收失败")
		return
	}
	e := conn.MasterEngine().NewSession()
	defer e.Close()
	e.Begin()
	_, err := e.Exec("delete from sd_user_job where job_id in( select id from sd_job where department_id = ?)", department.Id)
	if err != nil {
		e.Rollback()
		log.Print(err)
		log.Printf("数据库插入错误")
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	_, err = e.Exec("delete from sd_job where department_id = ?", department.Id)
	if err != nil {
		e.Rollback()
		log.Print(err)
		log.Printf("数据库插入错误")
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	_, err = e.Exec("delete from sd_department where id = ?", department.Id)
	if err != nil {
		e.Rollback()
		log.Print(err)
		log.Printf("数据库插入错误")
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	e.Commit()
	utils.MakeSuccessRes(ctx, model.Success, nil)
}
