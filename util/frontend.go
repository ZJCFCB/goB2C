package util

import (
	"goB2C/model"

	"github.com/gin-gonic/gin"
)

func FrontendAuth(Ctx *gin.Context) {
	//前台用户有没有登陆
	user := model.User{}
	ok := model.Cookie.Get(Ctx, "userinfo", &user)
	if ok == false {
		Ctx.Redirect(302, "/auth/login")
	}
}
