package service

import (
	"context"
	"strings"

	"github.com/sunho/gorani-reader/server/pkg/dbh"
	"github.com/sunho/gorani-reader/server/pkg/util"
	pb "github.com/sunho/gorani-reader/server/proto/etl"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func protoWordToDbhWord(word *pb.Word) (out dbh.Word) {
	out = dbh.Word{
		Word:          word.Word,
		Pronunciation: util.BlankToNil(word.Pronunciation),
		Definitions:   []dbh.Definition{},
	}
	for _, def := range word.Definitions {
		odef := dbh.Definition{
			Definition: def.Definition,
			POS:        util.BlankToNil(def.Pos),
			Examples:   []dbh.Example{},
		}
		for _, example := range def.Examples {
			oexample := dbh.Example{
				Foreign: example.Foreign,
				Native:  util.BlankToNil(example.Native),
			}
			odef.Examples = append(odef.Examples, oexample)
		}
		out.Definitions = append(out.Definitions, odef)
	}
	return
}

func dbhWordToProtoWord(word dbh.Word) (out *pb.Word) {
	out = &pb.Word{
		Id:            int32(word.Id),
		Word:          word.Word,
		Pronunciation: util.NilToBlank(word.Pronunciation),
		Definitions:   []*pb.Definition{},
	}
	for _, def := range word.Definitions {
		odef := &pb.Definition{
			Id:         int32(def.Id),
			WordId:     int32(word.Id),
			Definition: def.Definition,
			Pos:        util.NilToBlank(def.POS),
			Examples:   []*pb.Example{},
		}
		for _, example := range def.Examples {
			oexample := &pb.Example{
				DefinitionId: int32(def.Id),
				Foreign:      example.Foreign,
				Native:       util.NilToBlank(example.Native),
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
		out = append(out, &pb.Word{
			Id:            int32(word.Id),
			Word:          word.Word,
			Pronunciation: util.NilToBlank(word.Pronunciation),
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

	return &pb.Empty{}, nil
}

func (s *Service) GetWordById(c context.Context, req *pb.GetWordByIdRequest) (*pb.Word, error) {
	word, err := dbh.GetWordById(s.e.Mysql, int(req.Id))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			err = status.Error(codes.NotFound, err.Error())
		} else {
			err = status.Error(codes.Internal, err.Error())
		}
		return nil, err
	}
	pword := dbhWordToProtoWord(word)
	return pword, nil
}

func (s *Service) GetWordByWord(c context.Context, req *pb.GetWordByWordRequest) (*pb.Word, error) {
	word, err := dbh.GetWordByWord(s.e.Mysql, req.Word)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			err = status.Error(codes.NotFound, err.Error())
		} else {
			err = status.Error(codes.Internal, err.Error())
		}
		return nil, err
	}
	pword := dbhWordToProtoWord(word)
	return pword, nil
}

func (s *Service) DeleteWord(c context.Context, req *pb.DeleteWordRequest) (*pb.Empty, error) {
	word, err := dbh.GetWordById(s.e.Mysql, int(req.Id))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			err = status.Error(codes.NotFound, err.Error())
		} else {
			err = status.Error(codes.Internal, err.Error())
		}
		return nil, err
	}

	err = word.Delete(s.e.Mysql)
	if err != nil {
		err = status.Error(codes.Internal, err.Error())
		return nil, err
	}

	return &pb.Empty{}, nil
}
