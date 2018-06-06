package service

import (
	"context"

	"github.com/sunho/gorani-reader/server/etl/relword"
	pb "github.com/sunho/gorani-reader/server/proto/etl"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) CalculateRelevantWords(c context.Context, req *pb.CalculateRelevantWordsRequest) (*pb.Empty, error) {
	err := relword.Calculate(s.e.Mysql, 3, req.Reltype)
	if err != nil {
		err = status.Error(codes.Internal, err.Error())
		return nil, err
	}
	return &pb.Empty{}, nil
}
