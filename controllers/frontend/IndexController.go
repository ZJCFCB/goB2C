package frontend

import (
	"fmt"
	"goB2C/dao"
	"goB2C/model"
	"time"
)

//首页控制器

type IndexController struct {
	BaseController
}

func (I *IndexController) Get() {
	I.BaseInit()
	startTime := time.Now().UnixNano()

	banner := []model.Banner{}
	if hasBanner := dao.RedisGet("banner", &banner); hasBanner == true {
		I.Ctx.Set("topMenuList", banner)
	} else {
		dao.DB.Where("status=1 AND banner_type=1").Order("sort desc").Find(&banner)
		I.Ctx.Set("bannerList", banner)
		dao.RedisSet("banner", banner)
	}

	//获取手机商品列表
	redisPhone := []model.Product{}
	if hasPhone := dao.RedisGet("phone", &redisPhone); hasPhone == true {
		I.Ctx.Set("phoneList", redisPhone)
	} else {
		phone := model.GetProductByCategory(1, "hot", 8)
		I.Ctx.Set("phoneList", phone)
		dao.RedisSet("phone", phone)
	}

	//获取电视商品列表
	redisTv := []model.Product{}
	if hasTv := dao.RedisGet("tv", &redisTv); hasTv == true {
		I.Ctx.Set("tvList", redisTv)
	} else {
		tv := model.GetProductByCategory(4, "best", 8)
		I.Ctx.Set("tvList", tv)
		dao.RedisSet("tv", tv)
	}

	//结束时间
	endTime := time.Now().UnixNano()

	fmt.Println("执行时间", endTime-startTime)

	//c.TplName = "frontend/index/index.html"
}
