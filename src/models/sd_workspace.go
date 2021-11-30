package model

type SdWorkspace struct {
	Id        int    `xorm:"not null pk autoincr INT(11)"`
	Name      string `xorm:"not null VARCHAR(50)"`
	Phone     string `xorm:"VARCHAR(20)"`
	Remark    string `xorm:"VARCHAR(255)"`
	IsDeleted int    `xorm:"not null default 0 TINYINT(1)"`
}
