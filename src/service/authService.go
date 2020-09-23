package service

import (
	// "github.com/rs/xid"

	"context"
	"fmt"
	"sso/src/model"
	"sso/src/utils"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var ctx = context.Background()

func LoginService(user model.User) (string, error) {
	err := utils.LdapValid(user.Username, user.Password)
	if err != nil {
		// log.Error(err)
		return "", err
	}
	roles, err := user.GetRolesByUser(user)
	if err != nil {
		return "", err
	}
	var roleidarr []string
	for i, role := range roles {
		roleidarr[i] = role.RoleID
	}

	token, err := utils.JwtEncode(user.Username, roleidarr)
	return token, err

}

// 验证token是否有效, 有效则返回nil username roleidarr
func CheckJwtService(token string) (string, []string, error) {
	// 查看token是否是被登出的但未过期的
	result, err := utils.RedisClient.Exists(ctx, token).Result()
	if err != nil {
		return "", []string{}, err
	}
	if result == 1 {
		return "", []string{}, fmt.Errorf("此token已经被用户执行过登出操作")
	}

	username, roleidarr, err := utils.JwtDecode(viper.GetString("token.secret"), token)
	if err != nil {
		return "", []string{}, err
	}
	return username, roleidarr, nil
}

func LogoutService(token string) error {

	_, _, err := utils.JwtDecode(viper.GetString("token.secret"), token)
	if err != nil {
		log.Info("登出：无效的token")
		return nil
	}

	err = utils.RedisClient.Set(ctx, token, "-", time.Minute*viper.GetDuration("token.expireTime")).Err()
	if err != nil {
		return err
	}
	return nil
}

func RefreshToken(token string) (string, error) {
	// 验证token是否有效
	result, err := utils.RedisClient.Exists(ctx, token).Result()
	if err != nil {
		return "", err
	}
	if result == 1 {
		return "", fmt.Errorf("此token已经被用户执行过登出操作")
	}
	username, roleidarr, err := utils.JwtDecode(viper.GetString("token.secret"), token)
	if err != nil {
		return "", err
	}
	// 生成新的token
	token, err = utils.JwtEncode(username, roleidarr)

	return token, err
}
