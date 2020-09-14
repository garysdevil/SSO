package model

import "time"

type CommonModel struct {
	// ID        uint `gorm:"autoIncrement"`
	CreatedAt time.Time  `swaggerignore:"true"`
	UpdatedAt time.Time  `swaggerignore:"true"`
	DeletedAt *time.Time `swaggerignore:"true"`
}

type System struct {
	CommonModel
	SystemID   string `json:"system_id" gorm:"column:system_id;primaryKey;type:varchar(255)"`
	SystemName string `json:"system_name" gorm:"column:system_name"`
	API        []API  `json:"apis" gorm:"foreignKey:SystemID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Menu       []Menu `json:"menus" gorm:"foreignKey:SystemID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
type API struct {
	CommonModel
	APIID    string `json:"api_id" gorm:"column:api_id;primaryKey;type:varchar(255)"`
	APIName  string `json:"api_name" gorm:"column:api_name"`
	Path     string `json:"path" gorm:"column:path"`
	Action   string `json:"action" gorm:"column:action"`
	SystemID string `json:"system_id" gorm:"column:system_id;type:varchar(255);"`
}
type Menu struct {
	CommonModel
	MenuID   string `json:"menu_id" gorm:"column:menu_id;primaryKey;type:varchar(255)"`
	MenuName string `json:"menu_name" gorm:"column:menu_name"`
	ParentID string `json:"parent_id" gorm:"column:parent_id;type:varchar(255)"`
	SystemID string `json:"system_id" gorm:"column:system_id;type:varchar(255);"`
}

type Role struct {
	CommonModel
	RoleID   string `json:"role_id" gorm:"column:role_id;primary_key;type:varchar(255)" swaggerignore:"true"`
	RoleName string `json:"roleName" gorm:"role_name"`
	Menus    []Menu `gorm:"many2many:role_menu;JOINTABLE_FOREIGNKEY:role_id;ASSOCIATION_JOINTABLE_FOREIGNKEY:menu_id" swaggerignore:"true"`
	API      []API  `gorm:"many2many:role_api;JOINTABLE_FOREIGNKEY:role_id;ASSOCIATION_JOINTABLE_FOREIGNKEY:api_id" swaggerignore:"true"`
}

type Group struct {
	CommonModel
	GroupID   string `json:"group_id" gorm:"column:group_id;primary_key;type:varchar(255)"`
	Groupname string `json:"name" gorm:"column:groupname"`
	Role      []Role `json:"group_role" gorm:"many2many:group_role;JOINTABLE_FOREIGNKEY:group_id;ASSOCIATION_JOINTABLE_FOREIGNKEY:role_id"`
}

type User struct {
	CommonModel
	UserID    string  `json:"user_id" gorm:"column:user_id;primary_key;type:varchar(255)" swaggerignore:"true"`
	Comment   string  `json:"comment" gorm:"column:comment"`
	Username  string  `json:"username" gorm:"column:username"`
	Password  string  `json:"password" gorm:"column:password"`
	CreatedBy uint    `json:"created_by" gorm:"column:created_by"`
	Group     []Group `json:"user_group" gorm:"many2many:user_group;JOINTABLE_FOREIGNKEY:user_id;ASSOCIATION_JOINTABLE_FOREIGNKEY:group_id"`
	Roles     []Role  `json:"-" gorm:"-"`
}
