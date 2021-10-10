package svc

import (
	"github.com/tal-tech/go-zero/rest"
	"user/service/search/internal/config"
	"user/service/search/internal/middleware"
)

type ServiceContext struct {
	Config     config.Config
	CheckLogin rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		CheckLogin: middleware.NewCheckLoginMiddleware().Handle,
	}
}
