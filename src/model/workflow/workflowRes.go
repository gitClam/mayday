package workflow

import (
	"encoding/json"
	"mayday/src/utils"
)

type WorkflowRes struct {
	Id           int             `validate:"required" example:"睡觉流程"`
	Name         string          `validate:"required" example:"睡觉流程"`
	CreateUser   int             `validate:"required" example:"睡觉流程"`
	CreateTime   utils.LocalTime `validate:"required" example:"睡觉流程"`
	IsStart      int             `validate:"required" example:"睡觉流程"`
	CeilingCount int             `validate:"required" example:"睡觉流程"`
	Structure    json.RawMessage `validate:"required" example:"睡觉流程"`
	Tables       json.RawMessage `validate:"required" example:"睡觉流程"`
	Remarks      string          `validate:"required" example:"睡觉流程"`
}

func GetWorkflowRes(sd []SdWorkflow) (res []WorkflowRes) {
	for _, Workflow := range sd {
		workflowRes := WorkflowRes{}

		workflowRes.Id = Workflow.Id
		workflowRes.Name = Workflow.Name
		workflowRes.CreateUser = Workflow.CreateUser
		workflowRes.CreateTime = Workflow.CreateTime
		workflowRes.IsStart = Workflow.IsStart
		workflowRes.CeilingCount = Workflow.CeilingCount
		workflowRes.Structure = Workflow.Structure
		workflowRes.Tables = Workflow.Tables
		workflowRes.Remarks = Workflow.Remarks

		res = append(res, workflowRes)
	}
	return
}
