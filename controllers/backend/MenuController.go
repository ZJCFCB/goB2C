package backend

import (
	"goB2C/dao"
	"goB2C/model"
	"goB2C/util"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MenuController struct {
	BaseController
}

func (c *MenuController) Get(Ctx *gin.Context) {
	//当前页
	pageTemp := Ctx.Query("page")
	page, _ := strconv.Atoi(pageTemp)
	if page == 0 {
		page = 1
	}
	//每一页显示的数量
	pageSize := 5
	//查询数据
	menu := []model.Menu{}
	dao.DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&menu)
	//查询menu表里面的数量
	var count int64
	dao.DB.Table("menu").Count(&count)
	// if len(menu) == 0 {
	// 	prvPage := page - 1
	// 	if prvPage == 0 {
	// 		prvPage = 1
	// 	}
	// 	c.Goto("/menu?page=" + strconv.Itoa(prvPage))
	// }
	Ctx.HTML(200, "menu_index.html", gin.H{
		"menuList":   menu,
		"totalPages": math.Ceil(float64(count) / float64(pageSize)),
		"page":       page,
	})

}

func (c *MenuController) Add(Ctx *gin.Context) {
	Ctx.HTML(200, "menu_add.html", gin.H{})
}

func (c *MenuController) GoAdd(Ctx *gin.Context) {
	title := Ctx.PostForm("title")
	link := Ctx.PostForm("link")
	positionTemp := Ctx.PostForm("position")
	position, _ := strconv.Atoi(positionTemp)
	isOpennewTemp := Ctx.PostForm("is_opennew")
	isOpennew, _ := strconv.Atoi(isOpennewTemp)
	relation := Ctx.PostForm("relation")
	sortTemp := Ctx.PostForm("sort")
	sort, _ := strconv.Atoi(sortTemp)
	statusTemp := Ctx.PostForm("status")
	status, _ := strconv.Atoi(statusTemp)

	menu := model.Menu{
		Title:     title,
		Link:      link,
		Position:  position,
		IsOpennew: isOpennew,
		Relation:  relation,
		Sort:      sort,
		Status:    status,
		AddTime:   int(util.GetUnix()),
	}

	err := dao.DB.Create(&menu).Error
	if err != nil {
		c.Error(Ctx, "增加数据失败", "/menu/add")
	} else {
		dao.RedisDel("middleMenu")
		dao.RedisDel("topMenu")
		c.Success(Ctx, "增加成功", "/menu")
	}
}

func (c *MenuController) Edit(Ctx *gin.Context) {
	idTemp := Ctx.Query("id")
	id, err := strconv.Atoi(idTemp)
	if err != nil {
		c.Error(Ctx, "传入参数错误", "/menu")
		return
	}
	menu := model.Menu{Id: id}
	dao.DB.Find(&menu)
	Ctx.HTML(200, "menu_edit.html", gin.H{
		"menu":     menu,
		"prevPage": Ctx.Request.Referer(),
	})
}

func (c *MenuController) GoEdit(Ctx *gin.Context) {

	idTemp := Ctx.PostForm("id")
	id, err1 := strconv.Atoi(idTemp)
	if err1 != nil {
		c.Error(Ctx, "传入参数错误", "/menu")
		return
	}
	title := Ctx.PostForm("title")
	link := Ctx.PostForm("link")
	positionTemp := Ctx.PostForm("position")
	position, _ := strconv.Atoi(positionTemp)
	isOpennewTemp := Ctx.PostForm("is_opennew")
	isOpennew, _ := strconv.Atoi(isOpennewTemp)
	relation := Ctx.PostForm("relation")
	sortTemp := Ctx.PostForm("sort")
	sort, _ := strconv.Atoi(sortTemp)
	statusTemp := Ctx.PostForm("status")
	status, _ := strconv.Atoi(statusTemp)
	prevPage := Ctx.PostForm("prevPage")

	//修改
	menu := model.Menu{Id: id}
	dao.DB.Find(&menu)
	menu.Title = title
	menu.Link = link
	menu.Position = position
	menu.IsOpennew = isOpennew
	menu.Relation = relation
	menu.Sort = sort
	menu.Status = status

	err2 := dao.DB.Save(&menu).Error
	if err2 != nil {
		c.Error(Ctx, "修改数据失败", "/menu/edit?id="+strconv.Itoa(id))
	} else {
		dao.RedisDel("middleMenu")
		dao.RedisDel("topMenu")
		c.Success(Ctx, "修改数据成功", prevPage)
	}

}

func (c *MenuController) Delete(Ctx *gin.Context) {
	idTemp := Ctx.Query("id")
	id, err := strconv.Atoi(idTemp)
	if err != nil {
		c.Error(Ctx, "传入参数错误", "/menu")
		return
	}
	menu := model.Menu{Id: id}
	dao.DB.Delete(&menu)
	dao.RedisDel("middleMenu")
	dao.RedisDel("topMenu")
	c.Success(Ctx, "删除数据成功", Ctx.Request.Referer())
}
