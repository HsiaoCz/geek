package main

import (
	"log"
	"net/http"
	"os"
)

func HandleLY(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		file, err := os.OpenFile("./小樊宜的全盛时代.md", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			log.Fatal(err)
		}
		_, err = file.WriteString("小樊宜今天真好看!")
		if err != nil {
			log.Fatal(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("写入成功!"))
		file.Close()
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func main() {
	http.HandleFunc("/api/v1/ly", HandleLY)
	http.ListenAndServe("127.0.0.1:3002", nil)
}
