package model

//系统管理员列表

type Administrator struct {
	Id       int    `gorm:"column:id" json:"id"`
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
	Mobile   string `gorm:"column:mobile" json:"mobile"`
	Email    string `gorm:"column:email" json:"email"`
	Status   int8   `gorm:"column:status" json:"status"`
	RoleID   int    `gorm:"column:role_id" json:"role_id"`
	AddTime  int    `gorm:"column:add_time" json:"add_time"`
	IsSuper  int8   `gorm:"column:is_super" json:"is_super"`
}

func (Administrator) TableName() string {
	return "administrator"
}
