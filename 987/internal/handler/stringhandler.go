package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"primitive/internal/logic"
	"primitive/internal/svc"
)

func StringHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewStringLogic(r.Context(), ctx)
		resp, err := l.String()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
