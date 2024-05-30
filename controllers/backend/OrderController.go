package backend

import (
	"goB2C/dao"
	"goB2C/model"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	BaseController
}

func (c *OrderController) Get(Ctx *gin.Context) {
	pageTemp := Ctx.Query("page")
	page, _ := strconv.Atoi(pageTemp)
	if page == 0 {
		page = 1
	}
	pageSize := 10
	keyword := Ctx.Query("keyword")
	order := []model.Order{}
	var count int64
	if keyword == "" {
		dao.DB.Table("order").Count(&count)
		dao.DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&order)
	} else {
		dao.DB.Where("phone=?", keyword).Offset((page - 1) * pageSize).Limit(pageSize).Find(&order)
		dao.DB.Where("phone=?", keyword).Table("order").Count(&count)
	}
	// if len(order) == 0 {
	// 	prvPage := page - 1
	// 	if prvPage == 0 {
	// 		prvPage = 1
	// 	}
	// 	c.Goto("/order?page=" + strconv.Itoa(prvPage))
	// }

	Ctx.HTML(200, "order_order.html", gin.H{
		"totalPages": math.Ceil(float64(count) / float64(pageSize)),
		"page":       page,
		"order":      order,
	})
}

func (c *OrderController) Detail(Ctx *gin.Context) {
	idTemp := Ctx.Query("id")
	id, err := strconv.Atoi(idTemp)
	if err != nil {
		c.Error(Ctx, "传入参数错误", "/order")
		return
	}
	order := model.Order{}
	dao.DB.Where("id=?", id).Find(&order)
	Ctx.HTML(200, "order_info.html", gin.H{
		"order": order,
	})
}
func (c *OrderController) Edit(Ctx *gin.Context) {
	idTemp := Ctx.Query("id")
	id, err := strconv.Atoi(idTemp)
	if err != nil {
		c.Error(Ctx, "传入参数错误", "/order")
		return
	}
	order := model.Order{}
	dao.DB.Where("id=?", id).Find(&order)
	Ctx.HTML(200, "order_edit.html", gin.H{
		"order": order,
	})
}
func (c *OrderController) GoEdit(Ctx *gin.Context) {
	idtemp := Ctx.PostForm("id")
	id, err := strconv.Atoi(idtemp)
	if err != nil {
		c.Error(Ctx, "传入参数错误", "/order")
		return
	}
	orderId := Ctx.PostForm("order_id")
	allPrice := Ctx.PostForm("all_price")
	name := Ctx.PostForm("name")
	phone := Ctx.PostForm("phone")
	address := Ctx.PostForm("address")
	zipcode := Ctx.PostForm("zipcode")
	paystatustemp := Ctx.PostForm("pay_status")
	payStatus, _ := strconv.Atoi(paystatustemp)
	paytypetemp := Ctx.PostForm("pay_type")
	payType, _ := strconv.Atoi(paytypetemp)
	orderstatustemp := Ctx.PostForm("order_status")
	orderStatus, _ := strconv.Atoi(orderstatustemp)
	order := model.Order{}
	dao.DB.Where("id=?", id).Find(&order)
	order.OrderId = orderId
	order.AllPrice, _ = strconv.ParseFloat(allPrice, 64)
	order.Name = name
	order.Phone = phone
	order.Address = address
	order.Zipcode = zipcode
	order.PayStatus = payStatus
	order.PayType = payType
	order.OrderStatus = orderStatus
	dao.DB.Save(&order)
	c.Success(Ctx, "订单修改成功", "/order")
}
func (c *OrderController) Delete(Ctx *gin.Context) {
	idTemp := Ctx.Query("id")
	id, err := strconv.Atoi(idTemp)
	if err != nil {
		c.Error(Ctx, "传入参数错误", "/order")
		return
	}
	order := model.Order{}
	dao.DB.Where("id=?", id).Delete(&order)
	c.Success(Ctx, "删除订单记录成功", "/order")
}
