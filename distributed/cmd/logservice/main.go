package main

import (
	"context"
	"fmt"
	stlog "log"

	"github.com/HsiaoCz/geek/distributed/log"
	"github.com/HsiaoCz/geek/distributed/registry"
	"github.com/HsiaoCz/geek/distributed/service"
)

func main() {
	log.Run("./distributed.log")
	host, port := "localhost", "8081"
	serviceAddr := fmt.Sprintf("http://%s:%s", host, port)
	r := registry.Registration{
		ServiceName: "log service",
		ServiceURL:  serviceAddr,
	}
	ctx, err := service.Start(context.Background(), r, host, port, log.RegisterHandlers)
	if err != nil {
		stlog.Fatalln(err)
	}

	<-ctx.Done()

	fmt.Println("shutdown service")
}
