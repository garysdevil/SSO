package main

import (
	"fmt"
	"os"
	"sso/cmd"

	// "sso/src/model"

	"sso/src/utils"
)

func initAll() {
	if err := cmd.InitConfig(""); err != nil {
		panic(err)
	}
	// model.InitDB()

	utils.InitRedisClient()
}
func args() {
	// arg1 := flag.String("arg", "", "whether init db tables")
	// flag.Parse()
	// fmt.Println(*arg1 + "=========")
	fmt.Println(os.Args)
	if len(os.Args) == 1 {
		cmd.Execute()
	}
	for _, arg := range os.Args {
		if arg == "initdb" {
			cmd.AutoMigrateTable()
			break
		}
		if arg == "start" {
			cmd.Execute()
			break
		}

	}
}
func main() {

	initAll()

	args()

}
