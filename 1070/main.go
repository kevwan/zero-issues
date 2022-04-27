package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func main() {
	logx.Disable()

	var c rest.RestConf
	conf.MustLoad("config.yaml", &c)

	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		return 499, err
	})
	srv := rest.MustNewServer(c)
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
			time.Sleep(time.Second)
			io.WriteString(w, "hello")
			io.WriteString(w, "world")
		},
	})
	defer srv.Stop()
	srv.Start()
}
