package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/rest/router"
)

func main() {
	logx.Disable()

	var c rest.RestConf
	conf.MustLoad("config.yaml", &c)

	r := router.NewRouter()
	srv := rest.MustNewServer(c, rest.WithRouter(r))
	srv.Use(func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("middleware")
			io.WriteString(w, "again")
			next(w, r)
		}
	})
	srv.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/hello",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			io.WriteString(w, "hello")
			io.WriteString(w, "world")
		},
	})
	defer srv.Stop()
	srv.Start()
}
