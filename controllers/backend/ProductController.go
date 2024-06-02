package backend

import (
	"fmt"
	"goB2C/dao"
	"goB2C/model"
	"goB2C/util"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

var wg sync.WaitGroup

type ProductController struct {
	BaseController
}

func (c *ProductController) Get(Ctx *gin.Context) {
	pageTemp := Ctx.Query("page")
	page, _ := strconv.Atoi(pageTemp)
	if page == 0 {
		page = 1
	}
	pageSize := 5
	keyword := Ctx.Query("keyword")
	where := "1=1"
	if len(keyword) > 0 {
		where += " AND title like \"%" + keyword + "%\""
	}
	productList := []model.Product{}
	dao.DB.Where(where).Offset((page - 1) * pageSize).Limit(pageSize).Find(&productList)
	var count int64
	dao.DB.Where(where).Table("product").Count(&count)
	Ctx.HTML(200, "product_index.html", gin.H{
		"productList": productList,
		"totalPages":  math.Ceil(float64(count) / float64(pageSize)),
		"page":        page,
	})

}

func (c *ProductController) Add(Ctx *gin.Context) {

	//获取商品分类
	productCate := []model.ProductCate{}
	dao.DB.Where("pid=?", 0).Preload("ProductCateItem").Find(&productCate)

	//获取颜色信息
	productColor := []model.ProductColor{}
	dao.DB.Find(&productColor)

	//获取商品类型信息
	productType := []model.ProductType{}
	dao.DB.Find(&productType)

	Ctx.HTML(200, "product_add.html", gin.H{
		"productCateList": productCate,
		"productColor":    productColor,
		"productType":     productType,
	})
}

func (c *ProductController) GoAdd(Ctx *gin.Context) {

	//1、获取表单提交过来的数据
	title := Ctx.PostForm("title")
	subTitle := Ctx.PostForm("sub_title")
	productSn := Ctx.PostForm("product_sn")
	cateTemp := Ctx.PostForm("cate_id")
	cateId, _ := strconv.Atoi(cateTemp)
	productnumTemp := Ctx.PostForm("product_number")
	productNumber, _ := strconv.Atoi(productnumTemp)
	marketPriceTemp := Ctx.PostForm("market_price")

	marketPrice, _ := strconv.ParseFloat(marketPriceTemp, 64)
	priceTemp := Ctx.PostForm("price")
	price, _ := strconv.ParseFloat(priceTemp, 64)
	relationProduct := Ctx.PostForm("relation_product")
	productAttr := Ctx.PostForm("product_attr")
	productVersion := Ctx.PostForm("product_version")
	productGift := Ctx.PostForm("product_gift")
	productFitting := Ctx.PostForm("product_fitting")
	productColor := Ctx.PostFormArray("product_color")
	productKeywords := Ctx.PostForm("product_keywords")
	productDesc := Ctx.PostForm("product_desc")
	productContent := Ctx.PostForm("product_content")
	isDeleteTemp := Ctx.PostForm("is_delete")
	isDelete, _ := strconv.Atoi(isDeleteTemp)
	isHotTemp := Ctx.PostForm("is_hot")
	isHot, _ := strconv.Atoi(isHotTemp)
	isBestTemp := Ctx.PostForm("is_best")
	isBest, _ := strconv.Atoi(isBestTemp)
	isNewTemp := Ctx.PostForm("is_new")
	isNew, _ := strconv.Atoi(isNewTemp)
	productTypeidTemp := Ctx.PostForm("product_type_id")
	productTypeId, _ := strconv.Atoi(productTypeidTemp)
	sortTemp := Ctx.PostForm("sort")
	sort, _ := strconv.Atoi(sortTemp)
	statusTemp := Ctx.PostForm("status")
	status, _ := strconv.Atoi(statusTemp)
	addTime := int(util.GetUnix())

	//2、获取颜色信息 把颜色转化成字符串
	productColorStr := strings.Join(productColor, ",")

	//3、上传图片   生成缩略图
	productImg, _ := c.UploadImg(Ctx, "product_img")

	//4、增加商品数据
	product := model.Product{
		Title:           title,
		SubTitle:        subTitle,
		ProductSn:       productSn,
		CateId:          cateId,
		ClickCount:      100,
		ProductNumber:   productNumber,
		MarketPrice:     marketPrice,
		Price:           price,
		RelationProduct: relationProduct,
		ProductAttr:     productAttr,
		ProductVersion:  productVersion,
		ProductGift:     productGift,
		ProductFitting:  productFitting,
		ProductKeywords: productKeywords,
		ProductDesc:     productDesc,
		ProductContent:  productContent,
		IsDelete:        isDelete,
		IsHot:           isHot,
		IsBest:          isBest,
		IsNew:           isNew,
		ProductTypeId:   productTypeId,
		Sort:            sort,
		Status:          status,
		AddTime:         addTime,
		ProductColor:    productColorStr,
		ProductImg:      productImg,
	}
	err1 := dao.DB.Create(&product).Error
	if err1 != nil {
		c.Error(Ctx, "增加失败", "/product/add")
	}
	//5、增加图库 信息
	wg.Add(1)
	go func() {
		productImageList := Ctx.PostFormArray("product_image_list")
		for _, v := range productImageList {
			productImgObj := model.ProductImage{}
			productImgObj.ProductId = product.Id
			productImgObj.ImgUrl = v
			productImgObj.Sort = 10
			productImgObj.Status = 1
			productImgObj.AddTime = int(util.GetUnix())
			dao.DB.Create(&productImgObj)
		}
		wg.Done()
	}()

	//6、增加规格包装
	wg.Add(1)
	go func() {
		attrIdList := Ctx.PostFormArray("attr_id_list")
		attrValueList := Ctx.PostFormArray("attr_value_list")
		for i := 0; i < len(attrIdList); i++ {
			productTypeAttributeId, _ := strconv.Atoi(attrIdList[i])
			productTypeAttributeObj := model.ProductTypeAttribute{Id: productTypeAttributeId}
			dao.DB.Find(&productTypeAttributeObj)

			productAttrObj := model.ProductAttr{}
			productAttrObj.ProductId = product.Id
			productAttrObj.AttributeTitle = productTypeAttributeObj.Title
			productAttrObj.AttributeType = productTypeAttributeObj.AttrType
			productAttrObj.AttributeId = productTypeAttributeObj.Id
			productAttrObj.AttributeCateId = productTypeAttributeObj.CateId
			productAttrObj.AttributeValue = attrValueList[i]
			productAttrObj.Status = 1
			productAttrObj.Sort = 10
			productAttrObj.AddTime = int(util.GetUnix())
			dao.DB.Create(&productAttrObj)
		}
		wg.Done()
	}()

	wg.Wait()
	c.Success(Ctx, "增加数据成功", "/product")

}
func (c *ProductController) Edit(Ctx *gin.Context) {

	// 1、获取商品数据
	idTemp := Ctx.Query("id")
	id, err1 := strconv.Atoi(idTemp)
	if err1 != nil {
		c.Error(Ctx, "非法请求", "/product")
	}
	product := model.Product{Id: id}
	dao.DB.Find(&product)

	//2、获取商品分类
	productCate := []model.ProductCate{}
	dao.DB.Where("pid=?", 0).Preload("ProductCateItem").Find(&productCate)

	// 3、获取所有颜色 以及选中的颜色
	productColorSlice := strings.Split(product.ProductColor, ",")
	productColorMap := make(map[string]string)
	for _, v := range productColorSlice {
		productColorMap[v] = v
	}
	//获取颜色信息
	productColor := []model.ProductColor{}
	dao.DB.Find(&productColor)
	for i := 0; i < len(productColor); i++ {
		_, ok := productColorMap[strconv.Itoa(productColor[i].Id)]
		if ok {
			productColor[i].Checked = true
		}
	}
	//4、商品的图库信息
	productImage := []model.ProductImage{}
	dao.DB.Where("product_id=?", product.Id).Find(&productImage)

	// 5、获取商品类型
	productType := []model.ProductType{}
	dao.DB.Find(&productType)

	//6、获取规格信息
	productAttr := []model.ProductAttr{}
	dao.DB.Where("product_id=?", product.Id).Find(&productAttr)

	var productAttrStr string
	for _, v := range productAttr {
		if v.AttributeType == 1 {
			productAttrStr += fmt.Sprintf(`<li><span>%v: 　</span>  <input type="hidden" name="attr_id_list" value="%v" />   <input type="text" name="attr_value_list" value="%v" /></li>`, v.AttributeTitle, v.AttributeId, v.AttributeValue)
		} else if v.AttributeType == 2 {
			productAttrStr += fmt.Sprintf(`<li><span>%v: 　</span><input type="hidden" name="attr_id_list" value="%v" />  <textarea cols="50" rows="3" name="attr_value_list">%v</textarea></li>`, v.AttributeTitle, v.AttributeId, v.AttributeValue)
		} else {

			// 获取 attr_value  获取可选值列表
			oneProductTypeAttribute := model.ProductTypeAttribute{Id: v.AttributeId}
			dao.DB.Find(&oneProductTypeAttribute)
			attrValueSlice := strings.Split(oneProductTypeAttribute.AttrValue, "\n")
			productAttrStr += fmt.Sprintf(`<li><span>%v: 　</span>  <input type="hidden" name="attr_id_list" value="%v" /> `, v.AttributeTitle, v.AttributeId)
			productAttrStr += fmt.Sprintf(`<select name="attr_value_list">`)
			for j := 0; j < len(attrValueSlice); j++ {
				if attrValueSlice[j] == v.AttributeValue {
					productAttrStr += fmt.Sprintf(`<option value="%v" selected >%v</option>`, attrValueSlice[j], attrValueSlice[j])
				} else {
					productAttrStr += fmt.Sprintf(`<option value="%v">%v</option>`, attrValueSlice[j], attrValueSlice[j])
				}
			}
			productAttrStr += fmt.Sprintf(`</select>`)
			productAttrStr += fmt.Sprintf(`</li>`)
		}
	}

	Ctx.HTML(200, "product_edit.html", gin.H{
		"productAttrStr": productAttrStr,
		//上一页地址
		"prevPage":        Ctx.Request.Referer(),
		"productType":     productType,
		"productImage":    productImage,
		"productColor":    productColor,
		"productCateList": productCate,
		"product":         product,
	})
}

func (c *ProductController) GoEdit(Ctx *gin.Context) {

	//1.获取要修改的商品数据
	idTemp := Ctx.PostForm("id")
	id, err1 := strconv.Atoi(idTemp)
	if err1 != nil {
		c.Error(Ctx, "非法请求", "/product")
	}

	title := Ctx.PostForm("title")
	subTitle := Ctx.PostForm("sub_title")
	productSn := Ctx.PostForm("product_sn")
	cateTemp := Ctx.PostForm("cate_id")
	cateId, _ := strconv.Atoi(cateTemp)
	productnumTemp := Ctx.PostForm("product_number")
	productNumber, _ := strconv.Atoi(productnumTemp)
	marketPriceTemp := Ctx.PostForm("market_price")
	marketPrice, _ := strconv.ParseFloat(marketPriceTemp, 64)
	priceTemp := Ctx.PostForm("price")
	price, _ := strconv.ParseFloat(priceTemp, 64)
	relationProduct := Ctx.PostForm("relation_product")
	productAttr := Ctx.PostForm("product_attr")
	productVersion := Ctx.PostForm("product_version")
	productGift := Ctx.PostForm("product_gift")
	productFitting := Ctx.PostForm("product_fitting")
	productColor := Ctx.PostFormArray("product_color")
	productKeywords := Ctx.PostForm("product_keywords")
	productDesc := Ctx.PostForm("product_desc")
	productContent := Ctx.PostForm("product_content")
	isDeleteTemp := Ctx.PostForm("is_delete")
	isDelete, _ := strconv.Atoi(isDeleteTemp)
	isHotTemp := Ctx.PostForm("is_hot")
	isHot, _ := strconv.Atoi(isHotTemp)
	isBestTemp := Ctx.PostForm("is_best")
	isBest, _ := strconv.Atoi(isBestTemp)
	isNewTemp := Ctx.PostForm("is_new")
	isNew, _ := strconv.Atoi(isNewTemp)
	productTypeidTemp := Ctx.PostForm("product_type_id")
	productTypeId, _ := strconv.Atoi(productTypeidTemp)
	sortTemp := Ctx.PostForm("sort")
	sort, _ := strconv.Atoi(sortTemp)
	statusTemp := Ctx.PostForm("status")
	status, _ := strconv.Atoi(statusTemp)

	prevPage := Ctx.PostForm("prevPage")
	//2.获取颜色信息 把颜色转化成字符串
	productColorStr := strings.Join(productColor, ",")
	product := model.Product{Id: id}
	dao.DB.Find(&product)
	product.Title = title
	product.SubTitle = subTitle
	product.ProductSn = productSn
	product.CateId = cateId
	product.ProductNumber = productNumber
	product.MarketPrice = marketPrice
	product.Price = price
	product.RelationProduct = relationProduct
	product.ProductAttr = productAttr
	product.ProductVersion = productVersion
	product.ProductGift = productGift
	product.ProductFitting = productFitting
	product.ProductKeywords = productKeywords
	product.ProductDesc = productDesc
	product.ProductContent = productContent
	product.IsDelete = isDelete
	product.IsHot = isHot
	product.IsBest = isBest
	product.IsNew = isNew
	product.ProductTypeId = productTypeId
	product.Sort = sort
	product.Status = status
	product.ProductColor = productColorStr
	//3.上传图片，生成缩略图
	productImg, err2 := c.UploadImg(Ctx, "product_img")
	if err2 == nil && len(productImg) > 0 {
		product.ProductImg = productImg
	}
	//4.执行修改商品
	err3 := dao.DB.Save(&product).Error
	if err3 != nil {
		c.Error(Ctx, "修改数据失败", "/product/edit?id="+strconv.Itoa(id))
		return
	}
	//5.修改图库数据 （增加）
	wg.Add(1)
	go func() {
		productImageList := Ctx.PostFormArray("product_image_list")
		for _, v := range productImageList {
			productImgObj := model.ProductImage{}
			productImgObj.ProductId = product.Id
			productImgObj.ImgUrl = v
			productImgObj.Sort = 10
			productImgObj.Status = 1
			productImgObj.AddTime = int(util.GetUnix())
			dao.DB.Create(&productImgObj)
		}
		wg.Done()
	}()

	//6.修改商品类型属性数据         1、删除当前商品id对应的类型属性  2、执行增加

	//删除当前商品id对应的类型属性
	productAttrObj := model.ProductAttr{}
	dao.DB.Where("product_id=?", product.Id).Delete(&productAttrObj)
	//执行增加
	wg.Add(1)
	go func() {
		attrIdList := Ctx.PostFormArray("attr_id_list")
		attrValueList := Ctx.PostFormArray("attr_value_list")
		for i := 0; i < len(attrIdList); i++ {
			productTypeAttributeId, _ := strconv.Atoi(attrIdList[i])
			productTypeAttributeObj := model.ProductTypeAttribute{Id: productTypeAttributeId}
			dao.DB.Find(&productTypeAttributeObj)

			productAttrObj := model.ProductAttr{}
			productAttrObj.ProductId = product.Id
			productAttrObj.AttributeTitle = productTypeAttributeObj.Title
			productAttrObj.AttributeType = productTypeAttributeObj.AttrType
			productAttrObj.AttributeId = productTypeAttributeObj.Id
			productAttrObj.AttributeCateId = productTypeAttributeObj.CateId
			productAttrObj.AttributeValue = attrValueList[i]
			productAttrObj.Status = 1
			productAttrObj.Sort = 10
			productAttrObj.AddTime = int(util.GetUnix())
			dao.DB.Create(&productAttrObj)
		}
		wg.Done()
	}()

	wg.Wait()
	c.Success(Ctx, "修改数据成功", prevPage)
}
func (c *ProductController) Delete(Ctx *gin.Context) {
	col := Ctx.Query("id")
	id, err1 := strconv.Atoi(col)
	if err1 != nil {
		c.Error(Ctx, "传入参数错误", "/product")
		return
	}
	product := model.Product{Id: id}
	dao.DB.Find(&product)
	path, _ := os.Getwd()
	path = strings.ReplaceAll(path, "\\", "/")
	address := path + "/" + product.ProductImg
	os.Remove(address)
	err2 := dao.DB.Delete(&product).Error
	if err2 != nil {
		productAttr := model.ProductAttr{ProductId: id}
		dao.DB.Delete(&productAttr)
		productImage := model.ProductImage{ProductId: id}
		dao.DB.Delete(&productImage)
	}
	c.Success(Ctx, "删除商品成功", Ctx.Request.Referer())
}

func (c *ProductController) GoUpload(Ctx *gin.Context) {
	fmt.Println("Called go upload")
	savePath, err := c.UploadImg(Ctx, "file")
	if err != nil {
		Ctx.JSON(200, gin.H{
			"link": "",
		})
	} else {
		//返回json数据
		Ctx.JSON(200, gin.H{
			"link": "/" + savePath,
		})

	}

}

// 获取商品类型属性
func (c *ProductController) GetProductTypeAttribute(Ctx *gin.Context) {
	col := Ctx.Query("cate_id")
	cate_id, err1 := strconv.Atoi(col)
	ProductTypeAttribute := []model.ProductTypeAttribute{}
	err2 := dao.DB.Where("cate_id=?", cate_id).Find(&ProductTypeAttribute).Error
	if err1 != nil || err2 != nil {

		Ctx.JSON(200, gin.H{
			"result":  "",
			"success": false,
		})

	} else {
		Ctx.JSON(200, gin.H{
			"result":  ProductTypeAttribute,
			"success": true,
		})

	}

}

// 修改图片对应颜色信息
func (c *ProductController) ChangeProductImageColor(Ctx *gin.Context) {
	col := Ctx.Query("color_id")
	colorId, err1 := strconv.Atoi(col)
	p := Ctx.Query("product_image_id")
	productImageId, err2 := strconv.Atoi(p)

	productImage := model.ProductImage{Id: productImageId}
	dao.DB.Find(&productImage)
	productImage.ColorId = colorId
	err3 := dao.DB.Save(&productImage).Error

	if err1 != nil || err2 != nil || err3 != nil {
		Ctx.JSON(200, gin.H{
			"result":  "更新失败",
			"success": false,
		})

	} else {
		Ctx.JSON(200, gin.H{
			"result":  "更新成功",
			"success": true,
		})

	}
}

// 删除图库
func (c *ProductController) RemoveProductImage(Ctx *gin.Context) {
	p := Ctx.Query("product_image_id")
	productImageId, err1 := strconv.Atoi(p)
	productImage := model.ProductImage{Id: productImageId}
	err2 := dao.DB.Delete(&productImage).Error
	os.Remove(productImage.ImgUrl)

	if err1 != nil || err2 != nil {
		Ctx.JSON(200, gin.H{
			"result":  "删除失败",
			"success": false,
		})
	} else {
		//删除图片
		Ctx.JSON(200, gin.H{
			"result":  "删除",
			"success": true,
		})
	}

}
