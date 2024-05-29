package backend

import (
	"fmt"
	"goB2C/dao"
	"goB2C/model"
	"goB2C/util"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type AdministratorController struct {
	BaseController
}

func (c *AdministratorController) Get(Ctx *gin.Context) {
	administrator := []model.Administrator{}
	dao.DB.Preload("Role").Find(&administrator)

	Ctx.HTML(200, "administrator_index.html", gin.H{
		"administratorList": administrator,
	})
}

func (c *AdministratorController) Add(Ctx *gin.Context) {
	role := []model.Role{}
	dao.DB.Find(&role)
	Ctx.HTML(200, "administrator_add.html", gin.H{
		"roleList": role,
	})

}

func (c *AdministratorController) GoAdd(Ctx *gin.Context) {
	username := strings.Trim(Ctx.PostForm("username"), "")
	password := strings.Trim(Ctx.PostForm("password"), "")
	mobile := strings.Trim(Ctx.PostForm("mobile"), "")
	email := strings.Trim(Ctx.PostForm("email"), "")
	roleidtemp := Ctx.PostForm("role_id")
	roleId, err1 := strconv.Atoi(roleidtemp)
	if err1 != nil {
		c.Error(Ctx, "非法请求", "/administrator/add")
	}
	if len(username) < 2 || len(password) < 6 {
		c.Error(Ctx, "用户名或密码长度不合法", "/administrator/add")
		return
	} else if util.VerifyEmail(email) == false {
		c.Error(Ctx, "邮箱格式不正确，请重新填写!", "/administrator/add")
		return
	}
	administratorList := []model.Administrator{}
	dao.DB.Where("username=?", username).Find(&administratorList)
	if len(administratorList) > 0 {
		c.Error(Ctx, "用户名已存在", "/administrator/add")
		return
	}

	administrator := model.Administrator{}
	administrator.Username = username
	administrator.Password = util.Md5(password)
	administrator.Mobile = mobile
	administrator.Email = email
	administrator.Status = 1
	administrator.AddTime = int(util.GetUnix())
	administrator.RoleId = roleId
	err := dao.DB.Create(&administrator).Error
	if err != nil {
		c.Error(Ctx, "增加管理员失败", "/administrator/add")
		return
	}
	c.Success(Ctx, "增加管理员成功", "/administrator")
}

func (c *AdministratorController) Edit(Ctx *gin.Context) {
	idtemp := Ctx.Query("id")
	id, err := strconv.Atoi(idtemp)
	if err != nil {
		c.Error(Ctx, "传入参数错误", "/administrator")
		return
	}
	administrator := model.Administrator{Id: id}
	dao.DB.Find(&administrator)
	role := []model.Role{}
	dao.DB.Find(&role)

	Ctx.HTML(200, "administrator_edit.html", gin.H{
		"administrator": administrator,
		"roleList":      role,
	})
}

func (c *AdministratorController) GoEdit(Ctx *gin.Context) {
	idtemp := Ctx.PostForm("id")
	id, err := strconv.Atoi(idtemp)
	if err != nil {
		fmt.Println(id)
		fmt.Println(err)
		c.Error(Ctx, "传入参数错误", "/administrator")
		return
	}
	username := strings.Trim(Ctx.PostForm("Username"), "")
	password := strings.Trim(Ctx.PostForm("Password"), "")
	mobile := strings.Trim(Ctx.PostForm("Mobile"), "")
	email := strings.Trim(Ctx.PostForm("Email"), "")
	roleidtemp := Ctx.PostForm("role_id")
	roleId, err1 := strconv.Atoi(roleidtemp)
	if err1 != nil {
		c.Error(Ctx, "非法请求", "/administrator")
		return
	}
	if password != "" {
		if len(password) < 6 {
			c.Error(Ctx, "密码长度不合法！", "/administrator/add?id="+strconv.Itoa(id))
			return
		} else if util.VerifyEmail(email) == false {
			c.Error(Ctx, "邮箱格式不正确，请重新填写!", "/administrator/add?id="+strconv.Itoa(id))
			return
		}
		password = util.Md5(password)
	}
	administrator := model.Administrator{Id: id}
	dao.DB.Find(&administrator)
	administrator.Username = username
	administrator.Password = password
	administrator.Mobile = mobile
	administrator.Email = email
	administrator.RoleId = roleId
	err2 := dao.DB.Save(&administrator).Error
	if err2 != nil {
		c.Error(Ctx, "修改管理员失败", "/administrator/edit?id="+strconv.Itoa(id))
	} else {
		c.Success(Ctx, "修改管理员成功", "/administrator")
	}
}

func (c *AdministratorController) Delete(Ctx *gin.Context) {
	idtemp := Ctx.Query("id")
	id, err := strconv.Atoi(idtemp)
	if err != nil {
		c.Error(Ctx, "传入参数错误", "/role")
		return
	}
	administrator := model.Administrator{Id: id}
	dao.DB.Delete(&administrator)
	c.Success(Ctx, "删除管理员成功", "/administrator")
}
