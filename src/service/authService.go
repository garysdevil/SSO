package service

import (
	// "github.com/rs/xid"

	"fmt"
	"sso/src/model"
	"sso/src/utils"
)

func LoginService(user model.User) (string, error) {
	fmt.Println("===========1")
	err := utils.LdapValid(user.Username, user.Password)
	fmt.Println("===========11")
	if err != nil {
		// log.Error(err)
		return "", err
	}
	fmt.Println("===========2")
	roles, err := user.GetRolesByUser(user)
	if err != nil {
		return "", err
	}
	var roleidarr []string
	for i, role := range roles {
		roleidarr[i] = role.RoleID
	}
	fmt.Println("===========3")
	token, err := utils.JwtEncode(user.Username, roleidarr)
	return token, err

}

// /*
// 验证接口可以不用,现在验证放在客户端
// */
// //func CheckJwtService(token string) bool {
// //	secret,err :=utils.RedisClient().Get(token).Result()
// //	if err !=nil{
// //		fmt.Println(err)
// //		panic(err)
// //	}
// //	result,_ :=utils.JwtDecode(token,secret)
// //	//判断过期时间
// //	//if result && existTime<viper.GetInt64("token.expireTime")-viper.GetInt64("token.refreshTime") {
// //	//	return result,true
// //	//}
// //	return result
// //}

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
