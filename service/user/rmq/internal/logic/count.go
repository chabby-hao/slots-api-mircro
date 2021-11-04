package logic

import (
	"context"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/threading"
	"rmq/internal/svc"
)

type Counter struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Counter {
	return &Counter{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *Counter) Start() {
	logx.Infof("start consumer")

	threading.GoSafe(func() {
		mp := make(map[int]int)
		for v := range l.svcCtx.CountCh {
			_, ok := mp[v]
			if !ok {
				mp[v] = 1
			} else {
				mp[v] += 1
			}

			for k, v := range mp {
				if v != 2 {
					fmt.Printf("%v: %v\n", k, v)
				}
			}
		}
	})
}

func (l *Counter) Stop() {
	logx.Infof("stop consumer")
}
