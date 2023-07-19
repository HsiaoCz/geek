package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/HsiaoCz/geek/distributed/registry"
)

func main() {
	http.Handle("/services", &registry.RegistryService{})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	srv := http.Server{
		Addr: registry.ServerPort,
	}

	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		fmt.Println("Registry service started,press any key to stop")
		var s string
		fmt.Scanln(&s)
		srv.Shutdown(ctx)
		cancel()
	}()
	<-ctx.Done()
	fmt.Println("Shutting down registry service")
}
