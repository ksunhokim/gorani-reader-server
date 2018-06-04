package service

import (
	"context"

	pb "github.com/sunho/gorani-reader/server/proto/etl"
)

func (s *Service) BuildSqlite(c context.Context, em *pb.Empty) (*pb.BuildSqliteResponse, error) {
	return nil, nil
}

func (s *Service) InspectIsbn(c context.Context, req *pb.InsepectIsbnRequest) (*pb.InsepectIsbnResponse, error) {
	return nil, nil
}

func (s *Service) InsertBook(c context.Context, req *pb.InsertBookRequest) (*pb.Empty, error) {
	return nil, nil
}
