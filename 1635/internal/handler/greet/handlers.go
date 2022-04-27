package handler

import (
	"net/http"

	"1635/internal/logic/greet"
	"1635/internal/svc"
	"1635/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func HelloHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewHelloLogic(r.Context(), ctx)
		resp, err := l.Hello(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func GoodByeHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGoodByeLogic(r.Context(), ctx)
		resp, err := l.GoodBye(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
