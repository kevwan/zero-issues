package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"primitive/internal/logic"
	"primitive/internal/svc"
)

func IntsHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewIntsLogic(r.Context(), ctx)
		resp, err := l.Ints()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
