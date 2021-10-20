package service

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gogo/gateway"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/michaelmass/hellomicro/api"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

var e = &emptypb.Empty{}

type MetadataFunc func(c context.Context, r *http.Request) metadata.MD

var metadataFuncMap = map[string]MetadataFunc{
	"/v1/echo": echoMetadataFunc,
}

type Service struct {
	marshaler *gateway.JSONPb
	endpoint  string
	port      string
}

func New(endpoint string, port string) *Service {
	return &Service{
		endpoint: endpoint,
		port:     port,
		marshaler: &gateway.JSONPb{
			EmitDefaults: true,
			Indent:       "  ",
			OrigName:     true,
		},
	}
}

// Client returns a client for hellomicro
func Client(endpoint string, insecure bool, port string) api.HellomicroClient {
	var conn *grpc.ClientConn

	if insecure {
		conn, _ = grpc.Dial(endpoint, grpc.WithInsecure())
	} else {
		creds := credentials.NewClientTLSFromCert(nil, "")
		conn, _ = grpc.Dial(endpoint, grpc.WithTransportCredentials(creds))
	}

	return api.NewHellomicroClient(conn)
}

// Run starts the service
func (service *Service) Run() {
	port := fmt.Sprintf(":%s", service.port)
	l, err := net.Listen("tcp", port)

	if err != nil {
		logrus.WithError(errors.Wrapf(err, "failed to listen to %s", port)).Fatal()
	}

	m := cmux.New(l)
	grpcL := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	gatewayL := m.Match(cmux.Any())

	go service.startGRPC(grpcL)
	go service.startHTTP(gatewayL)
	go m.Serve()

	defer m.Close()

	select {}
}

func (service *Service) startGRPC(l net.Listener) {
	s := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_recovery.StreamServerInterceptor(),
			grpc_validator.StreamServerInterceptor(),
			grpc_prometheus.StreamServerInterceptor,
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(),
			grpc_validator.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
		)),
	)

	grpc_prometheus.EnableHandlingTimeHistogram(
		grpc_prometheus.WithHistogramBuckets(prometheus.DefBuckets),
	)

	reflection.Register(s)
	grpc_prometheus.Register(s)

	api.RegisterHellomicroServer(s, service)

	err := s.Serve(l)

	if err != nil {
		logrus.WithError(errors.Wrap(err, "failed to serve grpc")).Fatal()
	}
}

func (service *Service) startHTTP(l net.Listener) {
	gwmux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, service.marshaler),
		runtime.WithProtoErrorHandler(runtime.DefaultHTTPProtoErrorHandler),
		runtime.WithMetadata(func(c context.Context, r *http.Request) metadata.MD {
			metadataFunc, ok := metadataFuncMap[r.URL.Path]

			if !ok {
				return nil
			}

			return metadataFunc(c, r)
		}),
	)

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithUserAgent("gateway-loopback"),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:    time.Second * 10,
			Timeout: time.Millisecond * 100,
		}),
	}

	err := api.RegisterHellomicroHandlerFromEndpoint(context.Background(), gwmux, fmt.Sprintf("%s:%s", "127.0.0.1", service.port), opts)

	if err != nil {
		logrus.WithError(errors.Wrap(err, "failed to serve http gateway")).Fatal()
	}

	http.Serve(l, service.handler(gwmux))
}
