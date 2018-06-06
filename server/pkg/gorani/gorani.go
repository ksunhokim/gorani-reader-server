package gorani

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	minio "github.com/minio/minio-go"
)

type Gorani struct {
	Config Config
	Mysql  *gorm.DB
	S3     *minio.Client
}

func New(conf Config) (*Gorani, error) {
	mysql, err := createMysqlConn(conf)
	if err != nil {
		return nil, err
	}

	s, err := createS3(conf)
	if err != nil {
		return nil, err
	}

	gorn := &Gorani{
		Config: conf,
		Mysql:  mysql,
		S3:     s,
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
	db.DB().SetMaxOpenConns(conf.MysqlConnectionLimit)
	db.Exec(`SET @@session.time_zone = '+00:00';`)

	return db, nil
}

func createS3(conf Config) (*minio.Client, error) {
	m, err := minio.New(conf.S3EndPoint, conf.S3Id, conf.S3Secret, conf.S3Ssl)
	if err != nil {
		return nil, err
	}

	exists, _ := m.BucketExists("dict")
	if !exists {
		return nil, fmt.Errorf("We don't own bucket: dict")
	}

	exists, _ = m.BucketExists("picture")
	if !exists {
		return nil, fmt.Errorf("We don't own bucket: picture")
	}

	return m, err
}
