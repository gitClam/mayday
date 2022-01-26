package model

import (
	"encoding/json"
	"mayday/src/utils"
	//"time"
)

type SdWorkflow struct {
	Id           int             `xorm:"not null pk autoincr INT(11)"`
	Name         string          `xorm:"VARCHAR(50)"`
	CreateUser   int             `xorm:"not null index INT(11)"`
	CreateTime   utils.LocalTime `xorm:"not null DATETIME"`
	IsStart      int             `xorm:"not null TINYINT(1)"`
	CeilingCount int             `xorm:"INT(11)"`
	IsDeleted    int             `xorm:"not null default 0 TINYINT(1)"`
	Structure    json.RawMessage `xorm:"not null JSON"`
	Tables       json.RawMessage `xorm:"not null JSON"`
	Remarks      string          `xorm:"VARCHAR(100)"`
}
