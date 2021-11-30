package model

import (
	"encoding/json"
	//"time"
)

type SdOrderTable struct {
	Id             int             `xorm:"not null pk autoincr INT(11)"`
	OrderHistoryId int             `xorm:"index INT(11)"`
	FormStructure  json.RawMessage `xorm:"JSON"`
	FormData       json.RawMessage `xorm:"JSON"`
}
