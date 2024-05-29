package main

import (
	"fmt"
	"goB2C/dao"
	"goB2C/model"
	"goB2C/util"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/spf13/viper"
)

type TestS struct {
	Id  int    `json:"id"`
	Pwd string `json:"pwd`
}

func main() {
	fmt.Println(util.TransToDay(1716993222))
	viper.SetConfigFile("../conf/app.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read config file: %w", err))
	}
	dao.MysqlInit()
	dao.RedisInit()
	//fmt.Println(model.GetProductByCategory(1, "best", 8))

	r := gin.Default()
	r.GET("/cap", model.CapTest)
	r.LoadHTMLFiles("../view/frontend/public/page_footer.html")
	r.GET("/foot", func(ctx *gin.Context) {
		var Cap *base64Captcha.Captcha

		id, base64, value, err := Cap.Generate()
		ctx.Set(id, value)
		if err != nil {
			fmt.Println("get capte failed", err)
		}
		// create html
		ctx.String(200, fmt.Sprintf(`<input type="hidden" name="captcha_id" value="%s">`+
			`<img id="captcha-img" src="data:image/png;base64,%s"`+
			` alt="Captcha" onclick="refreshCaptcha()">`, id, base64))
	})
	r.Run(":" + viper.GetString("server.port"))
}
