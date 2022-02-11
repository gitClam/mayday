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

//删除流程
func DeleteWorkflow(ctx iris.Context, id []int) {
	//TODO 验证权限
	sdWorkflow := new(WorkflowModel.SdWorkflow)
	e := global.GVA_DB
	affected, err := e.Id(id).Delete(sdWorkflow)
	if affected <= 0 || err != nil {
		utils.Responser.FailWithMsg(ctx, "流程删除失败", err)
		return
	}
	utils.Responser.Ok(ctx)
}

//删除流程草稿
func DeleteWorkflowDraft(ctx iris.Context, id []int) {

	user := ctx.Values().Get("user").(UserModel.SdUser)
	var sdWorkflowDrafts []WorkflowModel.SdWorkflowDraft
	e := global.GVA_DB

	err := e.Id(id).Find(&sdWorkflowDrafts)
	if err != nil {
		utils.Responser.FailWithMsg(ctx, "流程草稿不存在", err)
		return
	}

	// 只能删除自己的草稿
	for _, sdWorkflow := range sdWorkflowDrafts {
		if sdWorkflow.OwnerId != user.Id {
			utils.Responser.FailWithMsg(ctx, "非法请求", err)
			return
		}
	}

	affected, err := e.Id(id).Delete(new(WorkflowModel.SdWorkflowDraft))
	if affected <= 0 || err != nil {
		utils.Responser.FailWithMsg(ctx, "流程删除失败", err)
		return
	}
	utils.Responser.Ok(ctx)
}
