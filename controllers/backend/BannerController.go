package backend

import (
	"fmt"
	"goB2C/dao"
	"goB2C/model"
	"goB2C/util"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BannerController struct {
	BaseController
}

func (c *BannerController) Get(Ctx *gin.Context) {
	banner := []model.Banner{}
	dao.DB.Find(&banner)
	Ctx.HTML(200, "banner_index.html", gin.H{
		"bannerList": banner,
	})
}

func (c *BannerController) Add(Ctx *gin.Context) {
	Ctx.HTML(200, "banner_add.html", gin.H{})
}

func (c *BannerController) GoAdd(Ctx *gin.Context) {

	bannerTypeTemp := Ctx.PostForm("banner_type")
	bannerType, err1 := strconv.Atoi(bannerTypeTemp)
	title := Ctx.PostForm("title")
	link := Ctx.PostForm("link")

	sortTemp := Ctx.PostForm("sort")
	sort, err2 := strconv.Atoi(sortTemp)

	statuetemp := Ctx.PostForm("status")
	status, err3 := strconv.Atoi(statuetemp)
	if err1 != nil || err3 != nil {
		c.Error(Ctx, "非法请求", "/banner")
		return
	}
	if err2 != nil {
		c.Error(Ctx, "排序表单里面输入的数据不合法", "/banner/add")
		return
	}
	bannerImgSrc, err4 := c.UploadImg(Ctx, "banner_img")
	if err4 == nil {
		banner := model.Banner{
			Title:      title,
			BannerType: bannerType,
			BannerImg:  bannerImgSrc,
			Link:       link,
			Sort:       sort,
			Status:     status,
			AddTime:    int64(util.GetUnix()),
		}
		dao.DB.Create(&banner)
		c.Success(Ctx, "增加轮播图成功", "/banner")
	} else {
		c.Error(Ctx, "增加轮播图失败", "/banner/add")
		return
	}
}

func (c *BannerController) Edit(Ctx *gin.Context) {
	idTemp := Ctx.Query("id")
	id, err := strconv.Atoi(idTemp)
	if err != nil {
		c.Error(Ctx, "非法请求", "/banner")
		return
	}
	banner := model.Banner{Id: id}
	dao.DB.Find(&banner)
	Ctx.HTML(200, "banner_edit.html", gin.H{
		"banner": banner,
	})
}

func (c *BannerController) GoEdit(Ctx *gin.Context) {
	idTemp := Ctx.PostForm("id")
	id, err := strconv.Atoi(idTemp)
	bannerTypeTemp := Ctx.PostForm("banner_type")
	bannerType, err1 := strconv.Atoi(bannerTypeTemp)
	title := Ctx.PostForm("title")
	link := Ctx.PostForm("link")

	sortTemp := Ctx.PostForm("sort")
	sort, err2 := strconv.Atoi(sortTemp)

	statuetemp := Ctx.PostForm("status")
	status, err3 := strconv.Atoi(statuetemp)
	if err != nil || err1 != nil || err3 != nil {
		c.Error(Ctx, "非法请求", "/banner")
		return
	}
	if err2 != nil {
		c.Error(Ctx, "排序表单里面输入的数据不合法", "/banner/edit?id="+strconv.Itoa(id))
		return
	}
	bannerImgSrc, _ := c.UploadImg(Ctx, "banner_img")
	banner := model.Banner{Id: id}
	dao.DB.Find(&banner)
	banner.Title = title
	banner.BannerType = bannerType
	banner.Link = link
	banner.Sort = sort
	banner.Status = status
	if bannerImgSrc != "" {
		banner.BannerImg = bannerImgSrc
	}
	err5 := dao.DB.Save(&banner).Error
	if err5 != nil {
		c.Error(Ctx, "修改轮播图失败", "/banner/edit?id="+strconv.Itoa(id))
		return
	}
	c.Success(Ctx, "修改轮播图成功", "/banner")
}

func (c *BannerController) Delete(Ctx *gin.Context) {
	idTemp := Ctx.Query("id")
	id, err := strconv.Atoi(idTemp)
	if err != nil {
		c.Error(Ctx, "传入参数错误", "/banner")
		return
	}
	banner := model.Banner{Id: id}
	dao.DB.Find(&banner)
	address := banner.BannerImg
	if address != "" {
		test := os.Remove(address)
		if test != nil {
			fmt.Println(test)
			c.Error(Ctx, "删除物理机上图片错误", "/banner")
			return
		}
	}

	dao.DB.Delete(&banner)
	c.Success(Ctx, "删除轮播图成功", "/banner")
}
