package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"user/service/search/internal/logic"
	"user/service/search/internal/svc"
	"user/service/search/internal/types"
)

func searchHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewSearchLogic(r.Context(), ctx)
		resp, err := l.Search(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
