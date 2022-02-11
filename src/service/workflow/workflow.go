package workflow

import (
	"github.com/kataras/iris/v12"
	"mayday/src/global"
	UserModel "mayday/src/model/user"
	WorkflowModel "mayday/src/model/workflow"
	"mayday/src/utils"
	"time"
)

//创建流程
func CreateWorkflow(ctx iris.Context, workflowReq WorkflowModel.WorkflowReq) {
	//TODO 待检验权限问题
	user := ctx.Values().Get("user").(UserModel.SdUser)

	sdWorkflow := workflowReq.GetSdWorkflow()
	sdWorkflow.CreateTime = utils.LocalTime(time.Now())
	sdWorkflow.CreateUser = user.Id

	//方便测试
	e := global.GVA_DB
	effect, err := e.Insert(sdWorkflow)
	if effect <= 0 || err != nil {
		utils.Responser.FailWithMsg(ctx, "流程创建失败", err)
		return
	}

	//正确写法
	//session := e.NewSession()
	//
	//effect, err := session.Insert(&sdWorkflow)
	//if effect <= 0 || err != nil {
	//	utils.Responser.FailWithMsg(ctx, "流程创建失败", err)
	//	if err := session.Rollback(); err != nil {
	//		global.GVA_LOG.Error("数据库事务回滚失败", zap.Error(err))
	//	}
	//	return
	//}
	//SdWorkflowApplication := model.SdWorkflowApplication{WorkflowId: sdWorkflow.Id, ApplicationId: workflowReq.ApplicationId}
	//effect, err := session.Insert(SdWorkflowApplication)
	//if effect <= 0 || err != nil {
	//	utils.Responser.FailWithMsg(ctx, "流程创建失败", err)
	//	if err := session.Rollback(); err != nil {
	//		global.GVA_LOG.Error("数据库事务回滚失败", zap.Error(err))
	//	}
	//	return
	//}
	//
	//if err := session.Commit(); err != nil {
	//	utils.Responser.FailWithMsg(ctx, "流程创建失败", err)
	//	global.GVA_LOG.Error("数据库事务提交失败", zap.Error(err))
	//	return
	//}

	utils.Responser.Ok(ctx)
}

//创建流程草稿
func CreateWorkflowDraft(ctx iris.Context, workflowDraftReq WorkflowModel.WorkflowDraftReq) {

	user := ctx.Values().Get("user").(UserModel.SdUser)

	sdWorkflowDraft := workflowDraftReq.GetSdWorkflowDraft()
	sdWorkflowDraft.OwnerId = user.Id

	e := global.GVA_DB
	effect, err := e.Insert(sdWorkflowDraft)
	if effect <= 0 || err != nil {
		utils.Responser.FailWithMsg(ctx, "流程草稿创建失败", err)
		return
	}
	utils.Responser.Ok(ctx)
}
