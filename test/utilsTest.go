package main

import (
	"fmt"

	"sso/cmd"
	"sso/src/utils"
)

func main1() {
	if err := cmd.InitConfig("../config/settings.dev.yaml"); err != nil {
		panic(err)
	}
	ldap()
	jwt()
	fmt.Println("end")
}

func ldap() {
	err := utils.ValidLdap("xieshigang", "123456")
	if err == nil {
		fmt.Println("登陆成功")
	}
}

func jwt() {
	token, err := utils.JwtEncode("谢石冈", []int{1, 2, 3})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(token)
	}

	a, err := utils.JwtDecode("123456", token)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(a)
	}
}
