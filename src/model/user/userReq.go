package user

import "mayday/src/utils"

type UserReq struct {
	Name       string          `example:"小明"`
	Password   string          `example:"小明"`
	Realname   string          `example:"小明"`
	Age        int             `example:"10"`
	Birthday   utils.LocalTime `example:"0001-01-01 00:00:00"`
	Sex        string          `example:"小明"`
	Wechat     string          `example:"小明"`
	Qqnumber   string          `example:"小明"`
	Info       string          `example:"小明"`
	Mail       string          `example:"小明"`
	Company    string          `example:"小明"`
	Vocation   string          `example:"小明"`
	Department string          `example:"小明"`
}
