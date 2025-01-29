package service

import (
	"context"
	"time"

	api "github.com/michaelmass/hellomicro/api/proto"
	empty "google.golang.org/protobuf/types/known/emptypb"
)

func (service *Service) Latency(ctx context.Context, request *api.LatencyReq) (*empty.Empty, error) {
	time.Sleep(time.Duration(request.Duration) * time.Millisecond)
	return e, nil
}
