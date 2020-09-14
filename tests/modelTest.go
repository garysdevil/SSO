package main

import (
	"fmt"

	"sso/src/model"

	"github.com/rs/xid"
)

func modelTest() {

	// test4V2Batch()
	// test3V2AutoMigrate()

	// test2FirstUser()
	test1CreateTable()

	fmt.Println("\nEnd modelTest()")
}
func test4V2Batch() {
	type Test struct {
		Name string
		model.CommonModel
	}
	// batch
	model.DB.AutoMigrate(&Test{})
	test := []Test{
		{Name: "fdsf"},
		{Name: "fdsfaaa"},
	}
	model.DB.Create(&test)

	// drop table
	model.DB.Migrator().DropTable(&Test{})

}
func test3V2AutoMigrate() {
	type Test struct {
		Name string
		model.CommonModel
	}

	// new grammer
	type CreditCard struct {
		Number    string `gorm:"primaryKey"`
		UserRefer string `gorm:"type:varchar(255);not null"`
	}
	type User struct {
		AA          string       `gorm:"primaryKey;type:varchar(255);not null"`
		CreditCards []CreditCard `gorm:"foreignKey:UserRefer"`
	}

	model.DB.Debug().AutoMigrate(&User{})
	model.DB.Debug().AutoMigrate(&CreditCard{})

	// drop table
	model.DB.Migrator().DropTable(&CreditCard{})
	model.DB.Migrator().DropTable(&User{})
}

func test2FirstUser() {
	user := model.User{
		UserID:    xid.New().String(),
		Username:  "ldap",
		CreatedBy: 0,
	}
	model.DB.Create(&user)

}

func test1CreateTable() {

	model.DB.AutoMigrate(&model.System{})
	model.DB.AutoMigrate(&model.API{}) //AddForeignKey("system_id", "systems(system_id)", "SET NULL", "CASCADE")  v1
	model.DB.AutoMigrate(&model.Menu{})

	model.DB.AutoMigrate(&model.Role{})
	model.DB.AutoMigrate(&model.User{})
	model.DB.AutoMigrate(&model.Group{})

	// model.DB.Migrator().DropTable(&model.System{}, &model.API{}, &model.Menu{}, &model.Role{}, &model.User{}, &model.Group{}, "role_apis", "role_menus", "user_groups", "group_roles")

}
