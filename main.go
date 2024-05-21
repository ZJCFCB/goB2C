package main

import (
	"fmt"
	"goB2C/dao"
	"goB2C/service"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	//viper.AddConfigPath("./conf")
	viper.SetConfigFile("conf/app.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read config file: %w", err))
	}
	dao.MysqlInit()
	dao.RedisInit()
	fmt.Println(service.GetByAuthId(1))
	r := gin.Default()
	r.Run(":" + viper.GetString("server.port"))
}
