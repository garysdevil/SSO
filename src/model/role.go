package model

// 重新定义表名
// func (Role) TableName() string {
// 	return "t_sso_role_info"
// }

//角色增删改查
func (role *Role) CreateRole() error {
	return DB.Model(&role).Create(&role).Error
}

func (role *Role) UpdateRole() error {
	//menus :=role.Menus
	//if menus !=nil{
	//	DB.Debug().Model(&role).Association("Roles").Clear()
	//	for _,menu :=range menus{
	//		fmt.Println(menu.MenuId)
	//		DB.Debug().Model(&role).Association("Roles").Replace(Menu{MenuId:menu.MenuId})
	//	}
	//}
	menus := role.Menus
	DB.Model(&role).Association("Menus").Clear()
	DB.Model(&role).Association("APIs").Clear()
	// DB.Model(&role).Save(&role)
	role.Menus = menus
	return DB.Set("gorm:association_autocreate", false).Set("gorm:association_autoupdate", false).
		Model(&role).Update("role_name", role.RoleName).Error
}

// func (role *Role) DeleteRole() error {
// 	DB.Model(&role).Association("Menus").Clear()
// 	return DB.Where("role_id = ?", role.RoleId).Delete(&role).Error
// }

// func (role *Role) ListRole(pageNo, pageSize int) ([]Role, int32) {
// 	var count int32
// 	DB.Model(&role).Count(&count)
// 	var roleList []Role
// 	DB.Model(&role).Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(&roleList)
// 	return roleList, count
// }

// ////角色菜单关联
// //func (role *Role) RoleMenuRelate()  {
// //	menus :=role.Menus
// //	DB.Debug().Preload("Menus").First(&role)
// //	for _,menu :=range menus{
// //		//fmt.Println(roleId)
// //		DB.Model(&role).Association("Roles").Append(Menu{MenuId:menu.MenuId})
// //	}
// //}

// func (role Role) GetRole() (Role, error) {
// 	DB.First(&role)
// 	var menus []Menu
// 	err := DB.Model(&role).Related(&menus, "Menus").Error
// 	role.Menus = menus
// 	return role, err
// }
