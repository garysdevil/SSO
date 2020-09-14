package handler

import (
	"net/http"

	"sso/src/handler/exception"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//response统一返回调用方法
func SendResponse(c *gin.Context, custom exception.CustomCode, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    custom.Code,
		Message: custom.Message,
		Data:    data,
	})
}
