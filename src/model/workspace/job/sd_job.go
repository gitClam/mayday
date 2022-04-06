package job

type SdJob struct {
	Id           int    `xorm:"not null pk autoincr INT(11)"`
	DepartmentId int    `xorm:"not null index INT(11)"`
	Name         string `xorm:"not null VARCHAR(50)"`
	IsDelete     int    `xorm:"not null default 0 TINYINT(1)"`
}
