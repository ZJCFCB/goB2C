package backend

import (
	"goB2C/dao"
	"goB2C/model"
	"strings"

	"github.com/gin-gonic/gin"
)

type SettingController struct {
	BaseController
}

func (c *SettingController) Get(Ctx *gin.Context) {
	setting := model.Setting{}
	dao.DB.Order("id desc").First(&setting)

	Ctx.HTML(200, "setting_index.html", gin.H{
		"setting": setting,
	})
}

func (c *SettingController) GoEdit(Ctx *gin.Context) {
	var setting1 model.Setting
	dao.DB.Order("id desc").First(&setting1)

	var setting2 model.Setting
	Ctx.ShouldBind(&setting2)

	setting2.Ip = strings.Split(Ctx.Request.RemoteAddr, ":")[0]
	siteLogo, err := c.UploadImg(Ctx, "site_logo")
	if len(siteLogo) > 0 && err == nil {
		setting2.SiteLogo = siteLogo
	} else {
		setting2.SiteLogo = setting1.SiteLogo
	}
	noPicture, err := c.UploadImg(Ctx, "no_picture")
	if len(noPicture) > 0 && err == nil {
		setting2.NoPicture = noPicture
	} else {
		setting2.NoPicture = setting1.NoPicture
	}
	err = dao.DB.Save(&setting2).Error
	if err != nil {
		c.Error(Ctx, "修改数据失败", "/setting")
		return
	}
	c.Success(Ctx, "修改数据成功", "/setting")
}
