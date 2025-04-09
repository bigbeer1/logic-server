package logic

import (
	"github.com/valyala/fasthttp"
	"logic-server/service/websocket/svc"
)

type Websocket struct {
	svcCtx *svc.ServiceContext
}

func NewCronScheduler(svcCtx *svc.ServiceContext) *Websocket {
	return &Websocket{
		svcCtx: svcCtx,
	}
}

func (l *Websocket) Register() (requestHandler fasthttp.RequestHandler) {

	requestHandler = func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/":
			l.EchoView(ctx)
		case "/home":
			l.Home(ctx)
		default:
			ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		}
	}
	return requestHandler
}
