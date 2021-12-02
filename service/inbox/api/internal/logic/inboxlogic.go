package logic

import (
	"context"
	"github.com/tal-tech/go-zero/core/logx"
	"inbox/api/internal/svc"
	"inbox/api/internal/types"
	"inbox/rpc/inbox"
)

type InboxLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInboxLogic(ctx context.Context, svcCtx *svc.ServiceContext) InboxLogic {
	return InboxLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InboxLogic) Inbox(req types.Request) (*types.Response, error) {
	// todo: add your logic here and delete this line

	res, err := l.svcCtx.UserRpc.Ping(l.ctx, &inbox.Request{
		Ping: req.Name,
	})
	if err != nil {
		return nil, err
	}

	return &types.Response{
		Message: res.Pong,
	}, nil
}
