package main

import (
	"flag"
	"fmt"
	"net/http"

	"issue/internal/config"
	"issue/internal/handler"
	"issue/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/search-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf, rest.WithUnauthorizedCallback(
		func(w http.ResponseWriter, r *http.Request, err error) {
			httpx.WriteJson(w, http.StatusUnauthorized, "oops")
		}))
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
