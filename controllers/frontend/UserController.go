package frontend

import (
	"goB2C/dao"
	"goB2C/model"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

func (U *UserController) Get(Ctx *gin.Context) {
	U.BaseInit(Ctx)
	user := model.User{}
	model.Cookie.Get(Ctx, "userinfo", &user)

	time := time.Now().Hour()
	var hello string
	if time >= 12 && time <= 18 {
		hello = "尊敬的用户下午好"
	} else if time >= 6 && time < 12 {
		hello = "尊敬的用户上午好！"
	} else {
		hello = "深夜了，注意休息哦！"
	}

	order := []model.Order{} //用户订单
	dao.DB.Where("uid=?", user.Id).Find(&order)
	var wait_pay int
	var wait_rec int
	for i := 0; i < len(order); i++ {
		if order[i].PayStatus == 0 {
			wait_pay += 1
		}
		if order[i].OrderStatus >= 2 && order[i].OrderStatus < 4 {
			wait_rec += 1
		}
	}

	Ctx.HTML(200, "user_welcome.html", gin.H{
		"user":            user,
		"Hello":           hello,
		"wait_pay":        wait_pay,
		"wait_rec":        wait_rec,
		"userinfo":        U.UserInfo,
		"topMenuList":     U.TopMenu,
		"productCateList": U.ProductCate,
		"middleMenuList":  U.MiddleMenu,
	})
}

func (c *UserController) OrderList(Ctx *gin.Context) {
	//1、获取当前用户
	user := model.User{}
	model.Cookie.Get(Ctx, "userinfo", &user)
	//2、获取当前用户下面的订单信息 并分页
	tempPage := Ctx.Query("page")
	page, _ := strconv.Atoi(tempPage)
	if page == 0 {
		page = 1
	}
	pageSize := 5
	//3、获取搜索关键词
	where := "uid=?"
	keywords := Ctx.Query("keywords")
	if keywords != "" {
		orderitem := []model.OrderItem{}
		dao.DB.Where("product_title like ?", "%"+keywords+"%").Find(&orderitem)
		var str string
		for i := 0; i < len(orderitem); i++ {
			if i == 0 {
				str += orderitem[i].OrderID
			} else {
				str += "," + orderitem[i].OrderID
			}
		}
		where += " AND id in (" + str + ")"
	}
	//获取筛选条件
	tempOrderStatus := Ctx.Query("order_status")
	orderStatus, err := strconv.Atoi(tempOrderStatus)
	if err != nil {
		orderStatus = 0
	}
	//3、总数量
	var count int64
	dao.DB.Where("uid=?", user.Id).Table("order").Count(&count)
	order := []model.Order{}
	dao.DB.Where("uid=?", user.Id).Offset((page - 1) * pageSize).Limit(pageSize).Preload("OrderItems").Order("add_time desc").Find(&order)

	Ctx.HTML(200, "user_order.html", gin.H{
		"order":       order,
		"totalPages":  math.Ceil(float64(count) / float64(pageSize)),
		"page":        page,
		"keywords":    keywords,
		"orderStatus": orderStatus,
		"userinfo":    GetUserInfo(user),
		"topMenuList": GetTopMenuList(),
	})
}
func (c *UserController) OrderInfo(Ctx *gin.Context) {
	tempId := Ctx.Query("id")
	id, _ := strconv.Atoi(tempId)
	user := model.User{}
	model.Cookie.Get(Ctx, "userinfo", &user)
	order := model.Order{}
	dao.DB.Where("id=? AND uid=?", id, user.Id).Preload("OrderItems").Find(&order)

	if order.OrderID == "" {
		Ctx.Redirect(http.StatusFound, "/mainPage")
	}
	Ctx.HTML(200, "user_order_info.html", gin.H{
		"order":           order,
		"userinfo":        GetUserInfo(user),
		"productCateList": GetProductCateList(),
		"topMenuList":     GetTopMenuList(),
	})
}
