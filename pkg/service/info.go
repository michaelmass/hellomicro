package service

import (
	"context"
	"runtime"

	"github.com/michaelmass/hellomicro/api"
	empty "google.golang.org/protobuf/types/known/emptypb"
)

func (service *Service) Info(ctx context.Context, request *empty.Empty) (*api.InfoRes, error) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	return &api.InfoRes{
		Os:           runtime.GOOS,
		Version:      runtime.Version(),
		NumCpu:       int32(runtime.NumCPU()),
		NumGoRoutine: int32(runtime.NumGoroutine()),
		NumCgoCall:   runtime.NumCgoCall(),
		Memory: &api.InfoRes_Memory{
			Alloc:      m.Alloc,
			TotalAlloc: m.TotalAlloc,
			Sys:        m.Sys,
			NumGc:      m.NumGC,
		},
	}, nil
}
