package test

import (
	"context"
	"log"
	"logic-server/common/dify"
	"strings"
	"testing"
)

var (
	host         = "http://localhost/v1"
	apiSecretKey = "app-57CZKK39N8CQHimHWsm5i9rz"
)

func TestGet1009(t *testing.T) {

	var (
		ctx = context.Background()
		c   = dify.NewClient(host, apiSecretKey)

		req = &dify.ChatMessageRequest{
			Query: "你好吗",
			User:  "nima",
		}

		ch  chan dify.ChatMessageStreamChannelResponse
		err error
	)

	if ch, err = c.Api().ChatMessagesStream(ctx, req); err != nil {
		return
	}

	var strBuilder strings.Builder

	for {
		select {
		case <-ctx.Done():
			return
		case streamData, isOpen := <-ch:
			if err = streamData.Err; err != nil {
				log.Println(err.Error())
				return
			}
			if !isOpen {
				log.Println(strBuilder.String())
				return
			}

			strBuilder.WriteString(streamData.Answer)
		}
	}

}

func TestGet1005(t *testing.T) {

}

type ReData struct {
	Type        string      `json:"type"`
	State       string      `json:"state"`
	Text        string      `json:"text"`
	SessionId   string      `json:"session_id"`
	Emotion     string      `json:"emotion,omitempty"`
	Version     string      `json:"version,omitempty"`
	Transport   string      `json:"transport,omitempty"`
	AudioParams interface{} `json:"audio_params,omitempty"`
}
