package cmd

import (
	"net/http"
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
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Access-Token")

		//放行所有OPTIONS方法
		method := c.Request.Method
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}

// func CrosHandler() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		method := c.Request.Method

// 		c.Header("Access-Control-Allow-Origin", "*")
// 		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
// 		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
// 		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
// 		c.Header("Access-Control-Allow-Credentials", "true")
// 		if method == "OPTIONS" {
// 			// c.AbortWithStatus(http.StatusNoContent)
// 			data := map[string]string{"data": "Options Request!"}
// 			c.JSON(http.StatusOK, data)
// 		}
// 		c.Next()
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
