package utils

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var ctx = context.Background()
var RedisClient *redis.Client

func getRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.server"),   // use default Addr
		Password: viper.GetString("redis.password"), // no password set
		DB:       viper.GetInt("redis.db"),          // use default DB
	})
}
func InitRedisClient() {
	RedisClient = getRedisClient()
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		fmt.Println("redis连接失败")
		panic(err)
	} else {
		fmt.Println("redis连接成功")
	}

}
