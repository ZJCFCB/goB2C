package frontend

import (
	"goB2C/dao"
	"goB2C/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AddressController struct {
	BaseController
}

func (c *AddressController) AddAddress(Ctx *gin.Context) {
	user := model.User{}
	model.Cookie.Get(Ctx, "userinfo", &user)
	name := Ctx.PostForm("name")
	phone := Ctx.PostForm("phone")
	address := Ctx.PostForm("address")
	zipcode := Ctx.PostForm("zipcode")
	var addressCount int64
	dao.DB.Where("uid=?", user.Id).Table("address").Count(&addressCount)
	if addressCount > 10 {
		Ctx.JSON(200, gin.H{
			"success": false,
			"message": "增加收货地址失败，收货地址数量超过限制",
		})
		return
	}
	dao.DB.Table("address").Where("uid=?", user.Id).Updates(map[string]interface{}{"default_address": 0})
	addressResult := model.Address{
		Uid:            user.Id,
		Name:           name,
		Phone:          phone,
		Address:        address,
		Zipcode:        zipcode,
		DefaultAddress: 1,
	}
	dao.DB.Create(&addressResult)
	allAddressResult := []model.Address{}
	dao.DB.Where("uid=?", user.Id).Find(&allAddressResult)
	Ctx.JSON(200, gin.H{
		"success": true,
		"result":  allAddressResult,
	})
}

// 点击以后，获取某一个id的地址信息
func (c *AddressController) GetOneAddressList(Ctx *gin.Context) {
	addressIdtemp := Ctx.Query("address_id")
	addressId, err := strconv.Atoi(addressIdtemp)
	if err != nil {
		Ctx.JSON(200, gin.H{
			"success": false,
			"message": "传入参数错误",
		})
		return
	}
	address := model.Address{}
	dao.DB.Where("id=?", addressId).Find(&address)
	Ctx.JSON(200, gin.H{
		"success": true,
		"result":  address,
	})
	return
}

// 修改地址
func (c *AddressController) GoEditAddressList(Ctx *gin.Context) {
	user := model.User{}
	model.Cookie.Get(Ctx, "userinfo", &user)
	addressIdtemp := Ctx.PostForm("address_id")
	addressId, err := strconv.Atoi(addressIdtemp)
	if err != nil {
		Ctx.JSON(200, gin.H{
			"success": false,
			"message": "传入参数错误",
		})
		return
	}
	name := Ctx.PostForm("name")
	phone := Ctx.PostForm("phone")
	address := Ctx.PostForm("address")
	zipcode := Ctx.PostForm("zipcode")
	dao.DB.Table("address").Where("uid=?", user.Id).Updates(map[string]interface{}{"default_address": 0})
	addressModel := model.Address{}
	dao.DB.Where("id=?", addressId).Find(&addressModel)
	addressModel.Name = name
	addressModel.Phone = phone
	addressModel.Address = address
	addressModel.Zipcode = zipcode
	addressModel.DefaultAddress = 1
	dao.DB.Save(&addressModel)
	// 查询当前用户的所有收货地址并返回
	allAddressResult := []model.Address{}
	dao.DB.Where("uid=?", user.Id).Order("default_address desc").Find(&allAddressResult)
	Ctx.JSON(200, gin.H{
		"success": true,
		"result":  allAddressResult,
	})
}

// 修改默认地址
func (c *AddressController) ChangeDefaultAddress(Ctx *gin.Context) {
	user := model.User{}
	model.Cookie.Get(Ctx, "userinfo", &user)
	addressIdtemp := Ctx.Query("address_id")
	addressId, err := strconv.Atoi(addressIdtemp)
	if err != nil {
		Ctx.JSON(200, gin.H{
			"success": false,
			"message": "传入参数错误",
		})
		return
	}
	dao.DB.Table("address").Where("uid=?", user.Id).Updates(map[string]interface{}{"default_address": 0})
	dao.DB.Table("address").Where("id=?", addressId).Updates(map[string]interface{}{"default_address": 1})
	Ctx.JSON(200, gin.H{
		"success": true,
		"result":  "更新默认收获地址成功",
	})

}
