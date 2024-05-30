package backend

import (
	"goB2C/dao"
	"goB2C/model"
	"goB2C/util"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ProductTypeAttrController struct {
	BaseController
}

func (c *ProductTypeAttrController) Get(Ctx *gin.Context) {
	cateIdTemp := Ctx.Query("cate_id")
	cateId, err1 := strconv.Atoi(cateIdTemp)
	if err1 != nil {
		c.Error(Ctx, "非法请求", "/productType")
	}
	//获取当前的类型
	productType := model.ProductType{Id: cateId}
	dao.DB.Find(&productType)

	//查询当前类型下面的商品类型属性
	productTypeAttr := []model.ProductTypeAttribute{}
	dao.DB.Where("cate_id=?", cateId).Find(&productTypeAttr)

	Ctx.HTML(200, "productTypeAttribute_index.html", gin.H{
		"productType":         productType,
		"productTypeAttrList": productTypeAttr,
	})
}

func (c *ProductTypeAttrController) Add(Ctx *gin.Context) {

	cateIdTemp := Ctx.Query("cate_id")
	cateId, err1 := strconv.Atoi(cateIdTemp)
	if err1 != nil {
		c.Error(Ctx, "非法请求", "/productType")
	}

	productType := []model.ProductType{}
	dao.DB.Find(&productType)

	Ctx.HTML(200, "productTypeAttribute_add.html", gin.H{
		"productTypeList": productType,
		"cateId":          cateId,
	})
}

func (c *ProductTypeAttrController) GoAdd(Ctx *gin.Context) {
	cateIdTemp := Ctx.PostForm("cate_id")
	cateId, err1 := strconv.Atoi(cateIdTemp)

	title := Ctx.PostForm("title")
	attrTypeTemp := Ctx.PostForm("attr_type")
	attrType, err2 := strconv.Atoi(attrTypeTemp)
	attrValue := Ctx.PostForm("attr_value")
	sortTemp := Ctx.PostForm("sort")
	sort, err4 := strconv.Atoi(sortTemp)
	if err1 != nil || err2 != nil {
		c.Error(Ctx, "非法请求", "/productType")
		return
	}
	if strings.Trim(title, " ") == "" {
		c.Error(Ctx, "商品类型属性名称不能为空", "/productTypeAttribute/add?cate_id="+strconv.Itoa(cateId))
		return
	}
	if err4 != nil {
		c.Error(Ctx, "排序值错误", "/productTypeAttribute/add?cate_id="+strconv.Itoa(cateId))
		return
	}
	productTypeAttr := model.ProductTypeAttribute{
		Title:     title,
		CateId:    cateId,
		AttrType:  attrType,
		AttrValue: attrValue,
		Status:    1,
		AddTime:   int(util.GetUnix()),
		Sort:      sort,
	}
	dao.DB.Create(&productTypeAttr)
	c.Success(Ctx, "增加成功", "/productTypeAttribute?cate_id="+strconv.Itoa(cateId))

}

func (c *ProductTypeAttrController) Edit(Ctx *gin.Context) {
	IdTemp := Ctx.Query("id")
	id, err1 := strconv.Atoi(IdTemp)
	if err1 != nil {
		c.Error(Ctx, "非法请求", "/goodType")
		return
	}
	productTypeAttr := model.ProductTypeAttribute{Id: id}
	dao.DB.Find(&productTypeAttr)
	productType := []model.ProductType{}
	dao.DB.Find(&productType)

	Ctx.HTML(200, "productTypeAttribute_edit.html", gin.H{
		"productTypeAttr": productTypeAttr,
		"productTypeList": productType,
	})
}

func (c *ProductTypeAttrController) GoEdit(Ctx *gin.Context) {

	cateIdTemp := Ctx.PostForm("cate_id")
	cateId, err1 := strconv.Atoi(cateIdTemp)

	title := Ctx.PostForm("title")
	attrTypeTemp := Ctx.PostForm("attr_type")
	attrType, err2 := strconv.Atoi(attrTypeTemp)
	attrValue := Ctx.PostForm("attr_value")
	sortTemp := Ctx.PostForm("sort")
	sort, err4 := strconv.Atoi(sortTemp)

	IdTemp := Ctx.PostForm("id")
	id, err := strconv.Atoi(IdTemp)

	if err != nil || err1 != nil || err2 != nil {
		c.Error(Ctx, "非法请求", "/productTypeAttribute")
		return
	}
	if strings.Trim(title, " ") == "" {
		c.Error(Ctx, "商品类型属性名称不能为空", "/productTypeAttribute/edit?cate_id="+strconv.Itoa(id))
		return
	}
	if err4 != nil {
		c.Error(Ctx, "排序值错误", "/productTypeAttribute/edit?cate_id="+strconv.Itoa(id))
		return
	}
	productTypeAttr := model.ProductTypeAttribute{Id: id}
	dao.DB.Find(&productTypeAttr)
	productTypeAttr.Title = title
	productTypeAttr.CateId = cateId
	productTypeAttr.AttrType = attrType
	productTypeAttr.AttrValue = attrValue
	productTypeAttr.Sort = sort
	err3 := dao.DB.Save(&productTypeAttr).Error
	if err3 != nil {
		c.Error(Ctx, "修改数据失败", "/productTypeAttribute/edit?cate_id="+strconv.Itoa(id))
	}
	c.Success(Ctx, "修改数据成功", "/productTypeAttribute?cate_id="+strconv.Itoa(cateId))
}
func (c *ProductTypeAttrController) Delete(Ctx *gin.Context) {
	IdTemp := Ctx.Query("id")
	id, err := strconv.Atoi(IdTemp)
	cateIdTemp := Ctx.Query("cate_id")
	cateId, err1 := strconv.Atoi(cateIdTemp)
	if err != nil {
		c.Error(Ctx, "传入参数错误", "/productTypeAttribute?cate_id="+strconv.Itoa(cateId))
		return
	}
	if err1 != nil {
		c.Error(Ctx, "非法请求", "/productType")
	}
	productTypeAttr := model.ProductTypeAttribute{Id: id}
	dao.DB.Delete(&productTypeAttr)
	c.Success(Ctx, "删除数据成功", "/productTypeAttribute?cate_id="+strconv.Itoa(cateId))
}
