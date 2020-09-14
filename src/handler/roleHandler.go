package handler

import (
	"encoding/json"
	"sso/src/handler/exception"
	"sso/src/model"
	"sso/src/service"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// @Summary 创建角色接口
// @Tags 角色管理
// @Accept  json
// @Produce  json
// @Param user body model.Role true "create"
// @Success 200 {object} Response "{"code":0,"data":{},"msg":"success"}"
// @Router /sso/role/create [post]
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
