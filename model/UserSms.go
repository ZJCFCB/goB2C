package model

type UserSms struct {
	ID        int    `gorm:"column:id" json:"id"`
	IP        string `gorm:"column:ip" json:"ip"`
	Phone     string `gorm:"column:phone" json:"phone"`
	SendCount int    `gorm:"column:send_count" json:"send_count"`
	AddDay    string `gorm:"column:add_day" json:"add_day"`
	AddTime   int    `gorm:"column:add_time" json:"add_time"`
	Sign      string `gorm:"column:sign" json:"sign"`
}
