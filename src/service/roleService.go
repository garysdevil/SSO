package service

import (
	"sso/src/model"

	"github.com/rs/xid"
	log "github.com/sirupsen/logrus"
)

func CreateRole(role model.Role) error {

	role.RoleID = xid.New().String()

	err := role.CreateRole()

	if err == nil {
		log.Info("角色：" + role.RoleName + "创建成功")
	}

	return err
}

// func UpdateRole(role model.Role) error {
// 	err := role.UpdateRole()
// 	if err != nil {
// 		log.Info("角色：" + role.RoleID + "更新成功")
// 	}
// 	return err
// }

// func DeleteRole(id string) error {
// 	var role model.Role
// 	role.RoleId = id
// 	err := role.DeleteRole()
// 	if err != nil {
// 		log.Error(err)
// 		return err
// 	}
// 	log.Info("角色：" + role.RoleId + "删除成功！")
// 	return nil
// }

// func ListRole(pageNo, pageSize int) *schema.RoleList {
// 	var role model.Role
// 	roleList, total := role.ListRole(pageNo, pageSize)
// 	users := &schema.RoleList{
// 		RoleList: roleList,
// 		Total:    total,
// 	}
// 	return users
// }

// ////角色菜单关联代码
// //func RoleMenuRelate(role model.Role)  {
// //	role.RoleMenuRelate()
// //}

// func GetRole(id string) (model.Role, error) {
// 	var role model.Role
// 	role.RoleId = id
// 	role, err := role.GetRole()
// 	if err != nil {
// 		log.Error(err)
// 		return role, err
// 	}
// 	return role, nil
// }
