package service

import "goB2C/dao"

type RoleAuth struct {
	AuthId int `gorm:"column:auth_id`
	RoleId int `gorm:"column:role_id`
}

func (R RoleAuth) TableName() string {
	return "role_auth"
}

func GetByAuthId(id int) RoleAuth {
	var x RoleAuth
	dao.DB.Model(&RoleAuth{}).Where("auth_id = ?", id).First(&x)
	return x
}
