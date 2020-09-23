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

	// 验证此用户是否已经有token
	token, err := utils.RedisClient.Get(ctx, user.Username).Result()
	if token != "" {
		return token, nil
	}
	if err.Error() != "redis: nil" {
		return "", err
	}

	// 生成新的token
	roles, err := user.GetRolesByUser(user)
	if err != nil {
		return "", err
	}
	var roleidarr []string
	for i, role := range roles {
		roleidarr[i] = role.RoleID
	}
	token, err = utils.JwtEncode(user.Username, roleidarr)
	if err != nil {
		return "", err
	}
	_, err = utils.RedisClient.Set(ctx, user.Username, token, time.Minute*viper.GetDuration("token.expireTime")).Result()
	if err != nil {
		return "", err
	}
	return token, nil

}

// 验证token是否有效, 有效则返回nil
func CheckJwtService(token string) error {
	result, err := utils.RedisClient.Exists(ctx, token).Result()
	if err != nil {
		return err
	}

	if result != 1 {
		return fmt.Errorf("redis: token为空; result=" + string(result))
	}

	// username, err := utils.JwtDecode(viper.GetString("token.secret"), token)

	return nil
}

func LogoutService(token string) error {

	_, err := utils.JwtDecode(viper.GetString("token.secret"), token)
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

func LogoutService(token string) error {

	_, err := utils.JwtDecode(viper.GetString("token.secret"), token)
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

// func RefreshToken(token string) (string, string, error) {
// 	oldSecret, err := utils.RedisClient().Get(token).Result()
// 	if err != nil {
// 		log.Info(err)
// 		return "", "", err
// 	}
// 	user, err := utils.JwtDecode(token, oldSecret)
// 	if err != nil {
// 		return "", "", err
// 	}
// 	newToken, newSecret, err := utils.JwtEncode(user)
// 	if err != nil {
// 		log.Info("jwt编码错误：" + err.Error())
// 		return "", "", err
// 	}
// 	//60秒，后续从配置读取key过期时间
// 	_, err = utils.RedisClient().Set(newToken, newSecret, 60*time.Minute).Result()

// 	return token, user, err
// }
