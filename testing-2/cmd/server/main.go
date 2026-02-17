package main

import (
	"context"
	"net"
	"net/http"
	httpCaller "net/http"
	"os"

	"github.com/go-kit/kit/log/level"
	"github.com/marco-kit/kit-home-service/pkg/pb/protocols/testing/testing"
	"github.com/oklog/run"
	"github.com/testing-2/pkg/endpoint"
	"github.com/testing-2/pkg/service"
	"github.com/testing-2/transports"
	"google.golang.org/grpc"
)

func main() {
	cfg := HandleCfg()
	logger := SetupLogger(cfg)
	level.Info(logger).Log("msg", "server started")

	var (
		grpcServer *grpc.Server

		svc         = service.NewService(logger)
		endpoints   = endpoint.NewEndpointSetup(svc, logger)
		grpcHandler = transports.NewGRPCServer(*endpoints)
		httpHandler = transports.NewHTTPServer(*endpoints, logger)

		httpServer *httpCaller.Server
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
	{
		httpListener, err := net.Listen("tcp", cfg.HttpAddr)
		if err != nil {
			level.Error(logger).Log("msg", "failed to listen on http address", "err", err)
			os.Exit(1)
		}

		g.Add(func() error {
			strip := http.StripPrefix("/api", httpHandler)

			httpServer = &httpCaller.Server{
				Handler: strip,
			}

			return httpServer.Serve(httpListener)
		}, func(error) {
			level.Error(logger).Log("msg", "failed to listen on http address", "err", err)
		})
	}

	level.Info(logger).Log("msg", "starting servers")
	if err := g.Run(); err != nil {
		level.Error(logger).Log("msg", "servers failed", "err", err)
		os.Exit(1)
	}
}
