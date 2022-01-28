package user

// UserDetailsRes 前端需要的数据结构
type UserDetailsRes struct {
	Id         int    `example:"1"`
	Name       string `example:"M.Salah"`
	RealName   string `example:"罗智"`
	Age        int    `example:"3"`
	Wechat     string `example:"M.Salah"`
	QqNumber   string `example:"123456789"`
	Birthday   string `example:"2021-01-01 00:00:00"`
	Sex        string `example:"男" enums:"男,女"`
	Info       string `example:"今晚点喂"`
	Mail       string `example:"123456@abc.com"`
	Company    string `example:"罗智地产有限公司"`
	Department string `example:"tnnd怎么还不点"`
	Vocation   string `example:"包工头"`
	Phone      string `example:"12345678912"`
	CreateDate string `example:"2021-01-01 00:00:00"`
	Token      string `example:"NDOAIIF@!Afaad21dAONF24b78B9b23br9B(HRbnv8020Bv893htb08BbivB082"`
}

type UserAbstractRes struct {
	Id    int    `example:"1"`
	Name  string `example:"M.Salah"`
	Age   int    `example:"3"`
	Sex   string `example:"男" enums:"男,女"`
	Info  string `example:"今晚点喂"`
	Mail  string `example:"123456@abc.com"`
	Phone string `example:"12345678912"`
}

// GetUserDetailsResWithToken 携带token
func GetUserDetailsResWithToken(token string, user *SdUser) (userDetailsRes UserDetailsRes) {
	userDetailsRes.Id = user.Id
	userDetailsRes.Name = user.Name
	userDetailsRes.RealName = user.Realname
	userDetailsRes.Age = user.Age
	userDetailsRes.Wechat = user.Wechat
	userDetailsRes.QqNumber = user.Qqnumber
	userDetailsRes.Birthday = user.Birthday.String("2006-01-02")
	userDetailsRes.Sex = user.Sex
	userDetailsRes.Info = user.Info
	userDetailsRes.Mail = user.Mail
	userDetailsRes.Company = user.Company
	userDetailsRes.Department = user.Department
	userDetailsRes.Vocation = user.Vocation
	userDetailsRes.Phone = user.Phone
	userDetailsRes.CreateDate = user.CreateDate.String("2006-01-02 15:04:05")
	userDetailsRes.Token = token
	return
}

// GetUserAbstractResList 用户列表，不带啊token
func GetUserAbstractResList(userList []SdUser) (UserAbstractResList []UserAbstractRes) {
	for _, sdUser := range userList {
		userAbstractRes := UserAbstractRes{}

		userAbstractRes.Id = sdUser.Id
		userAbstractRes.Name = sdUser.Name
		userAbstractRes.Age = sdUser.Age
		userAbstractRes.Sex = sdUser.Sex
		userAbstractRes.Info = sdUser.Info
		userAbstractRes.Mail = sdUser.Mail
		userAbstractRes.Phone = sdUser.Phone

		UserAbstractResList = append(UserAbstractResList, userAbstractRes)
	}
	return
}

// GetUserDetailsResWithOutToken 用户，不带啊token
func GetUserDetailsResWithOutToken(user *SdUser) (userVO UserDetailsRes) {
	userVO.Id = user.Id
	userVO.Name = user.Name
	userVO.RealName = user.Realname
	userVO.Age = user.Age
	userVO.Wechat = user.Wechat
	userVO.QqNumber = user.Qqnumber
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
