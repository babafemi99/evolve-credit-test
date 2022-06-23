package web

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

var RouterChi = chi.NewRouter()

type chiRouter struct {
}

func (c *chiRouter) GET(uri string, f func(rw http.ResponseWriter, r *http.Request)) {
	RouterChi.Get(uri, f)
}

func (c *chiRouter) POST(uri string, f func(rw http.ResponseWriter, r *http.Request)) {
	RouterChi.Post(uri, f)
}

func (c *chiRouter) SERVE(port string) error {
	log.Println("Chi router listening on port - ", port)
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

func newChiRouter() RouterInterface {
	return &chiRouter{}
}
