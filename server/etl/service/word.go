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
	var pron *string
	if word.Pronunciation != "" {
		pron = &word.Pronunciation
	}
	out = dbh.Word{
		Word:          word.Word,
		Pronunciation: pron,
		Definitions:   []dbh.Definition{},
	}
	for _, def := range word.Definitions {
		var pos *string
		if def.Pos != "" {
			pos = &def.Pos
		}
		odef := dbh.Definition{
			Definition: def.Definition,
			POS:        pos,
			Examples:   []dbh.Example{},
		}
		for _, example := range def.Examples {
			var native *string
			if example.Native != "" {
				native = &example.Native
			}
			oexample := dbh.Example{
				Foreign: example.Foreign,
				Native:  native,
			}
			odef.Examples = append(odef.Examples, oexample)
		}
		out.Definitions = append(out.Definitions, odef)
	}
	return
}

func dbhWordToProtoWord(word dbh.Word) (out *pb.Word) {
	pron := ""
	if word.Pronunciation != nil {
		pron = *word.Pronunciation
	}
	out = &pb.Word{
		Id:            word.Id,
		Word:          word.Word,
		Pronunciation: pron,
		Definitions:   []*pb.Definition{},
	}
	for _, def := range word.Definitions {
		pos := ""
		if def.POS != nil {
			pos = *def.POS
		}
		odef := &pb.Definition{
			Id:         def.Id,
			WordId:     word.Id,
			Definition: def.Definition,
			Pos:        pos,
			Examples:   []*pb.Example{},
		}
		for _, example := range def.Examples {
			nat := ""
			if example.Native != nil {
				nat = *example.Native
			}
			oexample := &pb.Example{
				DefinitionId: def.Id,
				Foreign:      example.Foreign,
				Native:       nat,
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

	return &pb.Empty{}, nil
}

func (s *Service) GetWordById(c context.Context, req *pb.GetWordByIdRequest) (*pb.Word, error) {
	word, err := dbh.GetWordById(s.e.Mysql, req.Id)
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
	word, err := dbh.GetWordById(s.e.Mysql, req.Id)
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
