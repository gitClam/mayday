package workspace_routes

import (
	"github.com/kataras/iris/v12/core/router"
	workspcaeApi "mayday/src/api/v1/workspace"
)

func InitJobRouter(workspace router.Party) {

	job := workspace.Party("/job")
	{
		job.Get("/select", workspcaeApi.JobSelect)                      //查询职位信息
		job.Get("/select-user", workspcaeApi.SelectUserByJobId)         //查询职位信息
		job.Get("/select/department", workspcaeApi.JobSelectDepartment) //查询职位信息
		job.Get("/select/user", workspcaeApi.JobSelectUser)             //查询职位信息
		job.Post("/create", workspcaeApi.JobCreate)                     //创建职位
		job.Post("/editor", workspcaeApi.JobEditor)                     //修改职位信息
		job.Delete("/delete", workspcaeApi.JobDelete)                   //删除职位
		job.Post("/insert-user", workspcaeApi.JobInsert)                //添加用户
		job.Post("/delete-user", workspcaeApi.JobDeleteUser)            //删除用户
	}
}
