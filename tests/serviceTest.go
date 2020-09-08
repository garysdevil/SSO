package main

import (
	"sso/src/model"
	"sso/src/service"
)

func serviceTest() {
	service.InsertUserFromLdap(model.User{})
}
