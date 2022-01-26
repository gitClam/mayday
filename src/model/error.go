package model

const (
	SuccessCode    int = 1200
	NotFoundCode   int = 1404
	DbErrorCode    int = 1403
	OtherErrorCode int = 1402
)
const (
	// key定义
	CODE string = "code"
	MSG  string = "msg"
	DATA string = "data"

	// msg define
	Success           string = "成功"
	OptionSuccess     string = "操作成功"
	OptionFailur      string = "操作失败"
	ParseParamsFailur string = "解析参数失败"

	RegisteSuccess     string = "注册用户成功"
	RegisteFailur      string = "注册失败"
	LoginSuccess       string = "登录成功"
	LoginFailur        string = "登录失败"
	DeleteUsersSuccess string = "删除用户成功"
	DeleteUsersFailur  string = "删除用户错误"

	DeleteRolesSuccess string = "删除角色成功"
	DeleteRolesFailur  string = "删除角色错误"

	UsernameFailur             string = "用户名错误"
	PasswordFailur             string = "密码错误"
	TokenCreateFailur          string = "生成token错误"
	TokenExactFailur           string = "token不存在或header设置不正确"
	TokenExpire                string = "回话已过期"
	TokenParseFailur           string = "token解析错误"
	TokenParseFailurAndEmpty   string = "解析错误,token为空"
	TokenParseFailurAndInvalid string = "解析错误,token无效"
	NotFound                   string = "您请求的url不存在"
	PermissionsLess            string = "权限不足"

	RoleCreateFailur  string = "创建角色失败"
	RoleCreateSuccess string = "创建角色成功"

	// value define

)
