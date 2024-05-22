package model

// 购物车模型

type Cart struct {
	Id             int     `gorm:"column:id" json:"id"`
	Title          string  `gorm:"column:title" json:"title"`
	Price          float64 `gorm:"column:price" json:"price"`
	GoodsVersion   string  `gorm:"column:goods_version" json:"goods_version"`
	Num            int     `gorm:"column:num" json:"num"`
	ProductGift    string  `gorm:"column:product_gift" json:"product_gift"`
	ProductFitting string  `gorm:"column:product_fitting" json:"product_fitting"`
	ProductColor   string  `gorm:"column:product_color" json:"product_color"`
	ProductImg     string  `gorm:"column:product_img" json:"product_img"`
	ProductAttr    string  `gorm:"column:product_attr" json:"product_attr"`
}

func (c Cart) TableName() string {
	return "cart"
}

// 判断购物车里面有没有当前数据
func CartHasData(cartList []Cart, currentData Cart) bool {
	for i := 0; i < len(cartList); i++ {
		if cartList[i].Id == currentData.Id &&
			cartList[i].ProductColor == currentData.ProductColor &&
			cartList[i].ProductAttr == currentData.ProductAttr {
			return true
		}
	}
	return false
}
