package main

import (
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/service"
	"os"
	"os/signal"
	"rmq/internal/config"
	"rmq/internal/handler"
	"rmq/internal/svc"
	"syscall"
	"time"
)

var configFile = flag.String("f", "etc/rmq.yaml", "the config file")

func main() {
	for {
		fmt.Println("hello: " + time.Now().String())
		time.Sleep(3 * time.Second)
	}

	flag.Parse()

	//配置
	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	//注册job
	group := service.NewServiceGroup()
	handler.RegisterJob(ctx, group)

	//捕捉信号
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-ch
		logx.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			fmt.Printf("stop group")
			group.Stop()
			logx.Info("job exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
