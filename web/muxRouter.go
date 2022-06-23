package web

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var router = mux.NewRouter()

type muxRouter struct {
}

func (m *muxRouter) GET(uri string, f func(rw http.ResponseWriter, r *http.Request)) {
	router.HandleFunc(uri, f).Methods("GET")
}

func (m *muxRouter) POST(uri string, f func(rw http.ResponseWriter, r *http.Request)) {
	router.HandleFunc(uri, f).Methods("GET")
}

func (m *muxRouter) SERVE(port string) error {
	log.Println("mux router listening on port - ", port)
	s := &http.Server{
		Addr:    port,
		Handler: router,
	}
	err := s.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func NewRouterMux() RouterInterface {
	return &muxRouter{}
}
