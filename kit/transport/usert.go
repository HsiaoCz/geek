package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/HsiaoCz/geek/kit/endpoints"
)

func DecodeUserRequest(c context.Context, r *http.Request) (interface{}, error) {
	uid, err := strconv.Atoi(r.URL.Query().Get("uid"))
	return endpoints.UserRequest{UID: uid}, err
}

func EncodeUserResponse(c context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(response)
}
