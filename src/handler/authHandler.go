package handler

import (
	"fmt"
	"sso/src/handler/exception"
	"sso/src/model"
	"sso/src/service"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @Summary 登入接口
// @Tags 登陆管理
// @Accept  json
// @Produce  json
// @Param user body exception.LoginUser true "login"
// @Success 200 {object} Response "{"code":0,"data":{},"msg":"success"}"
// @Router /sso/login [post]
func LoginHandler(c *gin.Context) {
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
			Message: exception.LoginError.Message}, err)
		return
	}

	c.SetCookie("token", token, int(time.Minute*viper.GetDuration("token.expireTime")), "/", "wx.bc", false, true)
	log.Info("登录成功，用户：" + user.Username + "登陆时间:" + time.Now().String())
	data := map[string]string{"token": token}
	SendResponse(c, exception.CustomCode{Code: exception.OK.Code, Message: exception.OK.Message}, data)
}

// @Summary 验证token接口
// @Tags 登陆管理
// @Accept  json
// @Produce  json
// @Param user body exception.Token true "token"
// @Success 200 {object} Response "{"code":0,"data":{},"msg":"success"}"
// @Router /sso/check [post]
func CheckJwtHandler(c *gin.Context) {
	var token exception.Token
	err := c.Bind(&token)
	if err != nil {
		log.Info("绑定结构体错误：" + err.Error())

		SendResponse(c, exception.CustomCode{Code: exception.LoginError.Code,
			Message: exception.LoginError.Message}, err.Error())
		return
	}
	fmt.Println(token.Token + "---------")
	flag, err := service.CheckJwtService(token.Token)
	if err != nil {
		log.Info(err)
	}
	data := map[string]bool{"flag": flag}
	SendResponse(c, exception.CustomCode{Code: exception.OK.Code, Message: exception.OK.Message}, data)
}

// @Summary 登出接口
// @Tags 登陆管理
// @Accept  json
// @Produce  json
// @Param user body exception.Token false "token"
// @Success 200 {object} Response "{"code":0,"data":{},"msg":"success"}"
// @Router /sso/logout [post]
func LogoutHandler(c *gin.Context) {
	var token exception.Token
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
	data := map[string]bool{"flag": true}
	SendResponse(c, exception.CustomCode{Code: exception.OK.Code, Message: exception.OK.Message}, data)
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
