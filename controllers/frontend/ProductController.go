package frontend

import (
	"goB2C/dao"
	"goB2C/model"
	"goB2C/util"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	BaseController
}

// 公共功能
func (p *ProductController) CategoryList(Ctx *gin.Context) {
	//调用公共功能
	p.BaseInit(Ctx)
	user := model.User{}
	model.Cookie.Get(Ctx, "userinfo", &user)

	path := Ctx.Request.URL.Path
	re := regexp.MustCompile(`/category_(\d+)\.html`)
	id := re.FindStringSubmatch(path)[1]
	cateId, _ := strconv.Atoi(id)

	curretProductCate := model.ProductCate{}
	subProductCate := []model.ProductCate{}
	dao.DB.Where("id=?", cateId).Find(&curretProductCate)
	//当前页
	tempPage := Ctx.Query("page")
	page, _ := strconv.Atoi(tempPage)
	if page == 0 {
		page = 1
	}
	//每一页显示的数量
	pageSize := 5

	var tempSlice []int
	//pid 代表子类
	if curretProductCate.Pid == 0 { //顶级分类
		//二级分类
		dao.DB.Where("pid=?", curretProductCate.Id).Find(&subProductCate)
		for i := 0; i < len(subProductCate); i++ {
			tempSlice = append(tempSlice, subProductCate[i].Id)
		}
	} else {
		//获取当前二级分类对应的同级分类
		dao.DB.Where("pid=?", curretProductCate.Pid).Find(&subProductCate)
	}
	tempSlice = append(tempSlice, cateId)
	where := "cate_id in (?)"
	product := []model.Product{}
	dao.DB.Where(where, tempSlice).Select("id,title,price,product_img,sub_title").Offset((page - 1) * pageSize).Limit(pageSize).Order("sort desc").Find(&product)
	//查询product表里面的数量
	var count int64
	dao.DB.Where(where, tempSlice).Table("product").Count(&count)

	//指定分类模板
	tpl := curretProductCate.Template
	if tpl == "" {
		tpl = "product_list.html"
	}
	Ctx.HTML(200, tpl, gin.H{
		"productList":       product,
		"subProductCate":    subProductCate,
		"curretProductCate": curretProductCate,
		"totalPages":        math.Ceil(float64(count) / float64(pageSize)),
		"page":              page,
		"userinfo":          p.UserInfo,
		"topMenuList":       p.TopMenu,
		"productCateList":   p.ProductCate,
	})
}

