package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"github.com/marco-kit/kit-home-service/pkg/pb/protocols/testing/testing"
	"github.com/testing/pkg/service"
)

type EndpointSetup struct {
	Test endpoint.Endpoint
}

func NewEndpointSetup(s service.Service, logger log.Logger) *EndpointSetup {
	var testEndpoint endpoint.Endpoint
	{
		testEndpoint = MakeTestEndpoint(s)
		logger.Log("Endpoint value", "ok")
	}

	return &EndpointSetup{
		Test: testEndpoint,
	}
}

func MakeTestEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r := request.(*testing.TestRequest)
		rpcRequest := &testing.TestRequest{
			Name: r.Name,
		}

		fc, err := s.Test(ctx, rpcRequest)

		if err != nil {
			return &Resp{
				Error: err,
			}, nil
		}
		return &Resp{
			Items: fc,
		}, nil
	}
}

type Resp struct {
	Error  error       `json:"error,omitempty"`
	Items  interface{} `json:"items,omitempty"`
	Total  int64       `json:"total,omitempty"`
	Cursor string      `json:"cursor,omitempty"`
}
