package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type WebCustomCode struct {
	Code    int
	Message string
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var (
	OK              = &WebCustomCode{Code: 0, Message: "OK"}
	LoginError      = &WebCustomCode{Code: 10001, Message: "登入失败"}
	LogoutError     = &WebCustomCode{Code: 10002, Message: "登出失败"}
	FreshJwtError   = &WebCustomCode{Code: 10003, Message: "刷新token失败"}
	GetCookieError  = &WebCustomCode{Code: 10004, Message: "cookie获取失败"}
	CreateRoleError = &WebCustomCode{Code: 10005, Message: "角色创建失败"}
	DeleteRoleError = &WebCustomCode{Code: 10006, Message: "角色删除失败"}
	UpdateRoleError = &WebCustomCode{Code: 10007, Message: "角色更新失败"}
	GetRoleError    = &WebCustomCode{Code: 10008, Message: "角色查询失败"}
	ForbidUserError = &WebCustomCode{Code: 10010, Message: "用户禁用失败"}
	DeleteUserError = &WebCustomCode{Code: 10011, Message: "用户删除失败"}
	UpdateUserError = &WebCustomCode{Code: 10012, Message: "用户更新失败"}
	GetUserError    = &WebCustomCode{Code: 10013, Message: "用户查询失败"}
	CreateMenuError = &WebCustomCode{Code: 10020, Message: "菜单创建失败"}
	UpdateMenuError = &WebCustomCode{Code: 10021, Message: "菜单更新失败"}
	DeleteMenuError = &WebCustomCode{Code: 10022, Message: "菜单删除失败"}
)

//response统一返回调用方法
func WebSendResponse(c *gin.Context, custom WebCustomCode, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    custom.Code,
		Message: custom.Message,
		Data:    data,
	})
}
