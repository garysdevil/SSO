package router

import (
	_ "sso/src/docs" // swagger docs文档的位置
	"sso/src/handler"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"
)

func Router(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/sso/login", handler.LoginHandler)
	router.POST("/sso/check", handler.CheckJwtHandler)
	// router.GET("/sso/logout", handler.LogoutHandler)
	// router.GET("/sso/freshtoken", handler.RefreshTokenHandler)

	r := router.Group("/sso/role/")
	{
		//r.POST("/menu/add",handler.RoleMenuRelateHandler)
		r.POST("/create", handler.CreateRoleHandler)
		// r.POST("/update",handler.UpdateRoleHandler)
		// r.GET("/list",handler.ListRoleHandler)
		// r.GET("/get/:id",handler.GetRoleHandler)
		// r.DELETE("/delete/:id",handler.DeleteRoleHandler)
	}

	g := router.Group("/sso/group/")
	{
		//r.POST("/menu/add",handler.RoleMenuRelateHandler)
		g.POST("/create", handler.CreateGroupHandler)
		// r.POST("/update",handler.UpdateRoleHandler)
		// r.GET("/list",handler.ListRoleHandler)
		// r.GET("/get/:id",handler.GetRoleHandler)
		// r.DELETE("/delete/:id",handler.DeleteRoleHandler)
	}
}
