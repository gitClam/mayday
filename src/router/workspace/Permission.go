package workspace_routes

import (
	"github.com/kataras/iris/v12/core/router"
	workSpaceApi "mayday/src/api/v1/workspace"
)

func InitPermissionRouter(workspace router.Party) {

	permission := workspace.Party("/permission")
	{
		permission.Get("/select", workSpaceApi.SelectAdmin)    //获取工作空间管理员（必须本身是管理员）
		permission.Post("/create", workSpaceApi.SetAdmin)      //设置工作空间管理员（必须本身是管理员）
		permission.Delete("/delete", workSpaceApi.RemoveAdmin) //删除工作空间管理员（必须本身是管理员）
	}
}
