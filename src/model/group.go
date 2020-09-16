package model

//用户组增删改查
func (group *Group) CreateGroup() error {
	return DB.Create(&group).Error
}

// // 用户与用户组进行关联
// func (group *User) UserGroupRelate() {
// 	groups := user.Groups
// 	for _, group := range groups {
// 		groupid := group.GroupID
// 		fmt.Println(groupid)
// 		a := Group{GroupID: groupid}
// 		DB.Debug().Model(&user).Association("Groups").Append([]Group{a})

// 	}
// 	// groupsa := []Group{}
// 	// fmt.Println(DB.Debug().Model(&user).Association("Groups").Find(&groupsa))
// }