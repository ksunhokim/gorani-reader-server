package api

import (
	"github.com/go-redis/redis"
	"github.com/sunho/gorani-reader/server/api/services"
	"github.com/sunho/gorani-reader/server/pkg/auth"
	"github.com/sunho/gorani-reader/server/pkg/gorani"
)

type Api struct {
	*gorani.Gorani
	Config   Config
	Services auth.Services
	Redis    *redis.Client
}

func New(gorn *gorani.Gorani, conf Config) (*Api, error) {
	r, err := createRedisConn(conf)
	if err != nil {
		return nil, err
	}

	ap := &Api{
		Gorani:   gorn,
		Config:   conf,
		Services: services.New(),
		Redis:    r,
	}
	ap.Config.Config = gorn.Config

	return ap, nil
}

func createRedisConn(conf Config) (*redis.Client, error) {
	opt, err := redis.ParseURL(conf.RedisURL)
	if err != nil {
		return nil, err
	}

	opt.PoolSize = conf.RedisConnectionPoolSize

	client := redis.NewClient(opt)
	_, err = client.Ping().Result()

	return client, err
}
