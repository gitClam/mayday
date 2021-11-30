package model
import (
	"log"
	"github.com/kataras/iris/v12"
)

type Order struct {
	Id       			int
	User_id     		int
	Workflow_id     	int
	Create_time      	string
	Serial_number     	int
	State    			string
}

type Record struct {
	Id       			int
	User_id     		int
	Order_id     		int
	Time      		string
	Table     		string
}
