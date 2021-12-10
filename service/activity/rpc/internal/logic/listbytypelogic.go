package logic

import (
	"context"
	"math/rand"
	"time"

	"activity/activity"
	"activity/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type ListByTypeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListByTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListByTypeLogic {
	return &ListByTypeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListByTypeLogic) ListByType(in *activity.ListByTypeRequest) (*activity.ListByTypeResponse, error) {

	var lst []*activity.ActivityInfo
	for i := 0; i < 5; i++ {
		lst = append(lst, &activity.ActivityInfo{
			Id:      int64(i),
			Type:    in.Type,
			Level:   rand.Int63n(3),
			Status:  rand.Int63n(3),
			StartAt: int64(time.Now().Unix()),
			EndAt:   int64(time.Now().AddDate(0, 0, 1).Unix()),
		})
	}

	return &activity.ListByTypeResponse{
		List: lst,
	}, nil
}
