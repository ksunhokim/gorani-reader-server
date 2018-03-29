package dbs

import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"github.com/sunho/engbreaker/pkg/config"
	"github.com/sunho/engbreaker/pkg/dbs"
)

var RDB redis.UniversalClient

func initRedis() {
	addr := config.GetString("REDIS_ADDR", "127.0.0.1:6379")
	pw := config.GetString("REDIS_PW", "")
	db := config.GetInt("REDIS_DB", 0)
	if config.GetString("REDIS_CLUSTER", "false") == "false" {
		RDB = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: pw,
			DB:       db,
		})
	} else {
		// cluster
	}
}

func GenerateRedisNonce(identifier string) string {
	s := Nonce()
	b, err := dbs.RDB.Exists(identifier + s).Result()
	if err != nil {
		logrus.Error(err)
	}
	if b == 1 {
		return GenerateRedisNonce(identifier)
	}
	return s
}
