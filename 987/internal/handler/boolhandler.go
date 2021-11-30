package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"primitive/internal/logic"
	"primitive/internal/svc"
)

func BoolHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewBoolLogic(r.Context(), ctx)
		resp, err := l.Bool()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
