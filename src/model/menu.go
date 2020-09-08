package model

// //菜单增删改查
// type Menu struct {
// 	MenuId    string `json:"menuId" gorm:"column:menu_id;primary_key:true"`
// 	MenuName  string `json:"menu_name" gorm:"column:menu_name"`
// 	Pid       string `json:"pid" gorm:"column:pid"`
// 	DeletedAt string `json:"-" gorm:"column:is_delete;DEFAULT:null"`
// }

// //级联菜单测试
// type MenuTest struct {
// 	MenuId    string     `json:"menuId" gorm:"column:menu_id;primary_key:true"`
// 	MenuName  string     `json:"menu_name" gorm:"column:menu_name"`
// 	Pid       string     `json:"pid" gorm:"column:pid"`
// 	Children  []MenuTest `json:"children" gorm:"-"`
// 	DeletedAt string     `json:"-" gorm:"column:is_delete;DEFAULT:null"`
// }

// func (MenuTest) TableName() string {
// 	return "t_sso_menu_info"
// }

// func (Menu) TableName() string {
// 	return "t_sso_menu_info"
// }

// func (menu *Menu) CreateMenu() error {
// 	return DB.Create(&menu).Error
// }

// func (menu *Menu) DeleteMenu() error {
// 	return DB.Where("menu_id = ?", menu.MenuId).Delete(&menu).Error
// }

// func (menu *Menu) UpdateMenu() error {
// 	return DB.Omit("is_delete").Save(&menu).Error
// }

// func (menu *Menu) ListMenu(pageNo, pageSize int) ([]Menu, int32) {
// 	var count int32
// 	DB.Model(&menu).Count(&count)
// 	var menuList []Menu
// 	DB.Model(&menu).Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(&menuList)
// 	return menuList, count
// }
