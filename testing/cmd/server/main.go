package main

import (
	"context"
	"net"
	"os"

	"github.com/go-kit/kit/log/level"
	"github.com/marco-kit/kit-home-service/pkg/pb/protocols/testing/testing"
	"github.com/oklog/run"
	"github.com/testing/pkg/endpoint"
	"github.com/testing/pkg/service"
	"github.com/testing/transports"
	"google.golang.org/grpc"
)

func main() {
	cfg := HandleCfg()
	logger := SetupLogger(cfg)
	level.Info(logger).Log("msg", "server started")

	var (
		grpcServer *grpc.Server

		service     = service.NewService(logger)
		endpoints   = endpoint.NewEndpointSetup(service, logger)
		grpcHandler = transports.NewGRPCServer(*endpoints, nil)
	)

	grpcServer = grpc.NewServer()
	testing.RegisterTestingServiceServer(grpcServer, grpcHandler)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var g run.Group
	{
		grpcListener, err := net.Listen("tcp", cfg.GRPCAddr)
		if err != nil {
			level.Error(logger).Log("msg", "failed to listen on grpc address", "err", err)
			os.Exit(1)
		}

		g.Add(func() error {
			return grpcServer.Serve(grpcListener)
		}, func(error) {
			grpcServer.GracefulStop()
		})
	}

}
