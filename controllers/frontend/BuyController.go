package frontend

import (
	"fmt"
	"goB2C/dao"
	"goB2C/model"
	"goB2C/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BuyController struct {
	BaseController
}

// 结算页面
func (B *BuyController) Checkout(Ctx *gin.Context) {

	B.BaseInit(Ctx)
	//1.获取要结算的商品
	cartList := []model.Cart{}
	orderList := []model.Cart{} //要结算的商品

	//Todo 这是在干啥?

	model.Cookie.Get(Ctx, "cartList", &cartList)

	var allPrice float64
	//2.执行计算总价
	for i := 0; i < len(cartList); i++ {
		if cartList[i].Checked {
			allPrice += cartList[i].Price * float64(cartList[i].Num)
			orderList = append(orderList, cartList[i])
		}
	}
	//3.判断去结算页面有没有要结算的商品
	if len(orderList) == 0 {
		Ctx.Redirect(302, "/mainPage")
		return
	}

	//3.获取收货地址
	user := model.User{}
	model.Cookie.Get(Ctx, "userinfo", &user)
	addressList := []model.Address{}
	dao.DB.Where("uid=?", user.Id).Order("default_address desc").Find(&addressList)

	//4.防止重复提交订单 生成签名
	orderSign := util.Md5(util.GetRandomNum())

	key := "orderSign" + string(user.Phone) + string(user.Password)
	dao.RedisSetString(key, orderSign)

	Ctx.HTML(200, "buy_checkout.html", gin.H{
		"orderList":   orderList,
		"allPrice":    allPrice,
		"addressList": addressList,
		"orderSign":   orderSign,
		"topMenuList": B.TopMenu,
		"userinfo":    B.UserInfo,
	})
}

/*
提交订单
1、获取收货地址信息
2、获取购买商品的信息
3、把订单信息放在订单表，把商品信息放在商品表
4、删除购物车里面的选中数据
*/
func (B *BuyController) GoOrder(Ctx *gin.Context) {
	//0、防止重复提交订单
	user := model.User{}
	model.Cookie.Get(Ctx, "userinfo", &user)

	orderSign := Ctx.PostForm("orderSign")
	key := "orderSign" + string(user.Phone) + string(user.Password)
	sessionOrderSign, _ := dao.RedisGetString(key)
	if sessionOrderSign != orderSign {
		Ctx.Redirect(302, "/mainPage")
		return
	}
	dao.RedisDel(key)

	// 1、获取收货地址信息

	addressResult := []model.Address{}
	dao.DB.Where("uid=? AND default_address=1", user.Id).Find(&addressResult)

	if len(addressResult) > 0 {

		// 2、获取购买商品的信息   orderList就是要购买的商品信息
		cartList := []model.Cart{}
		orderList := []model.Cart{} //要结算的商品
		model.Cookie.Get(Ctx, "cartList", &cartList)
		var allPrice float64
		for i := 0; i < len(cartList); i++ {
			if cartList[i].Checked {
				allPrice += cartList[i].Price * float64(cartList[i].Num)
				orderList = append(orderList, cartList[i])
			}
		}

		//  3、把订单信息放在订单表，把商品信息放在商品表
		order := model.Order{
			OrderID:     util.GenerateOrderId(),
			Uid:         user.Id,
			AllPrice:    allPrice,
			Phone:       addressResult[0].Phone,
			Name:        addressResult[0].Name,
			Address:     addressResult[0].Address,
			Zipcode:     addressResult[0].Zipcode,
			PayStatus:   0,
			PayType:     0,
			OrderStatus: 0,
			AddTime:     int(util.GetUnix()),
		}
		err := dao.DB.Create(&order).Error
		if err == nil {
			for i := 0; i < len(orderList); i++ {
				orderItem := model.OrderItem{
					OrderId:        string(order.Id),
					Uid:            user.Id,
					ProductTitle:   orderList[i].Title,
					ProductId:      orderList[i].Id,
					ProductImg:     orderList[i].ProductImg,
					ProductPrice:   orderList[i].Price,
					ProductNum:     orderList[i].Num,
					ProductVersion: orderList[i].ProductVersion,
					ProductColor:   orderList[i].ProductColor,
					AddTime:        int(util.GetUnix()),
				}
				err := dao.DB.Create(&orderItem).Error
				if err != nil {
					fmt.Println(err)
				}
			}
			// 4、删除购物车里面的选中数据

			noSelectedCartList := []model.Cart{}
			for i := 0; i < len(cartList); i++ {
				if !cartList[i].Checked {
					noSelectedCartList = append(noSelectedCartList, cartList[i])
				}
			}
			model.Cookie.Set(Ctx, "cartList", noSelectedCartList)
			Ctx.Redirect(302, "/buy/confirm?id="+strconv.Itoa(order.Id))

		} else {
			//非法请求
			Ctx.Redirect(302, "/")
		}
	} else {
		//非法请求
		Ctx.Redirect(302, "/")
	}

}

// 去结算
func (B *BuyController) Confirm(Ctx *gin.Context) {
	B.BaseInit(Ctx)
	idtemp := Ctx.Query("id")
	id, err := strconv.Atoi(idtemp)
	if err != nil {
		Ctx.Redirect(302, "/")
		return
	}
	//获取用户信息
	user := model.User{}
	model.Cookie.Get(Ctx, "userinfo", &user)

	//获取主订单信息
	order := model.Order{}
	dao.DB.Where("id=?", id).Find(&order)

	//判断当前数据是否合法
	if user.Id != order.Uid {
		Ctx.Redirect(302, "/")
		return
	}

	//获取主订单下面的商品信息
	orderItem := []model.OrderItem{}
	dao.DB.Where("order_id=?", id).Find(&orderItem)

	Ctx.HTML(200, "buy_confirm.html", gin.H{
		"order":       order,
		"orderItem":   orderItem,
		"topMenuList": B.TopMenu,
	})
}

// 获取订单状态
func (B *BuyController) OrderPayStatus(Ctx *gin.Context) {
	//1、获取订单号
	idtemp := Ctx.Query("id")
	id, err := strconv.Atoi(idtemp)
	if err != nil {
		Ctx.JSON(200, gin.H{
			"success": false,
			"message": "传入参数错误",
		})

		return
	}
	//2、查询订单
	order := model.Order{}
	dao.DB.Where("id=?", id).Find(&order)

	user := model.User{}
	model.Cookie.Get(Ctx, "userinfo", &user)
	//3、判断当前数据是否合法
	if user.Id != order.Uid {
		Ctx.JSON(200, gin.H{
			"success": false,
			"message": "传入参数错误",
		})
		return
	}

	//4、判断订单的支付状态
	if order.PayStatus == 1 && order.OrderStatus == 1 {
		Ctx.JSON(200, gin.H{
			"success": true,
			"message": "已支付",
		})

	} else {
		Ctx.JSON(200, gin.H{
			"success": false,
			"message": "未支付",
		})
	}
}
