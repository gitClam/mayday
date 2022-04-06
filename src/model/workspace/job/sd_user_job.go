package job

type SdUserJob struct {
	Id     int `xorm:"not null pk autoincr INT(11)"`
	UserId int `xorm:"not null index INT(11)"`
	JobId  int `xorm:"not null index INT(11)"`
}
