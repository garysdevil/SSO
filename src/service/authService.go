package service

import (
	// "github.com/rs/xid"

	"sso/src/model"
	"sso/src/utils"

	"github.com/spf13/viper"
)

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

// 验证token是否过期
func CheckJwtService(token string) (bool, error) {
	// secret,err :=utils.RedisClient().Get(token).Result()
	// if err !=nil{
	// 	fmt.Println(err)
	// 	panic(err)
	// }
	username, err := utils.JwtDecode(viper.GetString("token.secret"), token)
	//判断过期时间
	//if result && existTime<viper.GetInt64("token.expireTime")-viper.GetInt64("token.refreshTime") {
	//	return result,true
	//}
	if username != "" {
		return true, nil
	}
	return false, err
}

// func LogoutService(token string) error {
// 	result := utils.RedisClient().Del(token)
// 	if result.Err() != nil {
// 		return result.Err()
// 	}
// 	return nil
// }

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
