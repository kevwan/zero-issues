package svc

import (
	"gorm.io/gorm"
	"user/service/user/api/db"
	"user/service/user/api/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Mysql  *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := db.Conn(c.MySql.DataSource)
	return &ServiceContext{
		Config: c,
		Mysql:  conn,
	}
}
