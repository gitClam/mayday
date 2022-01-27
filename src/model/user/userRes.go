package user

// UserRes 前端需要的数据结构
type UserRes struct {
	Id         int
	Name       string
	RealName   string
	Age        int
	Wechat     string
	QqNumber   string
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

// TransformUserVOToken 携带token
func TransformUserVOToken(token string, user *SdUser) (uVO UserRes) {
	uVO.Id = user.Id
	uVO.Name = user.Name
	uVO.RealName = user.Realname
	uVO.Age = user.Age
	uVO.Wechat = user.Wechat
	uVO.QqNumber = user.Qqnumber
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

// TransformUserVOList 用户列表，不带啊token
func TransformUserVOList(userList []SdUser) (userVOList []UserRes) {
	for _, sdUser := range userList {
		uVO := UserRes{}

		uVO.Id = sdUser.Id
		uVO.Name = sdUser.Name
		uVO.Age = sdUser.Age
		uVO.Sex = sdUser.Sex
		uVO.Mail = sdUser.Mail
		uVO.Phone = sdUser.Phone

		userVOList = append(userVOList, uVO)
	}
	return
}

// TransformUserVO 用户，不带啊token
func TransformUserVO(user *SdUser) (userVO UserRes) {
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