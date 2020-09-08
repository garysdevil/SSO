package model

import (
	"encoding/json"
	"fmt"
)

type TokenInfo struct {
	UserID   string   `json:"userid"`
	Username string   `json:"username"`
	RoleIDs  []string `json:"roleids"`
}

func (user User) GetTokenInfo() (TokenInfo, error) {
	DB.First(&user)

	var tokenInfo TokenInfo = TokenInfo{}

	var roles []Role

	err := DB.Model(&user).Association("Roles").Find(&user, "Roles")

	tokenInfo.UserID = user.UserID
	tokenInfo.Username = user.Username

	for index := range roles {
		tokenInfo.RoleIDs[index] = roles[index].RoleID
	}

	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
	return tokenInfo, err
}
