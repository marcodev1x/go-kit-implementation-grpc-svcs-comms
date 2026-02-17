package service

import (
	"context"
	"fmt"

	"github.com/go-kit/log"
	"github.com/marco-kit/kit-home-service/connectors"
	"github.com/marco-kit/kit-home-service/pkg/pb/protocols/testing/testing"
)

type Service interface {
	Test(ctx context.Context, request *testing.TestRequest) (*testing.TestResponse, error)
}

type service struct {
	logger  log.Logger
	testing testing.TestingServiceClient
}

func NewService(logger log.Logger) Service {
	return &service{
		logger:  logger,
		testing: testing.NewTestingServiceClient(connectors.Testing()),
	}
}

func (s *service) Test(ctx context.Context, request *testing.TestRequest) (*testing.TestResponse, error) {
	response, err := s.testing.Test(ctx, request)

	if err != nil {
		return nil, err
	}

	fmt.Println("test")

	return response, nil
}
