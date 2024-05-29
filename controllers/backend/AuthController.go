package backend

import (
	"fmt"
	"goB2C/dao"
	"goB2C/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	BaseController
}

func (c *AuthController) Get(Ctx *gin.Context) {
	auth := []model.Auth{}
	dao.DB.Preload("AuthItem").Where("module_id=0").Find(&auth)
	Ctx.HTML(200, "auth_index.html", gin.H{
		"authList": auth,
	})
}

func (c *AuthController) Add(Ctx *gin.Context) {
	auth := []model.Auth{}
	dao.DB.Where("module_id=0").Find(&auth)
	Ctx.HTML(200, "auth_add.html", gin.H{
		"authList": auth,
	})
}

func (c *AuthController) GoAdd(Ctx *gin.Context) {
	moduleName := Ctx.PostForm("module_name")

	TypeTemp := Ctx.PostForm("type")
	iType, err1 := strconv.Atoi(TypeTemp)

	actionName := Ctx.PostForm("action_name")
	url := Ctx.PostForm("url")

	moduleIdtemp := Ctx.PostForm("module_id")
	moduleId, err2 := strconv.Atoi(moduleIdtemp)

	sorttemp := Ctx.PostForm("sort")
	sort, err3 := strconv.Atoi(sorttemp)

	description := Ctx.PostForm("description")

	statustemp := Ctx.PostForm("status")
	status, err4 := strconv.Atoi(statustemp)

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		fmt.Println("err1", err1, "err2", err2, "err3", err3, "err4", err4)
		c.Error(Ctx, "传入参数错误", "/auth/add")
		return
	}
	auth := model.Auth{
		ModuleName:  moduleName,
		Type:        iType,
		ActionName:  actionName,
		Url:         url,
		ModuleId:    moduleId,
		Sort:        sort,
		Description: description,
		Status:      status,
	}
	err := dao.DB.Create(&auth).Error
	if err != nil {
		c.Error(Ctx, "增加数据失败", "/auth/add")
		return
	}
	c.Success(Ctx, "增加数据成功", "/auth")
}

func (c *AuthController) Edit(Ctx *gin.Context) {
	idtemp := Ctx.Query("id")
	id, err := strconv.Atoi(idtemp)
	if err != nil {
		c.Error(Ctx, "传入参数错误", "/auth")
		return
	}
	auth := model.Auth{Id: id}
	dao.DB.Find(&auth)
	authList := []model.Auth{}
	dao.DB.Where("module_id=0").Find(&authList)

	Ctx.HTML(200, "auth_edit.html", gin.H{
		"auth":     auth,
		"authList": authList,
	})
}

func (c *AuthController) GoEdit(Ctx *gin.Context) {
	moduleName := Ctx.PostForm("module_name")

	TypeTemp := Ctx.PostForm("type")
	iType, err1 := strconv.Atoi(TypeTemp)

	actionName := Ctx.PostForm("action_name")
	url := Ctx.PostForm("url")

	moduleIdtemp := Ctx.PostForm("module_id")
	moduleId, err2 := strconv.Atoi(moduleIdtemp)

	sorttemp := Ctx.PostForm("sort")
	sort, err3 := strconv.Atoi(sorttemp)

	description := Ctx.PostForm("description")

	statustemp := Ctx.PostForm("status")
	status, err4 := strconv.Atoi(statustemp)

	idtemp := Ctx.PostForm("id")
	id, err5 := strconv.Atoi(idtemp)
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil {
		c.Error(Ctx, "传入参数错误", "/auth")
		return
	}
	auth := model.Auth{Id: id}
	dao.DB.Find(&auth)
	auth.ModuleName = moduleName
	auth.Type = iType
	auth.ActionName = actionName
	auth.Url = url
	auth.ModuleId = moduleId
	auth.Sort = sort
	auth.Description = description
	auth.Status = status
	err6 := dao.DB.Save(&auth).Error
	if err6 != nil {
		c.Error(Ctx, "修改权限失败", "/auth/edit?id="+strconv.Itoa(id))
		return
	}
	c.Success(Ctx, "修改权限成功", "/auth")
}

func (c *AuthController) Delete(Ctx *gin.Context) {
	idtemp := Ctx.Query("id")
	id, err := strconv.Atoi(idtemp)
	if err != nil {
		c.Error(Ctx, "传入参数错误", "/role")
		return
	}
	auth := model.Auth{Id: id}
	dao.DB.Find(&auth)
	if auth.ModuleId == 0 {
		auth2 := []model.Auth{}
		dao.DB.Where("module_id=?", auth.Id).Find(&auth2)
		if len(auth2) > 0 {
			c.Error(Ctx, "请删除当前顶级模块下面的菜单或操作！", "/auth")
			return
		}
	}
	dao.DB.Delete(&auth)
	c.Success(Ctx, "删除成功", "/auth")
}
