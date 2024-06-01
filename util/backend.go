package util

import (
	"goB2C/dao"
	"goB2C/model"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

var TopPath string = "/admin"

// 后台权限判断
func BackendAuth(Ctx *gin.Context) {
	pathname := Ctx.Request.URL.String()
	userinfo := model.Administrator{}
	ok := model.Cookie.Get(Ctx, "adminUserinfo", &userinfo)
	if !(ok && userinfo.Username != "") { // 未登录 重定向到登陆
		if pathname != TopPath+"/login" &&
			pathname != TopPath+"/login/gologin" &&
			pathname != TopPath+"/login/verificode" {
			Ctx.Redirect(302, TopPath+"/login")
		}
	} else {
		pathname = strings.Replace(pathname, TopPath, "", 1)
		urlPath, _ := url.Parse(pathname)
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
			if _, ok := roleAuthMap[auth.Id]; !ok && ok { // 先不做鉴权
				Ctx.Writer.WriteString("没有权限")
				return
			}
		}
	}
}

// 检验路径权限
func excludeAuthPath(urlPath string) bool {
	excludeAuthPathSlice := strings.Split(TopPath, ",")
	for _, v := range excludeAuthPathSlice {
		if v == urlPath {
			return true
		}
	}
	return false
}
