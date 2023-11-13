package main

import (
	"encoding/json"
	"net/http"
)

func HandleSayHelloToCEOLittleFan(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("小樊总好!"))
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func main() {
	http.HandleFunc("/xiaofanzong", HandleSayHelloToCEOLittleFan)
	http.HandleFunc("/ceofan", HandleCEOFan)
	http.ListenAndServe("127.0.0.1:9091", nil)
}

func HandleCEOFan(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"Message": "小樊总今天真好看!",
		})
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}
