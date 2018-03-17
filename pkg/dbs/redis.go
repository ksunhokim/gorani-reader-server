package dbs

import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"github.com/sunho/engbreaker/pkg/config"
)

var RDB *redis.Client

func init() {
	url := config.GetString(config.REDISURL)
	ops, err := redis.ParseURL(url)
	if err != nil {
		logrus.Panic(err)
	}

	tdb := redis.NewClient(ops)
	_, err = tdb.Ping().Result()
	if err != nil {
		logrus.Panic(err)
	}

	RDB = tdb
}
