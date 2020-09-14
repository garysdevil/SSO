package main

import (
	"sso/src/model"
	"sso/src/service"
)

func serviceTest() {
	// 角色
	role := model.Role{
		RoleName: "aaa",
	}
	service.CreateRole(role)
}
