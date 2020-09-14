package handler

// import (
// 	"sso/src/model"
// 	"sso/src/service"
// 	"sso/src/exception"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	log "github.com/sirupsen/logrus"
// 	"github.com/spf13/viper"
// )

// // @Summary 登陆接口
// // @Tags 登陆管理
// // @Accept  json
// // @Produce  json
// // @Param user body schema.LoginUser true "login"
// // @Success 200 {object} model.Response "{"code":0,"data":{},"msg":"success"}"
// // @Router /sso/login [post]
// func LoginHandler(c *gin.Context) {
// 	var user model.User
// 	err := c.Bind(&user)
// 	if err != nil {
// 		log.Info("绑定结构体错误：" + err.Error())
// 		WebSendResponse(c, WebCustomCode{Code: LoginError.Code,
// 			Message: LoginError.Message}, err.Error)
// 		return
// 	}
// 	token, user, err := service.LoginService(user)
// 	if err != nil {
// 		log.Info(err.Error)
// 		WebSendResponse(c, WebCustomCode{Code: exception.LoginError.Code,
// 			Message: exception.LoginError.Message}, err.Error)
// 		return
// 	}
// 	c.SetCookie("token", token, int(time.Minute*viper.GetDuration("token.expireTime")), "/", "wx.bc", false, true)
// 	log.Info("登录成功，用户：" + user.Username + "登陆时间:" + time.Now().String())
// 	//data := model.User{
// 	//	Username: user.Username,
// 	//}
// 	WebSendResponse(c, WebCustomCode{Code: exception.OK.Code, Message: exception.OK.Message}, user)
// }

// //func CheckJwtHandler(c *gin.Context)  {
// //	token :=c.Query("jwt")
// //	result :=service.CheckJwtService(token)
// //	if result {
// //		//newToken,err :=service.RefreshToken(token)
// //		//if err!=nil{
// //		//	panic(err)
// //		//}
// //		//c.SetCookie("token",newToken,0,"/","wx.bc",false,true)
// //		c.String(http.StatusOK,"true")
// //	}
// //	//c.String(http.StatusOK,"true")
// //}

// // @Summary 登出接口
// // @Tags 登陆管理
// // @Accept  json
// // @Produce  json
// // @Success 200 {object} model.Response "{"code":0,"data":{},"msg":"success"}"
// // @Router /sso/logout [get]
// func LogoutHandler(c *gin.Context) {
// 	token, err := c.Cookie("token")
// 	if err != nil {
// 		log.Info("获取token错误：" + err.Error())
// 		WebSendResponse(c, WebCustomCode{Code: exception.GetCookieError.Code,
// 			Message: exception.GetCookieError.Message}, "")
// 		return
// 	}
// 	err = service.LogoutService(token)
// 	if err != nil {
// 		log.Info("登出失败：" + err.Error())
// 		WebSendResponse(c, WebCustomCode{Code: exception.LogoutError.Code,
// 			Message: exception.LogoutError.Message}, "")
// 		return
// 	}
// 	c.SetCookie("token", "", -1, "/", "wx.bc", false, true)
// 	log.Info("登出成功")
// 	WebSendResponse(c, WebCustomCode{Code: exception.OK.Code, Message: exception.OK.Message}, "")
// }

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
// 		WebSendResponse(c, WebCustomCode{Code: exception.GetCookieError.Code,
// 			Message: exception.GetCookieError.Message}, "")
// 		return
// 	}
// 	newToken, user, err := service.RefreshToken(token)
// 	if err != nil {
// 		WebSendResponse(c, WebCustomCode{Code: exception.FreshJwtError.Code,
// 			Message: exception.FreshJwtError.Message}, "")
// 		return
// 	}

// 	c.SetCookie("token", newToken, int(time.Minute*viper.GetDuration("token.expireTime")), "/", "wx.bc", false, true)
// 	log.Info("用户：" + user + "刷新token成功")
// 	WebSendResponse(c, WebCustomCode{Code: exception.OK.Code,
// 		Message: exception.OK.Message}, "")
// }
