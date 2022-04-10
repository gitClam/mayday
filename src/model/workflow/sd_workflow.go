package workflow

import (
	"encoding/json"
	"mayday/src/model/common/timedecoder"
	//"time"
)

type SdWorkflow struct {
	Id           int                   `xorm:"not null pk autoincr INT(11)"`
	Name         string                `xorm:"VARCHAR(50)"`
	CreateUser   int                   `xorm:"not null index INT(11)"`
	CreateTime   timedecoder.LocalTime `xorm:"not null DATETIME"`
	IsStart      int                   `xorm:"not null TINYINT(1)"`
	CeilingCount int                   `xorm:"INT(11)"`
	IsDeleted    int                   `xorm:"not null default 0 TINYINT(1)"`
	Structure    json.RawMessage       `xorm:"not null JSON"`
	Tables       json.RawMessage       `xorm:"not null JSON"`
	Remarks      string                `xorm:"VARCHAR(100)"`
}
type WorkflowSimpleRes struct {
	Id           int                   `xorm:"not null pk autoincr INT(11)"`
	Name         string                `xorm:"VARCHAR(50)"`
	CreateUser   int                   `xorm:"not null index INT(11)"`
	CreateTime   timedecoder.LocalTime `xorm:"not null DATETIME"`
	IsStart      int                   `xorm:"not null TINYINT(1)"`
	CeilingCount int                   `xorm:"INT(11)"`
	IsDeleted    int                   `xorm:"not null default 0 TINYINT(1)"`
	Remarks      string                `xorm:"VARCHAR(100)"`
}

func (sd *SdWorkflow) ToWorkflowSimpleRes() (res WorkflowSimpleRes) {
	res.Id = sd.Id
	res.Name = sd.Name
	res.CreateUser = sd.CreateUser
	res.CreateTime = sd.CreateTime
	res.IsStart = sd.IsStart
	res.CeilingCount = sd.CeilingCount
	res.IsDeleted = sd.IsDeleted
	res.Remarks = sd.Remarks
	return
}
