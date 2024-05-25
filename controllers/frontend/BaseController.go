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
}

func (B *BaseController) BaseInit(Ctx *gin.Context) {

	//顶部分类
	topMenu := []model.Menu{}
	if hasTopMenu := dao.RedisGet("topMenu", &topMenu); hasTopMenu == true {
		Ctx.Set("topMenuList", topMenu)
	} else {
		dao.DB.Where("status=1 AND position=1").Order("sort desc").Find(&topMenu)
		Ctx.Set("topMenuList", topMenu)
		dao.RedisSet("topMenu", topMenu)
	}

	//Todo
	//获取轮播图，这里还有bug，一时半会不知道怎么改

	//左侧分类（预加载）
	productCate := []model.ProductCate{}

	if hasproductCate := dao.RedisGet("productCate", &productCate); hasproductCate == true {
		Ctx.Set("productCateList", productCate)
	} else {
		//注意这里的查询语句，预加载的相关知识
		dao.DB.Preload("ProductCateItem",
			func(db *gorm.DB) *gorm.DB {
				return db.Where("product_cate.status=1").
					Order("product_cate.sort DESC")
			}).Where("pid=0 AND status=1").Order("sort desc").
			Find(&productCate)
		Ctx.Set("productCateList", productCate)
		dao.RedisSet("productCate", productCate)
	}

	//获取中间导航的数据
	middleMenu := []model.Menu{}

	if hasMiddleMenu := dao.RedisGet("middleMenu", &middleMenu); hasMiddleMenu == true {
		Ctx.Set("middleMenuList", middleMenu)
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
		Ctx.Set("middleMenuList", middleMenu)
		dao.RedisSet("middleMenu", middleMenu)
	}
	//判断用户是否登录

	//在top list中，如果判断用户是已经登录状态（cookie），那么显示的就是登录的前端页面
	//如果判断是非登录状态，显示的就是登录、注册页面

	user := model.User{}
	model.Cookie.Get(Ctx, "userinfo", &user)
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
		Ctx.Set("userinfo", str)
	} else {
		str := fmt.Sprintf(`<ul>
			<li><a href="/auth/login" target="_blank">登录</a></li>
			<li>|</li>
			<li><a href="/auth/registerStep1" target="_blank" >注册</a></li>
		</ul>`)

		Ctx.Set("userinfo", str)
	}
	//把path也放进去
	urlpath, _ := url.Parse(Ctx.Copy().Request.URL.String())
	Ctx.Set("pathname", urlpath)
}
