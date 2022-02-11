package workflow

import (
	"encoding/json"
)

type WorkflowReq struct {
	ApplicationId int    `validate:"required" example:"12" extensions:"现在为了方便可以暂时不填"`
	Name          string `validate:"required" example:"请假流程"`
	IsStart       int    `example:"0" enums:"0,1" default:"0"`
	Structure     string `validate:"required" example:"流程的JSON文件"`
	Tables        string `validate:"required" example:"[\"1001\", \"1002\"]"`
	Remarks       string `example:"请假流程1"`
}

type WorkflowDraftReq struct {
	Name      string `validate:"required" example:"睡觉流程"`
	Structure string `validate:"required" example:"流程的JSON文件"`
	Tables    string `validate:"required" example:"[\"1001\", \"1002\"]"`
	Remarks   string `example:"请假流程2"`
}

func (req *WorkflowDraftReq) GetSdWorkflowDraft() (sd SdWorkflowDraft) {

	sd.Name = req.Name
	sd.Structure = json.RawMessage(req.Structure)
	sd.Tables = json.RawMessage(req.Tables)
	sd.Remarks = req.Remarks

	return
}
func (req *WorkflowReq) GetSdWorkflow() (sd SdWorkflow) {

	sd.Name = req.Name
	sd.IsStart = req.IsStart
	sd.Structure = json.RawMessage(req.Structure)
	sd.Tables = json.RawMessage(req.Tables)
	sd.Remarks = req.Remarks

	return
}
