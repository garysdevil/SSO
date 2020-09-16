package handler

import (
	"encoding/json"
	"sso/src/handler/exception"
	"sso/src/model"
	"sso/src/service"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// @Summary 创建用户组组接口
// @Tags 用户组组管理
// @Accept  json
// @Produce  json
// @Param user body model.Group true "create"
// @Success 200 {object} Response "{"code":0,"data":{},"msg":"success"}"
// @Router /sso/group/create [post]
func CreateGroupHandler(c *gin.Context) {
	var group model.Group
	err := c.Bind(&group)
	if err != nil {
		log.Error(err)
		SendResponse(c, exception.CustomCode{Code: exception.CreateGroupError.Code,
			Message: exception.CreateGroupError.Message}, "")
		return
	}
	data, _ := json.Marshal(group)
	log.Info(string(data))
	err = service.CreateGroup(group)
	if err != nil {
		log.Error(err)
		SendResponse(c, exception.CustomCode{Code: exception.CreateGroupError.Code,
			Message: exception.CreateGroupError.Message}, "")
		return
	}
	SendResponse(c, exception.CustomCode{Code: exception.OK.Code, Message: exception.OK.Message}, group)

}
