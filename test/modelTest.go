package main

import (
	"fmt"

	"sso/cmd"
	"sso/src/model"

	"github.com/jinzhu/gorm"
)

func main() {
	if err := cmd.InitConfig("../config/settings.dev.yaml"); err != nil {
		panic(err)
	}
	model.InitDB()
	test1()

	fmt.Println("end")
}

func test() {
	user := model.User{}
	user.GetTokenInfo()
}

//菜单增删改查
type MenuAbc struct {
	MenuId    string `json:"menuId" gorm:"column:menu_id;primary_key:true"`
	MenuName  string `json:"menu_name" gorm:"column:menu_name"`
	Pid       string `json:"pid" gorm:"column:pid"`
	DeletedAt string `json:"-" gorm:"column:is_delete;DEFAULT:null"`
}
type RoleAbc struct {
	//ID	uint	`gorm:"-"`
	RoleId    string    `json:"roleId" gorm:"column:role_id;primary_key:true"`
	RoleName  string    `json:"roleName" gorm:"role_name"`
	DeletedAt string    `json:"-" gorm:"column:is_delete;DEFAULT:null"`
	Menus     []MenuAbc `gorm:"many2many:t_sso_role_menuaaa;JOINTABLE_FOREIGNKEY:role_id;ASSOCIATION_JOINTABLE_FOREIGNKEY:menu_id"`
}
type UserABC struct {
	UserId    string    `json:"userId,omitempty" gorm:"column:user_id;primary_key:true"`
	Username  string    `json:"username" gorm:"column:user_name"`
	Password  string    `json:"password" gorm:"-"`
	Status    string    `json:"status" gorm:"column:status;DEFAULT:'active'"`
	DeletedAt string    `json:"-" gorm:"column:is_delete;DEFAULT:null"`
	Roles     []RoleAbc `json:"roles" gorm:"many2many:t_sso_user_roleaaa;JOINTABLE_FOREIGNKEY:user_id;ASSOCIATION_JOINTABLE_FOREIGNKEY:role_id"`
	Menus     []MenuAbc `json:"menus" gorm:"-"`
}

type User struct {
	gorm.Model
	Name    string
	Company []Company
}

type Company struct {
	gorm.Model
	Job    string
	User   User
	UserID int
}

func test1() {
	model.DB.AutoMigrate(&User{})
	model.DB.AutoMigrate(&Company{})

	var user User
	var company []Company

	user.ID = 1

	fmt.Println("====")
	model.DB.Model(&user).Related(&company)
	model.DB.Model(&user).Association("company").Find(&company)

}
