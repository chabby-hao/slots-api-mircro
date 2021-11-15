package svc

import "gitlab.haloapps.com/batatagames/slots/backend/slots-api-micro/user/rpc/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
