package service

import (
	"context"

	"github.com/sunho/gorani-reader/server/pkg/dbh"
	pb "github.com/sunho/gorani-reader/server/proto/etl"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) CalculateRelevantWords(c context.Context, req *pb.CalculateRelevantWordsRequest) (*pb.Empty, error) {
	words, err := dbh.GetWords(s.e.Mysql)
	if err != nil {
		err = status.Error(codes.Internal, err.Error())
		return nil, err
	}

	err = relcal.Calculate(s.e.Mysql, req.Reltype, words, 3)
	if err != nil {
		err = status.Error(codes.Internal, err.Error())
		return nil, err
	}

	return &pb.Empty{}, nil
}
