package model

type SdWorkflowNodeDraft struct {
	Id              int    `xorm:"not null pk autoincr INT(11)"`
	Name            string `xorm:"VARCHAR(50)"`
	WorkflowId      int    `xorm:"not null index INT(11)"`
	TableId         int    `xorm:"not null index INT(11)"`
	SerialNumber    int    `xorm:"not null INT(11)"`
	WorkflowGroupId int    `xorm:"INT(11)"`
	Permissions     string `xorm:"not null VARCHAR(255)"`
	IsRemind        int    `xorm:"not null default 0 TINYINT(1)"`
	IsDeleted       int    `xorm:"not null default 0 TINYINT(1)"`
}
