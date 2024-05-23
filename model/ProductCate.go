package model

type ProductCate struct {
	Id          int    `gorm:"column:id" json:"id"`
	Title       string `gorm:"column:title" json:"title"`
	CateImg     string `gorm:"column:cate_img" json:"cate_img"`
	Link        string `gorm:"column:link" json:"link"`
	Template    string `gorm:"column:template" json:"template"`
	Pid         int    `gorm:"column:pid" json:"pid"`
	SubTitle    string `gorm:"column:sub_title" json:"sub_title"`
	Keywords    string `gorm:"column:keywords" json:"keywords"`
	Description string `gorm:"column:description" json:"description"`
	Sort        int    `gorm:"column:sort" json:"sort"`
	Status      int8   `gorm:"column:status" json:"status"`
	AddTime     int    `gorm:"column:add_time" json:"add_time"`
}

func (ProductCate) TableName() string {
	return "product_cate"
}
