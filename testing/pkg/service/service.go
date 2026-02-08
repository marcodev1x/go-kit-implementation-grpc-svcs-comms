package service

import (
	"context"

	"github.com/go-kit/log"
	"github.com/marco-kit/kit-home-service/pkg/pb/protocols/testing/testing"
)

type Service interface {
	Test(ctx context.Context, request *testing.TestRequest) (*testing.TestResponse, error)
}

type service struct {
	logger log.Logger
}

func NewService(logger log.Logger) Service {
	return &service{
		logger: logger,
	}
}

func (s *service) Test(ctx context.Context, request *testing.TestRequest) (*testing.TestResponse, error) {
	return &testing.TestResponse{
		Message: "Hello " + request.Name,
	}, nil
}
