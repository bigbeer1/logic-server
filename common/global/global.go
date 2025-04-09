package global

import (
	"github.com/fasthttp/websocket"
	"logic-server/common/speech/speech"
	"sync"
	"time"
)

// 遥性解译
type YxInterpreting struct {
	Key   string `json:"key"`   // 键值
	Value string `json:"value"` // 对应显示
}

var (
	SConnection     = make(map[string]*SConnectionData) // socket列表 key为ip value为 会话conn
	OnlineList      = make(map[string][]*OnlineData)    // key为uid  value 为[]ip
	Ulock           sync.Mutex
	ShangHaiTime, _ = time.LoadLocation("Asia/Shanghai") //上海
	SpeechClient    speech.Speech
)

var WebSocketUserPrefix = "websocketUser:"

type SConnectionData struct {
	Uid  string          `json:"uid"`  // uid
	Conn *websocket.Conn `json:"conn"` // IP
}

type OnlineData struct {
	Ip                string `json:"ip"`                  // IP
	Device            string `json:"device"`              // 设备
	RegisterTime      int64  `json:"register_time"`       // 注册时间
	WebsocketClientId string `json:"websocket_client_id"` // 登录websocket订阅ID
}

type SocketMonitor struct {
	Id            int64  `json:"id"`
	ResultValue   string `json:"result_value"`   // 监测值
	UpdateTime    int64  `json:"update_time"`    // 更新时间
	Level         int64  `json:"level"`          // 告警等级  0正常
	RuleType      int64  `json:"rule_type"`      // 告警状态  1
	PointCategory int64  `json:"point_category"` // 类别：1:遥信/2:遥测/3:遥脉
	Unit          string `json:"unit"`           // 单位
	ModelKey      int64  `json:"model_key"`      //模型对应的Key
}

type SocketMonitorImage struct {
	Id             int64  `json:"id"`
	ResultValue    string `json:"result_value"`     // 监测值
	UpdateTime     int64  `json:"update_time"`      // 更新时间
	PointImageType int64  `json:"point_image_type"` // 类别：1:超声波
	ModelKey       int64  `json:"model_key"`        //模型对应的Key
}

type SocketMonitorGroup struct {
	MonitorPoints       []*SocketMonitor
	MonitorCombinations []*SocketMonitor
}

type WebsocketChatData struct {
	// 发送人ID
	PromoterId string `json:"promoter_id"`
	// 发送人IP
	PromoterIp string `json:"promoter_ip"`
	// 发送人名称
	PromoterName string `json:"promoter_name"`
	// 接收人ID
	RecipientId string `json:"recipient_id"`
	// 接收人ip和端口
	RecipientAddress string `json:"recipient_address"`
	// 内容
	Data string `json:"data"`
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
