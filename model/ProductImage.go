package model

type ProductImage struct {
	Id        int    `gorm:"column:id" json:"id"`
	ProductID int    `gorm:"column:product_id" json:"product_id"`
	ImgUrl    string `gorm:"column:img_url" json:"img_url"`
	ColorID   int    `gorm:"column:color_id" json:"color_id"`
	Sort      int    `gorm:"column:sort" json:"sort"`
	AddTime   int    `gorm:"column:add_time" json:"add_time"`
	Status    int8   `gorm:"column:status" json:"status"`
}

func (ProductImage) TableName() string {
	return "product_image"
}
