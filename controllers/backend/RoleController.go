package backend

import (
	"goB2C/dao"
	"goB2C/model"
	"goB2C/util"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	BaseController
}

func (c *RoleController) Get(Ctx *gin.Context) {
	role := []model.Role{}
	dao.DB.Find(&role)
	Ctx.HTML(200, "role_index.html", gin.H{
		"rolelist": role,
	})

}

func (c *RoleController) Add(Ctx *gin.Context) {
	Ctx.HTML(200, "role_add.html", gin.H{})
}

func (c *RoleController) GoAdd(Ctx *gin.Context) {
	title := strings.Trim(Ctx.PostForm("title"), "")
	description := strings.Trim(Ctx.PostForm("description"), "")
	if title == "" {
		c.Error(Ctx, "标题不能为空", "/role/add")
		return
	}
	roleList := []model.Role{}
	dao.DB.Where("title=?", title).Find(&roleList)
	if len(roleList) != 0 {
		c.Error(Ctx, "该部门已存在！", "/role/add")
		return
	}
	role := model.Role{}
	role.Title = title
	role.Description = description
	role.Status = 1
	role.AddTime = int(util.GetUnix())
	err := dao.DB.Create(&role).Error
	if err != nil {
		c.Error(Ctx, "增加部门失败", "/role/add")
	} else {
		c.Success(Ctx, "增加部门成功", "/role")
	}
}

func (c *RoleController) Edit(Ctx *gin.Context) {
	idtemp := Ctx.Query("id")
	id, err := strconv.Atoi(idtemp)
	if err != nil {
		c.Error(Ctx, "传入参数错误", "/role")
		return
	}
	role := model.Role{Id: id}
	dao.DB.Find(&role)
	Ctx.HTML(200, "role_edit.html", gin.H{
		"role": role,
	})

}

func (c *RoleController) GoEdit(Ctx *gin.Context) {
	title := strings.Trim(Ctx.PostForm("title"), "")
	description := strings.Trim(Ctx.PostForm("description"), "")
	idtemp := Ctx.PostForm("id")
	id, err := strconv.Atoi(idtemp)
	if err != nil {
		c.Error(Ctx, "传入参数错误", "/role")
		return
	}
	role := model.Role{Id: id}
	dao.DB.Find(&role)
	role.Title = title
	role.Description = description
	err2 := dao.DB.Save(&role).Error
	if err2 != nil {
		c.Error(Ctx, "修改部门失败", "/role/edit?id="+strconv.Itoa(id))
	} else {
		c.Success(Ctx, "修改部门成功", "/role")
	}
}

func (c *RoleController) Delete(Ctx *gin.Context) {
	idtemp := Ctx.Query("id")
	id, err := strconv.Atoi(idtemp)
	if err != nil {
		c.Error(Ctx, "传入参数错误", "/role")
		return
	}
	role := model.Role{Id: id}
	administrator := []model.Administrator{}
	roleAuth := model.RoleAuth{}
	dao.DB.Where("role_id=?", id).Delete(&roleAuth)
	dao.DB.Preload("Role").Where("role_id=?", id).Find(&administrator)
	if len(administrator) > 0 {
		c.Error(Ctx, "该部门还有未处理的员工，无法删除该部门", "/role")
		return
	}
	dao.DB.Delete(&role)
	c.Success(Ctx, "删除部门成功", "/role")
}

func (c *RoleController) Auth(Ctx *gin.Context) {
	idtemp := Ctx.Query("id")
	roleId, err := strconv.Atoi(idtemp)
	if err != nil {
		c.Error(Ctx, "传入参数错误", "/role")
		return
	}
	auth := []model.Auth{}
	dao.DB.Preload("AuthItem").Where("module_id=0").Find(&auth)
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
	Ctx.HTML(200, "role_auth.html", gin.H{
		"accessList": auth,
		"roleId":     roleId,
	})

}

func (c *RoleController) GoAuth(Ctx *gin.Context) {
	idtemp := Ctx.PostForm("id")
	roleId, err := strconv.Atoi(idtemp)
	if err != nil {
		c.Error(Ctx, "传入参数错误", "/role")
		return
	}

	authNode := Ctx.PostFormArray("auth_node")
	roleAuth := model.RoleAuth{}
	dao.DB.Where("role_id=?", roleId).Delete(&roleAuth)
	for _, v := range authNode {
		authId, _ := strconv.Atoi(v)
		roleAuth.AuthId = authId
		roleAuth.RoleId = roleId
		dao.DB.Create(&roleAuth)
	}
	c.Success(Ctx, "授权成功", "/role/auth?id="+strconv.Itoa(roleId))
}
