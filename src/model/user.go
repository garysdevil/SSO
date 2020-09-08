package model

//用户增删改查
func (user *User) CreateUser() error {
	return DB.Create(&user).Error
}

func (user *User) CreateUsers(users *[]User) error {
	return DB.Create(&users).Error
}

// func (user *User) CreateUsers(users []User) error {
// 	for use := range users {
// 		return DB.Create(&use).Error
// 	}
// }

// func (user *User) UpdateUser() error {
// 	roles := user.Roles
// 	DB.Model(&user).Association("Roles").Clear()
// 	//DB.Model(&role).Save(&role)
// 	user.Roles = roles
// 	return DB.Set("gorm:association_autocreate", false).Set("gorm:association_autoupdate", false).
// 		Model(&user).Update("user_name", user.Username).Error
// }

// func (user *User) DeleteUser() error {
// 	//先清引用
// 	DB.Model(&user).Association("Roles").Clear()
// 	return DB.Where("user_id = ?", user.ID).Delete(&user).Error
// }

// func (user *User) ForbidUser() error {
// 	return DB.Model(&user).Where("user_name=?", user.Username).Update("status", "inactive").Error
// }

// func (user *User) FindUser() (*User, []MenuTest, error) {
// 	var u User
// 	db := DB.Where("user_name=?", user.Username).First(&user)
// 	//db :=DB.Model(&user).Preload("Roles").First(&user).Scan(&u)
// 	if db.RecordNotFound() {
// 		return nil, nil, nil
// 	}
// 	u, menus, err := user.GetUser()
// 	if err != nil {
// 		return &u, nil, err
// 	}
// 	return &u, menus, db.Error
// }

// func (user *User) ListUser(pageNo, pageSize int) ([]User, int32) {
// 	var count int32
// 	DB.Model(&user).Count(&count)
// 	var userList []User
// 	DB.Model(&user).Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(&userList)

// 	return userList, count
// }

// //用户角色关联（这块代码需要改进）
// //func (user *User) UserRoleRelate()  {
// //	roles :=user.Roles
// //	//DB.Preload("Roles").First(&user)
// //	for _,role :=range roles{
// //		roleId := role.RoleId
// //		//fmt.Println(roleId)
// //		DB.Debug().Model(&user).Association("Roles").Append(Role{RoleId:roleId})
// //	}
// //}

// func (user User) GetUser() (User, []MenuTest, error) {
// 	//DB.Raw("select * from t_sso_user_info a,t_sso_role_info b,t_sso_user_role c where a.user_id = c.user_id and " +
// 	//	"b.role_id = c.role_id and a.user_id=?",user.ID).Scan(&user)
// 	DB.First(&user)
// 	var roles []Role
// 	var menus []MenuTest
// 	err := DB.Model(&user).Related(&roles, "Roles").Model(&roles).Related(&menus, "Menus").Error
// 	//for i :=0;i<len(roles);i++{
// 	//	roles[i].Menus=menus
// 	//}
// 	//理解切片
// 	//for index :=range roles{
// 	//	roles[index].Menus=menus
// 	//}
// 	user.Roles = roles
// 	return user, menus, err
// }
