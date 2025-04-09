package main

import (
	"flag"
	"fmt"
	"github.com/valyala/fasthttp"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"golang.org/x/time/rate"
	"logic-server/service/websocket/config"
	"logic-server/service/websocket/logic"
	"logic-server/service/websocket/svc"
	"os"
	"runtime"
	"time"
)

var configFile = flag.String("f", "etc/websocket.yaml", "Specify the config file")

func main() {
	flag.Parse()
	var c config.Config

	conf.MustLoad(*configFile, &c)

	logx.DisableStat()
	if err := c.SetUp(); err != nil {
		panic(err)
	}

	cpuNum := runtime.NumCPU() //获得当前设备的cpu核心数
	fmt.Println("任务使用,cpu核心数:", cpuNum)
	runtime.GOMAXPROCS(cpuNum) //设置需要用到的cpu数量

	// 设置日志输出 接口慢时间  rpc
	zrpc.SetServerSlowThreshold(time.Second * 999000)

	ctx := svc.NewServiceContext(c)

	limit := rate.Every(time.Duration(c.TimeLimit) * time.Millisecond)
	ctx.TaskMaxLimit = rate.NewLimiter(limit, c.TaskMaxLimit)
	fmt.Println(fmt.Sprintf("限流器任务数:%v,速率:%v", c.TaskMaxLimit, c.TimeLimit))

	websocket := logic.NewCronScheduler(ctx)
	requestHandler := websocket.Register()

	server := fasthttp.Server{
		Name:        c.Name,
		Handler:     requestHandler,
		ReadTimeout: time.Minute,
	}

	if err := server.ListenAndServe(fmt.Sprintf("%v:%v", c.Host, c.Port)); err != nil {
		logx.Errorf("!!!WebsocketError: !!!  run err:%+v", err)
		os.Exit(1)
	}

}
