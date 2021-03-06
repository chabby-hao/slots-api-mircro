package main

import (
	"flag"
	"fmt"
	"gitlab.haloapps.com/batatagames/slots/backend/slots-api-micro/user/rpc/internal/config"
	"gitlab.haloapps.com/batatagames/slots/backend/slots-api-micro/user/rpc/internal/server"
	"gitlab.haloapps.com/batatagames/slots/backend/slots-api-micro/user/rpc/internal/svc"
	"gitlab.haloapps.com/batatagames/slots/backend/slots-api-micro/user/rpc/rpc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/service"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/rpc.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewRpcServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		rpc.RegisterRpcServer(grpcServer, srv)

		switch c.Mode {
		case service.DevMode, service.TestMode:
			reflection.Register(grpcServer)
		default:
		}

	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
