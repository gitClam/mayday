package model

type SdApplication struct {
	Id          int    `xorm:"not null pk autoincr INT(11)"`
	WorkspaceId int    `xorm:"not null index INT(11)"`
	Name        string `xorm:"VARCHAR(50)"`
	Remark      string `xorm:"VARCHAR(255)"`
	IsDeleted   int    `xorm:"not null default 0 TINYINT(1)"`
}
