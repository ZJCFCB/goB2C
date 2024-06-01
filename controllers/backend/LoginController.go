package backend

import (
	"goB2C/dao"
	"goB2C/model"
	"goB2C/util"
	"strings"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	BaseController
}

func (L *LoginController) Login(Ctx *gin.Context) {
	Ctx.HTML(200, "login_login.html", gin.H{
		"getadminPath": util.TopPath,
	})
}

func (L *LoginController) GoLogin(Ctx *gin.Context) {
	//var flag = model.Cap.VerifyReq(Ctx.Request)
	var flag = true
	if flag {
		username := strings.Trim(Ctx.PostForm("username"), "")
		password := util.Md5(strings.Trim(Ctx.PostForm("password"), ""))
		administrator := []model.Administrator{}
		dao.DB.Where("username=? AND password=? AND status=1", username, password).Find(&administrator)
		if len(administrator) == 1 {
			userinfo := model.Administrator{}
			userinfo.Username = administrator[0].Username
			userinfo.Password = administrator[0].Password
			userinfo.RoleId = administrator[0].RoleId
			userinfo.IsSuper = administrator[0].IsSuper

			model.Cookie.Set(Ctx, "adminUserinfo", userinfo)
			Ctx.Redirect(302, util.TopPath+"/mainBack")
			L.Success(Ctx, "登陆成功", "")

		} else {
			L.Error(Ctx, "无登陆权限或用户名密码错误", "/login")
		}
	} else {
		L.Error(Ctx, "验证码错误", "/login")
	}
}

func (c *LoginController) LoginOut(Ctx *gin.Context) {
	model.Cookie.Remove(Ctx, "adminUserinfo", "")
	c.Success(Ctx, "退出登录成功,将返回登陆页面！", "/login")
}
