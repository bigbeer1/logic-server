package speech

import (
	"context"
	"github.com/zeromicro/go-zero/zrpc"
	"logic-server/common/global"
	"logic-server/common/speech/speech"
	"logic-server/common/speech/speechclient"
)

func Speech(ctx context.Context, speechRpc zrpc.RpcClientConf, msg string) (string, error) {

	if global.SpeechClient == nil {

		speechRpc.Timeout = 50000

		global.SpeechClient = speech.NewSpeech(zrpc.MustNewClient(speechRpc))
	}

	res, err := global.SpeechClient.PaddleSpeech(ctx, &speechclient.PaddleSpeechReq{
		Text: msg,
	})

	if err != nil {
		return "", err
	}

	return res.Address, nil
}
