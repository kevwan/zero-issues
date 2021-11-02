package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/service"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/rest/httpx"
)

var port = flag.Int("port", 3333, "the port to listen")

type Request struct {
	User int `form:"user,default=1"`
	Age  int `form:"age,default=18"`
}

func handle(w http.ResponseWriter, r *http.Request) {
	var req Request
	err := httpx.Parse(r, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	httpx.OkJson(w, fmt.Sprintf("user: %d, age: %d", req.User, req.Age))
}

func main() {
	flag.Parse()

	logx.DisableStat()
	srv := rest.MustNewServer(rest.RestConf{
		ServiceConf: service.ServiceConf{
			Log: logx.LogConf{
				Mode: "console",
			},
		},
		Port: *port,
	})
	defer srv.Stop()

	srv.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/hello",
		Handler: handle,
	})
	srv.Start()
}
