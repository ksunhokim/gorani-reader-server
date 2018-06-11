package main

import (
	"context"
	"io/ioutil"
	"time"

	"github.com/go-redis/redis"
	pb "github.com/sunho/gorani-reader-server/go/pkg/proto"
)

func addBook(isbn string, epub string, addr string, redisurl string) error {
	buf, err := ioutil.ReadFile(epub)
	if err != nil {
		return err
	}

	cli, conn, err := makeClient(addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	opt, err := redis.ParseURL(redisurl)
	if err != nil {
		return err
	}

	red := redis.NewClient(opt)
	_, err = red.Ping().Result()
	if err != nil {
		return err
	}

	red.Set(isbn+"_book_temp", buf, time.Hour)

	_, err = cli.AddBook(context.Background(), &pb.AddBookRequest{
		RedisKey: isbn + "_book_temp",
		Isbn:     isbn,
	})

	if err != nil {
		return err
	}

	return nil
}
