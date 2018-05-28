package api

import (
	"io/ioutil"

	"github.com/go-redis/redis"
	"github.com/sunho/gorani-reader/server/pkg/auth"
	"github.com/sunho/gorani-reader/server/pkg/gorani"
)

type Api struct {
	Gorn     *gorani.Gorani
	Config   Config
	Services auth.Services
	Redis    *redis.Client
}

func New(gorn *gorani.Gorani, conf Config) (*Api, error) {
	s, err := createServices(conf)
	if err != nil {
		return nil, err
	}

	r, err := createRedisConn(conf)
	if err != nil {
		return nil, err
	}

	ap := &Api{
		Gorn:     gorn,
		Config:   conf,
		Services: s,
		Redis:    r,
	}
	return ap, nil
}

func createServices(conf Config) (auth.Services, error) {
	bytes, err := ioutil.ReadFile(conf.ServicesUrl)
	if err != nil {
		return auth.Services{}, err
	}

	services, err := auth.NewServices(bytes)
	if err != nil {
		return auth.Services{}, err
	}

	return services, nil
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
