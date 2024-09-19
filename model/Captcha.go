package model

import (
	"fmt"
	"goB2C/dao"
	"html/template"
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

func CreateCaptchaHTML() template.HTML {
	id, base64, value, err := Cap.Generate()
	dao.RedisSetString(id, value)
	if err != nil {
		fmt.Println("get capte failed", err)
		return ""
	}
	// create html
	return template.HTML(fmt.Sprintf(`<input type="hidden" name="captcha_id" value="%s">`+
		`<img id="captcha-img" src="%s"  width="120" height="50"`+
		` alt="Captcha" >`, id, base64))
	//onclick="refreshCaptcha()"
}

func CaptchaVerify(phoneCodeId string, phone_code string) bool {

	//把验证码校验去掉
	return true

	ans, err := dao.RedisGetString(phoneCodeId)
	if err != nil {
		return false
	}
	if ans != phone_code {
		return false
	}
	return true
}
