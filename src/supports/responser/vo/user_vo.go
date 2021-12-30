package vo

import (
	"mayday/src/models"
	//"log"
)

// 前端需要的数据结构
type UserVO struct {
	Id         int       
	Name       string    
	Realname   string    
	Age        int       
	Wechat     string    
	Qqnumber   string    
	Birthday   string 
	Sex        string    
	Info       string    
	Mail       string    
	Company    string 
	Department string
	Vocation   string    
	Phone      string       
	CreateDate string 
	Token      string 
}


// 携带token
func TansformUserVOToken(token string, user *model.SdUser) (uVO UserVO) {
	uVO.Id = user.Id
	uVO.Name = user.Name
	uVO.Realname = user.Realname
	uVO.Age = user.Age
	uVO.Wechat = user.Wechat
	uVO.Qqnumber = user.Qqnumber
	uVO.Birthday = user.Birthday.String("2006-01-02")
	uVO.Sex = user.Sex
	uVO.Info = user.Info
	uVO.Mail = user.Mail
	uVO.Company = user.Company
	uVO.Department = user.Department
	uVO.Vocation = user.Vocation
	uVO.Phone = user.Phone
	uVO.CreateDate = user.CreateDate.String("2006-01-02 15:04:05")	
	uVO.Token = token
	
	return
}

// 用户列表，不带啊token
func TansformUserVOList(userList []model.SdUser) (userVOList []UserVO) {
	for _, user := range userList {	
		uVO := UserVO{}
		
		uVO.Id = user.Id
		uVO.Name = user.Name
		uVO.Age = user.Age
		uVO.Sex = user.Sex
		uVO.Mail = user.Mail
		uVO.Phone = user.Phone
		userVOList = append(userVOList, uVO)
	}
	return
}
// 用户，不带啊token
func TansformUserVO(user *model.SdUser) (userVO UserVO) {	

	userVO.Id = user.Id
	userVO.Name = user.Name
	userVO.Realname = user.Realname
	userVO.Age = user.Age
	userVO.Wechat = user.Wechat
	userVO.Qqnumber = user.Qqnumber
	userVO.Birthday = user.Birthday.String("2006-01-02")
	userVO.Sex = user.Sex
	userVO.Info = user.Info
	userVO.Mail = user.Mail
	userVO.Company = user.Company
	userVO.Department = user.Department
	userVO.Vocation = user.Vocation
	userVO.Phone = user.Phone
	userVO.CreateDate = user.CreateDate.String("2006-01-02 15:04:05")		
	return
}
