package endpoints

import (
	"context"

	"github.com/HsiaoCz/geek/kit-test/services"
	"github.com/go-kit/kit/endpoint"
)

// SumRequest Sum方法的参数.
type SumRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

// SumResponse Sum方法的响应
type SumResponse struct {
	V   int    `json:"v"`
	Err string `json:"err,omitempty"`
}

// ConcatRequest Concat方法的参数.
type ConcatRequest struct {
	A string `json:"a"`
	B string `json:"b"`
}

// ConcatResponse  Concat方法的响应.
type ConcatResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}

func MakeSumEndpoint(adds services.AddS) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SumRequest)
		v, err := adds.Sum(ctx, req.A, req.B)
		if err != nil {
			return SumResponse{V: v, Err: err.Error()}, nil
		}
		return SumResponse{V: v}, nil
	}
}

func MakeConcatEndpoint(adds services.AddS) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ConcatRequest)
		v, err := adds.Concat(ctx, req.A, req.B)
		if err != nil {
			return ConcatResponse{V: v, Err: err.Error()}, nil
		}
		return ConcatResponse{V: v}, nil
	}
}
