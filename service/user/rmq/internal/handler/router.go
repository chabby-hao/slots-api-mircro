package handler

import (
	"context"
	"github.com/tal-tech/go-zero/core/service"
	"rmq/internal/logic"
	"rmq/internal/svc"
)

func RegisterJob(serverCtx *svc.ServiceContext, group *service.ServiceGroup) {

	group.Add(logic.NewProducerLogic(context.Background(), serverCtx))
	group.Add(logic.NewConsumerLogic(context.Background(), serverCtx))

	group.Start()

}
