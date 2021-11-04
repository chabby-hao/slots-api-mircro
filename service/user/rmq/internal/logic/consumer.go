package logic

import (
	"context"
	"fmt"
	"github.com/tal-tech/go-queue/kq"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/queue"
	"github.com/tal-tech/go-zero/core/threading"
	"rmq/internal/svc"
)

type Consumer struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	q queue.MessageQueue
}

func NewConsumerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Consumer {
	fmt.Println("%+v", svcCtx.Config.KqConf.Brokers)
	return &Consumer{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		q: kq.MustNewQueue(svcCtx.Config.KqConf, kq.WithHandle(func(key, value string) error {
			logx.Infof("consumer job  %s: %s", key, value)
			return nil
		})),
	}
}

func (l *Consumer) Start() {
	logx.Infof("start consumer")

	threading.GoSafe(func() {
		l.q.Start()
	})
}

func (l *Consumer) Stop() {
	logx.Infof("stop consumer")

	l.q.Stop()
}
