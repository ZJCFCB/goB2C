package model

type RoleAuth struct {
	AuthID int `gorm:"column:auth_id;primary_key" json:"auth_id"`
	RoleID int `gorm:"column:role_id" json:"role_id"`
}
