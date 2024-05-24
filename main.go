package main

import (
	"fmt"
	"goB2C/controllers"
	"goB2C/dao"
	"goB2C/model"
	"goB2C/util"
	"text/template"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigFile("conf/app.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read config file: %w", err))
	}

	dao.MysqlInit()
	dao.RedisInit()
	r := gin.Default()

	//添加方法用于前端调用
	r.SetFuncMap(template.FuncMap{
		"timestampToDate": util.TimestampToData,
		"formatImage":     util.FormatImage,
		"mul":             util.Mul,
		"formatAttribute": util.FormatAttribute,
		"setting":         model.GetSettingByColumn,
		"substr":          util.SubString,
		"str2html":        util.Str2html,
	})

	// 设置跨域访问选项
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	//设置Gin的错误捕获,放在error.log中
	r.Use(gin.RecoveryWithWriter(util.GetLogWriter()))

	//加载渲染文件
	util.InitHTML(r)

	r.Static("static", "static")
	//注册方法
	controllers.RegistFunc(r)

	r.Run(":" + viper.GetString("server.port"))
}
