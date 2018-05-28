package gorani

import (
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sunho/gorani-reader/server/pkg/config"
)

type Gorani struct {
	Config config.Config
	Mysql  *gorm.DB
	Redis  *redis.Client
}

func New(conf config.Config) (*Gorani, error) {
	mysql, err := createMysqlConn(conf)
	if err != nil {
		return nil, err
	}

	redis, err := createRedisConn(conf)
	if err != nil {
		return nil, err
	}

	gorn := &Gorani{
		Config: conf,
		Mysql:  mysql,
		Redis:  redis,
	}

	return gorn, nil
}

func createMysqlConn(conf config.Config) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", conf.MysqlURL)
	if err != nil {
		return nil, err
	}

	if conf.Debug {
		db.LogMode(true)
	}

	db.DB().SetMaxIdleConns(conf.MysqlConnectionPoolSize)
	db.Exec(`SET @@session.time_zone = '+00:00';`)

	return db, nil
}

func createRedisConn(conf config.Config) (*redis.Client, error) {
	opt, err := redis.ParseURL(conf.RedisURL)
	if err != nil {
		return nil, err
	}

	opt.PoolSize = conf.RedisConnectionPoolSize

	client := redis.NewClient(opt)
	_, err = client.Ping().Result()

	return client, err
}
