package main

import (
	"GO-redis/proto"
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"log"
)

/**
* @Author: super
* @Date: 2020-08-17 17:17
* @Description:
**/

var redisOP proto.RedisOperationService

func main() {
	reg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("go.micro.service.redis.client"),
	)
	service.Init()

	redisOP = proto.NewRedisOperationService("go.micro.service.redis", service.Client())
	s, err := SetString("set1", "hello")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(s)
	s, err = GetString("set1")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(s)
}

func SetString(key string, value string) (string, error) {
	rsp, err := redisOP.SetString(context.Background(), &proto.SetStringRequest{Key: key, Value: value})
	if err != nil {
		return "", err
	}
	return rsp.Result, nil
}

func GetString(key string) (string, error) {
	rsp, err := redisOP.GetString(context.TODO(), &proto.GetStringRequest{Key: key})
	if err != nil {
		return "", nil
	}
	return rsp.Result, nil
}
