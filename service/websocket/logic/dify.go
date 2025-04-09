package logic

import (
	"context"
	"fmt"
	"github.com/fasthttp/websocket"
	"logic-server/common/dify"
	"logic-server/common/global"
	"logic-server/common/speech/speechclient"
	"logic-server/service/websocket/util"
	"regexp"
	"strings"
	"time"
)

func (l *Websocket) DifyMessage(conn *websocket.Conn, reData global.ReData) (err error) {

	// 用于接收http阻塞ch
	var ch chan dify.ChatMessageStreamChannelResponse

	chStr := make(chan *dify.VideoMessageStreamChannelResponse)

	ctx, _ := context.WithTimeout(context.Background(), 300*time.Second)

	req := &dify.ChatMessageRequest{
		Query:          reData.Text,
		ConversationID: reData.SessionId,
		User:           "my",
	}

	if ch, err = l.svcCtx.DifyClient.Api().ChatMessagesStream(ctx, req); err != nil {
		return
	}

	// 异步调用http请求聊天流
	go l.chatStream(conn, ch, chStr, reData)

	util.SendSocketReData(conn, global.ReData{
		Type:        "ttsUrl",
		State:       "start",
		SessionId:   reData.SessionId,
		Emotion:     "",
		Version:     "",
		Transport:   "",
		AudioParams: nil,
	})

	ctxA, _ := context.WithTimeout(context.Background(), 300*time.Second)

	for {
		select {
		case <-ctxA.Done():
			return
		case text := <-chStr:
			output := removePunctuationRegex(text.Text)
			if len(output) > 0 {
				// 调用语音合成RPC
				res, errf := l.svcCtx.SpeechRpc.PaddleSpeech(ctxA, &speechclient.PaddleSpeechReq{
					Text: output,
				})
				if errf == nil {
					util.SendSocketReData(conn, global.ReData{
						Type:        "ttsUrl",
						State:       "sentence_start",
						Text:        res.Address,
						SessionId:   reData.SessionId,
						Emotion:     "",
						Version:     "",
						Transport:   "",
						AudioParams: nil,
					})
				} else {
					// todo 录音失败 现在还没想好咋返回
					fmt.Println("录音失败" + output)
					//return
				}
			}
			// 判断是否结束 结束退出程序
			if text.State {
				util.SendSocketReData(conn, global.ReData{
					Type:        "ttsUrl",
					State:       "end",
					SessionId:   reData.SessionId,
					Emotion:     "",
					Version:     "",
					Transport:   "",
					AudioParams: nil,
				})
				return
			}

		}

	}

}

func (l *Websocket) chatStream(conn *websocket.Conn,
	ch chan dify.ChatMessageStreamChannelResponse, chStr chan *dify.VideoMessageStreamChannelResponse, reData global.ReData) {

	var strBuilder strings.Builder

	ctx, _ := context.WithTimeout(context.Background(), 300*time.Second)

	// 发送开始
	util.SendSocketReData(conn, global.ReData{
		Type:        "tts",
		State:       "start",
		SessionId:   reData.SessionId,
		Emotion:     "",
		Version:     "",
		Transport:   "",
		AudioParams: nil,
	})

	for {
		select {
		case <-ctx.Done():
			return
		case streamData, isOpen := <-ch:
			if err := streamData.Err; err != nil {
				// 错误也发送结束
				util.SendSocketReData(conn, global.ReData{
					Type:        "tts",
					State:       "end",
					SessionId:   reData.SessionId,
					Emotion:     "",
					Version:     "",
					Transport:   "",
					AudioParams: nil,
				})
				return
			}
			if !isOpen {
				strData := strings.Replace(strBuilder.String(), " ", "", -10)
				strData = strings.Replace(strData, "\n", "", -10)

				// 把剩余的发送出去
				util.SendSocketReData(conn, global.ReData{
					Type:        "tts",
					State:       "sentence_start",
					Text:        strData,
					SessionId:   reData.SessionId,
					Emotion:     "",
					Version:     "",
					Transport:   "",
					AudioParams: nil,
				})

				// 结束了 发送结束
				util.SendSocketReData(conn, global.ReData{
					Type:        "tts",
					State:       "end",
					SessionId:   reData.SessionId,
					Emotion:     "",
					Version:     "",
					Transport:   "",
					AudioParams: nil,
				})
				// 生成最后的语音
				chStr <- &dify.VideoMessageStreamChannelResponse{
					Text:  strData,
					State: true,
				}
				return
			}

			// 将接收的数据存到strBuilder
			strBuilder.WriteString(streamData.Answer)

			// 判断是否大于20个字符
			if len(strBuilder.String()) > 20 {
				strData := strings.Replace(strBuilder.String(), " ", "", -10)
				strData = strings.Replace(strData, "\n", "", -10)

				// 根据ReData 解析数据
				util.SendSocketReData(conn, global.ReData{
					Type:        "tts",
					State:       "sentence_start",
					Text:        strData,
					SessionId:   reData.SessionId,
					Emotion:     "",
					Version:     "",
					Transport:   "",
					AudioParams: nil,
				})
				// 调用语音生成
				chStr <- &dify.VideoMessageStreamChannelResponse{
					Text:  strData,
					State: false,
				}
				strBuilder = strings.Builder{}

			}
		}

	}

}

func removePunctuationRegex(s string) string {
	re := regexp.MustCompile(`[\p{P}\p{S}]`) // 匹配标点和符号（如数学符号）
	return re.ReplaceAllString(s, "")
}
