package frontend

import (
	"goB2C/model"
	"net/url"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
	TopMenu     []model.Menu
	MiddleMenu  []model.Menu
	ProductCate []model.ProductCate
	UserInfo    string
	PathName    *url.URL
}

func (B *BaseController) BaseInit(Ctx *gin.Context) {

	//顶部分类
	B.TopMenu = GetTopMenuList()

	//Todo
	//获取轮播图，这里还有bug，一时半会不知道怎么改

	//左侧分类（预加载）
	B.ProductCate = GetProductCateList()

	//获取中间导航的数据
	B.MiddleMenu = GetMiddleList()
	//判断用户是否登录

	//在top list中，如果判断用户是已经登录状态（cookie），那么显示的就是登录的前端页面
	//如果判断是非登录状态，显示的就是登录、注册页面
	user := model.User{}
	model.Cookie.Get(Ctx, "userinfo", &user)
	B.UserInfo = GetUserInfo(user)

	//把path也放进去
	urlpath, _ := url.Parse(Ctx.Copy().Request.URL.String())
	B.PathName = urlpath
}
