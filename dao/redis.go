package dao

import (
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/spf13/viper"
)

var RedisPool *redis.Pool

func RedisInit() {
	ip := viper.GetString("redis.host")
	port := viper.GetString("redis.port")
	db := viper.GetInt("redis.db")
	pool := &redis.Pool{
		MaxIdle:     10,                // 最大空闲连接数
		MaxActive:   0,                 // 最大活动连接数
		IdleTimeout: 300 * time.Second, // 空闲连接超时时间
		Dial: func() (redis.Conn, error) { // 创建连接的函数
			conn, err := redis.Dial("tcp", ip+":"+port, redis.DialDatabase(db))
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
	}
	RedisPool = pool
}
