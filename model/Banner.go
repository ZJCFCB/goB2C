package model

//焦点图表

type Banner struct {
	Id         int    `gorm:"column:id" json:"id"`
	Title      string `gorm:"column:title" json:"title"`
	BannerType int    `gorm:"column:banner_type" json:"banner_type"`
	BannerImg  string `gorm:"column:banner_img" json:"banner_img"`
	Link       string `gorm:"column:link" json:"link"`
	Sort       int    `gorm:"column:sort" json:"sort"`
	Status     int    `gorm:"column:status" json:"status"`
	AddTime    int64  `gorm:"column:add_time" json:"add_time"`
}

func (Banner) TableName() string {
	return "banner"
}
