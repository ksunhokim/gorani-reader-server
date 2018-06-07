package service

import (
	"bytes"
	"context"
	"fmt"

	"github.com/sunho/gorani-reader/server/etl/book"
	pb "github.com/sunho/gorani-reader/server/proto/etl"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) BuildSqlite(c context.Context, em *pb.Empty) (*pb.BuildSqliteResponse, error) {
	return nil, nil
}

func (s *Service) InspectIsbn(c context.Context, req *pb.InsepectIsbnRequest) (*pb.InsepectIsbnResponse, error) {
	return nil, nil
}

func (s *Service) InsertBook(c context.Context, req *pb.InsertBookRequest) (*pb.Empty, error) {
	// TODO implement grpc file transfer
	cmd := s.e.Redis.Get(req.RedisKey)
	buf, err := cmd.Bytes()
	if err != nil {
		err = status.Error(codes.Internal, err.Error())
		return nil, err
	}

	r := bytes.NewReader(buf)
	b, err := book.Parse("sdfsf", r, int64(len(buf)))
	if err != nil {
		err = status.Error(codes.Internal, err.Error())
		return nil, err
	}
	err = b.UploadCover(s.e.S3)
	if err != nil && err != book.ErrNoCover {
		err = status.Error(codes.Internal, err.Error())
		return nil, err
	}
	fmt.Println(b.Name)
	fmt.Println(b.Author)
	fmt.Println(b.Cover)

	return &pb.Empty{}, nil
}
