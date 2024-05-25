package model

type Menu struct {
	Id          int       `gorm:"column:id" json:"id"`
	Title       string    `gorm:"column:title" json:"title"`
	Link        string    `gorm:"column:link" json:"link"`
	Position    int       `gorm:"column:position" json:"position"`
	IsOpennew   int       `gorm:"column:is_opennew" json:"is_opennew"`
	Relation    string    `gorm:"column:relation" json:"relation"`
	Sort        int       `gorm:"column:sort" json:"sort"`
	Status      int       `gorm:"column:status" json:"status"`
	AddTime     int       `gorm:"column:add_time" json:"add_time"`
	ProductItem []Product `gorm:"-"` // 用来存相关联的产品
}

func (Menu) TableName() string {
	return "menu"
}
