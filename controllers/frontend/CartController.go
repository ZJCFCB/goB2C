package frontend

import (
	"goB2C/dao"
	"goB2C/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 购物车结构体
type CartController struct {
	BaseController
}

// 购物车展示
func (C *CartController) Get(Ctx *gin.Context) {
	C.BaseInit(Ctx)
	cartList := []model.Cart{}
	model.Cookie.Get(Ctx, "cartList", &cartList)

	var allPrice float64
	//执行计算总价
	for i := 0; i < len(cartList); i++ {
		if cartList[i].Checked {
			allPrice += cartList[i].Price * float64(cartList[i].Num)
		}
	}
	user := model.User{}
	model.Cookie.Get(Ctx, "userinfo", &user)

	Ctx.HTML(200, "cart_cart.html", gin.H{
		"cartList":    cartList,
		"allPrice":    allPrice,
		"topMenuList": C.TopMenu,
		"userinfo":    GetUserInfo(user),
	})
}

func LimitRate(Ctx *gin.Context) {

}

// 加入购物车
func (c *CartController) AddCart(Ctx *gin.Context) {
	c.BaseInit(Ctx)
	colorIdTemp := Ctx.Query("color_id")
	colorId, err1 := strconv.Atoi(colorIdTemp)
	productIdTemp := Ctx.Query("product_id")
	productId, err2 := strconv.Atoi(productIdTemp)

	product := model.Product{}
	productColor := model.ProductColor{}
	err3 := dao.DB.Where("id=?", productId).Find(&product).Error
	err4 := dao.DB.Where("id=?", colorId).Find(&productColor).Error

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		Ctx.Redirect(302, "/item_"+strconv.Itoa(product.Id)+".html")
		return
	}
	// 1.获取增加购物车的商品数据
	currentData := model.Cart{
		Id:             productId,
		Title:          product.Title,
		Price:          product.Price,
		ProductVersion: product.ProductVersion,
		Num:            1,
		ProductColor:   productColor.ColorName,
		ProductImg:     product.ProductImg,
		ProductGift:    product.ProductGift, //赠品
		ProductAttr:    "",                  //根据自己的需求拓展
		Checked:        true,                //默认选中
	}

	//2.判断购物车有没有数据（cookie）
	cartList := []model.Cart{}
	model.Cookie.Get(Ctx, "cartList", &cartList)
	if len(cartList) > 0 { //购物车有数据
		//4、判断购物车有没有当前数据
		if model.CartHasData(cartList, currentData) {
			for i := 0; i < len(cartList); i++ {
				if cartList[i].Id == currentData.Id && cartList[i].ProductColor == currentData.ProductColor && cartList[i].ProductAttr == currentData.ProductAttr {
					cartList[i].Num = cartList[i].Num + 1
				}
			}
		} else {
			cartList = append(cartList, currentData)
		}
		model.Cookie.Set(Ctx, "cartList", cartList)

	} else {
		//3.如果购物车没有任何数据，直接把当前数据写入cookie
		cartList = append(cartList, currentData)
		model.Cookie.Set(Ctx, "cartList", cartList)
	}
	user := model.User{}
	model.Cookie.Get(Ctx, "userinfo", &user)
	Ctx.HTML(200, "cart_addcart_success.html", gin.H{
		"product":         product,
		"userinfo":        GetUserInfo(user),
		"topMenuList":     GetTopMenuList(),
		"productCateList": GetProductCateList(),
	})
}

// 减少数量
func (c *CartController) DecCart(Ctx *gin.Context) {
	var flag bool
	var allPrice float64
	var currentAllPrice float64
	var num int

	productColor := Ctx.Query("product_color")
	productIdTemp := Ctx.Query("product_id")
	productId, _ := strconv.Atoi(productIdTemp)

	productAttr := ""

	cartList := []model.Cart{}
	model.Cookie.Get(Ctx, "cartList", &cartList)
	for i := 0; i < len(cartList); i++ {
		if cartList[i].Id == productId && cartList[i].ProductColor == productColor && cartList[i].ProductAttr == productAttr {
			if cartList[i].Num > 1 {
				cartList[i].Num = cartList[i].Num - 1
			}
			flag = true
			num = cartList[i].Num
			currentAllPrice = cartList[i].Price * float64(cartList[i].Num)
		}
		if cartList[i].Checked {
			allPrice += cartList[i].Price * float64(cartList[i].Num)
		}
	}

	if flag {
		model.Cookie.Set(Ctx, "cartList", cartList)
		Ctx.JSON(200, gin.H{
			"success":         true,
			"message":         "修改数量成功",
			"allPrice":        allPrice,
			"currentAllPrice": currentAllPrice,
			"num":             num,
		})

	} else {
		Ctx.JSON(200, gin.H{
			"success": false,
			"message": "传入参数错误",
		})
	}
}

