package utils

import (
	"fmt"

	"github.com/go-ldap/ldap/v3"
	"github.com/spf13/viper"
)

type LdapConf struct {
	Server   string `json:"server"`
	Port     int    `json:"port"`
	BaseDN   string `json:"baseDN"`
	Username string `json:"username"`
	Password string `json:"password"`
	Domain   string `json:"domain"`
}

func ValidLdap(username, password string) error {
	var ldapConf = LdapConf{
		Server:   viper.GetString("ldap.server"),
		Port:     viper.GetInt("ldap.port"),
		BaseDN:   viper.GetString("ldap.baseDN"),
		Username: viper.GetString("ldap.username"),
		Password: viper.GetString("ldap.password"),
		Domain:   viper.GetString("ldap.domain"),
	}
	return _ValidLdap(ldapConf, username, password)
}

func _ValidLdap(ldapConf LdapConf, username, password string) error {

	// 连接
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", ldapConf.Server, ldapConf.Port))
	if err != nil {
		return err
	}
	defer l.Close()
	// 绑定管理员用户
	err = l.Bind(ldapConf.Username, ldapConf.Password)
	if err != nil {
		return err
	}
	// 验证用户
	controls := []ldap.Control{}
	controls = append(controls, ldap.NewControlBeheraPasswordPolicy())
	bindRequest := ldap.NewSimpleBindRequest(username+ldapConf.Domain, password, controls)

	r, err := l.SimpleBind(bindRequest)
	if err != nil {
		return err
	}
	// 验证策略
	ppolicyControl := ldap.FindControl(r.Controls, ldap.ControlTypeBeheraPasswordPolicy)
	var ppolicy *ldap.ControlBeheraPasswordPolicy
	if ppolicyControl != nil {
		ppolicy = ppolicyControl.(*ldap.ControlBeheraPasswordPolicy)
	}
	// else {
	// 	log.Info("ppolicyControl response not available.\n")
	// }
	if ppolicy != nil {
		if ppolicy.Expire >= 0 {
			return fmt.Errorf(". Password expires in %d seconds\n", ppolicy.Expire)
		} else if ppolicy.Grace >= 0 {
			return fmt.Errorf(". Password expired, %d grace logins remain\n", ppolicy.Grace)
		}
	}
	return nil
}
