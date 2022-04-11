package workflow

import (
	"encoding/json"
)

type WorkflowReq struct {
	Id            int
	ApplicationId int             `validate:"required" example:"12"`
	Name          string          `validate:"required" example:"请假流程"`
	IsStart       int             `example:"0" enums:"0,1" default:"0"`
	Structure     json.RawMessage `validate:"required" example:"流程的JSON文件"`
	Tables        json.RawMessage `validate:"required" example:"[\"1001\", \"1002\"]"`
	Remarks       string          `example:"请假流程1"`
}

type WorkflowDraftReq struct {
	Name      string          `validate:"required" example:"睡觉流程"`
	Structure json.RawMessage `validate:"required" example:"流程的JSON文件"`
	Tables    json.RawMessage `validate:"required" example:"[\"1001\", \"1002\"]"`
	Remarks   string          `example:"请假流程2"`
}

func (req *WorkflowDraftReq) GetSdWorkflowDraft() (sd SdWorkflowDraft) {

	sd.Name = req.Name
	sd.Structure = req.Structure
	sd.Tables = req.Tables
	sd.Remarks = req.Remarks

	return
}
func (req *WorkflowReq) GetSdWorkflow() (sd SdWorkflow) {
	sd.Id = req.Id
	sd.Name = req.Name
	sd.IsStart = req.IsStart
	sd.Structure = req.Structure
	sd.Tables = req.Tables
	sd.Remarks = req.Remarks

	return
}
