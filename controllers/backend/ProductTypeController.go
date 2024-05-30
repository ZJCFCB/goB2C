package backend

import (
	"goB2C/dao"
	"goB2C/model"
	"goB2C/util"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ProductTypeController struct {
	BaseController
}

func (c *ProductTypeController) Get(Ctx *gin.Context) {
	productType := []model.ProductType{}
	dao.DB.Find(&productType)
	Ctx.HTML(200, "productType_index.html", gin.H{
		"productTypeList": productType,
	})

}

func (c *ProductTypeController) Add(Ctx *gin.Context) {
	Ctx.HTML(200, "productType_add.html", gin.H{})
}

func (c *ProductTypeController) GoAdd(Ctx *gin.Context) {
	title := strings.Trim(Ctx.PostForm("title"), "")
	description := strings.Trim(Ctx.PostForm("description"), "")
	statusTemp := Ctx.PostForm("status")
	status, err := strconv.Atoi(statusTemp)
	if title == "" {
		c.Error(Ctx, "标题不能为空", "/productType/add")
		return
	}
	if err != nil {
		c.Error(Ctx, "传入参数不正确", "/productType/add")
		return
	}
	productTypeList := []model.ProductType{}
	dao.DB.Where("title=?", title).Find(&productTypeList)
	if len(productTypeList) != 0 {
		c.Error(Ctx, "该商品已存在！", "/productType/add")
		return
	}
	productType := model.ProductType{}
	productType.Title = title
	productType.Description = description
	productType.Status = status
	productType.AddTime = int(util.GetUnix())
	err1 := dao.DB.Create(&productType).Error
	if err1 != nil {
		c.Error(Ctx, "增加商品类型失败", "/productType/add")
	} else {
		c.Success(Ctx, "增加商品类型成功", "/productType")
	}
}

func (c *ProductTypeController) Edit(Ctx *gin.Context) {
	idtemp := Ctx.Query("id")
	id, err := strconv.Atoi(idtemp)
	if err != nil {
		c.Error(Ctx, "传入参数错误", "/productType")
		return
	}
	productType := model.ProductType{Id: id}
	dao.DB.Find(&productType)

	Ctx.HTML(200, "productType_edit.html", gin.H{
		"productType": productType,
	})
}

func (c *ProductTypeController) GoEdit(Ctx *gin.Context) {
	title := strings.Trim(Ctx.PostForm("title"), "")
	description := strings.Trim(Ctx.PostForm("description"), "")
	statusTemp := Ctx.PostForm("status")
	status, err := strconv.Atoi(statusTemp)
	idtemp := Ctx.PostForm("id")
	id, err1 := strconv.Atoi(idtemp)
	if err != nil || err1 != nil {
		c.Error(Ctx, "传入参数错误", "/productType")
		return
	}
	if title == "" {
		c.Error(Ctx, "标题不能为空", "/productType/edit?id="+strconv.Itoa(id))
		return
	}
	productType := model.ProductType{Id: id}
	dao.DB.Find(&productType)
	productType.Title = title
	productType.Description = description
	productType.Status = status
	err2 := dao.DB.Save(&productType).Error
	if err2 != nil {
		c.Error(Ctx, "修改商品类型失败", "/productType/edit?id="+strconv.Itoa(id))
	} else {
		c.Success(Ctx, "修改商品类型成功", "/productType")
	}
}

func (c *ProductTypeController) Delete(Ctx *gin.Context) {
	idtemp := Ctx.Query("id")
	id, err := strconv.Atoi(idtemp)
	if err != nil {
		c.Error(Ctx, "传入参数错误", "/productType")
		return
	}
	productType := model.ProductType{Id: id}
	dao.DB.Delete(&productType)
	c.Success(Ctx, "删除数据成功", "/productType")
}
