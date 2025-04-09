.PHONY: build
build: linux_env



# linux 环境变量
linux_env:
	set GOARCH=amd64
	go env -w GOARCH=amd64
	set GOOS=windows
	go env -w GOOS=windows

# 管理端微服务
admin:
	go build -o deploy/golang/manage/out/admin/rpc/admin-rpc service/manage/admin/rpc/admin.go
	go build -o deploy/golang/manage/out/admin/api/admin-api service/manage/admin/api/admin.go

# 代理商端微服务
agent:
	go build -o deploy/golang/manage/out/agent/rpc/agent-rpc service/manage/agent/rpc/agent.go
	go build -o deploy/golang/manage/out/agent/api/agent-api service/manage/agent/api/agent.go

# 用户微服务
user:
	go build -o deploy/golang/manage/out/user/rpc/user-rpc service/manage/user/rpc/user.go
	go build -o deploy/golang/manage/out/user/api/user-api service/manage/user/api/user.go

# 工作流微服务
workflow:
	go build -o deploy/golang/manage/out/workflow/rpc/workflow-rpc service/manage/workflow/rpc/workflow.go
	go build -o deploy/golang/manage/out/workflow/api/workflow-api service/manage/workflow/api/workflow.go

# 日志微服务
archive:
	go build -o deploy/golang/manage/out/archive/rpc/archive-rpc service/manage/archive/rpc/archive.go
	go build -o deploy/golang/manage/out/archive/api/archive-api service/manage/archive/api/archive.go

# 鉴权微服务
authentication:
	go build -o deploy/golang/manage/out/authentication/authentication-rpc service/manage/authentication/authentication.go

# 上传微服务
upload:
	go build -o deploy/golang/manage/out/upload/rpc/upload-rpc service/manage/upload/rpc/upload.go
	go build -o deploy/golang/manage/out/upload/api/upload-api service/manage/upload/api/upload.go

# 下载微服务
download:
	go build -o deploy/golang/manage/out/download/api/download-api service/manage/download/download.go


# 资产微服务
asset:
	go build -o deploy/golang/twy/out/app/rpc/asset-rpc service/twy/asset/rpc/asset.go
	go build -o deploy/golang/twy/out/app/api/asset-api service/twy/asset/api/asset.go

# 运维检修微服务
operation:
	go build -o deploy/golang/twy/out/app/rpc/operation-rpc service/twy/operation/rpc/operation.go
	go build -o deploy/golang/twy/out/app/api/operation-api service/twy/operation/api/operation.go

# 在线监测微服务
monitor:
	go build -o deploy/golang/twy/out/app/rpc/monitor-rpc service/twy/monitor/rpc/monitor.go
	go build -o deploy/golang/twy/out/app/api/monitor-api service/twy/monitor/api/monitor.go

# 机器人微服务
robot:
	go build -o deploy/golang/twy/out/app/api/robot-api service/twy/robot/api/robot.go

# mqtt接收服务
mqtt:
	go build -o deploy/golang/twy/out/other/mqtt-client service/twy/mqtt/mqtt.go

# mqtt发送服务
mqtt-send:
	go build -o deploy/golang/twy/out/other/mqtt-send service/twy/mqttSend/mqttsend.go

# 导入服务
import:
	go build -o deploy/golang/twy/out/import/import service/twy/import/import.go

# 定时任务/异步任务服务
asynq-server:
	go build -o deploy/golang/twy/out/other/asynq-server service/twy/asynq/asynq-server/asynq-server.go
	go build -o deploy/golang/twy/out/other/scheduler service/twy/asynq/scheduler/scheduler.go

# WebSocket服务
websocket:
	go build -o deploy/golang/twy/out/websocket/websocket service/twy/socket/websocket/websocket.go

# 计量表计微服务
meter:
	go build -o deploy/golang/twy/out/app/rpc/meter-rpc service/twy/meter/rpc/meter.go
	go build -o deploy/golang/twy/out/app/api/meter-api service/twy/meter/api/meter.go
