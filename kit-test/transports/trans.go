package transports

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/HsiaoCz/geek/kit-test/endpoints"
)

func DecodeSumRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request endpoints.SumRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request endpoints.ConcatRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
