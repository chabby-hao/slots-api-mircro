package handler

import (
	"inbox/api/internal/logic"
	"inbox/api/internal/svc"
	"inbox/api/internal/types"
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func InboxHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewInboxLogic(r.Context(), ctx)
		resp, err := l.Inbox(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
