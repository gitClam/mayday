package initialize

import (
	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"
	_ "mayday/docs"
	"mayday/src/middleware"
	UserRouter "mayday/src/router/user"
	WorkflowRouter "mayday/src/router/workflow"
	WorkspaceRouter "mayday/src/router/workspace"
)

func Routers(app *iris.Application) {

	app.Get("/swagger/{any:path}", swagger.WrapHandler(swaggerFiles.Handler))

	main := app.Party("/")
	// Option 请求直接返回
	main.Options("/*", func(ctx iris.Context) {})
	//解决跨域和路由拦截
	main.Use(middleware.Cors, middleware.ServeHTTP)
	UserRouter.InitUserRouter(main)           //用户路由
	WorkflowRouter.InitWorkflowRouter(main)   //流程路由
	WorkflowRouter.InitTableRouter(main)      //表单路由
	WorkspaceRouter.InitWorkspaceRouter(main) //工作空间路由

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
