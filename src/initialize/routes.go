package initialize

import (
	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"
	_ "mayday/docs"
	"mayday/src/middleware"
	OrderRouter "mayday/src/router/order"
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
	OrderRouter.InitOrderRouter(main)         //流程实例(事件)路由

}
