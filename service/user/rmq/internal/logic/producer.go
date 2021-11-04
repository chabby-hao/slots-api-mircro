package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/tal-tech/go-queue/kq"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/threading"
	"log"
	"math/rand"
	"rmq/internal/model"
	"rmq/internal/svc"
	"strconv"
	"time"
)

type Producer struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	pusher *kq.Pusher
}

func NewProducerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Producer {
	kqConf := svcCtx.Config.KqConf
	return &Producer{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		pusher: kq.NewPusher(kqConf.Brokers, kqConf.Topic),
	}
}

func (l *Producer) Start() {
	logx.Infof("start Producer")

	threading.GoSafe(func() {
		ticker := time.NewTicker(time.Second)
		for round := 0; round < 10; round++ {
			select {
			case <-ticker.C:
				count := rand.Intn(100)
				m := model.Message{
					Key:     strconv.FormatInt(time.Now().UnixNano(), 10),
					Value:   fmt.Sprintf("%d,%d", round, count),
					Payload: fmt.Sprintf("%d,%d", round, count),
				}
				body, err := json.Marshal(m)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println("push: " + string(body))
				if err := l.pusher.Push(string(body)); err != nil {
					log.Fatal(err)
				}
			}
		}
	})
}

func (l *Producer) Stop() {
	logx.Infof("stop Producer")
}
