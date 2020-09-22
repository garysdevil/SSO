package exception

type CustomCode struct {
	Code    int
	Message string
}

var (
	OK             = &CustomCode{Code: 0, Message: "OK"}
	LoginError     = &CustomCode{Code: 10001, Message: "登入失败"}
	LogoutError    = &CustomCode{Code: 10002, Message: "登出失败"}
	CheckJwtError  = &CustomCode{Code: 10002, Message: "验证token失败"}
	FreshJwtError  = &CustomCode{Code: 10004, Message: "刷新token失败"}
	GetCookieError = &CustomCode{Code: 10005, Message: "cookie获取失败"}

	CreateRoleError = &CustomCode{Code: 10010, Message: "角色创建失败"}
	DeleteRoleError = &CustomCode{Code: 10011, Message: "角色删除失败"}
	UpdateRoleError = &CustomCode{Code: 10012, Message: "角色更新失败"}
	GetRoleError    = &CustomCode{Code: 10013, Message: "角色查询失败"}

	ForbidUserError = &CustomCode{Code: 10020, Message: "用户禁用失败"}
	CreateUserError = &CustomCode{Code: 10021, Message: "用户创建失败"}
	DeleteUserError = &CustomCode{Code: 10022, Message: "用户删除失败"}
	UpdateUserError = &CustomCode{Code: 10023, Message: "用户更新失败"}
	GetUserError    = &CustomCode{Code: 10024, Message: "用户查询失败"}

	CreateGroupError = &CustomCode{Code: 10030, Message: "用户组创建失败"}
	DeleteGroupError = &CustomCode{Code: 10031, Message: "用户组删除失败"}
	UpdateGroupError = &CustomCode{Code: 10032, Message: "用户组更新失败"}
	GetGroupError    = &CustomCode{Code: 10033, Message: "用户组查询失败"}

	CreateMenuError = &CustomCode{Code: 10040, Message: "菜单创建失败"}
	DeleteMenuError = &CustomCode{Code: 10041, Message: "菜单删除失败"}
	UpdateMenuError = &CustomCode{Code: 10042, Message: "菜单更新失败"}
	GetMenuError    = &CustomCode{Code: 10043, Message: "菜单查询失败"}
)
