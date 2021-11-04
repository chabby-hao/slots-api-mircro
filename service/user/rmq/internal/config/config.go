package config

import (
	"github.com/tal-tech/go-queue/kq"
	"github.com/tal-tech/go-zero/core/service"
)

type Config struct {
	service.ServiceConf
	KqConf kq.KqConf
	//DqConf dq.DqConf
}
