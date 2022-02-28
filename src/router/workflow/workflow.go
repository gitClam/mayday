package workflow_routes

import (
	"github.com/kataras/iris/v12/core/router"
	workflowApi "mayday/src/api/v1/workflow"
)

func InitWorkflowRouter(Router router.Party) {

	workflow := Router.Party("/workflow")
	{
		get := workflow.Party("/get")
		{
			get.Get("/workflow", workflowApi.GetWorkflowById)                   //查询流程
			get.Get("/workflow-draft/user", workflowApi.GetWorkflowDraftByUser) //查询本人所有的流程草稿
			get.Get("/workflow-draft/id", workflowApi.GetWorkflowDraftById)     //查询流程草稿详细信息
		}
		create := workflow.Party("/create")
		{
			create.Post("/workflow", workflowApi.CreateWorkflow)            //创建流程
			create.Post("/workflow-draft", workflowApi.CreateWorkflowDraft) //创建流程草稿
		}

		update := workflow.Party("/update")
		{
			update.Post("/workflow", workflowApi.UpdateWorkflow)            //修改流程
			update.Post("/workflow-state", workflowApi.UpdateWorkflowState) //修改流程状态
			update.Post("/workflow-draft", workflowApi.UpdateWorkflowDraft) //修改流程草稿
		}

		delete := workflow.Party("/delete")
		{
			delete.Delete("/workflow", workflowApi.DeleteWorkflow)            //删除流程
			delete.Delete("/workflow-draft", workflowApi.DeleteWorkflowDraft) //删除流程草稿
		}
	}
}
