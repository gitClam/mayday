package model

import (
	//"time"
	"mayday/src/utils"
)

type SdUser struct {
	Id         int             `xorm:"not null pk autoincr INT(11)"`
	Name       string          `xorm:"not null VARCHAR(50)"`
	Password   string          `xorm:"not null VARCHAR(50)"`
	Realname   string          `xorm:"VARCHAR(30)"`
	Age        int             `xorm:"INT(11)"`
	Wechat     string          `xorm:"VARCHAR(100)"`
	Qqnumber   string          `xorm:"VARCHAR(100)"`
	Birthday   utils.LocalTime `xorm:"DATETIME"`
	Sex        string          `xorm:"ENUM('女','男')"`
	Info       string          `xorm:"VARCHAR(255)"`
	Mail       string          `xorm:"not null VARCHAR(50)"`
	Company    string          `xorm:"VARCHAR(20)"`
	Department string          `xorm:"VARCHAR(50)"`
	Vocation   string          `xorm:"VARCHAR(30)"`
	Phone      string          `xorm:"VARCHAR(20)"`
	Photo      string          `xorm:"VARCHAR(50)"`
	CreateDate utils.LocalTime `xorm:"not null DATETIME"`
	IsDeleted  int             `xorm:"not null default 0 TINYINT(1)"`
}
