package registry

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

const ServerPort = ":3001"
const ServicesURL = "http://localhost" + ServerPort + "/services"

type registry struct {
	registerations []Registration
	mutex          *sync.Mutex
}

func (r *registry) add(reg Registration) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.registerations = append(r.registerations, reg)
	return nil
}

var reg = registry{
	registerations: make([]Registration, 0),
	mutex:          new(sync.Mutex),
}

type RegistryService struct{}

func (s RegistryService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("request recived")
	switch r.Method {
	case http.MethodPost:
		dec := json.NewDecoder(r.Body)
		var r Registration
		err := dec.Decode(&r)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Printf("adding service:%v with Url:%s\n", r.ServiceName, r.ServiceURL)
		err = reg.add(r)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
