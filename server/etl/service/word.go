package service

import (
	"context"
	"strings"

	"github.com/sunho/gorani-reader/server/pkg/dbh"
	pb "github.com/sunho/gorani-reader/server/proto/etl"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func protoWordToDbhWord(word *pb.Word) (out dbh.Word) {
	out = dbh.Word{
		Word:          word.Word,
		Pronunciation: &word.Pronunciation,
		Definitions:   []dbh.Definition{},
	}
	for _, def := range word.Definitions {
		odef := dbh.Definition{
			Definition: def.Definition,
			POS:        &def.Pos,
			Examples:   []dbh.Example{},
		}
		for _, example := range def.Examples {
			oexample := dbh.Example{
				Foreign: example.Foreign,
				Native:  &example.Native,
			}
			odef.Examples = append(odef.Examples, oexample)
		}
		out.Definitions = append(out.Definitions, odef)
	}
	return
}

func (s *Service) GetWords(c context.Context, em *pb.Empty) (*pb.GetWordsResponse, error) {
	words, err := dbh.GetWords(s.e.Mysql)
	if err != nil {
		err = status.Error(codes.Internal, err.Error())
		return nil, err
	}

	out := make([]*pb.Word, 0, len(words))
	for _, word := range words {
		pron := ""
		if word.Pronunciation != nil {
			pron = *word.Pronunciation
		}

		out = append(out, &pb.Word{
			Id:            word.Id,
			Word:          word.Word,
			Pronunciation: pron,
		})
	}

	resp := &pb.GetWordsResponse{
		Words: out,
	}
	return resp, nil
}

func (s *Service) AddWord(c context.Context, req *pb.AddWordRequest) (*pb.Empty, error) {
	dword := protoWordToDbhWord(req.Word)
	err := dbh.AddWord(s.e.Mysql, &dword)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate") {
			err = status.Error(codes.AlreadyExists, err.Error())
		} else {
			err = status.Error(codes.Internal, err.Error())
		}
		return nil, err
	}

	return nil, nil
}

func (s *Service) GetWordById(c context.Context, req *pb.GetWordByIdRequest) (*pb.Word, error) {
	return nil, nil
}

func (s *Service) GetWordByWord(c context.Context, req *pb.GetWordByWordRequest) (*pb.Word, error) {
	return nil, nil
}

func (s *Service) DeleteWord(c context.Context, req *pb.DeleteWordRequest) (*pb.Empty, error) {
	return nil, nil
}
