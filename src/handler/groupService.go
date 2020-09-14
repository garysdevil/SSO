package handler

import (
	"encoding/json"
	"sso/src/handler/exception"
	"sso/src/model"
	"sso/src/service"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// @Summary 创建用户组接口
// @Tags 用户组管理
// @Accept  json
// @Produce  json
// @Param user body model.Group true "create"
// @Success 200 {object} Response "{"code":0,"data":{},"msg":"success"}"
// @Router /sso/group/create [post]
func CreateRoleHandler(c *gin.Context) {
	var role model.Role
	err := c.Bind(&role)
	if err != nil {
		log.Error(err)
		SendResponse(c, exception.CustomCode{Code: exception.CreateRoleError.Code,
			Message: exception.CreateRoleError.Message}, "")
		return
	}
	data, _ := json.Marshal(role)
	log.Info(string(data))
	err = service.CreateRole(role)
	if err != nil {
		log.Error(err)
		SendResponse(c, exception.CustomCode{Code: exception.CreateRoleError.Code,
			Message: exception.CreateRoleError.Message}, "")
		return
	}
	SendResponse(c, exception.CustomCode{Code: exception.OK.Code, Message: exception.OK.Message}, role)

}
