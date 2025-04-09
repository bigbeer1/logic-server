package speech

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
	"logic-server/common/speech/speech"
	"logic-server/common/speech/speechclient"
	"testing"
	"time"
)

func TestConnectSpeech(t *testing.T) {

	cc := zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts:              []string{"127.0.0.1:30000"},
			Key:                "speech.rpc",
			User:               "",
			Pass:               "",
			CertFile:           "",
			CertKeyFile:        "",
			CACertFile:         "",
			InsecureSkipVerify: false,
		},
		Endpoints: nil,
		Target:    "",
		App:       "",
		Token:     "",
		NonBlock:  false,
		Timeout:   30000,
	}

	client := speech.NewSpeech(zrpc.MustNewClient(cc))

	if client == nil {
		fmt.Println(11111)
	}
	// 创建30秒协程写入日志
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	msg, err := client.PaddleSpeech(ctx, &speechclient.PaddleSpeechReq{
		Text: "世界和平世界和平世界和平世界和平世界和平世界和平世界和平世界和平世界和平世界和平世界和平世界和平世界和平世界和平世界和平世界和平世界和平世界和平世界和平世界和平",
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(msg.Address)

}
