package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"golang.org/x/time/rate"
	"logic-server/common/dify"
	"logic-server/common/speech/speech"
	"logic-server/service/websocket/config"
)

type ServiceContext struct {
	Config config.Config
	// websocket最大任务数 限流器
	TaskMaxLimit *rate.Limiter

	DifyClient *dify.Client

	SpeechRpc speech.Speech
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config:     c,
		DifyClient: dify.NewClient(c.DifyHost, c.DifyApiSecretKey),
		SpeechRpc:  speech.NewSpeech(zrpc.MustNewClient(c.SpeechRpc)),
	}
}
