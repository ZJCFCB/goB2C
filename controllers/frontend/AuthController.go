package frontend

import (
	"goB2C/dao"
	"goB2C/model"
	"goB2C/util"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

//这里是用户注册模块

type AuthController struct {
	BaseController
}

func (A AuthController) RegisterStep1(Ctx *gin.Context) {

	Ctx.HTML(200, "auth_register_step1.html", gin.H{})
}

func (A AuthController) RegisterStep2(Ctx *gin.Context) {
	sign := Ctx.Query("sign")
	phone_code := Ctx.Query("phone_code")
	phoneCodeId := Ctx.Query("phoneCodeId")
	//验证图形验证码和前面是否正确
	sessionPhotoCode, err := dao.RedisGetString(phoneCodeId)
	if err != nil || phone_code != sessionPhotoCode {
		Ctx.Redirect(http.StatusFound, "/auth/registerStep1")
		return
	}
	userTemp := []model.UserSms{}
	dao.DB.Where("sign=?", sign).Find(&userTemp)
	if len(userTemp) > 0 {
		Ctx.HTML(200, "auth_register_step2.html", gin.H{
			"phone":      userTemp[0].Phone,
			"phone_code": phone_code,
			"sign":       sign,
		})
	} else {
		Ctx.Redirect(http.StatusFound, "/auth/registerStep1")
		return
	}
}

func (c *AuthController) RegisterStep3(Ctx *gin.Context) {
	sign := Ctx.Query("sign")
	sms_code := Ctx.Query("sms_code")
	//sessionSmsCode := c.GetSession("sms_code")
	if sms_code != "5259" {
		Ctx.Redirect(http.StatusFound, "/auth/registerStep1")
		return
	}
	userTemp := []model.UserSms{}
	dao.DB.Where("sign=?", sign).Find(&userTemp)
	if len(userTemp) > 0 {
		Ctx.HTML(200, "auth_register_step3.html", gin.H{
			"sign":     sign,
			"sms_code": sms_code,
		})
	} else {
		Ctx.Redirect(http.StatusFound, "/auth/registerStep1")
		return
	}
}

func (c *AuthController) SendCode(Ctx *gin.Context) {
	phone := Ctx.Query("phone")
	phone_code := Ctx.Query("phone_code")
	phoneCodeId := Ctx.Query("phoneCodeId")
	_, err := strconv.Atoi(phone_code)
	if err != nil {
		//session里面验证验证码是否合法
		Ctx.JSON(200, gin.H{
			"success": false,
			"msg":     "输入的图形验证码不正确,非法请求",
		})
		return
	}
	ans, err := dao.RedisGetString(phoneCodeId)
	if err != nil || phone_code != ans {
		Ctx.JSON(200, gin.H{
			"success": false,
			"msg":     "输入的图形验证码不正确",
		})
		return
	}
	pattern := `^[\d]{11}$`
	reg := regexp.MustCompile(pattern)
	if !reg.MatchString(phone) {
		Ctx.JSON(200, gin.H{
			"success": false,
			"msg":     "手机号格式不正确",
		})
		return
	}
	user := []model.User{}
	dao.DB.Where("phone=?", phone).Find(&user)
	if len(user) > 0 {
		Ctx.JSON(200, gin.H{
			"success": false,
			"msg":     "此用户已存在",
		})
		return
	}

	add_day := util.FormatDay()
	ip := strings.Split(Ctx.Request.RemoteAddr, ":")[0]
	sign := util.Md5(phone + add_day) //签名
	sms_code := util.GetRandomNum()
	userTemp := []model.UserSms{}
	dao.DB.Where("add_day=? AND phone=?", add_day, phone).Find(&userTemp)
	var sendCount int64
	dao.DB.Where("add_day=? AND ip=?", add_day, ip).Table("user_sms").Count(&sendCount)
	//验证IP地址今天发送的次数是否合法
	if sendCount <= 10 {
		if len(userTemp) > 0 {
			//验证当前手机号今天发送的次数是否合法
			if userTemp[0].SendCount < 5 {
				util.SendMsg(sms_code)
				//Ctx.Set("sms_code", sms_code)
				oneUserSms := model.UserSms{}
				dao.DB.Where("id=?", userTemp[0].Id).Find(&oneUserSms)
				oneUserSms.SendCount += 1
				dao.DB.Save(&oneUserSms)
				Ctx.JSON(200, gin.H{
					"success":  true,
					"msg":      "短信发送成功",
					"sign":     sign,
					"sms_code": sms_code,
				})
				return
			} else {
				Ctx.JSON(200, gin.H{
					"success": false,
					"msg":     "当前手机号今天发送短信数已达上限",
				})
				return
			}

		} else {
			util.SendMsg(sms_code)
			//Ctx.Set("sms_code", sms_code)
			//发送验证码 并给userTemp写入数据
			oneUserSms := model.UserSms{
				Ip:        ip,
				Phone:     phone,
				SendCount: 1,
				AddDay:    add_day,
				AddTime:   int(util.GetUnix()),
				Sign:      sign,
			}
			dao.DB.Create(&oneUserSms)
			Ctx.JSON(200, gin.H{
				"success":  true,
				"msg":      "短信发送成功",
				"sign":     sign,
				"sms_code": sms_code,
			})

			return
		}
	} else {
		Ctx.JSON(200, gin.H{
			"success": false,
			"msg":     "此IP今天发送次数已经达到上限，明天再试",
		})
		return
	}

}

func (c *AuthController) ValidateSmsCode(Ctx *gin.Context) {
	sign := Ctx.Query("sign")
	sms_code := Ctx.Query("sms_code")

	userTemp := []model.UserSms{}
	dao.DB.Where("sign=?", sign).Find(&userTemp)
	if len(userTemp) == 0 {
		Ctx.JSON(200, gin.H{
			"success": false,
			"msg":     "参数错误",
		})
		return
	}

	//sessionSmsCode := c.GetSession("sms_code")
	if sms_code != "5259" {
		Ctx.JSON(200, gin.H{
			"success": false,
			"msg":     "输入的短信验证码错误",
		})
		return
	}

	nowTime := util.GetUnix()
	if (nowTime-int64(userTemp[0].AddTime))/1000/60 > 15 {
		Ctx.JSON(200, gin.H{
			"success": false,
			"msg":     "验证码已过期",
		})
		return
	}
	Ctx.JSON(200, gin.H{
		"success": true,
		"msg":     "验证成功",
	})
}

func (c *AuthController) GoRegister(Ctx *gin.Context) {
	sign := Ctx.PostForm("sign")
	sms_code := Ctx.PostForm("sms_code")
	password := Ctx.PostForm("password")
	rpassword := Ctx.PostForm("rpassword")
	//sessionSmsCode := c.GetSession("sms_code")
	if sms_code != "5259" {
		Ctx.Redirect(http.StatusFound, "/auth/registerStep1")
		return
	}
	if len(password) < 6 {
		Ctx.Redirect(http.StatusFound, "/auth/registerStep1")
	}
	if password != rpassword {
		Ctx.Redirect(http.StatusFound, "/auth/registerStep1")
	}
	userTemp := []model.UserSms{}
	dao.DB.Where("sign=?", sign).Find(&userTemp)
	ip := strings.Split(Ctx.Request.RemoteAddr, ":")[0]
	if len(userTemp) > 0 {
		user := model.User{
			Phone:    userTemp[0].Phone,
			Password: util.Md5(password),
			LastIP:   ip,
		}
		dao.DB.Create(&user)

		model.Cookie.Set(Ctx, "userinfo", user)
		Ctx.Redirect(http.StatusFound, "/mainPage")
	} else {
		Ctx.Redirect(http.StatusFound, "/auth/registerStep1")
	}

}
