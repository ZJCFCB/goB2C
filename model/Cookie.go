package model

type cookie struct{}

/*
func (c cookie) Set(ctx *gin.Context, key string, value interface{}) {
	bytes, _ := json.Marshal(value)
	ctx.SetSecureCookie(viper.GetString("secureCookie"), key, string(bytes), 3600*24*30, "/", viper.GetString("domain"), nil, true)
}
*/
