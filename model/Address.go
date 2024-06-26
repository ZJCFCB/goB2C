package model

//用户地址模型

type Address struct {
	Id             int    `gorm:"column:id" json:"id"`
	Uid            int    `gorm:"column:uid" json:"uid"`
	Phone          string `gorm:"column:phone" json:"phone"`
	Name           string `gorm:"column:name" json:"name"`
	Zipcode        string `gorm:"column:zipcode" json:"zipcode"`
	Address        string `gorm:"column:address" json:"address"`
	DefaultAddress int    `gorm:"column:default_address" json:"default_address"`
	AddTime        int    `gorm:"column:add_time" json:"add_time"`
}

func (Address) TableName() string {
	return "address"
}
