package operation

import (
	"GO-redis/pool"
	"github.com/garyburd/redigo/redis"
)

/**
* @Author: super
* @Date: 2020-08-12 20:56
* @Description:
**/

func SetString(key string, value string) (string, error) {
	c := pool.GetConn()

	str, err := redis.String(c.Do("set", key, value))
	if err != nil {
		return "", err
	}
	return str, err
}

func GetString(key string) (string, error) {
	c := pool.GetConn()
	str, err := redis.String(c.Do("get", key))
	if err != nil {
		return "", err
	}
	return str, err
}

func AddElementToSet(key string, value string) (int, error) {
	c := pool.GetConn()

	result, err := redis.Int(c.Do("sadd", key, value))
	if err != nil {
		return -1, err
	}
	return result, err
}

func ElementIsInSet(key string, value string) (bool, error) {
	c := pool.GetConn()
	result, err := redis.Int(c.Do("sismember", key, value))
	if err != nil {
		return false, err
	}
	if result == 1 {
		return true, err
	}
	return false, err
}
