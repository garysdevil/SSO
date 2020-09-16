package service

import (
	"sso/src/model"

	"github.com/rs/xid"
	log "github.com/sirupsen/logrus"
)

func CreateGroup(group model.Group) (err error){

	group.GroupID = xid.New().String()
	
	err = group.CreateGroup()

	if err == nil {
		log.Info("用户组：" + group.Groupname + "创建成功")
	}

	return
}