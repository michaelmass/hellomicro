package service

import (
	"context"
	"io"
	"net/http"
	"strings"

	"github.com/michaelmass/hellomicro/api"
	"github.com/pkg/errors"
)

func (service *Service) Request(ctx context.Context, request *api.RequestReq) (*api.RequestRes, error) {
	req, err := http.NewRequestWithContext(context.Background(), request.Method, request.Url, strings.NewReader(request.Body))

	if err != nil {
		return nil, errors.Wrap(err, "creating http request")
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, errors.Wrap(err, "sending http request")
	}

	content, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, errors.Wrap(err, "reading http request body")
	}

	return &api.RequestRes{
		Body:       string(content),
		StatusCode: int32(res.StatusCode),
		Headers:    MapSAToMapApiSA(res.Header),
	}, nil
}
