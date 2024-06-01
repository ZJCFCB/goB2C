package backend

import (
	"goB2C/dao"
	"goB2C/model"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MainController struct {
	BaseController
}

func (M *MainController) Get(Ctx *gin.Context) {
	userinfo := model.Administrator{}
	ok := model.Cookie.Get(Ctx, "adminUserinfo", &userinfo)
	if ok {
		roleId := userinfo.RoleId
		auth := []model.Auth{}
		dao.DB.Preload("AuthItem", func(db *gorm.DB) *gorm.DB {
			return db.Order("auth.sort DESC")
		}).Order("sort desc").Where("module_id=?", 0).Find(&auth)
		//获取当前部门拥有的权限，并把权限ID放在一个MAP对象里面
		roleAuth := []model.RoleAuth{}
		dao.DB.Where("role_id=?", roleId).Find(&roleAuth)
		roleAuthMap := make(map[int]int)
		for _, v := range roleAuth {
			roleAuthMap[v.AuthId] = v.AuthId
		}
		for i := 0; i < len(auth); i++ {
			if _, ok := roleAuthMap[auth[i].Id]; ok {
				auth[i].Checked = true
			}
			for j := 0; j < len(auth[i].AuthItem); j++ {
				if _, ok := roleAuthMap[auth[i].AuthItem[j].Id]; ok {
					auth[i].AuthItem[j].Checked = true
				}
			}
		}
		Ctx.HTML(200, "main_index.html", gin.H{
			"username": userinfo.Username,
			"authList": auth,
			"isSuper":  userinfo.IsSuper,
		})
		return
	} else {
		M.Error(Ctx, "请先登录", "/login")
		return
	}

}

func (L *MainController) Welcome(Ctx *gin.Context) {
	Ctx.HTML(200, "main_welcome.html", gin.H{})
}

// 修改公共状态
func (M *MainController) ChangeStatus(Ctx *gin.Context) {
	idtemp := Ctx.Query("id")
	id, err := strconv.Atoi(idtemp)
	if err != nil {
		Ctx.JSON(200, gin.H{
			"success": false,
			"msg":     "非法请求",
		})
		return
	}
	table := Ctx.Query("table")
	field := Ctx.Query("field")
	err1 := dao.DB.Exec("update "+table+" set "+field+"=ABS("+field+"-1) where id=?", id).Error
	if err1 != nil {
		Ctx.JSON(200, gin.H{
			"success": false,
			"msg":     "更新数据失败",
		})

		return
	}
	Ctx.JSON(200, gin.H{
		"success": true,
		"msg":     "更新数据成功",
	})
}

func (M *MainController) EditNum(Ctx *gin.Context) {
	id := Ctx.Query("id")
	table := Ctx.Query("table")
	field := Ctx.Query("field")
	num := Ctx.Query("num")
	err1 := dao.DB.Exec("update " + table + " set " + field + "=" + num + " where id=" + id).Error
	if err1 != nil {
		Ctx.JSON(200, gin.H{
			"success": false,
			"msg":     "修改数量失败",
		})

		return
	}
	Ctx.JSON(200, gin.H{
		"success": true,
		"msg":     "修改数量成功",
	})

}
