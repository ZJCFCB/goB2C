package model

type ProductTypeAttribute struct {
	Id        int    `gorm:"column:id" json:"id"`
	CateId    int    `gorm:"column:cate_id" json:"cate_id"`
	Title     string `gorm:"column:title" json:"title"`
	AttrType  int    `gorm:"column:attr_type" json:"attr_type"`
	AttrValue string `gorm:"column:attr_value" json:"attr_value"`
	Status    int    `gorm:"column:status" json:"status"`
	Sort      int    `gorm:"column:sort" json:"sort"`
	AddTime   int    `gorm:"column:add_time" json:"add_time"`
}

func (ProductTypeAttribute) TableName() string {
	return "product_type_attribute"
}
