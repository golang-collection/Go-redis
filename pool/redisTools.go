package pool

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/spf13/viper"
	"time"
)

/**
* @Author: super
* @Date: 2020-08-12 20:54
* @Description:
**/

var pool *redis.Pool

//创建redis连接池
func newRedisPool() *redis.Pool{
	return &redis.Pool{
		MaxIdle:50,
		MaxActive:30,
		IdleTimeout:300*time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", viper.GetString("url"))
			if err != nil{
				fmt.Println(err.Error())
				return nil, err
			}
			return conn, nil
		},
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute{
				return nil
			}
			_, err := conn.Do("PING")
			return err
		},
	}
}

func init(){
	pool = newRedisPool()
}

func GetConn() redis.Conn{
	return pool.Get()
}