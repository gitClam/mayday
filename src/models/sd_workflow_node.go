package model

type SdWorkflowNode struct {
	Id              int    `xorm:"not null pk autoincr INT(11)"`
	Name            string `xorm:"VARCHAR(50)"`
	WorkflowId      int    `xorm:"not null index INT(11)"`
	TableId         int    `xorm:"not null index INT(11)"`
	SerialNumber    int    `xorm:"not null INT(11)"`
	WorkflowGroupId int    `xorm:"index INT(11)"`

	permissions string `xorm:"not null VARCHAR(255)"`
	IsRemind    int    `xorm:"not null default 0 TINYINT(1)"`
	IsDeleted   int    `xorm:"not null default 0 TINYINT(1)"`
}
