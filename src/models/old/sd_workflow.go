package model

import (
	"log"
	"github.com/kataras/iris/v12"
)

type Workflow struct {
	Id                 	int
	Name               	string
	Create_user        	string
	Create_time        	string
	Workspace_id       	int
	Is_start           	string
	Start_time    		string
	End_time         	string	
	Day_start_time      string
	Day_end_time        string
	Ceiling_count     	int
}

type Workflow_node struct{
	Id 				int
	Name 			string
	workflow_id 		string
	table_id 			string
	serial_number 		int
	workflow_group_id 	int
	permissions 		string
	is_remind 		int
}

type Table struct{
	Id 				int
	Workspace_id 		int
	data 			string
	name 			string
}

type Workflow_draft struct {
	Id          		int
	Name        		string
	Owner_id    		int
	Is_start      		string
	Start_time    		string
	End_time         	string	
	Day_start_time      string
	Day_end_time        string
	Ceiling_count     	int
}

type Workflow_node_draft struct{
	Id 				int
	Name 			string
	workflow_id 		string
	table_id 			string
	serial_number 		int
	workflow_group_id 	int
	permissions 		string
	is_remind 		int
}

type Workflow_group struct{
	Id 				int
	Name 			string
}
