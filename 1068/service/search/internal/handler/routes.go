// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"user/service/search/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckLogin},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/api/search",
					Handler: searchHandler(serverCtx),
				},
			}...,
		),
	)
}