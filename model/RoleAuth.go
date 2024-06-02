package model

type RoleAuth struct {
	AuthId int `gorm:"column:auth_id" json:"auth_id"`
	RoleId int `gorm:"column:role_id" json:"role_id"`
}

func (RoleAuth) TableName() string {
	return "role_auth"
}
