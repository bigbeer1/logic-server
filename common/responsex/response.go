package responsex

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

// http返回
func HttpResult(r *http.Request, w http.ResponseWriter, req interface{}, resp interface{}, err error) {
	//var isRequest int64
	// 请求判断
	if err != nil {
		// 失败
		//isRequest = 0
		logx.WithContext(r.Context()).Errorf("【API-ERR】 : %+v ", err)
		resp = err
	} else {
		// 成功
		//isRequest = 1
	}

	httpx.WriteJson(w, http.StatusOK, resp)

}
