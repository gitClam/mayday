package initialize

import (
	"mayday/src/middleware"
	"mayday/src/router/user"
	"mayday/src/router/workflow"
	"mayday/src/router/workspace"

	"github.com/kataras/iris/v12"
	//"github.com/iris-contrib/middleware/cors"
)

func Routers(app *iris.Application) {
	main := app.Party("/")
	main.Options("/*", func(ctx iris.Context) {})
	main.Use(middleware.ServeHTTP)
	user := main.Party("/user")
	{
		user.Post("/registe", user_routes.UserRegister)            //用户注册
		user.Post("/login", user_routes.UserLogin)                 //用户登录
		user.Delete("/cancellation", user_routes.UserCancellation) //用户注销
		user.Post("/editor/message", user_routes.SetUserMessage)   //修改用户信息
		user.Get("/message", user_routes.UserMessage)              //获取用户信息
		user.Get("/photo/{id:int}", user_routes.UserPhoto)         //获取用户头像
		user.Post("/set_photo", user_routes.SetUserPhoto)          //设置用户头像头像
	}
	workflow := main.Party("/workflow")
	{
		select1 := workflow.Party("/select")
		{
			select1.Get("/workflow", workflow_routes.WorkflowSelectWorkflow)                    //查询流程（已发布）
			select1.Post("/table", workflow_routes.WorkflowSelectTable)                         //查询表单（已发布）
			select1.Get("/workflow-draft", workflow_routes.WorkflowSelectWorkflowDraft)         //查询流程（草稿）
			select1.Post("/table-draft", workflow_routes.WorkflowSelectTableDraft)              //查询表单（草稿）
			select1.Get("/workflow-byId", workflow_routes.WorkflowSelectWorkflowById)           //查询流程（已发布）
			select1.Post("/table-workSpace", workflow_routes.WorkflowSelectTableDraftWorkspace) //查询表单（草稿）
			select1.Post("/table-draft-user", workflow_routes.WorkflowSelectTableDraftUser)     //查询表单1（草稿）
		}
		create := workflow.Party("/create")
		{
			create.Post("/workflow", workflow_routes.WorkflowCreateWorkflow)            //创建流程（发布）
			create.Post("/table", workflow_routes.WorkflowCreateTable)                  //创建表单（发布）
			create.Post("/workflow-draft", workflow_routes.WorkflowCreateWorkflowDraft) //创建流程（草稿）
			create.Post("/table-draft", workflow_routes.WorkflowCreateTableDraft)       //创建表单（草稿）
		}

		editor := workflow.Party("/editor")
		{
			editor.Post("/workflow", workflow_routes.WorkflowEditorWorkflow)            //修改流程（已发布）
			editor.Post("/workflow-state", workflow_routes.WorkflowEditorWorkflowState) //修改流程状态
			editor.Post("/table", workflow_routes.WorkflowEditorTable)                  //修改表单（已发布）
			editor.Post("/workflow-draft", workflow_routes.WorkflowEditorWorkflowDraft) //修改流程（草稿）
			editor.Post("/table-draft", workflow_routes.WorkflowEditorTableDraft)       //修改表单（草稿）
		}

		delete := workflow.Party("/delete")
		{
			delete.Post("/workflow", workflow_routes.WorkflowDeleteWorkflow)            //删除流程（已发布）
			delete.Post("/table", workflow_routes.WorkflowDeleteTable)                  //删除表单（已发布）
			delete.Post("/workflow-draft", workflow_routes.WorkflowDeleteWorkflowDraft) //删除流程（草稿）
			delete.Post("/table-draft", workflow_routes.WorkflowDeleteTableDraft)       //删除表单（草稿）
		}

		participate := workflow.Party("/order")
		{
			participate.Post("/create-order", workflow_routes.WorkflowOrderCreateOrder) //创建流程申请
			participate.Post("/fill-table", workflow_routes.WorkflowOrderFillTable)     //填写表单（会修改流程状态）
			participate.Get("/notification", workflow_routes.WorkflowOrderNotification) //获取消息提醒
			participate.Get("/order-state", workflow_routes.WorkflowOrderOrderState)    //获取流程状态
		}

	}
	workspace := main.Party("/workspace")
	{
		select1 := workspace.Party("/select")
		{
			select1.Get("/user", workspace_routes.WorkspaceSelectWorkspaceUserid) //根据用户ID获取工作空间信息
			select1.Post("/workspace", workspace_routes.WorkspaceSelectWorkspace) //直接根据ID查询
		}

		create := workspace.Party("/create")
		{
			create.Post("/workspace", workspace_routes.WorkspaceCreate) //创建工作空间
		}

		editor := workspace.Party("/editor")
		{
			editor.Post("/workspace", workspace_routes.WorkspaceEditor) //修改工作空间信息
		}

		delete := workspace.Party("/delete")
		{
			delete.Post("/workspace", workspace_routes.WorkspaceDelete) //删除工作空间
		}

		department := workspace.Party("/department")
		{
			department.Post("/select", workspace_routes.DepartmentSelect)                    //查询部门信息
			department.Post("/select/workspace", workspace_routes.DepartmentSelectWorkspace) //查询部门信息
			department.Get("/select/user", workspace_routes.DepartmentSelectUser)            //查询部门信息
			department.Post("/create", workspace_routes.DepartmentCreate)                    //创建部门
			department.Post("/editor", workspace_routes.DepartmentEditor)                    //修改部门信息
			department.Post("/delete", workspace_routes.DepartmentDelete)                    //删除部门
		}

		job := workspace.Party("/job")
		{
			job.Post("/select", workspace_routes.JobSelect)                        //查询职位信息
			job.Post("/select-user", workspace_routes.JobSelectUser)               //查询职位信息
			job.Post("/select/department", workspace_routes.Job_select_department) //查询职位信息
			job.Get("/select/user", workspace_routes.Job_select_user)              //查询职位信息
			job.Post("/create", workspace_routes.Job_create)                       //创建职位
			job.Post("/editor", workspace_routes.Job_editor)                       //修改职位信息
			job.Post("/delete", workspace_routes.Job_delete)                       //删除职位
			job.Post("/insert-user", workspace_routes.JobInsert)                   //添加用户
			job.Post("/delete-user", workspace_routes.JobDeleteUser)               //删除
		}

		application := workspace.Party("/application")
		{
			application.Post("/select", workspace_routes.ApplicationSelect)                    //查询应用信息
			application.Post("/select/workspace", workspace_routes.ApplicationSelectWorkspace) //查询应用信息
			application.Post("/create", workspace_routes.ApplicationCreate)                    //创建应用
			application.Post("/editor", workspace_routes.ApplicationEditor)                    //修改应用信息
			application.Delete("/delete", workspace_routes.ApplicationDelete)                  //删除应用
			application.Delete("/delete-workflow", workspace_routes.ApplicationDeleteWorkflow) //删除流程
			application.Post("/insert-workflow", workspace_routes.ApplicationInsert)           //添加流程
		}
	}
}
