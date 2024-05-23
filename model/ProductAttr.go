package model

type ProductAttr struct {
	Id              int    `gorm:"column:id" json:"id"`
	ProductID       int    `gorm:"column:product_id" json:"product_id"`
	AttributeCateID int    `gorm:"column:attribute_cate_id" json:"attribute_cate_id"`
	AttributeID     int    `gorm:"column:attribute_id" json:"attribute_id"`
	AttributeTitle  string `gorm:"column:attribute_title" json:"attribute_title"`
	AttributeType   int    `gorm:"column:attribute_type" json:"attribute_type"`
	AttributeValue  string `gorm:"column:attribute_value" json:"attribute_value"`
	Sort            int    `gorm:"column:sort" json:"sort"`
	AddTime         int    `gorm:"column:add_time" json:"add_time"`
	Status          int8   `gorm:"column:status" json:"status"`
}

func (ProductAttr) TableName() string {
	return "product_attr"
}
