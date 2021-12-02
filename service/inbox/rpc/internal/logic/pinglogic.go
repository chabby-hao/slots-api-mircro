package logic

import (
	"context"
	"inbox/rpc/inbox"
	"inbox/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *inbox.Request) (*inbox.Response, error) {
	// todo: add your logic here and delete this line

	return &inbox.Response{
		Pong: in.Ping,
	}, nil
}