func (p *ProductController) ProductItem(Ctx *gin.Context) {
	p.BaseInit(Ctx)
	path := Ctx.Request.URL.Path
	re := regexp.MustCompile(`/item_(\d+)\.html`)
	tempid := re.FindStringSubmatch(path)[1]
	id, _ := strconv.Atoi(tempid)
	//获取商品信息
	product := model.Product{}
	dao.DB.Where("id=?", id).Find(&product)

	//获取关联商品  RelationProduct
	relationProduct := []model.Product{}
	product.RelationProduct = strings.ReplaceAll(product.RelationProduct, "，", ",")
	relationIds := strings.Split(product.RelationProduct, ",")
	dao.DB.Where("id in (?)", relationIds).Select("id,title,price,product_version").Find(&relationProduct)

	//获取关联赠品 ProductGift
	productGift := []model.Product{}
	product.ProductGift = strings.ReplaceAll(product.ProductGift, "，", ",")
	giftIds := strings.Split(product.ProductGift, ",")
	dao.DB.Where("id in (?)", giftIds).Select("id,title,price,product_img").Find(&productGift)

	//获取关联颜色 ProductColor
	productColor := []model.ProductColor{}
	product.ProductColor = strings.ReplaceAll(product.ProductColor, "，", ",")
	colorIds := strings.Split(product.ProductColor, ",")
	dao.DB.Where("id in (?)", colorIds).Find(&productColor)

	//获取关联配件 ProductFitting
	productFitting := []model.Product{}
	product.ProductFitting = strings.ReplaceAll(product.ProductFitting, "，", ",")
	fittingIds := strings.Split(product.ProductFitting, ",")
	dao.DB.Where("id in (?)", fittingIds).Select("id,title,price,product_img").Find(&productFitting)

	//获取商品关联的图片 ProductImage
	productImage := []model.ProductImage{}
	dao.DB.Where("product_id=?", product.Id).Find(&productImage)

	//获取规格参数信息 ProductAttr
	productAttr := []model.ProductAttr{}
	dao.DB.Where("product_id=?", product.Id).Find(&productAttr)

	//是否被收藏
	user := model.User{}
	ok := model.Cookie.Get(Ctx, "userinfo", &user)

	if ok != true {
		Ctx.JSON(200, gin.H{
			"success": false,
			"msg":     "请先登陆",
		})
		return
	}

	collect := model.ProductCollect{}
	isExist := dao.DB.Where("user_id=? AND product_id=?", user.Id, id).First(&collect)
	is_collect := true
	if isExist.RowsAffected == 0 {
		is_collect = false
	}

	Ctx.HTML(200, "product_item.html", gin.H{
		"product":         product,
		"userinfo":        p.UserInfo,
		"topMenuList":     p.TopMenu,
		"productCateList": p.ProductCate,
		"relationProduct": relationProduct,
		"productGift":     productGift,
		"productColor":    productColor,
		"productFitting":  productFitting,
		"productImage":    productImage,
		"productAttr":     productAttr,
		"collectStatus":   is_collect,
	})
}

// 收藏
func (c *ProductController) Collect(Ctx *gin.Context) {
	tempProductid := Ctx.Query("product_id")
	productId, err := strconv.Atoi(tempProductid)
	if err != nil {
		Ctx.JSON(200, gin.H{
			"success": false,
			"msg":     "传参错误",
		})
		return
	}
	user := model.User{}
	ok := model.Cookie.Get(Ctx, "userinfo", &user)
	if ok != true {
		Ctx.JSON(200, gin.H{
			"success": false,
			"msg":     "请先登陆",
		})
		return
	}
	isExist := dao.DB.First(&user)
	if isExist.RowsAffected == 0 {
		Ctx.JSON(200, gin.H{
			"success": false,
			"msg":     "非法用户",
		})
		return
	}

	goodCollect := model.ProductCollect{}
	isExist = dao.DB.Where("user_id=? AND product_id=?", user.Id, productId).First(&goodCollect)
	if isExist.RowsAffected == 0 {
		goodCollect.UserId = user.Id
		goodCollect.ProductId = productId
		goodCollect.AddTime = util.FormatDay()
		dao.DB.Create(&goodCollect)
		Ctx.JSON(200, gin.H{
			"success": true,
			"msg":     "收藏成功",
		})
	} else {
		dao.DB.Delete(&goodCollect)
		Ctx.JSON(200, gin.H{
			"success": true,
			"msg":     "取消收藏成功",
		})
	}

}

func (c *ProductController) GetImgList(Ctx *gin.Context) {
	tempcolorId := Ctx.Query("color_id")
	colorId, err1 := strconv.Atoi(tempcolorId)

	tempproductId := Ctx.Query("product_id")
	productId, err2 := strconv.Atoi(tempproductId)

	//查询商品图库信息
	productImage := []model.ProductImage{}
	err3 := dao.DB.Where("color_id=? AND product_id=?", colorId, productId).Find(&productImage).Error

	if err1 != nil || err2 != nil || err3 != nil {
		Ctx.JSON(200, gin.H{
			"result":  "失败",
			"success": false,
		})
	} else {
		if len(productImage) == 0 {
			dao.DB.Where("product_id=?", productId).Find(&productImage)
		}
		Ctx.JSON(200, gin.H{
			"result":  productImage,
			"success": true,
		})
	}
}
