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
			get.Get("/workflow/{id:int}", workflowApi.WorkflowSelectWorkflow)            //查询流程
			get.Get("/workflow-draft", workflowApi.WorkflowSelectWorkflowDraft)          //查询本人所有的流程草稿
			get.Get("/workflow-draft/{id:int}", workflowApi.WorkflowSelectWorkflowDraft) //查询流程草稿详细信息
		}
		create := workflow.Party("/create")
		{
			create.Post("/workflow", workflowApi.WorkflowCreateWorkflow)            //创建流程
			create.Post("/workflow-draft", workflowApi.WorkflowCreateWorkflowDraft) //创建流程草稿
		}

		update := workflow.Party("/update")
		{
			update.Post("/workflow", workflowApi.WorkflowEditorWorkflow)            //修改流程
			update.Post("/workflow-state", workflowApi.WorkflowEditorWorkflowState) //修改流程状态
			update.Post("/workflow-draft", workflowApi.WorkflowEditorWorkflowDraft) //修改流程草稿
		}

		delete := workflow.Party("/delete")
		{
			delete.Delete("/workflow/{id:int}", workflowApi.WorkflowDeleteWorkflow)            //删除流程
			delete.Delete("/workflow-draft/{id:int}", workflowApi.WorkflowDeleteWorkflowDraft) //删除流程草稿
		}
	}
}
