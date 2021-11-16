package aws_kq

import "github.com/tal-tech/go-zero/core/service"

type AwsKqConf struct {
	service.ServiceConf
	Brokers []string
	Group   string
	Topic   string
}
