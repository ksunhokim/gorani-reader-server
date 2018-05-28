package gorani

import (
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Gorani struct {
	Config Config
	Mysql  *gorm.DB
	Redis  *redis.Client
}

func New(conf Config) (*Gorani, error) {
	mysql, err := createMysqlConn(conf)
	if err != nil {
		return nil, err
	}

	gorn := &Gorani{
		Config: conf,
		Mysql:  mysql,
	}

	return gorn, nil
}

func createMysqlConn(conf Config) (*gorm.DB, error) {
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
