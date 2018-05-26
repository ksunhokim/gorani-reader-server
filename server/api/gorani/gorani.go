package gorani

import (
	"fmt"
	"io/ioutil"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sunho/gorani-reader/server/api/auth"
	"github.com/sunho/gorani-reader/server/api/config"
	"github.com/sunho/gorani-reader/server/api/log"
)

type Gorani struct {
	Config   config.Config
	Mysql    *gorm.DB
	Redis    *redis.Client
	Logger   log.Logger
	Services auth.Services
}

func NewGorani(conf config.Config) (*Gorani, error) {
	mysql, err := createMysqlConn(conf)
	if err != nil {
		return nil, err
	}

	redis, err := createRedisConn(conf)
	if err != nil {
		return nil, err
	}

	l, err := createLogger(conf)
	if err != nil {
		return nil, err
	}

	s, err := createServices(conf)

	gorn := &Gorani{
		Config:   conf,
		Mysql:    mysql,
		Redis:    redis,
		Logger:   l,
		Services: s,
	}

	return gorn, nil
}

func createServices(conf config.Config) (auth.Services, error) {
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

func createLogger(conf config.Config) (log.Logger, error) {
	switch conf.LoggerType {
	case config.LoggerTypeStdout:
		return log.NewStdoutLogger(), nil

	case config.LoggerTypeFluent:
		l, err := log.NewFluentLogger(conf.FluentHost, conf.FluentPort)
		if err != nil {
			return nil, err
		}

		return l, nil

	case config.LoggerTypeBoth:
		l, err := log.NewFluentLogger(conf.FluentHost, conf.FluentPort)
		if err != nil {
			return nil, err
		}
		s := log.NewStdoutLogger()

		return log.NewBothLogger(l, s), nil

	default:
		return nil, fmt.Errorf("Allowed logger types are stdout, fluent, both")
	}
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
