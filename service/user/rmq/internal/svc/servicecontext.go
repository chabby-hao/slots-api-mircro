package svc

import (
	"rmq/internal/config"
)

type ServiceContext struct {
	Config  config.Config
	CountCh chan int
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		CountCh: make(chan int, 10),
	}
}
