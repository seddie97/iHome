package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/consul"
	"github.com/micro/micro/v3/service/logger"
	"user/handler"
	"user/model"
	pb "user/proto"
)

func main() {

	//初始化 mysql连接池
	model.InitDb()

	//初始化redis连接池
	model.InitRedis()

	//初始化consul
	consulReg := consul.NewRegistry()

	// Create service
	srv := micro.NewService(
		micro.Address("localhost:1136"), //防止随机生成port
		micro.Name("user"),
		micro.Registry(consulReg),
		micro.Version("latest"),
	)

	// Register handler
	pb.RegisterUserHandler(srv.Server(), new(handler.User))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
