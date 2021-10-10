package handler

import (
	"net/http"
	"user/service/user/api/internal/logic"
	"user/service/user/api/internal/svc"
	"user/service/user/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func UserHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserLogic(r.Context(), ctx)
		resp, err := l.User(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
