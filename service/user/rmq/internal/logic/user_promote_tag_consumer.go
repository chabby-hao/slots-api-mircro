package logic

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/threading"
	"github.com/twmb/franz-go/pkg/kgo"
	"rmq/ext/aws_kq"
	"rmq/ext/user_api_client"
	"rmq/internal/model"
	"rmq/internal/svc"
	"time"
)

type UserPromoteTagConsumer struct {
	logx.Logger
	ctx       context.Context
	svcCtx    *svc.ServiceContext
	q         *aws_kq.Consumer
	apiClient *user_api_client.ApiClient
}

func NewUserPromoteTagConsumerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserPromoteTagConsumer {

	c := &UserPromoteTagConsumer{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		//q:         aws_kq.MustNewQueue(svcCtx.Config.AwsKqConf, Consume),
		apiClient: user_api_client.NewUserApiClient(&svcCtx.Config.UserApiConf),
	}
	c.q = aws_kq.MustNewQueue(svcCtx.Config.AwsKqConf, c.Consume)
	return c
}

func (l *UserPromoteTagConsumer) Start() {
	logx.Infof("start consumer")

	threading.GoSafe(func() {
		l.q.Start()
	})
}

func (l *UserPromoteTagConsumer) Consume(record *kgo.Record) error {

	logx.Infof("record: %v", string(record.Value))

	var info = new(model.UserPromoteTagInfo)
	err := json.Unmarshal(record.Value, info)
	if err != nil {
		// 不能 json 解析的都认为是可以忽略的
		logx.Error(err)
		return nil
	}

	id1 := info.Idfv
	id2 := info.Idfa
	// 如果都为空则说明是android
	if id1 == "" && id2 == "" {
		id1 = info.Gaid
	}
	err = l.apiClient.SetUserPromoteTag(id1, id2, info.AppsflyerId, info.UserTag)
	if err != nil {
		if errors.Is(err, user_api_client.ErrorUserNotFound) {
			// 用户不存在，晚点再试
			go func() {
				for i := 0; i < 3; i++ {
					time.Sleep(5 * time.Second)
					err = l.apiClient.SetUserPromoteTag(id1, id2, info.AppsflyerId, info.UserTag)
					if err == nil {
						break
					} else {
						logx.Error(err)
					}
				}
			}()
		}

		return err
	}

	return nil
}

func (l *UserPromoteTagConsumer) Stop() {
	logx.Infof("stop consumer")

	l.q.Stop()
}
