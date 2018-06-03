package etl

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

func (e *Etl) AddWord(c context.Context, word *pb.Word) (resp *pb.AddWordResponse, err error) {
	dword := protoWordToDbhWord(word)
	err = dbh.AddWord(e.Gorn.Mysql, &dword)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate") {
			err = status.Error(codes.AlreadyExists, err.Error())
		} else {
			err = status.Error(codes.Internal, err.Error())
		}
		return
	}

	resp = &pb.AddWordResponse{
		Id: dword.Id,
	}
	return
}
