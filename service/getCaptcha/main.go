package main

import (
	"getCaptcha/handler"
	pb "getCaptcha/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/consul"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	//初始化consul
	consulReg := consul.NewRegistry()

	// Create service
	srv := micro.NewService(
		micro.Address("localhost:1135"), //防止随机生成port
		micro.Name("getCaptcha"),
		micro.Registry(consulReg),
		micro.Version("latest"),
	)

	// Register handler
	//这里如果出现问题要修改getCaptcha.pb.micro.go中的import
	//client "github.com/micro/go-micro/client"
	//server "github.com/micro/go-micro/server"
	//api "github.com/micro/micro/v3/service/api"
	pb.RegisterGetCaptchaHandler(srv.Server(), new(handler.GetCaptcha))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
