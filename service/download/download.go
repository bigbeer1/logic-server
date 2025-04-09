package main

import (
	"flag"
	"fmt"
	zero_handler "github.com/zeromicro/go-zero/rest/handler"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/go-zero/zrpc"
	"logic-server/common"
	"logic-server/common/msg"
	"logic-server/service/download/internal/config"
	"logic-server/service/download/internal/handler"
	"logic-server/service/download/internal/svc"
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/download-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	server := rest.MustNewServer(c.RestConf,
		// token错误拦截
		rest.WithUnauthorizedCallback(func(w http.ResponseWriter, r *http.Request, err error) {
			httpx.WriteJson(w, http.StatusOK, common.NewCodeError(common.TokenErrorCode, msg.TokenError, err.Error()))
		}),
		// 请求方式错误拦截
		rest.WithNotAllowedHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			httpx.WriteJson(w, http.StatusOK, common.NewCodeError(common.ReqNotAllCode, msg.ReqNotAllError, nil))
		})),
		// 路由错误拦截
		rest.WithNotFoundHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			httpx.WriteJson(w, http.StatusOK, common.NewCodeError(common.ReqRoutesErrorCode, msg.ReqRoutesError, nil))
		})),
	)
	defer server.Stop()

	// 设置日志输出 接口慢时间
	zrpc.SetClientSlowThreshold(time.Second * 2000)
	zero_handler.SetSlowThreshold(time.Second * 2000)

	handler.RegisterHandlers(server, ctx)
	fmt.Println(fmt.Sprintf("%s %s:%v", c.Name, c.Host, c.Port))
	server.Start()
}
