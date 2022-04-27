// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	greet "1635/internal/handler/greet"
	"1635/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/hello/:name",
				Handler: greet.HelloHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/good-bye/:name",
				Handler: greet.GoodByeHandler(serverCtx),
			},
		},
	)
}
