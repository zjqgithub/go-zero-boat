package main

import (
	"flag"
	"fmt"

	"boat/common/interceptor"
	"boat/rpc/user/internal/config"
	"boat/rpc/user/internal/server"
	"boat/rpc/user/internal/svc"
	"boat/rpc/user/pb/user"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServer(grpcServer, server.NewUserServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	//rpc log
	s.AddUnaryInterceptors(interceptor.LoggerInterceptor)

	logx.DisableStat()

	fmt.Printf("Starting a rpc server at %s...\n", c.ListenOn)
	s.Start()
}
