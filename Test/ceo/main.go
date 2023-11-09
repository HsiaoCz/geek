package main

import "net/http"

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
	http.ListenAndServe("127.0.0.1:9091", nil)
}


