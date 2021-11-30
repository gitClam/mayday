package model

type SdWorkflowApplication struct {
	Id            int `xorm:"not null pk autoincr INT(11)"`
	WorkflowId    int `xorm:"index INT(11)"`
	ApplicationId int `xorm:"index INT(11)"`
}
