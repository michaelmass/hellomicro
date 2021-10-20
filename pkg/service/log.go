package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"github.com/michaelmass/hellomicro/api"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func (service *Service) Log(ctx context.Context, request *api.LogReq) (*empty.Empty, error) {
	uuid, err := uuid.NewRandom()

	if err != nil {
		return nil, errors.Wrap(err, "creating random uuid")
	}

	for i := 1; i <= int(request.Count); i++ {
		logrus.Infof("%s: %d", uuid.String(), i)
	}

	return e, nil
}
