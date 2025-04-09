package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	TaskMaxLimit int // 限流数量

	TimeLimit int //  限流时间

	DifyHost string // dify地址

	DifyApiSecretKey string // dify密钥

	SpeechRpc zrpc.RpcClientConf
}
