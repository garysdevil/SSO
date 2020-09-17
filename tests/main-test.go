package main

import (
	"context"
	"fmt"
	"sso/cmd"
	"sso/src/model"
	"time"

	"github.com/go-redis/redis/v8"
)

func countRuntime(f func()) {
	beginTime := time.Now()
	// fmt.Println("beginTime: " + beginTime.Format("2006-01-02 3:04:05.000 PM Mon Jan"))
	fmt.Println("beginTime: " + beginTime.Format("2006年01月02日 15点04分05秒"))
	fmt.Println()

	f()

	endTime := time.Now()
	fmt.Println()
	fmt.Println("endTime: " + endTime.Format("2006年01月02日 15点04分05秒"))
	fmt.Println("Spending time: " + endTime.Sub(beginTime).String())
}

func main() {
	countRuntime(personTest)
	//countRuntime(projectTest)
}

func projectTest() {
	if err := cmd.InitConfig("../config/settings.dev.yaml"); err != nil {
		panic(err)
	}
	model.InitDB()

	// modelTest()
	// utilsTest()
	// serviceTest()
}

func personTest() {

	fmt.Println("aaa")
	ExampleNewClient()
}

var ctx = context.Background()

func ExampleNewClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "10.200.79.81:63791",
		Password: "", // no password set
		DB:       1,  // use default DB
	})

	pong, err := rdb.Ping(ctx).Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
	// err = rdb.Set(ctx, "key2", "value1", 100000000000000000).Err()

	isKeyExit, err := rdb.Exists(ctx, "key2").Result()
	fmt.Println(isKeyExit)

}
