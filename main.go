package main

import (
	"GO-redis/config"
	"GO-redis/proto"
	"GO-redis/server"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"log"
	"time"
)

/**
* @Author: super
* @Date: 2020-08-17 16:59
* @Description:
**/

func main() {

	//init config
	if err := config.Init("remote"); err != nil {
		log.Println(err)
		panic(err)
	}

	reg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			config.CONSUL_URL,
		}
	})

	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("go.micro.service.redis"),
		micro.RegisterTTL(time.Second*10),
		micro.RegisterInterval(time.Second*5),
	)
	service.Init()

	// 注册处理器
	err := proto.RegisterRedisOperationHandler(service.Server(), new(server.Redis))
	if err != nil {
		log.Println(err)
		return
	}

	// 运行服务
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
