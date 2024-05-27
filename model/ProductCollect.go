package model

type ProductCollect struct {
	Id        int    `json:"id" gorm:"column:id;primary_key;auto_increment"`
	UserId    int    `json:"user_id" gorm:"column:user_id"`
	ProductId int    `json:"product_id" gorm:"column:product_id"`
	AddTime   string `json:"add_time" gorm:"column:add_time"`
}

func (ProductCollect) TableName() string {
	return "product_collect"
}
