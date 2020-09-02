package model

type Role struct {
	//ID	uint	`gorm:"-"`
	RoleId    string `json:"roleId" gorm:"column:role_id;primary_key:true"`
	RoleName  string `json:"roleName" gorm:"role_name"`
	DeletedAt string `json:"-" gorm:"column:is_delete;DEFAULT:null"`
	Menus     []Menu `gorm:"many2many:t_sso_role_menu;JOINTABLE_FOREIGNKEY:role_id;ASSOCIATION_JOINTABLE_FOREIGNKEY:menu_id"`
}

func (Role) TableName() string {
	return "t_sso_role_info"
}

//角色增删改查
func (role *Role) CreateRole() error {
	return DB.Set("gorm:association_autocreate", false).Set("gorm:association_autoupdate", false).
		Model(&role).Create(&role).Error
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
	//DB.Model(&role).Save(&role)
	role.Menus = menus
	return DB.Set("gorm:association_autocreate", false).Set("gorm:association_autoupdate", false).
		Model(&role).Update("role_name", role.RoleName).Error
}

func (role *Role) DeleteRole() error {
	DB.Model(&role).Association("Menus").Clear()
	return DB.Where("role_id = ?", role.RoleId).Delete(&role).Error
}

func (role *Role) ListRole(pageNo, pageSize int) ([]Role, int32) {
	var count int32
	DB.Model(&role).Count(&count)
	var roleList []Role
	DB.Model(&role).Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(&roleList)
	return roleList, count
}

////角色菜单关联
//func (role *Role) RoleMenuRelate()  {
//	menus :=role.Menus
//	DB.Debug().Preload("Menus").First(&role)
//	for _,menu :=range menus{
//		//fmt.Println(roleId)
//		DB.Model(&role).Association("Roles").Append(Menu{MenuId:menu.MenuId})
//	}
//}

func (role Role) GetRole() (Role, error) {
	DB.First(&role)
	var menus []Menu
	err := DB.Model(&role).Related(&menus, "Menus").Error
	role.Menus = menus
	return role, err
}
