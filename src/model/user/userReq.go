package user

import "mayday/src/utils"

type UserReq struct {
	Name       string          `example:"M.Salah"`
	Password   string          `example:"123456"`
	Realname   string          `example:"罗智"`
	Age        int             `example:"3"`
	Birthday   utils.LocalTime `example:"2021-01-01 00:00:00"`
	Sex        string          `example:"男" enums:"男,女"`
	Wechat     string          `example:"M.Salah"`
	Qqnumber   string          `example:"123456789"`
	Info       string          `example:"今晚点喂"`
	Mail       string          `example:"123456@abc.com"`
	Company    string          `example:"罗智地产有限公司"`
	Vocation   string          `example:"包工头"`
	Department string          `example:"tnnd怎么还不点"`
	Phone      string          `example:"12345678912"`
}
