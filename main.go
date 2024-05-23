package main

import (
	"fmt"
	"goB2C/dao"
	"goB2C/model"

	"github.com/spf13/viper"
)

type TestS struct {
	Id  int    `json:"id"`
	Pwd string `json:"pwd`
}

func main() {
	//viper.AddConfigPath("./conf")
	viper.SetConfigFile("conf/app.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read config file: %w", err))
	}
	dao.MysqlInit()
	dao.RedisInit()
	var add model.UserSms
	dao.DB.Model(&model.UserSms{}).First(&add)
	fmt.Println(add)
	//r := gin.Default()
	//r.Run(":" + viper.GetString("server.port"))
}
