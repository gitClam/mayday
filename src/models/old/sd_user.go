package model
import (
	"log"
	"github.com/kataras/iris/v12"
)

type User struct {
	Id          	int
	Name        	string
	Password       string
	Realname    	string
	Age         	int
	QQNumber    	string
	Wechat      	string
	Birthday    	string
	Sex         	string	
	Info        	string
	Mail        	string
	Company     	string
	Vocation    	string
	Phone       	string
	Create_date 	string	
}
