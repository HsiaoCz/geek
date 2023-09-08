package controller

import (
	"encoding/json"
	"net/http"
)

type H map[string]any

func ResponseJSON(w http.ResponseWriter, code int, message string, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(H{
		"Message": message,
		"Data":    data,
	})
}
