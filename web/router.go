package web

import "net/http"

type RouterInterface interface {
	GET(uri string, f func(rw http.ResponseWriter, r *http.Request))
	POST(uri string, f func(rw http.ResponseWriter, r *http.Request))
	SERVE(port string) error
}
