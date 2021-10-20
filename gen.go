package main

// service
//nolint:lll
//go:generate protoc proto/service.proto --gogo_out=plugins=grpc,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types:api --grpc-gateway_out=logtostderr=true:api --openapiv2_opt logtostderr=true --openapiv2_out api -I=proto;$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway;$GOPATH/src/github.com/srikrsna/protoc-gen-gotag;$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis;$GOPATH/src/github.com/gogo/protobuf;

// go-tags
//nolint:lll
//go:generate protoc-go-inject-tag -input=./api/service.pb.go -XXX_skip=yaml,xml,bson

// typescript client
//go:generate openapi-generator-cli generate
