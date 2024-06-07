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

	// var l backend.LoginController
	// r.GET("/backend/auth/login", l.Login)
	// r.Use(util.FrontendAuth)

	var u frontend.UserController
	userGroup := r.Group("/user")
	userGroup.Use(util.FrontendAuth)
	{
		userGroup.GET("/", u.Get)
		userGroup.GET("/order", u.OrderList)
		userGroup.GET("/orderinfo", u.OrderInfo)
		userGroup.POST("/userInfo", u.PassUserinfo)
		userGroup.POST("/userInfo/add", u.PassUserinfoAdd)
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

		//部门管理
		var rolecon backend.RoleController
		mainPath.GET("/role", rolecon.Get)
		mainPath.GET("/role/add", rolecon.Add)
		mainPath.POST("/role/goadd", rolecon.GoAdd)
		mainPath.GET("/role/edit", rolecon.Edit)
		mainPath.POST("/role/goedit", rolecon.GoEdit)
		mainPath.GET("/role/delete", rolecon.Delete)
		mainPath.GET("/role/auth", rolecon.Auth)
		mainPath.POST("/role/goauth", rolecon.GoAuth)

		//轮播图管理
		var ban backend.BannerController
		mainPath.GET("/banner", ban.Get)
		mainPath.GET("/banner/add", ban.Add)
		mainPath.GET("/banner/edit", ban.Edit)
		mainPath.POST("/banner/goadd", ban.GoAdd)
		mainPath.POST("/banner/goedit", ban.GoEdit)
		mainPath.GET("/banner/delete", ban.Delete)

		//导航管理
		var menu backend.MenuController
		mainPath.GET("/menu", menu.Get)
		mainPath.GET("/menu/add", menu.Add)
		mainPath.GET("/menu/edit", menu.Edit)
		mainPath.POST("/menu/goadd", menu.GoAdd)
		mainPath.POST("/menu/goedit", menu.GoEdit)
		mainPath.GET("/menu/delete", menu.Delete)

		//订单管理
		var order backend.OrderController
		mainPath.GET("/order", order.Get)
		mainPath.GET("/order/detail", order.Detail)
		mainPath.GET("/order/edit", order.Edit)
		mainPath.POST("/order/goEdit", order.GoEdit)
		mainPath.GET("/order/delete", order.Delete)

		//商品分类
		var pro backend.ProductCateController
		mainPath.GET("/productCate", pro.Get)
		mainPath.GET("/productCate/add", pro.Add)
		mainPath.GET("/productCate/edit", pro.Edit)
		mainPath.POST("/productCate/goadd", pro.GoAdd)
		mainPath.POST("/productCate/goedit", pro.GoEdit)
		mainPath.GET("/productCate/delete", pro.Delete)

		//商品类型属性
		var protypeatt backend.ProductTypeAttrController
		mainPath.GET("/productTypeAttribute", protypeatt.Get)
		mainPath.GET("/productTypeAttribute/add", protypeatt.Add)
		mainPath.GET("/productTypeAttribute/edit", protypeatt.Edit)
		mainPath.POST("/productTypeAttribute/goadd", protypeatt.GoAdd)
		mainPath.POST("/productTypeAttribute/goedit", protypeatt.GoEdit)
		mainPath.GET("/productTypeAttribute/delete", protypeatt.Delete)

		//商品类型管理
		var protype backend.ProductTypeController
		mainPath.GET("/productType", protype.Get)
		mainPath.GET("/productType/add", protype.Add)
		mainPath.GET("/productType/edit", protype.Edit)
		mainPath.POST("/productType/goadd", protype.GoAdd)
		mainPath.POST("/productType/goedit", protype.GoEdit)
		mainPath.GET("/productType/delete", protype.Delete)

		//商品管理
		var product backend.ProductController
		mainPath.GET("/product", product.Get)
		mainPath.GET("/product/add", product.Add)
		mainPath.GET("/product/edit", product.Edit)
		mainPath.POST("/product/goadd", product.GoAdd)
		mainPath.POST("/product/goedit", product.GoEdit)
		mainPath.GET("/product/delete", product.Delete)
		mainPath.POST("/product/goUpload", product.GoUpload)
		mainPath.GET("/product/getProductTypeAttribute", product.GetProductTypeAttribute)
		mainPath.GET("/product/changeProductImageColor", product.ChangeProductImageColor)
		mainPath.GET("/product/removeProductImage", product.RemoveProductImage)

		//系统设置
		var setting backend.SettingController
		mainPath.GET("/setting", setting.Get)
		mainPath.POST("/setting/goedit", setting.GoEdit)
	}
}
