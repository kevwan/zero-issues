// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"primitive/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/bool",
				Handler: BoolHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/bools",
				Handler: BoolsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/int",
				Handler: IntHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/ints",
				Handler: IntsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/string",
				Handler: StringHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/strings",
				Handler: StringsHandler(serverCtx),
			},
		},
	)
}