package config

import (
	"github.com/tal-tech/go-zero/core/service"
	"rmq/ext/aws_kq"
	"rmq/ext/user_api_client"
)

type Config struct {
	service.ServiceConf
	AwsKqConf   aws_kq.AwsKqConf
	UserApiConf user_api_client.UserApiConfig
}
