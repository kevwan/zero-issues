package main

import (
	"net/http"

	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/service"
	"github.com/tal-tech/go-zero/rest"
)

func main() {
	srv := rest.MustNewServer(rest.RestConf{
		ServiceConf: service.ServiceConf{
			Name: "adhoc",
			Log: logx.LogConf{
				Mode: "console",
			},
		},
		Host: "127.0.0.1",
		Port: 8888,
	})
	srv.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/panic",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			panic("bad things happen")
		},
	})
	srv.Use(func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if p := recover(); p != nil {
					logx.Info("recovered from panic")
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte("oops"))
				}
			}()

			next(w, r)
		}
	})
	defer srv.Stop()
	srv.Start()
}
