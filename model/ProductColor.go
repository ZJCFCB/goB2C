package model

type ProductColor struct {
	Id         int    `gorm:"column:id" json:"id"`
	ColorName  string `gorm:"column:color_name" json:"color_name"`
	ColorValue string `gorm:"column:color_value" json:"color_value"`
	Status     int8   `gorm:"column:status" json:"status"`
	Checked    bool   `gorm:"column:checked" json:"checked"`
}

func (ProductColor) TableName() string {
	return "product_color"
}
