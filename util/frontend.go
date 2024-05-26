package util

import (
	"goB2C/model"

	"github.com/gin-gonic/gin"
)

func FrontendAuth(Ctx *gin.Context) {
	//前台用户有没有登陆
	user := model.User{}
	model.Cookie.Get(Ctx, "userinfo", &user)
	if len(user.Phone) != 11 {
		Ctx.Redirect(302, "/auth/login")
	}
}
