package model

type UserInfo struct {
	ID       int    `gorm:"column:id" json:"id"`
	UserID   int    `gorm:"column:user_id" json:"user_id"`
	Gender   int    `gorm:"column:gender" json:"gender"`
	Avatar   string `gorm:"column:avatar" json:"avatar"`
	Phone    string `gorm:"column:phone" json:"phone"`
	Username string `gorm:"column:username" json:"username"`
	Email    string `gorm:"column:email" json:"email"`
	AddTime  int    `gorm:"column:add_time" json:"add_time"`
}

// 设置表名
func (UserInfo) TableName() string {
	return "user_info"
}
