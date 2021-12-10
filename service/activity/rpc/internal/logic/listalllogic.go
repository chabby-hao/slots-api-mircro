package logic

import (
	"context"
	"math/rand"
	"time"

	"activity/activity"
	"activity/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type ListAllLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListAllLogic {
	return &ListAllLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListAllLogic) ListAll(in *activity.ListAllRequest) (*activity.ListAllResponse, error) {

	var lst []*activity.ActivityInfo
	for i := 0; i < 5; i++ {
		lst = append(lst, &activity.ActivityInfo{
			Id:      int64(i),
			Type:    rand.Int63n(3),
			Level:   rand.Int63n(3),
			Status:  rand.Int63n(3),
			StartAt: int64(time.Now().Unix()),
			EndAt:   int64(time.Now().AddDate(0, 0, 1).Unix()),
		})
	}

	return &activity.ListAllResponse{
		List: lst,
	}, nil
}
