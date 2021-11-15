// Code generated by goctl. DO NOT EDIT!
// Source: rpc.proto

package rpcclient

import (
	"context"
	"gitlab.haloapps.com/batatagames/slots/backend/slots-api-micro/user/rpc/rpc"

	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Request  = rpc.Request
	Response = rpc.Response

	Rpc interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	}

	defaultRpc struct {
		cli zrpc.Client
	}
)

func NewRpc(cli zrpc.Client) Rpc {
	return &defaultRpc{
		cli: cli,
	}
}

func (m *defaultRpc) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := rpc.NewRpcClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}
