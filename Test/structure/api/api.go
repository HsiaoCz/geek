package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/HsiaoCz/geek/Test/structure/model"
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

func (a *APIServer) handleGetCount(w http.ResponseWriter, r *http.Request) {
	countID := r.URL.Query().Get("count_id")
	cId, err := strconv.Atoi(countID)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("please offer useful count_id"))
		return
	}
	w.Header().Set("Content-Type", "application")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&model.Count{
		FirstName: "bob",
		LastName:  "phoniex",
		Identity:  int64(cId),
		Money:     "12222",
		Dart: []model.Goods{
			{
				Name:      "三角内裤",
				Price:     "122",
				RepoCount: 12222,
				Describe:  "买了绝对的不亏",
			},
		},
	})
}

func (a *APIServer) hanldeModCount(w http.ResponseWriter, r *http.Request) {}

func (a *APIServer) handleDelCount(w http.ResponseWriter, r *http.Request) {}
