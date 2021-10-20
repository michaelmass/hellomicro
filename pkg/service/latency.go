package service

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/michaelmass/hellomicro/api"
)

func (service *Service) Latency(ctx context.Context, request *api.LatencyReq) (*empty.Empty, error) {
	time.Sleep(time.Duration(request.Duration) * time.Millisecond)
	return e, nil
}
