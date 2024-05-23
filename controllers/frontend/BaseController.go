package frontend

import (
	"fmt"
	"goB2C/dao"
	"goB2C/model"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BaseController struct {
	Ctx *gin.Context
}

func (B *BaseController) BaseInit() {

	//顶部分类
	topMenu := []model.Menu{}
	if hasTopMenu := dao.RedisGet("topMenu", &topMenu); hasTopMenu == true {
		B.Ctx.Set("topMenuList", topMenu)
	} else {
		dao.DB.Where("status=1 AND position=1").Order("sort desc").Find(&topMenu)
		B.Ctx.Set("topMenuList", topMenu)
		dao.RedisSet("topMenu", topMenu)
	}

	//左侧分类（预加载）
	productCate := []model.ProductCate{}

	if hasMiddleMenu := dao.RedisGet("productCate", &productCate); hasMiddleMenu == true {
		B.Ctx.Set("middleMenu", productCate)
	} else {
		//注意这里的查询语句，预加载的相关知识
		dao.DB.Preload("ProductCateItem",
			func(db *gorm.DB) *gorm.DB {
				return db.Where("product_cate.status=1").
					Order("product_cate.sort DESC")
			}).Where("pid=0 AND status=1").Order("sort desc").
			Find(&productCate)
		B.Ctx.Set("middleMenu", productCate)
		dao.RedisSet("productCate", productCate)
	}

	//获取中间导航的数据
	middleMenu := []model.Menu{}

	if hasMiddleMenu := dao.RedisGet("middleMenu", &middleMenu); hasMiddleMenu == true {
		B.Ctx.Set("middleMenu", middleMenu)
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
		B.Ctx.Set("middleMenuList", middleMenu)
		dao.RedisSet("middleMenuList", middleMenu)
	}

	//判断用户是否登录
	user := model.User{}
	model.Cookie.Get(B.Ctx, "userinfo", &user)
	if len(user.Phone) == 11 {
		str := fmt.Sprintf(`<ul>
			<li class="userinfo">
				<a href="#">%v</a>

				<i class="i"></i>
				<ol>
					<li><a href="/user">个人中心</a></li>

					<li><a href="#">我的收藏</a></li>

					<li><a href="/auth/loginOut">退出登录</a></li>
				</ol>

			</li>
		</ul> `, user.Phone)
		B.Ctx.Set("userinfo", str)
	} else {
		str := fmt.Sprintf(`<ul>
			<li><a href="/auth/login" target="_blank">登录</a></li>
			<li>|</li>
			<li><a href="/auth/registerStep1" target="_blank" >注册</a></li>
		</ul>`)

		B.Ctx.Set("userinfo", str)
	}
	//把path也放进去
	urlpath, _ := url.Parse(B.Ctx.Copy().Request.URL.String())
	B.Ctx.Set("pathname", urlpath)
}
