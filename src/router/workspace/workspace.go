package workspace_routes

import "github.com/kataras/iris/v12/core/router"

func InitWorkspaceRouter(Router router.Party) {
	workspace := Router.Party("/workspace")
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
	InitJobRouter(workspace)
	InitDepartmentRouter(workspace)
}
