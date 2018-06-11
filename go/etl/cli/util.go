package main

import (
	pb "github.com/sunho/gorani-reader-server/go/pkg/proto"
	"google.golang.org/grpc"
)

func makeClient(addr string) (pb.ETLServiceClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}
	cli := pb.NewETLServiceClient(conn)
	return cli, conn, nil
}
