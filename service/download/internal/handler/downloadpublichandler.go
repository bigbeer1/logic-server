package handler

import (
	"logic-server/common/responsex"
	"logic-server/service/download/internal/logic"
	"logic-server/service/download/internal/svc"
	"net/http"
)

func DownloadPublicHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewDownloadPublicLogic(r.Context(), svcCtx)
		resp, err := l.DownloadPublic(r.URL.Path)
		if err != nil {
			responsex.HttpResult(r, w, "", resp, err)
		} else {
			w.Write(resp.Data)
		}
	}
}
