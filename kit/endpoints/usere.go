package endpoints

import (
	"context"

	"github.com/HsiaoCz/geek/kit/service"
	"github.com/go-kit/kit/endpoint"
)

type UserRequest struct {
	UID int `json:"uid"`
}

type UserResponse struct {
	Result string `json:"result"`
}

func GetUserEndpoint(userService service.IUserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r := request.(UserRequest)
		result := userService.GetUsername(r.UID)
		return UserResponse{Result: result}, nil
	}
}
