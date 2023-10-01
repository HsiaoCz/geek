package main

import (
	"log"
	"net/http"
	"time"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}
func (a *APIServer) Run() {
	http.HandleFunc("/api/count", a.handleCount)
	srv := http.Server{
		Handler:      nil,
		Addr:         a.addr,
		WriteTimeout: 1500 * time.Millisecond,
		ReadTimeout:  1500 * time.Millisecond,
	}
	log.Println("the server is running on port:", a.addr)
	log.Fatal(srv.ListenAndServe())
}
func (a *APIServer) handleCount(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		a.handleGetCount(w, r)
	case "POST":
		a.handleCreateCount(w, r)
	case "PUT":
		a.hanldeModCount(w, r)
	case "DELETE":
		a.handleDelCount(w, r)
	}

}

func (a *APIServer) handleCreateCount(w http.ResponseWriter, r *http.Request) {}

func (a *APIServer) handleGetCount(w http.ResponseWriter, r *http.Request) {}

func (a *APIServer) hanldeModCount(w http.ResponseWriter, r *http.Request) {}

func (a *APIServer) handleDelCount(w http.ResponseWriter, r *http.Request) {}
