package backend

import (
	"goB2C/dao"
	"goB2C/model"
	"goB2C/util"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductCateController struct {
	BaseController
}

func (c *ProductCateController) Get(Ctx *gin.Context) {
	productCate := []model.ProductCate{}
	dao.DB.Preload("ProductCateItem").Where("pid=0").Find(&productCate)
	Ctx.HTML(200, "productCate_index.html", gin.H{
		"productCateList": productCate,
	})
}

func (c *ProductCateController) Add(Ctx *gin.Context) {
	productCate := []model.ProductCate{}
	dao.DB.Where("pid=0").Find(&productCate)
	Ctx.HTML(200, "productCate_add.html", gin.H{
		"productCateList": productCate,
	})
}

func (c *ProductCateController) GoAdd(Ctx *gin.Context) {
	title := Ctx.PostForm("title")
	pidTemp := Ctx.PostForm("pid")
	pid, err1 := strconv.Atoi(pidTemp)
	link := Ctx.PostForm("link")
	template := Ctx.PostForm("template")
	subTitle := Ctx.PostForm("sub_title")
	keywords := Ctx.PostForm("keywords")
	description := Ctx.PostForm("description")
	sortTemp := Ctx.PostForm("sort")
	sort, err2 := strconv.Atoi(sortTemp)
	statustemp := Ctx.PostForm("status")
	status, err3 := strconv.Atoi(statustemp)

	if err1 != nil || err3 != nil {
		c.Error(Ctx, "传入参数类型不正确", "/productCate/add")
		return
	}
	if err2 != nil {
		c.Error(Ctx, "排序值必须是整数", "/productCate/add")
		return
	}
	uploadDir, _ := c.UploadImg(Ctx, "cate_img")
	productCate := model.ProductCate{
		Title:       title,
		Pid:         pid,
		SubTitle:    subTitle,
		Link:        link,
		Template:    template,
		Keywords:    keywords,
		Description: description,
		CateImg:     uploadDir,
		Sort:        sort,
		Status:      status,
		AddTime:     int(util.GetUnix()),
	}
	err := dao.DB.Create(&productCate).Error
	if err != nil {
		c.Error(Ctx, "增加失败", "/productCate/add")
		return
	}
	dao.RedisDel("productCate")
	c.Success(Ctx, "增加成功", "/productCate")
}

func (c *ProductCateController) Edit(Ctx *gin.Context) {
	idTemp := Ctx.Query("id")
	id, err1 := strconv.Atoi(idTemp)
	if err1 != nil {
		c.Error(Ctx, "传入参数错误", "/productCate")
		return
	}
	productCate := model.ProductCate{Id: id}
	dao.DB.Find(&productCate)

	productCateList := []model.ProductCate{}
	dao.DB.Where("pid=0").Find(&productCateList)

	Ctx.HTML(200, "productCate_edit.html", gin.H{
		"productCate":     productCate,
		"productCateList": productCateList,
	})
}

func (c *ProductCateController) GoEdit(Ctx *gin.Context) {
	title := Ctx.PostForm("title")
	pidTemp := Ctx.PostForm("pid")
	pid, err1 := strconv.Atoi(pidTemp)
	link := Ctx.PostForm("link")
	template := Ctx.PostForm("template")
	subTitle := Ctx.PostForm("sub_title")
	keywords := Ctx.PostForm("keywords")
	description := Ctx.PostForm("description")
	sortTemp := Ctx.PostForm("sort")
	sort, err2 := strconv.Atoi(sortTemp)
	statustemp := Ctx.PostForm("status")
	status, err3 := strconv.Atoi(statustemp)
	idtemp := Ctx.PostForm("id")
	id, err := strconv.Atoi(idtemp)
	if err != nil || err1 != nil || err3 != nil {
		c.Error(Ctx, "传入参数类型不正确", "/productCate/edit")
		return
	}
	if err2 != nil {
		c.Error(Ctx, "排序值必须是整数", "/productCate/edit?id="+strconv.Itoa(id))
		return
	}
	uploadDir, _ := c.UploadImg(Ctx, "cate_img")
	productCate := model.ProductCate{Id: id}
	dao.DB.Find(&productCate)
	productCate.Title = title
	productCate.Pid = pid
	productCate.Link = link
	productCate.Template = template
	productCate.SubTitle = subTitle
	productCate.Keywords = keywords
	productCate.Description = description
	productCate.Sort = sort
	productCate.Status = status
	if uploadDir != "" {
		productCate.CateImg = uploadDir
	}
	err5 := dao.DB.Save(&productCate).Error
	if err5 != nil {
		c.Error(Ctx, "修改失败", "/productCate/edit?id="+strconv.Itoa(id))
		return
	}
	dao.RedisDel("productCate")
	c.Success(Ctx, "修改成功", "/productCate")
}

func (c *ProductCateController) Delete(Ctx *gin.Context) {
	idTemp := Ctx.Query("id")
	id, err := strconv.Atoi(idTemp)
	if err != nil {
		c.Error(Ctx, "传入参数错误", "/goodCate")
		return
	}
	productCate := model.ProductCate{Id: id}
	dao.DB.Find(&productCate)
	address := productCate.CateImg
	if address != "" {
		test := os.Remove(address)
		if test != nil {
			c.Error(Ctx, "删除物理机上图片错误", "/productCate")
			return
		}
	}
	if productCate.Pid == 0 {
		productCate2 := []model.ProductCate{}
		dao.DB.Where("pid=?", productCate.Id).Find(&productCate2)
		if len(productCate2) > 0 {
			c.Error(Ctx, "请先删除当前顶级分类下面的商品！", "/productCate")
			return
		}
	}
	dao.DB.Delete(&productCate)
	dao.RedisDel("productCate")
	c.Success(Ctx, "删除成功", "/productCate")
}
