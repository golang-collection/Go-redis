package server

import (
	"GO-redis/pool"
	"GO-redis/proto"
	"context"
	"github.com/garyburd/redigo/redis"
)

/**
* @Author: super
* @Date: 2020-08-17 16:53
* @Description:
**/

type Redis struct {
}

func (r *Redis) SetString(ctx context.Context, req *proto.SetStringRequest, res *proto.SetStringResponse) error {
	c := pool.GetConn()

	str, err := redis.String(c.Do("set", req.Key, req.Value))
	if err != nil {
		return err
	}
	res.Result = str
	return nil
}

func (r *Redis) GetString(ctx context.Context, req *proto.GetStringRequest, res *proto.GetStringResponse) error {
	c := pool.GetConn()
	str, err := redis.String(c.Do("get", req.Key))
	if err != nil {
		return err
	}
	res.Result = str
	return nil
}
