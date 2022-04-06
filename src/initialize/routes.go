package initialize

import (
	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"
	_ "mayday/docs"
	"mayday/src/middleware"
	UserRouter "mayday/src/router/user"
	WorkflowRouter "mayday/src/router/workflow"
)

func Routers(app *iris.Application) {

	app.Get("/swagger/{any:path}", swagger.WrapHandler(swaggerFiles.Handler))

	main := app.Party("/")
	// Option 请求直接返回
	main.Options("/*", func(ctx iris.Context) {})
	//解决跨域和路由拦截
	main.Use(middleware.Cors, middleware.ServeHTTP)
	UserRouter.InitUserRouter(main)         //用户路由
	WorkflowRouter.InitWorkflowRouter(main) //流程路由
	WorkflowRouter.InitTableRouter(main)    //表单路由

	//workspace := main.Party("/workspace")
	//{
	//	select1 := workspace.Party("/select")
	//	{
	//		select1.Get("/user", workspace_routes.WorkspaceSelectWorkspaceUserid) //根据用户ID获取工作空间信息
	//		select1.Post("/workspace", workspace_routes.WorkspaceSelectWorkspace) //直接根据ID查询
	//	}
	//
	//	create := workspace.Party("/create")
	//	{
	//		create.Post("/workspace", workspace_routes.WorkspaceCreate) //创建工作空间
	//	}
	//
	//	editor := workspace.Party("/editor")
	//	{
	//		editor.Post("/workspace", workspace_routes.WorkspaceEditor) //修改工作空间信息
	//	}
	//
	//	delete := workspace.Party("/delete")
	//	{
	//		delete.Post("/workspace", workspace_routes.WorkspaceDelete) //删除工作空间
	//	}
	//
	//	department := workspace.Party("/department")
	//	{
	//		department.Post("/select", workspace_routes.DepartmentSelect)                    //查询部门信息
	//		department.Post("/select/workspace", workspace_routes.DepartmentSelectWorkspace) //查询部门信息
	//		department.Get("/select/user", workspace_routes.DepartmentSelectUser)            //查询部门信息
	//		department.Post("/create", workspace_routes.DepartmentCreate)                    //创建部门
	//		department.Post("/editor", workspace_routes.DepartmentEditor)                    //修改部门信息
	//		department.Post("/delete", workspace_routes.DepartmentDelete)                    //删除部门
	//	}
	//
	//	job := workspace.Party("/job")
	//	{
	//		job.Post("/select", workspace_routes.JobSelect)                        //查询职位信息
	//		job.Post("/select-user", workspace_routes.JobSelectUser)               //查询职位信息
	//		job.Post("/select/department", workspace_routes.Job_select_department) //查询职位信息
	//		job.Get("/select/user", workspace_routes.Job_select_user)              //查询职位信息
	//		job.Post("/create", workspace_routes.Job_create)                       //创建职位
	//		job.Post("/editor", workspace_routes.JobEditor)                        //修改职位信息
	//		job.Post("/delete", workspace_routes.JobDelete)                        //删除职位
	//		job.Post("/insert-user", workspace_routes.JobInsert)                   //添加用户
	//		job.Post("/delete-user", workspace_routes.JobDeleteUser)               //删除
	//	}
	//
	//	application := workspace.Party("/application")
	//	{
	//		application.Post("/select", workspace_routes.ApplicationSelect)                    //查询应用信息
	//		application.Post("/select/workspace", workspace_routes.ApplicationSelectWorkspace) //查询应用信息
	//		application.Post("/create", workspace_routes.ApplicationCreate)                    //创建应用
	//		application.Post("/editor", workspace_routes.ApplicationEditor)                    //修改应用信息
	//		application.Delete("/delete", workspace_routes.ApplicationDelete)                  //删除应用
	//		application.Delete("/delete-workflow", workspace_routes.ApplicationDeleteWorkflow) //删除流程
	//		application.Post("/insert-workflow", workspace_routes.ApplicationInsert)           //添加流程
	//	}
	//}
}
