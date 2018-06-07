package service

import (
	"bytes"
	"context"
	"strings"

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

func (s *Service) AddBook(c context.Context, req *pb.AddBookRequest) (*pb.Empty, error) {
	// TODO implement grpc file transfer
	cmd := s.e.Redis.Get(req.RedisKey)
	buf, err := cmd.Bytes()
	if err != nil {
		err = status.Error(codes.Internal, err.Error())
		return nil, err
	}

	r := bytes.NewReader(buf)
	b, err := book.Parse(req.Isbn, r, int64(len(buf)))
	if err != nil {
		err = status.Error(codes.Internal, err.Error())
		return nil, err
	}
	err = b.UploadCover(s.e.S3)
	if err != nil && err != book.ErrNoCover {
		err = status.Error(codes.Internal, err.Error())
		return nil, err
	}

	err = b.AddToDB(s.e.Mysql)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate") {
			err = status.Error(codes.AlreadyExists, err.Error())
		} else {
			err = status.Error(codes.Internal, err.Error())
		}
		return nil, err
	}

	return &pb.Empty{}, nil
}
