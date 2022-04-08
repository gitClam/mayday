package workspace_routes

import (
	"github.com/kataras/iris/v12/core/router"
	workSpaceApi "mayday/src/api/v1/workspace"
)

func InitWorkspaceRouter(Router router.Party) {
	workspace := Router.Party("/workspace")
	select1 := workspace.Party("/select")
	{
		select1.Get("/user", workSpaceApi.WorkspaceSelectWorkspaceUserid) //根据用户ID获取工作空间信息
		select1.Get("/workspace", workSpaceApi.WorkspaceSelectWorkspace)  //直接根据ID查询
	}

	create := workspace.Party("/create")
	{
		create.Post("/workspace", workSpaceApi.WorkspaceCreate) //创建工作空间
	}

	editor := workspace.Party("/editor")
	{
		editor.Post("/workspace", workSpaceApi.WorkspaceEditor) //修改工作空间信息
	}

	delete := workspace.Party("/delete")
	{
		delete.Delete("/workspace", workSpaceApi.WorkspaceDelete) //删除工作空间
	}
	InitJobRouter(workspace)
	InitDepartmentRouter(workspace)
	InitApplicationRouter(workspace)
}
