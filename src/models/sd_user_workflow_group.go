package model

type SdUserWorkflowGroup struct {
	Id              int `xorm:"not null pk autoincr INT(11)"`
	UserId          int `xorm:"not null index INT(11)"`
	WorkflowGroupId int `xorm:"not null index INT(11)"`
}
