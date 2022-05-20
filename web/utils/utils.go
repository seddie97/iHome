package utils

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/consul"
)

//初始化micro客户端
func InitMicro() micro.Service {
	//指定服务发现
	//初始化consul
	consulReg := consul.NewRegistry()
	consulSrv := micro.NewService(
		micro.Registry(consulReg),
	)

	return consulSrv
}
