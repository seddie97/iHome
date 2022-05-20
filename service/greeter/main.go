package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"greeter/handler"
	"greeter/subscriber"

	greeter "greeter/proto/greeter"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.greeter"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	greeter.RegisterGreeterHandler(service.Server(), new(handler.Greeter))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.greeter", service.Server(), new(subscriber.Greeter))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
