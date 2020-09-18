package cmd

import (
	"fmt"
	"sso/src/model"
)

func AutoMigrateTable() {

	model.DB.AutoMigrate(&model.System{})
	model.DB.AutoMigrate(&model.API{})
	model.DB.AutoMigrate(&model.Menu{})

	model.DB.AutoMigrate(&model.Role{})
	model.DB.AutoMigrate(&model.Group{})
	model.DB.AutoMigrate(&model.User{})

	// model.DB.Migrator().DropTable(&model.System{}, &model.API{}, &model.Menu{}, &model.Role{}, &model.User{}, &model.Group{}, "role_apis", "role_menus", "user_groups", "group_roles")
	fmt.Println("初始化数据库表格成功")
}
