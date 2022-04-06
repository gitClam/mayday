package order

import (
	"mayday/src/model/common/timedecoder"
)

type SdOrderCirculationHistory struct {
	Id           int                   `xorm:"not null pk autoincr INT(10)"`
	OrderId      int                   `xorm:"not null index INT(11)"`
	State        string                `xorm:"not null VARCHAR(20)"`
	Title        string                `xorm:"VARCHAR(20)"`
	Source       string                `xorm:"VARCHAR(128)"`
	Target       string                `xorm:"VARCHAR(128)"`
	Circulation  string                `xorm:"VARCHAR(128)"`
	Status       int                   `xorm:"INT(11)"`
	Processor    string                `xorm:"VARCHAR(50)"`
	ProcessorId  int                   `xorm:"index INT(11)"`
	CostDuration timedecoder.LocalTime `xorm:"DATETIME"`
	Remarks      string                `xorm:"VARCHAR(200)"`
}
