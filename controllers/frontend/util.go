package frontend

import (
	"fmt"
	"goB2C/dao"
	"goB2C/model"
	"strings"

	"gorm.io/gorm"
)

func GetTopMenuList() []model.Menu {
	topMenu := []model.Menu{}
	if hasTopMenu := dao.RedisGet("topMenu", &topMenu); hasTopMenu == true {
		return topMenu
	} else {
		dao.DB.Where("status=1 AND position=1").Order("sort desc").Find(&topMenu)
		dao.RedisSet("topMenu", topMenu)
		return topMenu
	}
}
func GetProductCateList() []model.ProductCate {
	productCate := []model.ProductCate{}

	if hasproductCate := dao.RedisGet("productCate", &productCate); hasproductCate == true {
		return productCate
	} else {
		//注意这里的查询语句，预加载的相关知识
		dao.DB.Preload("ProductCateItem",
			func(db *gorm.DB) *gorm.DB {
				return db.Where("product_cate.status=1").
					Order("product_cate.sort DESC")
			}).Where("pid=0 AND status=1").Order("sort desc").
			Find(&productCate)

		dao.RedisSet("productCate", productCate)
		return productCate
	}
}

func GetMiddleList() []model.Menu {
	middleMenu := []model.Menu{}

	if hasMiddleMenu := dao.RedisGet("middleMenu", &middleMenu); hasMiddleMenu == true {
		return middleMenu
	} else {
		dao.DB.Where("status=1 AND position=2").Order("sort desc").Find(&middleMenu)
		for i := 0; i < len(middleMenu); i++ {
			//获取关联产品
			middleMenu[i].Relation = strings.ReplaceAll(middleMenu[i].Relation, "，", ",")
			relation := strings.Split(middleMenu[i].Relation, ",")
			product := []model.Product{}
			dao.DB.Where("id in (?)", relation).Limit(6).Order("sort ASC").
				Select("id,title,product_img,price").Find(&product)
			middleMenu[i].ProductItem = product
		}
		dao.RedisSet("middleMenu", middleMenu)
		return middleMenu
	}
}

func GetBanner() []model.Banner {
	banner := []model.Banner{}
	if hasBanner := dao.RedisGet("banner", &banner); hasBanner == true {
		return banner
	} else {
		dao.DB.Where("status=1 AND banner_type=1").Order("sort desc").Find(&banner)

		dao.RedisSet("banner", banner)
		return banner
	}

}

func GetUserInfo(user model.User) string {
	var str string
	if len(user.Phone) == 11 {
		str = fmt.Sprintf(`<ul>
			<li class="userinfo">
				<a href="#">%v</a>

				<i class="i"></i>
				<ol>
					<li><a href="/user">个人中心</a></li>

					<li><a href="/coll">我的收藏</a></li>

					<li><a href="/auth/loginOut">退出登录</a></li>
				</ol>

			</li>
		</ul> `, user.Phone)
	} else {
		str = fmt.Sprintf(`<ul>
			<li><a href="/auth/login" target="_blank">登录</a></li>
			<li>|</li>
			<li><a href="/auth/registerStep1" target="_blank" >注册</a></li>
		</ul>`)
	}
	return str
}
