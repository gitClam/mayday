package workspace_routes

import (
	"github.com/kataras/iris/v12/core/router"
	workSpaceApi "mayday/src/api/v1/workspace"
)

func InitJobRouter(workspace router.Party) {

	job := workspace.Party("/job")
	{
		job.Get("/select", workSpaceApi.JobSelect)                      //查询职位信息
		job.Get("/select-user", workSpaceApi.SelectUserByJobId)         //查询职位信息
		job.Get("/select/department", workSpaceApi.JobSelectDepartment) //查询职位信息
		job.Get("/select/user", workSpaceApi.JobSelectUser)             //查询职位信息
		job.Post("/create", workSpaceApi.JobCreate)                     //创建职位
		job.Post("/editor", workSpaceApi.JobEditor)                     //修改职位信息
		job.Delete("/delete", workSpaceApi.JobDelete)                   //删除职位
		job.Post("/insert-user", workSpaceApi.JobInsert)                //添加用户
		job.Post("/delete-user", workSpaceApi.JobDeleteUser)            //删除用户
	}
}
