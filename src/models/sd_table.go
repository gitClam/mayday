package model
import (
	"encoding/json"
	//"time"
)

type SdTable struct {
	Id          int    `xorm:"not null pk autoincr INT(11)"`
	WorkspaceId int    `xorm:"not null index INT(11)"`
	Data        json.RawMessage `xorm:"not null JSON"`
	Name        string `xorm:"not null VARCHAR(50)"`
	IsDeleted   int    `xorm:"not null TINYINT(1)"`
}
