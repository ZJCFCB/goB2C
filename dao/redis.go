package dao

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/spf13/viper"
)

var RedisPool *redis.Pool
var redisTime int
var enableRedis bool

func RedisInit() {
	enableRedis = viper.GetBool("enableRedis")
	redisTime, _ = strconv.Atoi(viper.GetString("redisTime"))
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

func RedisSet(key string, value interface{}) {
	data, err := json.Marshal(value)
	if err != nil {
		fmt.Println(err)
		return
	}
	conn := RedisPool.Get()
	fmt.Println("redistiem", redisTime)
	_, err = conn.Do("SET", key, string(data), "EX", redisTime) //有一个过期时间
	if err != nil {
		fmt.Println("redis写入数据失败", err)
	}
}

func RedisGet(key string, value interface{}) bool {
	conn := RedisPool.Get()
	str, err := redis.String(conn.Do("GET", key))
	if err != nil {
		fmt.Println("redis读取数据失败", err)
		return false
	}
	err = json.Unmarshal([]byte(str), value)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func RedisSetString(key string, value string) {
	conn := RedisPool.Get()
	_, err := conn.Do("SET", key, value, "EX", 60*3) //有一个过期时间
	if err != nil {
		fmt.Println("redis写入数据失败", err)
	}
}

func RedisGetString(key string) (string, error) {
	conn := RedisPool.Get()
	value, err := redis.String(conn.Do("GET", key))
	if err != nil {
		fmt.Println("redis读取数据失败", err)
		return value, err
	}
	return value, nil
}

func RedisDel(key string) error {
	conn := RedisPool.Get()
	_, err := conn.Do("DEL", key)
	if err != nil {
		fmt.Println("redis删除数据失败", err)
		return err
	}
	return nil
}
