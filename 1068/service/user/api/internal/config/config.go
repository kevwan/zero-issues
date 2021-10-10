package config

import "github.com/tal-tech/go-zero/rest"

type Config struct {
	rest.RestConf

	MySql struct {
		DataSource string //mysql连接地址
	}

	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
}
