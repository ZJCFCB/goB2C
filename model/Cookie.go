package model

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type cookie struct{}

var Cookie = &cookie{}

func (c cookie) Set(ctx *gin.Context, key string, value interface{}) {
	bytes, _ := json.Marshal(value)
	domain := viper.GetString("domain")
	if domain == "" {
		domain = ctx.Request.Host
	}
	ctx.SetCookie(key, string(bytes), 3600*60*7, "/", domain, false, true)
}

func (c cookie) Remove(ctx *gin.Context, key string, value interface{}) {
	bytes, _ := json.Marshal(value)
	domain := viper.GetString("domain")
	if domain == "" {
		domain = ctx.Request.Host
	}
	ctx.SetCookie(key, string(bytes), -1, "/", domain, false, true)
}
func (c cookie) Get(ctx *gin.Context, key string, obj interface{}) bool {
	temp, err := ctx.Cookie(key)
	if err != nil {
		fmt.Println("get cookie failed", err)
		return false
	}
	if err := json.Unmarshal([]byte(temp), obj); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
