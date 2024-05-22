package model

type ProductColor struct {
	ID         int    `gorm:"column:id" json:"id"`
	ColorName  string `gorm:"column:color_name" json:"color_name"`
	ColorValue string `gorm:"column:color_value" json:"color_value"`
	Status     int8   `gorm:"column:status" json:"status"`
	Checked    int8   `gorm:"column:checked" json:"checked"`
}
