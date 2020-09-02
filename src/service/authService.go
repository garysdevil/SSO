package service

import (
	"github.com/rs/xid"
	log "github.com/sirupsen/logrus"

	"sso/src/model"
	"sso/src/utils"
	"time"
)

func LoginService(user model.User) (string, model.User, error) {
	err := utils.ValidLdap(user.Username, user.Password)
	if err == nil {
		token, err := utils.JwtEncode(user.Username, user.Roles)
		if err != nil {
			log.Error(err)
			return "", user, err
		}
		u, menus, err := user.FindUser()
		if err != nil {
			log.Error(err)
			return "", user, err
		}
		if u == nil {
			user.UserId = xid.New().String()
			err = user.CreateUser()
			if err != nil {
				return "", user, err
			}
			return token, user, err
		}
		//测试代码 start
		var menuList []model.MenuTest
		for j := range menus {
			for v := range menus {
				if menus[j].Pid == menus[v].MenuId {
					menus[v].Children = append(menus[v].Children, menus[j])
				}
			}
		}
		for i := range menus {
			if menus[i].Pid == "0" {
				menuList = append(menuList, menus[i])
			}
		}
		//for j := range menus{
		//	for v:=range menuList{
		//		if menus[j].Pid == menuList[v].MenuId{
		//			menuList[v].Children = append(menuList[v].Children, menus[j])
		//			break
		//		}
		//	}
		//}
		u.Menus = menuList
		//end
		return token, *u, result.Err()
	} else {
		log.Error(err)
		return "", user, err
	}
}

/*
验证接口可以不用,现在验证放在客户端
*/
//func CheckJwtService(token string) bool {
//	secret,err :=utils.RedisClient().Get(token).Result()
//	if err !=nil{
//		fmt.Println(err)
//		panic(err)
//	}
//	result,_ :=utils.JwtDecode(token,secret)
//	//判断过期时间
//	//if result && existTime<viper.GetInt64("token.expireTime")-viper.GetInt64("token.refreshTime") {
//	//	return result,true
//	//}
//	return result
//}

func LogoutService(token string) error {
	result := utils.RedisClient().Del(token)
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}

func RefreshToken(token string) (string, string, error) {
	oldSecret, err := utils.RedisClient().Get(token).Result()
	if err != nil {
		log.Info(err)
		return "", "", err
	}
	user, err := utils.JwtDecode(token, oldSecret)
	if err != nil {
		return "", "", err
	}
	newToken, newSecret, err := utils.JwtEncode(user)
	if err != nil {
		log.Info("jwt编码错误：" + err.Error())
		return "", "", err
	}
	//60秒，后续从配置读取key过期时间
	_, err = utils.RedisClient().Set(newToken, newSecret, 60*time.Minute).Result()

	return token, user, err
}
