package handler

import (
	"fmt"
	"sso/src/model"
	"sso/src/service"
	"time"

	"sso/src/handler/exception"
	"sso/src/handler/schema"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @Summary 登入接口
// @Tags 登陆管理
// @Accept  json
// @Produce  json
// @Param user body schema.LoginUser true "login"
// @Success 200 {object} Response "{"code":0,"data":{},"msg":"success"}"
// @Router /v1/sso/login [post]
func LoginHandler(c *gin.Context) {
	// w := c.Writer
	// r := c.Request
	// // 处理js-ajax跨域问题
	// w.Header().Set("Access-Control-Allow-Credentials", "true")
	// w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	// w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, POST")
	// // w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
	// w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	// w.Header().Add("Access-Control-Allow-Headers", "Access-Token")

	var user model.User
	err := c.Bind(&user)
	if err != nil {
		log.Info("绑定结构体错误：" + err.Error())

		SendResponse(c, exception.CustomCode{Code: exception.LoginError.Code,
			Message: exception.LoginError.Message}, err.Error())
		return
	}
	token, err := service.LoginService(user)
	if err != nil {
		log.Error(err)
		SendResponse(c, exception.CustomCode{Code: exception.LoginError.Code,
			Message: exception.LoginError.Message}, err.Error())
		return
	}

	c.SetCookie("token", token, int(time.Minute*viper.GetDuration("token.expireTime")), viper.GetString("cookie.path"), viper.GetString("cookie.domain"), false, false)
	log.Info("登录成功，用户：" + user.Username + "登陆时间:" + time.Now().String())
	data := map[string]string{"tokenString": token}
	SendResponse(c, exception.CustomCode{Code: exception.OK.Code, Message: exception.OK.Message}, data)
}

// @Summary 验证token接口
// @Tags 登陆管理
// @Accept  json
// @Produce  json
// @Param user body schema.Token true "token"
// @Success 200 {object} Response "{"code":0,"data":{},"msg":"success"}"
// @Router /v1/sso/check [post]
func CheckJwtHandler(c *gin.Context) {
	var token schema.Token
	err := c.Bind(&token)
	if err != nil {
		log.Info("绑定结构体错误：" + err.Error())

		SendResponse(c, exception.CustomCode{Code: exception.CheckJwtError.Code,
			Message: exception.CheckJwtError.Message}, err.Error())
		return
	}
	username, _, err := service.CheckJwtService(token.Token)
	if err != nil {
		log.Info(err)
		SendResponse(c, exception.CustomCode{Code: exception.CheckJwtError.Code,
			Message: exception.CheckJwtError.Message}, err.Error())
		return
	}
	data := map[string]string{"username": username}
	SendResponse(c, exception.CustomCode{Code: exception.OK.Code, Message: exception.OK.Message}, data)
}

// @Summary 登出接口
// @Tags 登陆管理
// @Accept  json
// @Produce  json
// @Param user body schema.Token false "token"
// @Success 200 {object} Response "{"code":0,"data":{},"msg":"success"}"
// @Router /v1/sso/logout [post]
func LogoutHandler(c *gin.Context) {
	var token schema.Token
	token.Token, _ = c.Cookie("token")
	err := c.Bind(&token)

	fmt.Println(token.Token + "===")
	if err != nil {
		log.Info("获取token错误：" + err.Error())
		SendResponse(c, exception.CustomCode{Code: exception.GetCookieError.Code,
			Message: exception.GetCookieError.Message}, "")
		return
	}

	err = service.LogoutService(token.Token)
	if err != nil {
		log.Info("登出失败：" + err.Error())
		SendResponse(c, exception.CustomCode{Code: exception.LogoutError.Code,
			Message: exception.LogoutError.Message}, "")
		return
	}
	c.SetCookie("token", "", -1, "/", "wx.bc", false, true)
	log.Info("登出成功")

	SendResponse(c, exception.CustomCode{Code: exception.OK.Code, Message: exception.OK.Message}, "")
}

// // @Summary 刷新token接口
// // @Tags 登陆管理
// // @Accept  json
// // @Produce  json
// // @Success 200 {object} model.Response "{"code":0,"data":{},"msg":"success"}"
// // @Router /sso/freshtoken [get]
// func RefreshTokenHandler(c *gin.Context) {
// 	token, err := c.Cookie("token")
// 	if err != nil {
// 		log.Info("cookie获取错误：" + err.Error())
// 		SendResponse(c, exception.CustomCode{Code: exception.GetCookieError.Code,
// 			Message: exception.GetCookieError.Message}, "")
// 		return
// 	}
// 	newToken, user, err := service.RefreshToken(token)
// 	if err != nil {
// 		SendResponse(c, exception.CustomCode{Code: exception.FreshJwtError.Code,
// 			Message: exception.FreshJwtError.Message}, "")
// 		return
// 	}

// 	c.SetCookie("token", newToken, int(time.Minute*viper.GetDuration("token.expireTime")), "/", "wx.bc", false, true)
// 	log.Info("用户：" + user + "刷新token成功")
// 	SendResponse(c, exception.CustomCode{Code: exception.OK.Code,
// 		Message: exception.OK.Message}, "")
// }