// 增加数量
func (c *CartController) IncCart(Ctx *gin.Context) {
	var flag bool
	var allPrice float64
	var currentAllPrice float64
	var num int

	productColor := Ctx.Query("product_color")
	productIdTemp := Ctx.Query("product_id")
	productId, _ := strconv.Atoi(productIdTemp)
	productAttr := ""

	cartList := []model.Cart{}
	model.Cookie.Get(Ctx, "cartList", &cartList)
	for i := 0; i < len(cartList); i++ {
		if cartList[i].Id == productId && cartList[i].ProductColor == productColor && cartList[i].ProductAttr == productAttr {
			cartList[i].Num = cartList[i].Num + 1
			flag = true
			num = cartList[i].Num
			currentAllPrice = cartList[i].Price * float64(cartList[i].Num)
		}
		if cartList[i].Checked {
			allPrice += cartList[i].Price * float64(cartList[i].Num)
		}
	}

	if flag {
		model.Cookie.Set(Ctx, "cartList", cartList)
		Ctx.JSON(200, gin.H{
			"success":         true,
			"message":         "修改数量成功",
			"allPrice":        allPrice,
			"currentAllPrice": currentAllPrice,
			"num":             num,
		})

	} else {
		Ctx.JSON(200, gin.H{
			"success": false,
			"message": "传入参数错误",
		})
	}
}

func (c *CartController) ChangeOneCart(Ctx *gin.Context) {
	var flag bool
	var allPrice float64
	productColor := Ctx.Query("product_color")
	productIdTemp := Ctx.Query("product_id")
	productId, _ := strconv.Atoi(productIdTemp)
	productAttr := ""

	cartList := []model.Cart{}
	model.Cookie.Get(Ctx, "cartList", &cartList)

	for i := 0; i < len(cartList); i++ {
		if cartList[i].Id == productId && cartList[i].ProductColor == productColor && cartList[i].ProductAttr == productAttr {
			cartList[i].Checked = !cartList[i].Checked
			flag = true
		}
		if cartList[i].Checked {
			allPrice += cartList[i].Price * float64(cartList[i].Num)
		}
	}

	if flag {
		model.Cookie.Set(Ctx, "cartList", cartList)
		Ctx.JSON(200, gin.H{
			"success":  true,
			"message":  "修改状态成功",
			"allPrice": allPrice,
		})

	} else {
		Ctx.JSON(200, gin.H{
			"success": false,
			"message": "传入参数错误",
		})
	}
}

// 全选反选
func (c *CartController) ChangeAllCart(Ctx *gin.Context) {
	flagTemp := Ctx.Query("flag")
	flag, _ := strconv.Atoi(flagTemp)
	var allPrice float64
	cartList := []model.Cart{}
	model.Cookie.Get(Ctx, "cartList", &cartList)
	for i := 0; i < len(cartList); i++ {
		if flag == 1 {
			cartList[i].Checked = true
		} else {
			cartList[i].Checked = false
		}
		//计算总价
		if cartList[i].Checked {
			allPrice += cartList[i].Price * float64(cartList[i].Num)
		}
	}
	model.Cookie.Set(Ctx, "cartList", cartList)

	Ctx.JSON(200, gin.H{
		"success":  true,
		"allPrice": allPrice,
	})
}

// 删除购物车
func (c *CartController) DelCart(Ctx *gin.Context) {
	productColor := Ctx.Query("product_color")
	productIdTemp := Ctx.Query("product_id")
	productId, _ := strconv.Atoi(productIdTemp)
	productAttr := ""

	cartList := []model.Cart{}
	model.Cookie.Get(Ctx, "cartList", &cartList)
	for i := 0; i < len(cartList); i++ {
		if cartList[i].Id == productId && cartList[i].ProductColor == productColor && cartList[i].ProductAttr == productAttr {
			//执行删除
			cartList = append(cartList[:i], cartList[(i+1):]...)
		}
	}
	model.Cookie.Set(Ctx, "cartList", cartList)
	Ctx.Redirect(302, "/cart")
}
