package model

type Order struct {
	Id          int         `gorm:"column:id" json:"id"`
	OrderID     string      `gorm:"column:order_id" json:"order_id"`
	Uid         int         `gorm:"column:uid" json:"uid"`
	AllPrice    float64     `gorm:"column:all_price" json:"all_price"`
	Phone       string      `gorm:"column:phone" json:"phone"`
	Name        string      `gorm:"column:name" json:"name"`
	Address     string      `gorm:"column:address" json:"address"`
	Zipcode     string      `gorm:"column:zipcode" json:"zipcode"`
	PayStatus   int8        `gorm:"column:pay_status" json:"pay_status"`
	PayType     int8        `gorm:"column:pay_type" json:"pay_type"`
	OrderStatus int8        `gorm:"column:order_status" json:"order_status"`
	AddTime     int         `gorm:"column:add_time" json:"add_time"`
	OrderItem   []OrderItem `gorm:"foreignkey:OrderId;association_foreignkey:Id"`
}

func (Order) TableName() string {
	return "order"
}
