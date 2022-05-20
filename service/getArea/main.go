package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"iHome/service/getArea/handler"

	"github.com/micro/go-plugins/registry/consul"
	"iHome/service/getArea/model"
	getArea "iHome/service/getArea/proto/getArea"
)

func main() {

	model.InitDb()
	model.InitRedis()
	// New Service
	consulRegistry := consul.NewRegistry()

	//service := micro.NewService(
	//	micro.Name("go.micro.srv.getArea"),
	//	micro.Version("latest"),
	//	micro.Registry(consulRegistry),
	//)
	// Create service
	service := micro.NewService(
		micro.Address("localhost:1137"), //防止随机生成port
		micro.Name("go.micro.srv.getArea"),
		micro.Registry(consulRegistry),
	)

	// Initialise service
	service.Init()

	// Register Handler
	getArea.RegisterGetAreaHandler(service.Server(), new(handler.GetArea))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
