package resultcode

const (
	// msg define
	Success  int = 200 //"成功"
	Fail     int = 777 //"失败"
	NotFound int = 404 //"您请求的url不存在"

	RegisterFail             int = 1000 //"注册失败"
	LoginFail                int = 1001 //"登录失败"
	DeleteUsersFail          int = 1002 //"删除用户错误"
	DeleteRolesFail          int = 1003 //"删除角色错误"
	UsernameFail             int = 1004 //"用户名错误"
	PasswordFail             int = 1005 //"密码错误"
	TokenCreateFail          int = 1006 //"生成token错误"
	TokenExactFail           int = 1007 //"token不存在或header设置不正确"
	TokenExpire              int = 1008 //"回话已过期"
	TokenParseFail           int = 1009 //"token解析错误"
	TokenParseFailAndEmpty   int = 1010 //"解析错误,token为空"
	TokenParseFailAndInvalid int = 1011 //"解析错误,token无效"
	PermissionsLess          int = 1013 //"权限不足"
	RoleCreateFail           int = 1014 //"创建角色失败"
	EmptyMaliOrPassWord      int = 1015 //"用户名或密码为空"
	PhotoReadFail            int = 1016 //"图像文件读取错误"
	DataReceiveFail          int = 1017 //"数据接收失败"
	PhotoSaveFail            int = 1018 //"图像文件保存失败"
	PhotoUpdateFail          int = 1019 //"图像更新失败"
	CancellationFail         int = 1020 //"注销失败"
	DataUpdateFail           int = 1021 //"数据更新失败"
	DataCreateFail           int = 1022 //"数据创建失败"
	DataSelectFail           int = 1023 //"数据不存在或查询失败"
	DataDeleteFail           int = 1024 //"数据删除失败"
	OrderIsEnd               int = 1025 //"工单已经结束"

)

var MessageMap = map[int]string{
	Success:                  "成功",
	Fail:                     "失败",
	RegisterFail:             "注册失败",
	DeleteUsersFail:          "删除用户错误",
	DeleteRolesFail:          "删除角色错误",
	LoginFail:                "登录失败",
	UsernameFail:             "登录失败",
	PasswordFail:             "密码错误",
	TokenCreateFail:          "生成token错误",
	TokenExactFail:           "token不存在或header设置不正确",
	TokenExpire:              "回话已过期",
	TokenParseFail:           "token解析错误",
	TokenParseFailAndEmpty:   "解析错误,token为空",
	TokenParseFailAndInvalid: "解析错误,token无效",
	NotFound:                 "您请求的url不存在",
	PermissionsLess:          "权限不足",
	RoleCreateFail:           "创建角色失败",
	EmptyMaliOrPassWord:      "用户名或密码为空",
	PhotoReadFail:            "图像文件读取错误",
	DataReceiveFail:          "数据接收失败",
	PhotoSaveFail:            "图像文件保存失败",
	PhotoUpdateFail:          "图像更新失败",
	CancellationFail:         "注销失败",
	DataUpdateFail:           "数据更新失败",
	DataCreateFail:           "数据创建失败",
	DataSelectFail:           "数据不存在或查询失败",
	DataDeleteFail:           "数据删除失败",
	OrderIsEnd:               "工单已经结束",
}
