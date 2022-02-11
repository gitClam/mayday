package workflow

import "encoding/json"

type TableReq struct {
	WorkspaceId int    `validate:"required" example:"001"`
	Data        string `validate:"required" example:"JSON格式的数据"`
	Name        string `validate:"required" example:"请假表单"`
}

type TableDraftReq struct {
	Data string `validate:"required" example:"JSON格式的数据"`
	Name string `validate:"required" example:"请假表单"`
}

func (req *TableReq) GetSdTable() (sd SdTable) {

	sd.Name = req.Name
	sd.WorkspaceId = req.WorkspaceId
	sd.Data = json.RawMessage(req.Data)

	return
}

func (req *TableDraftReq) GetSdTableDraft() (sd SdTableDraft) {

	sd.Name = req.Name
	sd.Data = json.RawMessage(req.Data)

	return
}
