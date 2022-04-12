package workflowService

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
	"mayday/src/global"
	"mayday/src/model/common/resultcode"
	"mayday/src/model/common/timedecoder"
	UserModel "mayday/src/model/user"
	WorkflowModel "mayday/src/model/workflow"
	"mayday/src/model/workspace/application"
	"mayday/src/utils"
	"strings"
	"time"
)

//创建流程
func CreateWorkflow(ctx iris.Context, workflowReq WorkflowModel.WorkflowReq) {
	//TODO 待检验权限问题
	user := ctx.Values().Get("user").(UserModel.SdUser)

	sdWorkflow := workflowReq.GetSdWorkflow()
	sdWorkflow.CreateTime = timedecoder.LocalTime(time.Now())
	sdWorkflow.CreateUser = user.Id

	tables, err := sdWorkflow.Tables.MarshalJSON()
	tablesString := string(tables)
	strings.Trim(tablesString, `]`)
	strings.Trim(tablesString, `[`)
	var ids []string
	ids = strings.Split(tablesString, `,`)
	for _, id := range ids {
		id = `"` + id + `"`
	}
	tablesString = `[` + strings.Join(ids, `,`) + `]`
	fmt.Println(tablesString)
	sdWorkflow.Tables = []byte(tablesString)
	//正确写法
	session := global.GVA_DB.NewSession()
	effect, err := session.Insert(&sdWorkflow)
	if effect <= 0 || err != nil {
		utils.Responser.Fail(ctx, resultcode.DataCreateFail, err)
		if err := session.Rollback(); err != nil {
			global.GVA_LOG.Error("数据库事务回滚失败", zap.Error(err))
		}
		return
	}
	SdWorkflowApplication := application.SdWorkflowApplication{WorkflowId: sdWorkflow.Id, ApplicationId: workflowReq.ApplicationId}
	effect, err = session.Insert(SdWorkflowApplication)
	if effect <= 0 || err != nil {
		utils.Responser.Fail(ctx, resultcode.DataCreateFail, err)
		if err := session.Rollback(); err != nil {
			global.GVA_LOG.Error("数据库事务回滚失败", zap.Error(err))
		}
		return
	}

	if err := session.Commit(); err != nil {
		utils.Responser.Fail(ctx, resultcode.DataCreateFail, err)
		global.GVA_LOG.Error("数据库事务提交失败", zap.Error(err))
		return
	}

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
		utils.Responser.Fail(ctx, resultcode.DataCreateFail, err)
		return
	}
	utils.Responser.Ok(ctx)
}

//删除流程
func DeleteWorkflow(ctx iris.Context, id []int) {
	//TODO 验证权限
	e := global.GVA_DB.NewSession()
	err := e.Begin()
	if err != nil {
		utils.Responser.Fail(ctx, resultcode.DataDeleteFail, err)
		return
	}
	_, err = global.GVA_DB.In("workflow_id", id).Delete(new(application.SdWorkflowApplication))
	if err != nil {
		e.Rollback()
		utils.Responser.Fail(ctx, resultcode.DataDeleteFail, err)
		return
	}
	affected, err1 := e.In("id", id).Delete(new(WorkflowModel.SdWorkflow))
	if affected <= 0 || err1 != nil {
		e.Rollback()
		utils.Responser.Fail(ctx, resultcode.DataDeleteFail, err1)
		return
	}
	err = e.Commit()
	if err != nil {
		utils.Responser.Fail(ctx, resultcode.DataDeleteFail, err)
		return
	}
	utils.Responser.Ok(ctx)
}

//删除流程草稿
func DeleteWorkflowDraft(ctx iris.Context, id []int) {

	user := ctx.Values().Get("user").(UserModel.SdUser)
	var sdWorkflowDrafts []WorkflowModel.SdWorkflowDraft
	e := global.GVA_DB

	err := e.In("id", id).Find(&sdWorkflowDrafts)
	if err != nil {
		utils.Responser.Fail(ctx, resultcode.DataSelectFail, err)
		return
	}

	// 只能删除自己的草稿
	for _, sdWorkflow := range sdWorkflowDrafts {
		if sdWorkflow.OwnerId != user.Id {
			utils.Responser.Fail(ctx, resultcode.PermissionsLess, err)
			return
		}
	}

	affected, err := e.In("id", id).Delete(new(WorkflowModel.SdWorkflowDraft))
	if affected <= 0 || err != nil {
		utils.Responser.Fail(ctx, resultcode.DataDeleteFail, err)
		return
	}
	utils.Responser.Ok(ctx)
}

