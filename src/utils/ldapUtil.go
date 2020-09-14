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

var ldapConf LdapConf

func _LdapConn() (*ldap.Conn, error) {
	ldapConf = LdapConf{
		Server:   viper.GetString("ldap.server"),
		Port:     viper.GetInt("ldap.port"),
		BaseDN:   viper.GetString("ldap.baseDN"),
		Username: viper.GetString("ldap.username"),
		Password: viper.GetString("ldap.password"),
		Domain:   viper.GetString("ldap.domain"),
	}
	// 连接
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", ldapConf.Server, ldapConf.Port))
	if err != nil {
		return nil, err
	}

	// 绑定管理员用户
	err = l.Bind(ldapConf.Username, ldapConf.Password)
	if err != nil {
		return nil, err
	}
	return l, nil
}

// LdapValid 验证用户账户密码
func LdapValid(username, password string) error {
	// return _LdapValid(username, password, viper.GetString("ldap.domain"))
	return _LdapValid(username, password, ldapConf.Domain)
}
func _LdapValid(username, password, domain string) error {

	l, err := _LdapConn()
	defer l.Close()
	// 验证用户
	controls := []ldap.Control{}
	controls = append(controls, ldap.NewControlBeheraPasswordPolicy())
	bindRequest := ldap.NewSimpleBindRequest(username+domain, password, controls)

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

// LdapGetAllUser 获取ldap里面所有的用户名
func LdapGetAllUser() ([]string, error) {
	l, err := _LdapConn()
	if err != nil {
		return nil, err
	}
	defer l.Close()

	searchRequest := ldap.NewSearchRequest(
		"OU=上海万向区块链股份公司,OU=Horizon Users,OU=Horizon,DC=wx,DC=bc", // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(&(objectClass=organizationalPerson))", // The filter to apply
		[]string{"dn", "cn"},                    // A list attributes to retrieve
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		// log.Fatal(err)
		return nil, err
	}

	length := len(sr.Entries)
	ldapUsers := make([]string, length)
	for i, entry := range sr.Entries {
		ldapUsers[i] = entry.GetAttributeValue("cn")
		// fmt.Printf("%s:=== %v\n", entry.DN, entry.GetAttributeValue("cn"))
	}
	return ldapUsers, nil
}
