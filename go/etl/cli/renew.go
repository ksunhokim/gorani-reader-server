package main

import (
	"context"
	"log"
	"time"

	pb "github.com/sunho/gorani-reader-server/proto/etl"
)

func relevantWords(addr, reltype string) error {
	t := time.Now()

	cli, conn, err := makeClient(addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = cli.CalculateRelevantWords(context.Background(), &pb.CalculateRelevantWordsRequest{
		Reltype: reltype,
	})
	if err != nil {
		return err
	}

	t2 := time.Now().Sub(t)
	log.Println("elasped time: ", t2.Seconds(), " seconds")

	return nil
}
