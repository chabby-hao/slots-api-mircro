// Code generated by goctl. DO NOT EDIT!
// Source: inbox.proto

package inboxclient

import (
	"context"
	"inbox/rpc/inbox"

	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Request  = inbox.Request
	Response = inbox.Response

	Inbox interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	}

	defaultInbox struct {
		cli zrpc.Client
	}
)

func NewInbox(cli zrpc.Client) Inbox {
	return &defaultInbox{
		cli: cli,
	}
}

func (m *defaultInbox) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := inbox.NewInboxClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}