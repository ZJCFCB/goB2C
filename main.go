package main

import (
	"fmt"
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
	})

	// 设置跨域访问选项
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	r.Run(":" + viper.GetString("server.port"))
}
