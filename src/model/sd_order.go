package model

import (
	"encoding/json"
	"mayday/src/utils"
	//"time"
)

type SdOrder struct {
	Id            int             `xorm:"not null pk autoincr INT(11)"`
	UserId        int             `xorm:"not null index INT(11)"`
	WorkflowId    int             `xorm:"not null index INT(11)"`
	CreateTime    utils.LocalTime `xorm:"not null DATETIME"`
	Title         string          `xorm:"VARCHAR(20)"`
	UrgeLastTime  utils.LocalTime `xorm:"DATETIME"`
	UrgeCount     int             `xorm:"INT(11)"`
	RelatedPerson json.RawMessage `xorm:"not null JSON"`
	IsDenied      int             `xorm:"not null INT(11)"`
	IsEnd         int             `xorm:"not null INT(11)"`
	State         json.RawMessage `xorm:"JSON"`
	IsDeleted     int             `xorm:"not null default 0 TINYINT(1)"`
}
