package handler

import (
	"context"
	"github.com/tal-tech/go-zero/core/service"
	"rmq/internal/logic"
	"rmq/internal/svc"
)

func RegisterJob(serverCtx *svc.ServiceContext, group *service.ServiceGroup) {

	group.Add(logic.NewUserPromoteTagConsumerLogic(context.Background(), serverCtx))

	group.Start()
}
