package frontend

import (
	"goB2C/dao"
	"goB2C/model"
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
	Ctx.Set("user", user)

	time := time.Now().Hour()
	if time >= 12 && time <= 18 {
		Ctx.Set("Hello", "尊敬的用户下午好")
	} else if time >= 6 && time < 12 {
		Ctx.Set("Hello", "尊敬的用户上午好！")
	} else {
		Ctx.Set("Hello", "深夜了，注意休息哦！")
	}

	order := []model.Order{} //这个应该是用户订单
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
	Ctx.Set("wait_pay", wait_pay)
	Ctx.Set("wait_rec", wait_rec)
	//c.TplName = "frontend/user/welcome.html"
}
