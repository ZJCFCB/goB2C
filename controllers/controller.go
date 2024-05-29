package controllers

import (
	"goB2C/controllers/backend"
	"goB2C/controllers/frontend"
	"goB2C/util"

	"github.com/gin-gonic/gin"
)

func RegistFunc(r *gin.Engine) {

	var c frontend.IndexController
	r.GET("/mainPage", c.MainPage)
	r.GET("/", c.MainPage)

	var a frontend.AuthController
	r.GET("/auth/registerStep1", a.RegisterStep1)
	r.GET("/auth/sendCode", a.SendCode)
	r.GET("/auth/registerStep2", a.RegisterStep2)
	r.GET("/auth/validateSmsCode", a.ValidateSmsCode)
	r.GET("/auth/registerStep3", a.RegisterStep3)
	r.POST("/auth/doRegister", a.GoRegister)
	r.GET("/auth/login", a.Login)
	r.POST("/auth/goLogin", a.GoLogin)
	r.GET("/auth/loginOut", a.LoginOut)

	var l backend.LoginController
	r.GET("/backend/auth/login", l.Login)
	r.Use(util.FrontendAuth)

	var u frontend.UserController
	userGroup := r.Group("/user")
	userGroup.Use(util.FrontendAuth)
	{
		userGroup.GET("/", u.Get)
		userGroup.GET("/order", u.OrderList)
		userGroup.GET("/orderinfo", u.OrderInfo)
	}

	var p frontend.ProductController
	r.GET("/category_:id([0-9]+).html", p.CategoryList)
	r.GET("/item_:id([0-9]+).html", p.ProductItem)
	r.GET("/seckill/item_:id([0-9]+).html", p.ProductItem)
	r.GET("/product/collect", p.Collect)
	r.GET("/product/getImgList", p.GetImgList)

	var addr frontend.AddressController
	addressGroup := r.Group("/address")
	addressGroup.Use(util.FrontendAuth)
	{
		addressGroup.POST("/addAddress", addr.AddAddress)
		addressGroup.GET("/getOneAddressList", addr.GetOneAddressList)
		addressGroup.POST("/goEditAddressList", addr.GoEditAddressList)
		addressGroup.GET("/changeDefaultAddress", addr.ChangeDefaultAddress)
	}

	var buy frontend.BuyController
	buyGroup := r.Group("/buy")
	buyGroup.Use(util.FrontendAuth)

	//store := cookie.NewStore([]byte("secret"))
	//r.Use(sessions.Sessions("mysession", store))

	{
		buyGroup.GET("/checkout", buy.Checkout)
		buyGroup.POST("/doOrder", buy.GoOrder)
		buyGroup.GET("/confirm", buy.Confirm)
		buyGroup.GET("/orderPayStatus", buy.OrderPayStatus)
	}

	var cart frontend.CartController
	r.GET("/cart", cart.Get)
	r.GET("/cart/addCart", cart.AddCart)
	r.GET("/cart/incCart", cart.IncCart)
	r.GET("/cart/decCart", cart.DecCart)
	r.GET("/cart/delCart", cart.DelCart)
	r.GET("/cart/changeOneCart", cart.ChangeOneCart)
	r.GET("/cart/changeAllCart", cart.ChangeAllCart)

	var pay frontend.PayController
	//支付宝支付
	r.GET("/alipay", pay.Alipay)
	r.POST("/alipayNotify", pay.AlipayNotify)
	r.GET("/alipayReturn", pay.AlipayReturn)

	//微信支付
	r.GET("/wxpay", pay.WxPay)
	r.POST("/wxpay/notify", pay.WxPayNotify)
}

func RegisterBackendFunc(r *gin.Engine) {

	mainPath := r.Group(util.TopPath)
	mainPath.Use(util.BackendAuth)
	{
		var login backend.LoginController
		mainPath.GET("/login", login.Login)
		// beego.NSRouter("/login/verificode", &backend.LoginController{}, "get:SetYzm"),
		mainPath.POST("/login/gologin", login.GoLogin)
		mainPath.GET("/login/loginout", login.LoginOut)

		var maincon backend.MainController
		mainPath.GET("/", maincon.Get)
		mainPath.GET("/mainBack", maincon.Get)
		mainPath.GET("/welcome", maincon.Welcome)
		mainPath.GET("/main/changestatus", maincon.ChangeStatus)
		mainPath.GET("/main/editnum", maincon.EditNum)

		//adm
		var adm backend.AdministratorController
		mainPath.GET("/administrator", adm.Get)
		mainPath.GET("/administrator/add", adm.Add)
		mainPath.GET("/administrator/edit", adm.Edit)
		mainPath.POST("/administrator/goadd", adm.GoAdd)
		mainPath.POST("/administrator/goedit", adm.GoEdit)
		mainPath.GET("/administrator/delete", adm.Delete)

		//权限管理
		var auth backend.AuthController
		mainPath.GET("/auth", auth.Get)
		mainPath.GET("/auth/add", auth.Add)
		mainPath.GET("/auth/edit", auth.Edit)
		mainPath.POST("/auth/goadd", auth.GoAdd)
		mainPath.POST("/auth/goedit", auth.GoEdit)
		mainPath.GET("/auth/delete", auth.Delete)

		//轮播图管理
		var ban backend.BannerController
		mainPath.GET("/banner", ban.Get)
		mainPath.GET("/banner/add", ban.Add)
		mainPath.GET("/banner/edit", ban.Edit)
		mainPath.POST("/banner/goadd", ban.GoAdd)
		mainPath.POST("/banner/goedit", ban.GoEdit)
		mainPath.GET("/banner/delete", ban.Delete)
	}
}
