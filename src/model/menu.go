package model

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
// 	// return DB.Omit("is_delete").Save(&menu).Error
// 	return DB.Save(&menu).Error
// }

// // 分页获取menu
// func (menu *Menu) ListMenu(pageNo, pageSize int) ([]Menu, int32) {
// 	var count int32
// 	DB.Model(&menu).Count(&count)
// 	DB.Table()
// 	var menuList []Menu
// 	DB.Model(&menu).Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(&menuList)
// 	return menuList, count
// }
