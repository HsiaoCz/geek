package main

import (
	"context"
	"fmt"
	stlog "log"

	"github.com/HsiaoCz/geek/distributed/log"
	"github.com/HsiaoCz/geek/distributed/service"
)

func main() {
	log.Run("./distributed.log")
	host, port := "localhost", "8081"
	ctx, err := service.Start(context.Background(), "log service", host, port, log.RegisterHandlers)
	if err != nil {
		stlog.Fatalln(err)
	}

	<-ctx.Done()

	fmt.Println("shutdown service")
}