//获取流程信息
func GetWorkflowById(ctx iris.Context, id []int) {
	var result []struct {
		Id            int
		Name          string
		CreateUser    int
		CreateTime    timedecoder.LocalTime
		IsStart       int
		CeilingCount  int
		IsDeleted     int
		Structure     json.RawMessage
		Tables        json.RawMessage
		Remarks       string
		ApplicationId int
	}
	var SdWorkflows []WorkflowModel.SdWorkflow
	e := global.GVA_DB
	err := e.In("id", id).Find(&SdWorkflows)
	if err != nil {
		utils.Responser.Fail(ctx, resultcode.DataSelectFail, err)
		return
	}
	for _, SdWorkflow := range SdWorkflows {
		var sdWorkflowApplication application.SdWorkflowApplication
		_, err := global.GVA_DB.Where("workflow_id = ?", SdWorkflow.Id).Limit(1).Get(&sdWorkflowApplication)
		if err != nil {
			utils.Responser.Fail(ctx, resultcode.DataSelectFail, err)
			return
		}
		result = append(result, struct {
			Id            int
			Name          string
			CreateUser    int
			CreateTime    timedecoder.LocalTime
			IsStart       int
			CeilingCount  int
			IsDeleted     int
			Structure     json.RawMessage
			Tables        json.RawMessage
			Remarks       string
			ApplicationId int
		}{Id: SdWorkflow.Id,
			Name:          SdWorkflow.Name,
			CreateUser:    SdWorkflow.CreateUser,
			CreateTime:    SdWorkflow.CreateTime,
			IsStart:       SdWorkflow.IsStart,
			CeilingCount:  SdWorkflow.CeilingCount,
			IsDeleted:     SdWorkflow.IsDeleted,
			Structure:     SdWorkflow.Structure,
			Tables:        SdWorkflow.Tables,
			Remarks:       SdWorkflow.Remarks,
			ApplicationId: sdWorkflowApplication.ApplicationId,
		})
	}
	utils.Responser.OkWithDetails(ctx, result)
}

//获取用户的流程草稿
func GetWorkflowDraftByUser(ctx iris.Context) {
	user := ctx.Values().Get("user").(UserModel.SdUser)
	var workflowDrafts []WorkflowModel.SdWorkflowDraft

	e := global.GVA_DB
	err := e.Where("owner_id = ?", user.Id).Find(&workflowDrafts)
	if err != nil {
		utils.Responser.Fail(ctx, resultcode.DataSelectFail, err)
		return
	}

	utils.Responser.OkWithDetails(ctx, workflowDrafts)
}

//获取流程草稿详细信息
func GetWorkflowDraftById(ctx iris.Context, id []int) {
	//TODO 验证权限
	var SdWorkflows []WorkflowModel.SdWorkflowDraft
	e := global.GVA_DB
	err := e.In("id", id).Find(&SdWorkflows)
	if err != nil {
		utils.Responser.Fail(ctx, resultcode.DataSelectFail, err)
		return
	}
	utils.Responser.OkWithDetails(ctx, SdWorkflows)
}

//更新流程信息
func UpdateWorkflow(ctx iris.Context, workflowReq WorkflowModel.WorkflowReq) {
	//TODO 验证权限
	e := global.GVA_DB
	sdWorkflow := workflowReq.GetSdWorkflow()
	affected, err := e.Id(sdWorkflow.Id).Update(sdWorkflow)
	if affected <= 0 || err != nil {
		utils.Responser.Fail(ctx, resultcode.DataUpdateFail, err)
		return
	}
	utils.Responser.Ok(ctx)
}

//修改流程状态
func UpdateWorkflowState(ctx iris.Context, sdWorkflow WorkflowModel.SdWorkflow) {
	//TODO 验证权限
	e := global.GVA_DB
	has, err := e.Id(sdWorkflow.Id).Get(&sdWorkflow)
	if !has || err != nil {
		utils.Responser.Fail(ctx, resultcode.DataSelectFail, err)
		return
	}

	if sdWorkflow.IsStart == 0 {
		sdWorkflow.IsStart = 1
	} else {
		sdWorkflow.IsStart = 0
	}

	affected, err := e.Id(sdWorkflow.Id).Cols("is_start").Update(sdWorkflow)
	if affected <= 0 || err != nil {
		utils.Responser.Fail(ctx, resultcode.DataUpdateFail, err)
		return
	}
	utils.Responser.Ok(ctx)
}

//修改流程草稿
func UpdateWorkflowDraft(ctx iris.Context, workflowDraftReq WorkflowModel.WorkflowDraftReq) {
	//TODO 验证权限
	e := global.GVA_DB
	sdWorkflowDraft := workflowDraftReq.GetSdWorkflowDraft()
	affected, err := e.Id(sdWorkflowDraft.Id).Update(sdWorkflowDraft)
	if affected <= 0 || err != nil {
		utils.Responser.Fail(ctx, resultcode.DataUpdateFail, err)
		return
	}
	utils.Responser.Ok(ctx)
}
