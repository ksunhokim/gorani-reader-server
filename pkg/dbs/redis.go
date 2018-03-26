package dbs

import (
	"github.com/go-redis/redis"
	"github.com/sunho/engbreaker/pkg/config"
)

var RDB redis.UniversalClient

func initRedis() {
	addr := config.GetString("REDIS_ADDR", "")
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
