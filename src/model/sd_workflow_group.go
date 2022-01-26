package model

type SdWorkflowGroup struct {
	Id        int    `xorm:"not null pk autoincr INT(11)"`
	Name      string `xorm:"VARCHAR(50)"`
	IsDeleted int    `xorm:"not null TINYINT(1)"`
}
