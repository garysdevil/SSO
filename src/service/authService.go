package service

import (
	// "github.com/rs/xid"

	"context"
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

// 验证token是否有效, 有效则返回nil
func CheckJwtService(token string) error {
	// secret, err := utils.RedisClient().Get(token).Result()
	// if err != nil {
	// 	fmt.Println(err)
	// 	panic(err)
	// }
	if result, err := utils.RedisClient.Exists(ctx, token).Result(); err != nil {
		log.Info(err)
	} else {
		if result == 1 {
			return nil
		}
	}

	username, err := utils.JwtDecode(viper.GetString("token.secret"), token)

	if username != "" {
		return nil
	}
	return err
}

func LogoutService(token string) error {

	_, err := utils.JwtDecode(viper.GetString("token.secret"), token)
	if err != nil {
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
