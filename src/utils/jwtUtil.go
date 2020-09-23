package utils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

type TokenClaims struct {
	jwt.StandardClaims
	Username  string   `json:"username"`
	RoleIDArr []string `json:"roleid"`
}

// JwtEncode (用户，角色，有效期是多久分钟，密钥) 返回 token
func _JwtEncode(claims TokenClaims, secret string) (string, error) {

	token, error := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
	return token, error
}

// JwtEncode (用户，角色) 返回 token
func JwtEncode(username string, roleidarr []string) (string, error) {
	validTime := viper.GetInt64("token.expireTime")
	secret := viper.GetString("token.secret")
	claims := TokenClaims{
		Username:  username,
		RoleIDArr: roleidarr,
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Minute * time.Duration(validTime)).Unix()
	return _JwtEncode(claims, secret)
}

// JwtDecode （密钥，token）返回 用户名, 角色id[]
func JwtDecode(secret string, tokenString string) (string, []string, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return "", nil, fmt.Errorf("jwt验证错误：" + err.Error())
	}

	if !token.Valid {
		return "", nil, fmt.Errorf("jwt验证 token is not valid")
	}
	//existTime :=(int64(token.Claims.(jwt.MapClaims)["exp"].(float64))-time.Now().Unix())/60
	user := token.Claims.(jwt.MapClaims)["username"]
	roleidarr := token.Claims.(jwt.MapClaims)["roleidarr"]
	if roleidarr == nil {
		roleidarr = []string{}
	}
	//fmt.Println(existTime)
	return user.(string), roleidarr.([]string), nil
}
