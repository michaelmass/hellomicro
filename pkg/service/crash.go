package service

import (
	"context"
	"os"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/michaelmass/hellomicro/api"
)

func (service *Service) Crash(ctx context.Context, request *empty.Empty) (*empty.Empty, error) {
	go func() { panic("oops") }()
	return e, nil
}

func (service *Service) Panic(ctx context.Context, request *empty.Empty) (*empty.Empty, error) {
	panic("")
}

func (service *Service) Exit(ctx context.Context, request *api.ExitReq) (*empty.Empty, error) {
	os.Exit(int(request.Code))
	return e, nil
}
