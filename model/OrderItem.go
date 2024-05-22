package model

type OrderItem struct {
	ID           int     `gorm:"column:id" json:"id"`
	OrderID      int     `gorm:"column:order_id" json:"order_id"`
	UID          int     `gorm:"column:uid" json:"uid"`
	ProductTitle string  `gorm:"column:product_title" json:"product_title"`
	ProductID    int     `gorm:"column:product_id" json:"product_id"`
	ProductImg   string  `gorm:"column:product_img" json:"product_img"`
	ProductPrice float64 `gorm:"column:product_price" json:"product_price"`
	ProductNum   int     `gorm:"column:product_num" json:"product_num"`
	GoodsVersion string  `gorm:"column:goods_version" json:"goods_version"`
	GoodsColor   string  `gorm:"column:goods_color" json:"goods_color"`
	AddTime      int     `gorm:"column:add_time" json:"add_time"`
}
