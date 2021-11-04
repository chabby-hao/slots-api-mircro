package svc

import (
	"rmq/internal/config"
)

type ServiceContext struct {
	Config config.Config
	//Consumer dq.Consumer
	//Consumerkq kq.ConsumeHandler
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		//Consumer: dq.NewConsumer(c.DqConf),
		//Consumerkq: kq.MustNewQueue(c.KqConf, ),
	}
}
