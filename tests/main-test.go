package main

import (
	"fmt"
	"sso/cmd"
	"sso/src/model"
	"time"
)

func main() {
	beginTime := time.Now()
	// fmt.Println("beginTime: " + beginTime.Format("2006-01-02 3:04:05.000 PM Mon Jan"))
	fmt.Println("beginTime: " + beginTime.Format("2006年01月02日 15点04分05秒"))
	fmt.Println()
	if err := cmd.InitConfig("../config/settings.dev.yaml"); err != nil {
		panic(err)
	}
	model.InitDB()

	test()

	endTime := time.Now()
	fmt.Println()
	fmt.Println("endTime: " + endTime.Format("2006年01月02日 15点04分05秒"))
	fmt.Println("Spending time: " + endTime.Sub(beginTime).String())
}

func test() {
	// modelTest()
	// utilsTest()
	serviceTest()
}
