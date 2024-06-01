package util

import (
	"errors"
	"goB2C/dao"
	"goB2C/model"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

var TopPath string = "/admin"
var ExcludeAuthPath string = "/mainBack,/,/welcome,/login/loginout"

// 登录校验和鉴权
func BackendAuth(Ctx *gin.Context) {
	pathname := Ctx.Request.URL.String()
	userinfo := model.Administrator{}
	ok := model.Cookie.Get(Ctx, "adminUserinfo", &userinfo)
	if !(ok && userinfo.Username != "") { // 未登录 重定向到登陆
		if pathname != TopPath+"/login" &&
			pathname != TopPath+"/login/gologin" &&
			pathname != TopPath+"/login/verificode" { //verificode 做验证码校验的，暂时去掉
			Ctx.Redirect(302, TopPath+"/login")
		}
	} else {
		pathname = strings.Replace(pathname, TopPath, "", 1)
		urlPath, _ := url.Parse(pathname)
		//isSuper是超级管理员，不做权限校验
		//ExcludeAuthPath 里面的路径，做鉴权豁免
		if userinfo.IsSuper == 0 && !excludeAuthPath(string(urlPath.Path)) {
			roleId := userinfo.RoleId
			roleAuth := []model.RoleAuth{}
			dao.DB.Where("role_id=?", roleId).Find(&roleAuth)
			roleAuthMap := make(map[int]int)
			for _, v := range roleAuth {
				roleAuthMap[v.AuthId] = v.AuthId
			}
			auth := model.Auth{}
			dao.DB.Where("url=?", urlPath.Path).Find(&auth)
			if _, ok := roleAuthMap[auth.Id]; !ok {
				Ctx.HTML(500, "public_error.html", gin.H{
					"Redirect": Ctx.Request.Referer(),
					"Message":  "没有权限",
				})
				Ctx.AbortWithError(500, errors.New("没有权限"))
			}
		}
	}
}

// 检验路径权限
func excludeAuthPath(urlPath string) bool {
	excludeAuthPathSlice := strings.Split(ExcludeAuthPath, ",")
	for _, v := range excludeAuthPathSlice {
		if v == urlPath {
			return true
		}
	}
	return false
}
