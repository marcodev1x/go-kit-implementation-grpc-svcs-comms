package transports

import (
	"context"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/marco-kit/kit-home-service/decode"
	"github.com/marco-kit/kit-home-service/pkg/pb/protocols/testing/testing"
	"github.com/testing/pkg/endpoint"
)

type GRPCServer struct {
	testing.UnimplementedTestingServiceServer

	test grpctransport.Handler
}

func NewGRPCServer(endpoints endpoint.EndpointSetup, test grpctransport.Handler) testing.TestingServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerBefore(decode.GRPCParams),
	}

	return &GRPCServer{
		test: grpctransport.NewServer(
			endpoints.Test,
			decodeGRPCRequest,
			decodeGRPCResponse,
			options...,
		),
	}
}

func decodeGRPCRequest(ctx context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

func decodeGRPCResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	return resp, nil
}
