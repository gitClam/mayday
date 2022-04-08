package workspace_routes

import (
	"github.com/kataras/iris/v12/core/router"
	workSpaceApi "mayday/src/api/v1/workspace"
)

func InitDepartmentRouter(workspace router.Party) {
	department := workspace.Party("/department")
	{
		department.Get("/select", workSpaceApi.DepartmentSelect)                    //查询部门信息
		department.Get("/select/workspace", workSpaceApi.DepartmentSelectWorkspace) //查询部门信息
		department.Get("/select/user", workSpaceApi.DepartmentSelectUser)           //查询部门信息
		department.Post("/create", workSpaceApi.DepartmentCreate)                   //创建部门
		department.Post("/editor", workSpaceApi.DepartmentEditor)                   //修改部门信息
		department.Delete("/delete", workSpaceApi.DepartmentDelete)                 //删除部门
	}
}
