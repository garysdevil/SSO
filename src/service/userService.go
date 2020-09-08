package service

import (
	"fmt"
	"sso/src/model"
	"sso/src/utils"

	"github.com/rs/xid"
	log "github.com/sirupsen/logrus"
)

func InsertUserFromLdap(user model.User) bool {
	ldapUsernames, err := utils.LdapGetAllUser()
	if err != nil {
		log.Errorln(err)
		return false
	}
	var users []model.User = make([]model.User, len(ldapUsernames))
	for i, username := range ldapUsernames {

		users[i] = model.User{
			UserID:   xid.New().String(),
			Username: username,
		}

		fmt.Println(i, username)

	}

	if err := user.CreateUsers(&users); err != nil {
		return true
	} else {
		log.Error(err)
		return false
	}

}
