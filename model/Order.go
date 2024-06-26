package model

type Order struct {
	Id          int         `gorm:"column:id" json:"id"`
	OrderId     string      `gorm:"column:order_id" json:"order_id"`
	Uid         int         `gorm:"column:uid" json:"uid"`
	AllPrice    float64     `gorm:"column:all_price" json:"all_price"`
	Phone       string      `gorm:"column:phone" json:"phone"`
	Name        string      `gorm:"column:name" json:"name"`
	Address     string      `gorm:"column:address" json:"address"`
	Zipcode     string      `gorm:"column:zipcode" json:"zipcode"`
	PayStatus   int         `gorm:"column:pay_status" json:"pay_status"`     // 支付状态： 0 表示未支付     1 已经支付
	PayType     int         `gorm:"column:pay_type" json:"pay_type"`         // 支付类型： 0 alipay    1 wechat
	OrderStatus int         `gorm:"column:order_status" json:"order_status"` // 订单状态： 0 已下单  1 已付款  2 已配货  3、发货   4、交易成功   5、退货     6、取消
	AddTime     int         `gorm:"column:add_time" json:"add_time"`
	OrderItems  []OrderItem `gorm:"foreignkey:order_id;association_foreignkey:Id"`
}

func (Order) TableName() string {
	return "order"
}
