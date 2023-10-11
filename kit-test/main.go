package main

import (
	"log"
	"net/http"

	"github.com/HsiaoCz/geek/kit-test/endpoints"
	"github.com/HsiaoCz/geek/kit-test/services"
	"github.com/HsiaoCz/geek/kit-test/transports"
	httptransport "github.com/go-kit/kit/transport/http"
)

const (
	addr = "127.0.0.1:9911"
)

func main() {
	svc := services.AddS{}
	sumHandler := httptransport.NewServer(
		endpoints.MakeSumEndpoint(svc),
		transports.DecodeSumRequest,
		transports.EncodeResponse,
	)
	concatHandler := httptransport.NewServer(
		endpoints.MakeConcatEndpoint(svc),
		transports.DecodeCountRequest,
		transports.EncodeResponse,
	)
	http.Handle("/sum", sumHandler)
	http.Handle("/concat", concatHandler)
	log.Fatal(http.ListenAndServe(addr, nil))

}
