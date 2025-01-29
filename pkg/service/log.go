package service

import (
	"context"

	"github.com/google/uuid"
	api "github.com/michaelmass/hellomicro/api/proto"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	empty "google.golang.org/protobuf/types/known/emptypb"
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
