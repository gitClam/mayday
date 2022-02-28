package workflow_routes

import (
	"github.com/kataras/iris/v12/core/router"
	workflowApi "mayday/src/api/v1/workflow"
)

func InitTableRouter(Router router.Party) {

	table := Router.Party("/table")
	{
		get := table.Party("/get")
		{
			get.Get("/table", workflowApi.GetTableById)                   //查询表单
			get.Get("/table-draft/user", workflowApi.GetTableDraftByUser) //查询用户拥有的表单草稿
			get.Get("/table-draft/id", workflowApi.GetTableDraftById)     //查询表单草稿
		}
		create := table.Party("/create")
		{
			create.Post("/table", workflowApi.CreateTable)            //创建表单
			create.Post("/table-draft", workflowApi.CreateTableDraft) //创建表单草稿
		}

		update := table.Party("/update")
		{
			update.Post("/table", workflowApi.UpdateTable)            //修改表单
			update.Post("/table-draft", workflowApi.UpdateTableDraft) //修改表单草稿
		}

		delete := table.Party("/delete")
		{
			delete.Delete("/table", workflowApi.DeleteTable)            //删除表单
			delete.Delete("/table-draft", workflowApi.DeleteTableDraft) //删除表单草稿
		}
	}
}
