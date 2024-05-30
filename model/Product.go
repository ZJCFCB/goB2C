package model

import "goB2C/dao"

type Product struct {
	Id              int     `gorm:"column:id" json:"id"`
	Title           string  `gorm:"column:title" json:"title"`
	SubTitle        string  `gorm:"column:sub_title" json:"sub_title"`
	ProductSn       string  `gorm:"column:product_sn" json:"product_sn"`
	CateId          int     `gorm:"column:cate_id" json:"cate_id"`
	ClickCount      int     `gorm:"column:click_count" json:"click_count"`
	ProductNumber   int     `gorm:"column:product_number" json:"product_number"`
	Price           float64 `gorm:"column:price" json:"price"`
	MarketPrice     float64 `gorm:"column:market_price" json:"market_price"`
	RelationProduct string  `gorm:"column:relation_product" json:"relation_product"`
	ProductAttr     string  `gorm:"column:product_attr" json:"product_attr"`
	ProductVersion  string  `gorm:"column:product_version" json:"product_version"`
	ProductImg      string  `gorm:"column:product_img" json:"product_img"`
	ProductGift     string  `gorm:"column:product_gift" json:"product_gift"`
	ProductFitting  string  `gorm:"column:product_fitting" json:"product_fitting"`
	ProductColor    string  `gorm:"column:product_color" json:"product_color"`
	ProductKeywords string  `gorm:"column:product_keywords" json:"product_keywords"`
	ProductDesc     string  `gorm:"column:product_desc" json:"product_desc"`
	ProductContent  string  `gorm:"column:product_content" json:"product_content"`
	IsDelete        int     `gorm:"column:is_delete" json:"is_delete"`
	IsHot           int     `gorm:"column:is_hot" json:"is_hot"`
	IsBest          int     `gorm:"column:is_best" json:"is_best"`
	IsNew           int     `gorm:"column:is_new" json:"is_new"`
	ProductTypeId   int     `gorm:"column:product_type_id" json:"product_type_id"`
	Sort            int     `gorm:"column:sort" json:"sort"`
	Status          int     `gorm:"column:status" json:"status"`
	AddTime         int     `gorm:"column:add_time" json:"add_time"`
}

func (Product) TableName() string {
	return "product"
}

func GetProductByCategory(cateId int, productType string, limitNum int) []Product {
	product := []Product{}
	productCate := []ProductCate{}
	dao.DB.Where("pid=?", cateId).Find(&productCate)
	var tempSlice []int
	if len(productCate) > 0 {
		for i := 0; i < len(productCate); i++ {
			tempSlice = append(tempSlice, productCate[i].Id)
		}
	}
	tempSlice = append(tempSlice, cateId)
	where := "cate_id in (?)"
	switch productType {
	case "hot":
		where += "AND is_hot=1"
	case "best":
		where += "AND is_best=1"
	case "new":
		where += "AND is_new=1"
	default:
		break
	}
	dao.DB.Where(where, tempSlice).Select("id,title,price,product_img,sub_title").Limit(limitNum).Order("sort desc").Find(&product)
	return product
}
