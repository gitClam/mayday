package workspace_routes

import (
	"github.com/kataras/iris/v12/core/router"
	workSpaceApi "mayday/src/api/v1/workspace"
)

func InitApplicationRouter(workspace router.Party) {

	application := workspace.Party("/application")
	{
		application.Get("/select", workSpaceApi.ApplicationSelect)                              //查询应用信息
		application.Get("/select/workspace", workSpaceApi.ApplicationSelectWorkspace)           //查询应用信息
		application.Post("/create", workSpaceApi.ApplicationCreate)                             //创建应用
		application.Post("/editor", workSpaceApi.ApplicationEditor)                             //修改应用信息
		application.Delete("/delete", workSpaceApi.ApplicationDelete)                           //删除应用
		application.Post("/delete-workflow", workSpaceApi.ApplicationDeleteWorkflow)            //删除流程
		application.Post("/insert-workflow", workSpaceApi.ApplicationInsert)                    //添加流程
		application.Get("/select/workFlowByWorkspaceId", workSpaceApi.GetWorkflowByWorkspaceId) //根据工作空间直接查询流程信息
	}
}
