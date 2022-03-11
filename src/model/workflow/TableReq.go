package workflow

import "encoding/json"

type CreateTableReq struct {
	WorkspaceId int    `validate:"required" example:"001"`
	Data        string `validate:"required" example:"JSON格式的数据"`
	Name        string `validate:"required" example:"请假表单"`
}

type CreateTableDraftReq struct {
	Data string `validate:"required" extensions:"JSON格式的数据"`
	Name string `validate:"required" example:"请假表单"`
}

type UpdateTableReq struct {
	Id          int    `validate:"required" example:"001"`
	WorkspaceId int    `example:"001"`
	Data        string `example:"JSON格式的数据"`
	Name        string `example:"请假表单"`
}

type UpdateTableDraftReq struct {
	Id   int    `validate:"required" example:"001"`
	Data string `example:"JSON格式的数据"`
	Name string `example:"请假表单"`
}

func (req *CreateTableReq) GetSdTable() (sd SdTable) {

	sd.Name = req.Name
	sd.WorkspaceId = req.WorkspaceId
	sd.Data = json.RawMessage(req.Data)

	return
}

func (req *CreateTableDraftReq) GetSdTableDraft() (sd SdTableDraft) {

	sd.Name = req.Name
	sd.Data = json.RawMessage(req.Data)

	return
}

func (req *UpdateTableReq) GetSdTable() (sd SdTable) {
	sd.Id = req.Id
	sd.Name = req.Name
	sd.WorkspaceId = req.WorkspaceId
	sd.Data = json.RawMessage(req.Data)

	return
}

func (req *UpdateTableDraftReq) GetSdTableDraft() (sd SdTableDraft) {
	sd.Id = req.Id
	sd.Name = req.Name
	sd.Data = json.RawMessage(req.Data)

	return
}
