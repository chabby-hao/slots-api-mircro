package svc

import (
	"github.com/tal-tech/go-zero/zrpc"
	"inbox/api/internal/config"
	"inbox/rpc/inboxclient"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc inboxclient.Inbox
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: inboxclient.NewInbox(zrpc.MustNewClient(c.InboxRpc)),
	}
}
