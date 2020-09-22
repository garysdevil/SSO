package cmd

import (
	"sso/src/middleware"
	"sso/src/router"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func CrosHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		w := c.Writer
		r := c.Request
		// 处理js-ajax跨域问题
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, POST, GET, OPTIONS, PUT, DELETE,UPDATE")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Add("Access-Control-Allow-Headers", "Access-Token")
		c.Next()
	}
}

// //跨域访问：cross  origin resource share
// func CrosHandler() gin.HandlerFunc {
// 	return func(context *gin.Context) {
// 		// method := context.Request.Method
// 		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		context.Header("Access-Control-Allow-Origin", "*") // 设置允许访问所有域
// 		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
// 		context.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma,token,openid,opentoken")
// 		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
// 		context.Header("Access-Control-Max-Age", "172800")
// 		context.Header("Access-Control-Allow-Credentials", "false")
// 		context.Set("content-type", "application/json") // 设置返回格式是json

// 		// if method == "OPTIONS" {
// 		// 	context.JSON(http.StatusOK, result.Result{Code: result.OK, Data: "Options Request!"})
// 		// }

// 		//处理请求
// 		context.Next()
// 	}
// }

func Execute() {

	// gin.SetMode(gin.ReleaseMode)
	server := gin.New()

	server.Use(middleware.Logger(), gin.Recovery(), CrosHandler())

	router.Router(server)

	serverAddr := viper.GetString("server.addr")
	log.Println("server start at " + serverAddr)
	server.Run(serverAddr)

}
