package server

import (
	"GO-redis/proto"
	"GO-redis/tools"
	"context"
	"github.com/garyburd/redigo/redis"
)

/**
* @Author: super
* @Date: 2020-08-17 16:53
* @Description:
**/

type RedisStruct struct {
}

func (redisstruct *RedisStruct)SetString(ctx context.Context, req *proto.SetStringRequest, res *proto.SetStringResponse) error {
	c := tools.GetConn()

	str, err := redis.String(c.Do("set", req.Key, req.Value))
	if err != nil{
		return err
	}
	res.Result = str
	return nil
}

func (redisstruct *RedisStruct)GetString(ctx context.Context, req *proto.GetStringRequest, res *proto.GetStringResponse) error {
	c := tools.GetConn()
	str, err := redis.String(c.Do("get", req.Key))
	if err != nil{
		return err
	}
	res.Result = str
	return nil
}

func main() {

}