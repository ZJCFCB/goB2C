package model

type User struct {
	ID       int    `gorm:"column:id" json:"id"`
	Phone    string `gorm:"column:phone" json:"phone"`
	Password string `gorm:"column:password" json:"password"`
	AddTime  int    `gorm:"column:add_time" json:"add_time"`
	LastIP   string `gorm:"column:last_ip" json:"last_ip"`
	Email    string `gorm:"column:email" json:"email"`
	Status   int8   `gorm:"column:status" json:"status"`
}
