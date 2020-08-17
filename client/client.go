package main

import (
	"GO-redis/proto"
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
)

/**
* @Author: super
* @Date: 2020-08-17 17:17
* @Description:
**/

func main() {
	service := micro.NewService(
		micro.Name("go.micro.service.redis.client"),
	)
	service.Init()

	SetString(service)
	GetString(service)

}

func SetString(service micro.Service){
	redisOP := proto.NewRedisOperationService("go.micro.service.redis", service.Client())
	rsp, err := redisOP.SetString(context.TODO(), &proto.SetStringRequest{Key:"s1", Value:"lala"})
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(rsp.Result)
}

func GetString(service micro.Service){
	redisOP := proto.NewRedisOperationService("go.micro.service.redis", service.Client())
	rsp, err := redisOP.GetString(context.TODO(), &proto.GetStringRequest{Key:"s1"})
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(rsp.Result)
}