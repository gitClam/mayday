package model

import (
	"log"
	"github.com/kataras/iris/v12"
)

type Job struct {
	Id       			int
	Name 			string
	Can_editor     	int
}
type Department struct {
	Id       			int
	Workspace_id     	int
	Name     			string
	Phone      		string
	Remark     		string
}
type Workspace struct {
	Id       			int
	Name     			string
	Phone      		string
	Remark     		string
}
type Application struct {
	Id       			int
	Workspace_id     	string
	Name      		string
	Remark     		string
}
