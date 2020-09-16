package main

import (
	"fmt"

	"sso/src/utils"
)

func utilsTest() {

	// test1GetAllUser()
	jwt()

	fmt.Println("\nEnd utilsTest()")
}

func test1GetAllUser() {
	// err := utils.LdapValid("xieshigang", "123456")
	// if err == nil {
	// 	fmt.Println("登陆成功")
	// }
	a, err := utils.LdapGetAllUser()
	fmt.Println(a)
	fmt.Println("--------------")
	fmt.Println(err)
}

func jwt() {
	token, err := utils.JwtEncode("谢石冈", []string{"1", "2", "3"})
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
