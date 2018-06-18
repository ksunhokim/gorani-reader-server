package gorani

import (
	"context"
	"net/http"
	"time"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	minio "github.com/minio/minio-go"
	"github.com/sunho/gorani-reader-server/go/pkg/work"
	"github.com/yanyiwu/simplelog"
)

const (
	S3DictBucket    = "dict"
	S3PictureBucket = "picture"
	S3Location      = "ap-northeast-2"
)

type Gorani struct {
	Config    Config
	WorkQueue *work.Queue
	Mysql     *gorm.DB
	Redis     *redis.Client
	S3        *minio.Client
	server    *http.Server
}

func (g *Gorani) Start(addr string, handler http.Handler) {
	if g.Config.IsConsumer {
		g.WorkQueue.StartConsuming()
	}
	if g.Config.IsGarbageCollector {
		g.WorkQueue.StartGarbageCollecting()
	}

	g.server = &http.Server{
		Handler: handler,
		Addr:    addr,
	}

	go func() {
		simplelog.Info("http server started | addr: %s", addr)
		err := g.server.ListenAndServe()
		if err != nil {
			simplelog.Fatal("http server failed to listen | err: %v", err)
		}
	}()
}

func (g *Gorani) End() {
	err := g.WorkQueue.End()
	if err != nil {
		simplelog.Error("failed to gracefully shutdown work queue | err: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	err = g.server.Shutdown(ctx)
	if err != nil {
		simplelog.Error("failed to gracefully shutdown http server | err: %v", err)
	}
}

func New(conf Config) (*Gorani, error) {
	mysql, err := createMysqlConn(conf)
	if err != nil {
		return nil, err
	}

	r, err := createRedisConn(conf)
	if err != nil {
		return nil, err
	}

	s, err := createS3(conf)
	if err != nil {
		return nil, err
	}

	w := work.New(r)

	gorn := &Gorani{
		Config:    conf,
		WorkQueue: w,
		Mysql:     mysql,
		Redis:     r,
		S3:        s,
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

	db.DB().SetMaxIdleConns(MysqlConnectionPoolSize)
	db.DB().SetMaxOpenConns(MysqlConnectionLimit)
	db.Exec(`SET @@session.time_zone = '+00:00';`)

	return db, nil
}

func createBucketIfNotExists(m *minio.Client, name string) error {
	exists, err := m.BucketExists(name)
	if err != nil {
		return err
	}

	if !exists {
		err = m.MakeBucket(name, S3Location)
		return err
	}

	return nil
}

func createS3(conf Config) (*minio.Client, error) {
	m, err := minio.New(conf.S3EndPoint, conf.S3Id, conf.S3Secret, conf.S3Ssl)
	if err != nil {
		return nil, err
	}

	err = createBucketIfNotExists(m, S3DictBucket)
	if err != nil {
		return nil, err
	}

	err = createBucketIfNotExists(m, S3PictureBucket)
	if err != nil {
		return nil, err
	}

	return m, err
}

func createRedisConn(conf Config) (*redis.Client, error) {
	opt, err := redis.ParseURL(conf.RedisURL)
	if err != nil {
		return nil, err
	}

	opt.PoolSize = RedisConnectionPoolSize

	client := redis.NewClient(opt)
	_, err = client.Ping().Result()

	return client, err
}
