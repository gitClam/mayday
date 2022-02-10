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
func CreateWorkflow(ctx iris.Context, user UserModel.SdUser, workflow WorkflowModel.SdWorkflow) {
	//TODO 待检验权限问题
	workflow.CreateTime = utils.LocalTime(time.Now())
	workflow.CreateUser = user.Id

	e := global.GVA_DB
	effect, err := e.Insert(workflow)
	if effect <= 0 || err != nil {
		utils.Responser.FailWithMsg(ctx, "流程创建失败")
		return
	}
	utils.Responser.Ok(ctx)
	global.GVA_LOG.Info("用户：" + user.Name + " 流程创建成功")
}
