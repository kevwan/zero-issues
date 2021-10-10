package main

import (
	"fmt"
	"net/http"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/rest/router"
)

func main() {
	var c rest.RestConf
	conf.MustLoad("config.yaml", &c)

	r := router.NewRouter()
	r.SetNotAllowedHandler(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("not allowed")
	}))
	r.SetNotFoundHandler(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("not found")
	}))
	srv := rest.MustNewServer(c, rest.WithRouter(r))
	srv.AddRoute(rest.Route{
		Method: http.MethodPost,
		Path:   "/hello",
		Handler: func(rw http.ResponseWriter, r *http.Request) {
		},
	})
	defer srv.Stop()
	srv.Start()
}
