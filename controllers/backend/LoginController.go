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

func (c *LoginController) Login(Ctx *gin.Context) {
	Ctx.HTML(200, "login_login.html", gin.H{})
}

func (c *LoginController) GoLogin(Ctx *gin.Context) {
	//var flag = model.Cap.VerifyReq(Ctx.Request)
	var flag = true
	if flag {
		username := strings.Trim(Ctx.GetString("username"), "")
		password := util.Md5(strings.Trim(Ctx.GetString("password"), ""))
		administrator := []model.Administrator{}
		dao.DB.Where("username=? AND password=? AND status=1", username, password).Find(&administrator)
		if len(administrator) == 1 {
			//c.SetSession("userinfo", administrator[0])
			c.Success(Ctx, "登陆成功", "/")
		} else {
			c.Error(Ctx, "无登陆权限或用户名密码错误", "/login")
		}
	} else {
		c.Error(Ctx, "验证码错误", "/login")
	}
}

func (c *LoginController) LoginOut(Ctx *gin.Context) {
	c.Success(Ctx, "退出登录成功,将返回登陆页面！", "/login")
}
