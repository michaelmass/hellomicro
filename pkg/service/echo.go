package service

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"

	api "github.com/michaelmass/hellomicro/api/proto"
	"google.golang.org/grpc/metadata"
	empty "google.golang.org/protobuf/types/known/emptypb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

const echoMetadataKey = "echo"

type RequestEcho struct {
	Cookies []*http.Cookie
}

func echoMetadataFunc(c context.Context, r *http.Request) metadata.MD {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return nil
	}

	password, passwordSet := r.URL.User.Password()

	cookies := []*api.Cookie{}

	for _, cookie := range r.Cookies() {
		expires := timestamppb.New(cookie.Expires)

		cookies = append(cookies, &api.Cookie{
			Name:       cookie.Name,
			Value:      cookie.Value,
			Path:       cookie.Path,
			Domain:     cookie.Domain,
			Expires:    expires,
			RawExpires: cookie.RawExpires,
			MaxAge:     int32(cookie.MaxAge),
			Secure:     cookie.Secure,
			HttpOnly:   cookie.HttpOnly,
			SameSite:   api.Cookie_SameSite(cookie.SameSite),
			Raw:        cookie.Raw,
			Unparsed:   cookie.Unparsed,
		})
	}

	content, err := json.Marshal(&api.EchoRes{
		Method:        r.Method,
		Proto:         r.Proto,
		ContentLength: r.ContentLength,
		RemoteAddr:    r.RemoteAddr,
		RequestUri:    r.RequestURI,
		Host:          r.Host,
		UserAgent:     r.UserAgent(),
		Body:          string(body),
		Params:        MapSAToMapApiSA(r.URL.Query()),
		Headers:       MapSAToMapApiSA(r.Header),
		Cookies:       cookies,
		Url: &api.URL{
			Scheme:      r.URL.Scheme,
			Opaque:      r.URL.Opaque,
			Username:    r.URL.User.Username(),
			Password:    password,
			PasswordSet: passwordSet,
			Host:        r.URL.Host,
			Path:        r.URL.Path,
			RawPath:     r.URL.RawPath,
			RawQuery:    r.URL.RawQuery,
			Fragment:    r.URL.Fragment,
			RawFragment: r.URL.RawFragment,
			ForceQuery:  r.URL.ForceQuery,
		},
	})

	if err != nil {
		return nil
	}

	return metadata.Pairs(echoMetadataKey, string(content))
}

func (service *Service) Echo(ctx context.Context, request *empty.Empty) (*api.EchoRes, error) {
	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return nil, errors.New("unable to get metadata from incoming context")
	}

	values := md.Get(echoMetadataKey)

	if len(values) == 0 {
		return nil, errors.New("not request metadata found")
	}

	if len(values) > 1 {
		return nil, errors.New("more than one request metadata found")
	}

	res := &api.EchoRes{}

	err := json.Unmarshal([]byte(values[0]), &res)

	if err != nil {
		return nil, errors.Wrap(err, "unable to decode echo metadata")
	}

	return res, nil
}
