package handler

import (
	"logic-server/common/responsex"
	"logic-server/service/download/internal/logic"
	"logic-server/service/download/internal/svc"
	"net/http"
)

func DownloadWavHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewDownloadWavLogic(r.Context(), svcCtx)
		resp, err := l.DownloadWav(r.URL.Path)
		if err != nil {
			responsex.HttpResult(r, w, "", resp, err)
		} else {
			w.Write(resp.Data)
		}
	}
}
