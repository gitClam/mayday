package user

import "mayday/src/utils"

type UserReq struct {
	Name       string
	Password   string
	Realname   string
	Age        int
	Birthday   utils.LocalTime
	Sex        string
	Wechat     string
	Qqnumber   string
	Info       string
	Mail       string
	Company    string
	Vocation   string
	Department string
}
