package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"issue/internal/logic"
	"issue/internal/svc"
	"issue/internal/types"
)

func userLoginHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserLoginLogic(r.Context(), ctx)
		resp, err := l.UserLogin(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
