package frontend

import (
	"fmt"
	"goB2C/dao"
	"goB2C/model"
	"time"

	"github.com/gin-gonic/gin"
)

// 首页控制器

type IndexController struct {
	BaseController
}

func (I *IndexController) MainPage(Ctx *gin.Context) {
	I.BaseInit(Ctx)
	startTime := time.Now().UnixNano()

	banner := []model.Banner{}
	if hasBanner := dao.RedisGet("banner", &banner); hasBanner == true {
		Ctx.Set("bannerList", banner)
	} else {
		dao.DB.Where("status=1 AND banner_type=1").Order("sort desc").Find(&banner)
		Ctx.Set("bannerList", banner)
		dao.RedisSet("banner", banner)
	}

	//获取手机商品列表
	redisPhone := []model.Product{}
	if hasPhone := dao.RedisGet("phone", &redisPhone); hasPhone == true {
		Ctx.Set("phoneList", redisPhone)
	} else {
		phone := model.GetProductByCategory(1, "hot", 8)
		Ctx.Set("phoneList", phone)
		dao.RedisSet("phone", phone)
	}

	//获取电脑商品列表
	redisTv := []model.Product{}
	if hasTv := dao.RedisGet("tv", &redisTv); hasTv == true {
		Ctx.Set("tvList", redisTv)
	} else {
		tv := model.GetProductByCategory(2, "best", 8)
		Ctx.Set("tvList", tv)
		dao.RedisSet("tv", tv)
	}

	//结束时间
	endTime := time.Now().UnixNano()

	fmt.Println("执行时间", endTime-startTime)

	middleMenuList, _ := Ctx.Get("middleMenuList")
	productCateList, _ := Ctx.Get("productCateList")
	bannerList, _ := Ctx.Get("bannerList")
	phoneList, _ := Ctx.Get("phoneList")
	tvList, _ := Ctx.Get("tvList")
	userinfo, _ := Ctx.Get("userinfo")
	topMenuList, _ := Ctx.Get("topMenuList")
	Ctx.HTML(200, "index.html", gin.H{
		"middleMenuList":  middleMenuList,
		"productCateList": productCateList,
		"bannerList":      bannerList,
		"phoneList":       phoneList,
		"tvList":          tvList,
		"userinfo":        userinfo,
		"topMenuList":     topMenuList,
		"TestUse":         "TestUse",
	})
}
