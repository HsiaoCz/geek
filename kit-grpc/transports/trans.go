package transports

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/HsiaoCz/geek/kit-grpc/endpoints"
	"github.com/HsiaoCz/geek/kit-grpc/pb"
	"github.com/HsiaoCz/geek/kit-grpc/services"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
)

// gRPC的请求与响应
// decodeGRPCSumRequest 将Sum方法的gRPC请求参数转为内部的SumRequest
func decodeGRPCSumRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.SumRequest)
	return endpoints.SumRequest{A: int(req.A), B: int(req.B)}, nil
}

// decodeGRPCConcatRequest 将Concat方法的gRPC请求参数转为内部的ConcatRequest
func decodeGRPCConcatRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.ConcatRequest)
	return endpoints.ConcatRequest{A: req.A, B: req.B}, nil
}

// encodeGRPCSumResponse 封装Sum的gRPC响应
func encodeGRPCSumResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.SumResponse)
	return &pb.SumResponse{V: int64(resp.V), Err: resp.Err}, nil
}

// encodeGRPCConcatResponse 封装Concat的gRPC响应
func encodeGRPCConcatResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.ConcatResponse)
	return &pb.ConcatResponse{V: resp.V, Err: resp.Err}, nil
}

// gRPC
type grpcServer struct {
	pb.UnimplementedHelloServiceServer

	sum    grpctransport.Handler
	concat grpctransport.Handler
}

func (s grpcServer) Sum(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
	_, resp, err := s.sum.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.SumResponse), nil
}

func (s grpcServer) Concat(ctx context.Context, req *pb.ConcatRequest) (*pb.ConcatResponse, error) {
	_, resp, err := s.concat.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ConcatResponse), nil
}

// NewGRPCServer 构造函数
func NewGRPCServer(svc services.AddService) pb.HelloServiceServer {
	return &grpcServer{
		sum: grpctransport.NewServer(
			endpoints.MakeSumEndpoint(svc), // endpoint
			decodeGRPCSumRequest,
			encodeGRPCSumResponse,
		),
		concat: grpctransport.NewServer(
			endpoints.MakeSumEndpoint(svc),
			decodeGRPCConcatRequest,
			encodeGRPCConcatResponse,
		),
	}
}

// HTTP
func decodeSumRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request endpoints.SumRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeConcatRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request endpoints.ConcatRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// HTTP Server
func NewHTTPServer(svc services.AddService, logger log.Logger) http.Handler {
	// 添加日志
	sum := endpoints.MakeSumEndpoint(svc)
	// 使用log为sum添加日志
	sum = endpoints.LoggingMiddleware(log.With(logger, "method", "sum"))(sum)
	sumHandler := httptransport.NewServer(
		sum,
		decodeSumRequest,
		encodeResponse,
	)

	// 给conca添加logger
	concat := endpoints.MakeConcatEndpoint(svc)
	// 使用logger
	concat = endpoints.LoggingMiddleware(log.With(logger, "method", "concat"))(concat)
	concatHandler := httptransport.NewServer(
		concat,
		decodeConcatRequest,
		encodeResponse,
	)
	// use github.com/gorilla/mux
	r := mux.NewRouter()
	r.Handle("/sum", sumHandler).Methods("POST")
	r.Handle("/concat", concatHandler).Methods("POST")

	// use gin
	// r := gin.Default()
	// r.POST("/sum", gin.WrapH(sumHandler))
	// r.POST("/concat", gin.WrapH(concatHandler))
	return r

}
