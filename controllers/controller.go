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
}
