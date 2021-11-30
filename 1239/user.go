package main

import (
	"flag"
	"fmt"
	"net/http"

	"issue/internal/config"
	"issue/internal/handler"
	"issue/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/rest/httpx"
)

// Result std response struct
type Result struct {
	Status    int         `json:"status"`
	ErrorCode interface{} `json:"error_code"`
	Error     interface{} `json:"error"`
	Data      interface{} `json:"data,omitempty"`
}

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()
	logx.Disable()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(
		c.RestConf,
		rest.WithCustomCors(func(header http.Header) {
			header.Set("foo", "bar")
		}, func(w http.ResponseWriter) {
			httpx.OkJson(w, &Result{Status: 0, ErrorCode: 1001, Error: "ACCESS_DENY"})
		}),
		rest.WithNotFoundHandler(notFoundHandler()),
		rest.WithUnauthorizedCallback(unauthorizedHandler),
	)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

func notFoundHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		httpx.OkJson(w, &Result{Status: 0, ErrorCode: 1000, Error: "NOT_FOUND"})
	}
}

func unauthorizedHandler(w http.ResponseWriter, r *http.Request, err error) {
	httpx.OkJson(w, &Result{Status: 0, ErrorCode: 1001, Error: "UNAUTHORIZED", Data: err.Error()})
}
