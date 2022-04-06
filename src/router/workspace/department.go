package workspace_routes

import "github.com/kataras/iris/v12/core/router"

func InitDepartmentRouter(workspace router.Party) {
	department := workspace.Party("/department")
	{
		department.Post("/select", workspace_routes.DepartmentSelect)                    //查询部门信息
		department.Post("/select/workspace", workspace_routes.DepartmentSelectWorkspace) //查询部门信息
		department.Get("/select/user", workspace_routes.DepartmentSelectUser)            //查询部门信息
		department.Post("/create", workspace_routes.DepartmentCreate)                    //创建部门
		department.Post("/editor", workspace_routes.DepartmentEditor)                    //修改部门信息
		department.Post("/delete", workspace_routes.DepartmentDelete)                    //删除部门
	}
}
