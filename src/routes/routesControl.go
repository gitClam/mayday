package route_Controller

import (
	"mayday/middleware"
	"mayday/src/routes/user"
	"mayday/src/routes/workflow"
	"mayday/src/routes/workspace"

	"github.com/kataras/iris/v12"
	//"github.com/iris-contrib/middleware/cors"
)

func Hub(app *iris.Application) {
	main := app.Party("/")
	main.Options("/*", func(ctx iris.Context) {})
	main.Use(middleware.ServeHTTP)
	user := main.Party("/user")
	{
		user.Post("/registe", user_routes.User_registe)             //用户注册
		user.Post("/login", user_routes.User_login)                 //用户登录
		user.Delete("/cancellation", user_routes.User_cancellation) //用户注销
		user.Post("/editor/message", user_routes.Set_User_message)  //修改用户信息
		user.Get("/message", user_routes.User_message)              //获取用户信息
		user.Get("/photo/{id:int}", user_routes.User_photo)         //获取用户头像
		user.Post("/set_photo", user_routes.Set_user_photo)         //设置用户头像头像
	}
	workflow := main.Party("/workflow")
	{
		select1 := workflow.Party("/select")
		{
			select1.Get("/workflow", workflow_routes.Workflow_select_workflow)                      //查询流程（已发布）
			select1.Post("/table", workflow_routes.Workflow_select_table)                           //查询表单（已发布）
			select1.Get("/workflow-draft", workflow_routes.Workflow_select_workflow_draft)          //查询流程（草稿）
			select1.Post("/table-draft", workflow_routes.Workflow_select_table_draft)               //查询表单（草稿）
			select1.Get("/workflow-byId", workflow_routes.Workflow_select_workflow_byId)            //查询流程（已发布）
			select1.Post("/table-workSpace", workflow_routes.Workflow_select_table_draft_workSpace) //查询表单（草稿）
			select1.Post("/table-draft-user", workflow_routes.Workflow_select_table_draft_user)     //查询表单1（草稿）
		}
		create := workflow.Party("/create")
		{
			create.Post("/workflow", workflow_routes.Workflow_create_workflow)             //创建流程（发布）
			create.Post("/table", workflow_routes.Workflow_create_table)                   //创建表单（发布）
			create.Post("/workflow-draft", workflow_routes.Workflow_create_workflow_draft) //创建流程（草稿）
			create.Post("/table-draft", workflow_routes.Workflow_create_table_draft)       //创建表单（草稿）
		}

		editor := workflow.Party("/editor")
		{
			editor.Post("/workflow", workflow_routes.Workflow_editor_workflow)             //修改流程（已发布）
			editor.Post("/workflow-state", workflow_routes.Workflow_editor_workflow_state) //修改流程状态
			editor.Post("/table", workflow_routes.Workflow_editor_table)                   //修改表单（已发布）
			editor.Post("/workflow-draft", workflow_routes.Workflow_editor_workflow_draft) //修改流程（草稿）
			editor.Post("/table-draft", workflow_routes.Workflow_editor_table_draft)       //修改表单（草稿）
		}

		delete := workflow.Party("/delete")
		{
			delete.Post("/workflow", workflow_routes.Workflow_delete_workflow)             //删除流程（已发布）
			delete.Post("/table", workflow_routes.Workflow_delete_table)                   //删除表单（已发布）
			delete.Post("/workflow-draft", workflow_routes.Workflow_delete_workflow_draft) //删除流程（草稿）
			delete.Post("/table-draft", workflow_routes.Workflow_delete_table_draft)       //删除表单（草稿）
		}

		participate := workflow.Party("/order")
		{
			participate.Post("/create-order", workflow_routes.Workflow_order_create_order) //创建流程申请
			participate.Post("/fill-table", workflow_routes.Workflow_order_fill_table)     //填写表单（会修改流程状态）
			participate.Get("/notification", workflow_routes.Workflow_order_notification)  //获取消息提醒
			participate.Get("/order-state", workflow_routes.Workflow_order_order_state)    //获取流程状态
		}

	}
	workspace := main.Party("/workspace")
	{
		select1 := workspace.Party("/select")
		{
			select1.Get("/user", workspace_routes.Workspace_select_workspace_userId) //根据用户ID获取工作空间信息
			select1.Post("/workspace", workspace_routes.Workspace_select_workspace)  //直接根据ID查询
		}

		create := workspace.Party("/create")
		{
			create.Post("/workspace", workspace_routes.Workspace_create) //创建工作空间
		}

		editor := workspace.Party("/editor")
		{
			editor.Post("/workspace", workspace_routes.Workspace_editor) //修改工作空间信息
		}

		delete := workspace.Party("/delete")
		{
			delete.Post("/workspace", workspace_routes.Workspace_delete) //删除工作空间
		}

		department := workspace.Party("/department")
		{
			department.Post("/select", workspace_routes.Department_select)                     //查询部门信息
			department.Post("/select/workspace", workspace_routes.Department_select_workspace) //查询部门信息
			department.Get("/select/user", workspace_routes.Department_select_user)            //查询部门信息
			department.Post("/create", workspace_routes.Department_create)                     //创建部门
			department.Post("/editor", workspace_routes.Department_editor)                     //修改部门信息
			department.Post("/delete", workspace_routes.Department_delete)                     //删除部门
		}

		job := workspace.Party("/job")
		{
			job.Post("/select", workspace_routes.Job_select)                       //查询职位信息
			job.Post("/select-user", workspace_routes.Job_selectUser)              //查询职位信息
			job.Post("/select/department", workspace_routes.Job_select_department) //查询职位信息
			job.Get("/select/user", workspace_routes.Job_select_user)              //查询职位信息
			job.Post("/create", workspace_routes.Job_create)                       //创建职位
			job.Post("/editor", workspace_routes.Job_editor)                       //修改职位信息
			job.Post("/delete", workspace_routes.Job_delete)                       //删除职位
			job.Post("/insert-user", workspace_routes.Job_insert)                  //添加用户
			job.Post("/delete-user", workspace_routes.Job_delete_user)             //删除
		}

		application := workspace.Party("/application")
		{
			application.Post("/select", workspace_routes.Application_select)                     //查询应用信息
			application.Post("/select/workspace", workspace_routes.Application_select_workspace) //查询应用信息
			application.Post("/create", workspace_routes.Application_create)                     //创建应用
			application.Post("/editor", workspace_routes.Application_editor)                     //修改应用信息
			application.Delete("/delete", workspace_routes.Application_delete)                   //删除应用
			application.Delete("/delete-workflow", workspace_routes.Application_delete_workflow) //删除流程
			application.Post("/insert-workflow", workspace_routes.Application_insert)            //添加流程
		}
	}
}
