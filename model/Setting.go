package model

type Setting struct {
	ID              int    `gorm:"column:id" json:"id"`
	SiteTitle       string `gorm:"column:site_title" json:"site_title"`
	SiteLogo        string `gorm:"column:site_logo" json:"site_logo"`
	SiteKeywords    string `gorm:"column:site_keywords" json:"site_keywords"`
	SiteDescription string `gorm:"column:site_description" json:"site_description"`
	NoPicture       string `gorm:"column:no_picture" json:"no_picture"`
	SiteIcp         string `gorm:"column:site_icp" json:"site_icp"`
	SiteTel         string `gorm:"column:site_tel" json:"site_tel"`
	SearchKeywords  string `gorm:"column:search_keywords" json:"search_keywords"`
	TongjiCode      string `gorm:"column:tongji_code" json:"tongji_code"`
	Appid           string `gorm:"column:appid" json:"appid"`
	AppSecret       string `gorm:"column:app_secret" json:"app_secret"`
	EndPoint        string `gorm:"column:end_point" json:"end_point"`
	BucketName      string `gorm:"column:bucket_name" json:"bucket_name"`
	OssStatus       int8   `gorm:"column:oss_status" json:"oss_status"`
}
