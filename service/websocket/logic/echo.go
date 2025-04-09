package logic

import (
	"encoding/json"
	"github.com/fasthttp/websocket"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
	"logic-server/common/global"
	"logic-server/service/websocket/util"
)

var upgrader = websocket.FastHTTPUpgrader{
	CheckOrigin: func(ctx *fasthttp.RequestCtx) bool {
		return true
	},
}

func (l *Websocket) EchoView(ctx *fasthttp.RequestCtx) {

	err := upgrader.Upgrade(ctx, func(ws *websocket.Conn) {
		// ip和端口
		//remoteAddr := ws.RemoteAddr().String()
		// 设备

		defer func() {
			global.Ulock.Lock()
			ws.Close()
			global.Ulock.Unlock()
		}()

		err := util.SendSocketReData(ws, global.ReData{
			Type:        "hello",
			State:       "",
			Text:        "",
			SessionId:   uuid.New().String(),
			Emotion:     "",
			Version:     "",
			Transport:   "",
			AudioParams: nil,
		})
		if err != nil {
			logx.Errorf(err.Error())
		}

		for {
			_, message, err := ws.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}

			// 限流器
			if l.svcCtx.TaskMaxLimit.Allow() == false {
				err = util.SendSocket(util.SendMessage{ws, []byte("error|服务器任务数超过限制请稍后再试")})
				if err != nil {
					logx.Errorf(err.Error())
				}
				continue
			}

			// 根据ReData 解析数据
			var reData global.ReData

			err = json.Unmarshal(message, &reData)
			if err != nil {
				continue
				//err = util.SendSocket(util.SendMessage{ws, []byte("error|数据解析格式不对")})
			}

			switch reData.Type {
			case "ping":
				err = util.SendSocket(util.SendMessage{ws, []byte("pong|")})
				if err != nil {
					logx.Errorf(err.Error())
				}
			case "listen":
				err = l.DifyMessage(ws, reData)
				if err != nil {
					logx.Errorf(err.Error())
				}

			default:
				//err = util.SendSocket(util.SendMessage{ws, []byte("error|不支持指令|")})
				//if err != nil {
				//	logx.Errorf(err.Error())
				//}
			}
		}
	})

	if err != nil {
		if _, ok := err.(websocket.HandshakeError); ok {
			log.Println(err)
		}
		return
	}
}
