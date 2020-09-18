package cmd

import (
	"sso/src/middleware"
	"sso/src/router"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Execute() {

	// gin.SetMode(gin.ReleaseMode)
	server := gin.New()

	server.Use(middleware.Logger(), gin.Recovery())

	router.Router(server)

	serverAddr := viper.GetString("server.addr")
	log.Println("server start at " + serverAddr)
	server.Run(serverAddr)

}
