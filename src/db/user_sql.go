package sql

import (
	//"database/sql"
	"mayday/src/models"
	"log"
	//"github.com/go-sql-driver/mysql"
	"mayday/src/db/conn"
)
func Register(user model.User)                                                                                                                                                                                                                                                                                                                    string{
     
     stmt,err := conn.Getconn().Prepare(`insert sd_user(name,password) values(?,?)`)
	if err != nil {
		log.Print(err)
		return err.Error()
	}
	
	res, err := stmt.Exec(user.Name,user.Password)
	if err != nil {
		log.Print(err)
		return err.Error()
	}
	
	id, err := res.LastInsertId()
	if err != nil {
		log.Print(err)
		return err.Error()
	}
	
	log.Println(id)
	return "ok"
}

func Login(user model.User) model.User{
     
     var newuser model.User
	stmt := "select id,name,password from sd_user where name=?"
	log.Print("res:",user.Name)
	err := conn.Getconn().QueryRow(stmt,user.Name).Scan(&newuser.Id,&newuser.Name,&newuser.Password)
	if(err != nil){
		log.Print("err:",err)
	}
	log.Print(newuser)
	return newuser
}
