package workspace_routes

import (
	"github.com/kataras/iris/v12"
	"log"
	"mayday/src/global"
	"mayday/src/middleware/jwts"
	"mayday/src/model"
	"mayday/src/utils"
)

// swagger:operation GET /workspace/select/user workspace select_workspace_userId
// ---
// summary: 获取工作空间信息
// description: 根据用户ID获取工作空间信息
func WorkspaceSelectWorkspaceUserid(ctx iris.Context) {

	user, ok := jwts.ParseToken(ctx)
	if !ok {
		log.Printf("解析TOKEN出错，请重新登录")
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.TokenParseFailur, nil)
		return
	}

	var Workspace []model.SdWorkspace

	e := global.GVA_DB
	err := e.SQL("select * from sd_workspace where id in (select workspace_id from sd_department where id in (select department_id from sd_job where id in(select job_id from sd_user_job where user_id = ?)))", user.Id).Find(&Workspace)
	if err != nil {
		log.Printf("数据库查询错误")
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}

	utils.MakeSuccessRes(ctx, model.Success, Workspace)
}

// swagger:operation Post /workspace/select/workspace workspace select_workspace_workspaceId
// ---
// summary: 获取工作空间信息
// description: 根据ID获取工作空间信息
// parameters:
// - name: id
//   description: 工作空间id
//   type: string
//   required: true
func WorkspaceSelectWorkspace(ctx iris.Context) {
	var Workspace model.SdWorkspace
	if err := ctx.ReadForm(&Workspace); err != nil || Workspace.Id == 0 {
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		log.Print("数据接收失败")
		return
	}
	e := global.GVA_DB
	has, err := e.Where("is_deleted = 0").Get(&Workspace)
	if !has || err != nil {
		log.Printf("数据库查询错误")
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	utils.MakeSuccessRes(ctx, model.Success, Workspace)
}

// swagger:operation Post /workspace/create/workspace workspace create_workspace
// ---
// summary: 创建工作空间
// description: 创建工作空间
// parameters:
// - name: name
//   description: 工作空间名字
//   type: string
//   required: true
// - name: phone
//   description: 联系电话
//   type: string
//   required: false
// - name: remark
//   description: 备注
//   type: string
//   required: false
func WorkspaceCreate(ctx iris.Context) {
	user, ok := jwts.ParseToken(ctx)
	if !ok {
		log.Printf("解析TOKEN出错，请重新登录")
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.TokenParseFailur, nil)
		return
	}
	var Workspace model.SdWorkspace
	if err := ctx.ReadForm(&Workspace); err != nil {
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		log.Print("数据接收失败")
		return
	}
	Workspace.IsDeleted = 0
	e := global.GVA_DB.NewSession()
	defer e.Close()
	e.Begin()
	log.Print(Workspace)
	affect, err := e.Insert(&Workspace)
	if affect <= 0 || err != nil {
		e.Rollback()
		log.Print(err)
		log.Printf("数据库插入错误")
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	has, err1 := e.Get(&Workspace)
	if !has || err1 != nil {
		e.Rollback()
		log.Printf("数据库插入错误")
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	department := model.SdDepartment{
		WorkspaceId: Workspace.Id,
		Name:        "默认",
		Phone:       "0000000",
		Remark:      "昨木",
		IsDeleted:   0}
	affect, err = e.Insert(&department)
	if affect <= 0 || err != nil {
		e.Rollback()
		log.Printf("数据库插入错误")
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	has, err1 = e.Get(&department)
	if !has || err1 != nil {
		e.Rollback()
		log.Printf("数据库插入错误")
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	job := model.SdJob{
		DepartmentId: department.Id,
		Name:         "默认职位",
		IsDelete:     0}
	affect, err = e.Insert(&job)
	if affect <= 0 || err != nil {
		e.Rollback()
		log.Print(err)
		log.Printf("数据库插入错误")
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	has, err1 = e.Get(&job)
	if !has || err1 != nil {
		e.Rollback()
		log.Printf("数据库插入错误")
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	userJob := model.SdUserJob{
		UserId: user.Id,
		JobId:  job.Id}
	affect, err = e.Insert(&userJob)
	if affect <= 0 || err != nil {
		e.Rollback()
		log.Print(err)
		log.Printf("数据库插入错误")
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	e.Commit()
	utils.MakeSuccessRes(ctx, model.Success, nil)
}

// swagger:operation POST /workspace/editor/workspace workspace editor_workspace
// ---
// summary: 修改工作空间信息
// description: 修改工作空间信息
// parameters:
// - name: id
//   description: 工作空间id
//   type: string
//   required: true
// - name: name
//   description: 工作空间名字
//   type: string
//   required: false
// - name: phone
//   description: 联系电话
//   type: string
//   required: false
// - name: remark
//   description: 备注
//   type: string
//   required: false
func WorkspaceEditor(ctx iris.Context) {
	var Workspace model.SdWorkspace
	if err := ctx.ReadForm(&Workspace); err != nil || Workspace.Id == 0 {
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		log.Print("数据接收失败")
		return
	}
	e := global.GVA_DB
	affect, err := e.Id(Workspace.Id).Update(&Workspace)
	if affect <= 0 || err != nil {
		log.Printf("数据库插入错误")
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	utils.MakeSuccessRes(ctx, model.Success, nil)
}

// swagger:operation POST /workspace/delete/workspace workspace delete_Workspace
// ---
// summary: 删除工作空间
// description: 删除工作空间
// parameters:
// - name: id
//   description: ID
//   type: int
//   required: true
func WorkspaceDelete(ctx iris.Context) {
	var Workspace model.SdWorkspace
	if err := ctx.ReadForm(&Workspace); err != nil || Workspace.Id == 0 {
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		log.Print("数据接收失败")
		return
	}
	e := global.GVA_DB.NewSession()
	defer e.Close()
	e.Begin()
	_, err := e.Exec("delete from sd_user_job where job_id in( select id from sd_job where department_id in (select id from sd_department where workspace_id = ?))", Workspace.Id)
	if err != nil {
		e.Rollback()
		log.Print(err)
		log.Printf("数据库插入错误")
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	_, err = e.Exec("delete from sd_job where department_id in (select id from sd_department where workspace_id = ?)", Workspace.Id)
	if err != nil {
		e.Rollback()
		log.Print(err)
		log.Printf("数据库插入错误")
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	_, err = e.Exec("delete from sd_department where workspace_id = ?", Workspace.Id)
	if err != nil {
		e.Rollback()
		log.Print(err)
		log.Printf("数据库插入错误")
		utils.MakeErrorRes(ctx, iris.StatusInternalServerError, model.OptionFailur, nil)
		return
	}
	_, err = e.Exec("delete from sd_workspace where id = ?", Workspace.Id)
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
