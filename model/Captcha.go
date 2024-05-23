package model

import (
	"fmt"
	"image/color"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

var Cap *base64Captcha.Captcha

var (
	Height          = 70
	Width           = 240
	NoiseCount      = 0
	ShowLineOptions = base64Captcha.OptionShowHollowLine
	BgColor         = &color.RGBA{
		R: 144,
		G: 238,
		B: 144,
		A: 10,
	}
	FontsStorage base64Captcha.FontsStorage
	Fonts        []string
)

// 这里是用于管理和生成 图片验证码的代码
// 采用数字的形式
func init() {
	var store = base64Captcha.DefaultMemStore
	driver := base64Captcha.NewDriverMath(Height, Width, NoiseCount, ShowLineOptions, BgColor, FontsStorage, Fonts)
	//driver := base64Captcha.NewDriverDigit(Height, Width, 5, 0.5, 5)
	Cap = base64Captcha.NewCaptcha(driver, store)
}

func CapTest(c *gin.Context) {
	id, base64, ans, err := Cap.Generate()
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, gin.H{
		"id":   id,
		"base": base64,
		"ans":  ans,
	})
}
