package service

import (
	"context"
	"os"
	"strings"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/michaelmass/hellomicro/api"
)

func (service *Service) Environments(ctx context.Context, request *empty.Empty) (*api.EnvironmentsRes, error) {
	environments := map[string]string{}

	for _, line := range os.Environ() {
		keyval := strings.Split(line, "=")
		environments[keyval[0]] = keyval[1]
	}

	return &api.EnvironmentsRes{
		Environments: environments,
	}, nil
}
